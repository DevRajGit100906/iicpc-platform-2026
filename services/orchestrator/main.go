package main
import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
	botfleet "github.com/devrajdas/iicpc-platform-2026/services/bot-fleet"
	"github.com/devrajdas/iicpc-platform-2026/services/scoring"
	"github.com/devrajdas/iicpc-platform-2026/services/scoring/correctness"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/redis/go-redis/v9"
)
const enginePort = nat.Port("8080/tcp")
var rdb *redis.Client
func main() {
	rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	http.HandleFunc("/upload", handleUpload)
	fmt.Println("IICPC Orchestrator on :8082  (engine-as-SUT; correctness + load; gated composite)")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
func setStatus(runID, state string) {
	ctx := context.Background()
	rdb.HSet(ctx, "run:"+runID, "status", state, "updated_at", time.Now().Unix())
	if state == "BUILDING" {
		rdb.ZAdd(ctx, "live_leaderboard", redis.Z{Score: 0, Member: runID})
	}
	fmt.Printf("[STATE] %s -> %s\n", runID, state)
}
func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(64 << 20)
	file, _, err := r.FormFile("submission")
	if err != nil {
		file, _, err = r.FormFile("bot")
	}
	if err != nil {
		http.Error(w, "missing submission upload", http.StatusBadRequest)
		return
	}
	defer file.Close()
	tmpDir, _ := os.MkdirTemp("", "sandbox-*")
	zipPath := filepath.Join(tmpDir, "submission.zip")
	dst, _ := os.Create(zipPath)
	io.Copy(dst, file)
	dst.Close()
	unzip(zipPath, tmpDir)
	runID := fmt.Sprintf("run-%d", time.Now().UnixNano())
	imageName := fmt.Sprintf("iicpc-submission:%s", runID)
	setStatus(runID, "BUILDING")
	if err := exec.Command("docker", "build", "-t", imageName, tmpDir).Run(); err != nil {
		setStatus(runID, "FAILED")
		os.RemoveAll(tmpDir)
		http.Error(w, "docker build failed", http.StatusInternalServerError)
		return
	}
	go runBenchmark(imageName, runID, tmpDir)
	fmt.Fprintf(w, "Accepted. Run ID: %s\n", runID)
}
type scoringSink struct {
	inner botfleet.Sink
	mu    sync.Mutex
	lat   []int64
	errs  int64
}
func (s *scoringSink) Emit(t botfleet.Telemetry) error {
	s.mu.Lock()
	if t.Status == "error" {
		s.errs++
	} else {
		s.lat = append(s.lat, t.LatencyMicro)
	}
	s.mu.Unlock()
	if s.inner != nil {
		return s.inner.Emit(t)
	}
	return nil
}
func runBenchmark(imageName, runID, tmpDir string) {
	defer os.RemoveAll(tmpDir)
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		setStatus(runID, "FAILED")
		return
	}
	defer cli.Close()
	pids := int64(256)
	resp, err := cli.ContainerCreate(ctx,
		&container.Config{
			Image:        imageName,
			ExposedPorts: nat.PortSet{enginePort: struct{}{}},
		},
		&container.HostConfig{
			AutoRemove:   true,
			PortBindings: nat.PortMap{enginePort: []nat.PortBinding{{HostIP: "127.0.0.1", HostPort: ""}}},
			CapDrop:      []string{"ALL"},
			SecurityOpt:  []string{"no-new-privileges"},
			Resources: container.Resources{
				Memory:    256 * 1024 * 1024,
				NanoCPUs:  500_000_000,
				PidsLimit: &pids,
			},
		}, nil, nil, "")
	if err != nil {
		setStatus(runID, "FAILED")
		return
	}
	engineID := resp.ID
	if err := cli.ContainerStart(ctx, engineID, container.StartOptions{}); err != nil {
		setStatus(runID, "FAILED")
		return
	}
	defer cli.ContainerStop(context.Background(), engineID, container.StopOptions{})
	if out, err := cli.ContainerLogs(ctx, engineID,
		container.LogsOptions{ShowStdout: true, ShowStderr: true, Follow: true}); err == nil {
		go io.Copy(os.Stdout, out)
	}
	info, err := cli.ContainerInspect(ctx, engineID)
	if err != nil || len(info.NetworkSettings.Ports[enginePort]) == 0 {
		setStatus(runID, "FAILED")
		return
	}
	addr := "127.0.0.1:" + info.NetworkSettings.Ports[enginePort][0].HostPort
	orderURL := "http://" + addr + "/order"
	setStatus(runID, "DEPLOYING")
	if err := waitTCP(addr, 20*time.Second); err != nil {
		setStatus(runID, "FAILED")
		return
	}
	setStatus(runID, "CHECKING")
	orders := correctness.GenerateScenario(time.Now().UnixNano(), 1000)
	rep, err := correctness.Check(orders, correctness.NewHTTPClient(orderURL))
	if err != nil {
		setStatus(runID, "FAILED")
		return
	}
	setStatus(runID, "RUNNING")
	ksink := botfleet.NewKafkaSink("localhost:19092", "telemetry")
	defer ksink.Close()
	sink := &scoringSink{inner: ksink}
	runCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	start := time.Now()
	res := botfleet.Run(runCtx,
		botfleet.NewHTTPEngine(orderURL), sink,
		botfleet.Config{RunID: runID, Bots: 50, OrdersPerBot: 200, Seed: time.Now().UnixNano()})
	dur := time.Since(start)
	pcts := scoring.ComputePercentiles(sink.lat)
	tps := 0.0
	if dur.Seconds() > 0 {
		tps = float64(res.Sent) / dur.Seconds()
	}
	result := scoring.Compute(scoring.Input{
		Lat:               pcts,
		TPS:               tps,
		Correctness:       rep.Score,
		CorrectnessPassed: rep.Passed,
	}, scoring.DefaultConfig())
	rdb.ZAdd(ctx, "live_leaderboard", redis.Z{Score: result.Composite, Member: runID})
	rdb.HSet(ctx, "run:"+runID,
		"status", "FINISHED",
		"score", fmt.Sprintf("%.2f", result.Composite),
		"correctness", fmt.Sprintf("%.4f", rep.Score),
		"passed", boolStr(rep.Passed),
		"gated", boolStr(result.Gated),
		"p99_us", pcts.P99,
		"p50_us", pcts.P50,
		"tps", fmt.Sprintf("%.0f", tps),
		"orders", res.Sent,
		"errors", sink.errs,
		"updated_at", time.Now().Unix(),
	)
	fmt.Printf("[SCORE] run=%s composite=%.2f correctness=%.2f(passed=%v gated=%v) p99=%dus tps=%.0f\n",
		runID, result.Composite, rep.Score, rep.Passed, result.Gated, pcts.P99, tps)
	if rep.FirstViolation != nil {
		fmt.Printf("[CORRECTNESS] first violation at trade #%d (order seq %d): %+v\n",
			rep.FirstViolation.Index, rep.FirstViolation.TakerSeq, rep.FirstViolation)
	}
}
func boolStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
func waitTCP(addr string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if c, err := net.DialTimeout("tcp", addr, 500*time.Millisecond); err == nil {
			c.Close()
			return nil
		}
		time.Sleep(300 * time.Millisecond)
	}
	return fmt.Errorf("engine not ready at %s within %s", addr, timeout)
}
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
		outFile, _ := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		rc, _ := f.Open()
		io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
	}
	return nil
}
