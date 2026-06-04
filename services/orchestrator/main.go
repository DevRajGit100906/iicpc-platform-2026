package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func main() {
	// Initialize Redis Connection
	rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	http.HandleFunc("/upload", handleUpload)

	fmt.Println("=================================================")
	fmt.Println("🚀 IICPC Orchestrator Submission API Online")
	fmt.Println("📡 Listening for contestant ZIP uploads on :8082")
	fmt.Println("=================================================")

	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}

// updateMatchState acts as the central state machine publisher
func updateMatchState(runID string, state string) {
	ctx := context.Background()
	// Push the state to a Redis Hash
	err := rdb.HSet(ctx, "match:"+runID, "status", state, "updated_at", time.Now().Unix()).Err()
	if err != nil {
		fmt.Printf("[REDIS ERROR] Failed to push state: %v\n", err)
	} else {
		fmt.Printf("🔵 [STATE CHANGE] %s -> %s\n", runID, state)
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("bot")
	if err != nil {
		http.Error(w, "Failed to read bot upload.", http.StatusBadRequest)
		return
	}
	defer file.Close()

	tmpDir, _ := os.MkdirTemp("", "sandbox-*")
	defer os.RemoveAll(tmpDir)

	zipPath := filepath.Join(tmpDir, "bot.zip")
	dst, _ := os.Create(zipPath)
	io.Copy(dst, file)
	dst.Close()

	unzip(zipPath, tmpDir)

	runID := fmt.Sprintf("run-%d", time.Now().Unix())
	imageName := fmt.Sprintf("iicpc-contestant-bot:%s", runID)

	// STATE: BUILDING
	updateMatchState(runID, "BUILDING")

	cmd := exec.Command("docker", "build", "-t", imageName, tmpDir)
	if err := cmd.Run(); err != nil {
		updateMatchState(runID, "FAILED")
		http.Error(w, "Docker build failed.", http.StatusInternalServerError)
		return
	}

	// Spin up the sandbox
	containerID := runSandboxedContainer(imageName, runID)

	fmt.Fprintf(w, "✅ Deployed! Run ID: %s | Container: %s\n", runID, containerID)
}

func runSandboxedContainer(imageName string, runID string) string {
	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	defer cli.Close()

	containerConfig := &container.Config{Image: imageName}
	hostConfig := &container.HostConfig{
		AutoRemove: true,
		Resources: container.Resources{
			Memory:   256 * 1024 * 1024,
			NanoCPUs: 500000000,
		},
	}

	resp, _ := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	cli.ContainerStart(ctx, resp.ID, container.StartOptions{})

	// STATE: RUNNING
	updateMatchState(runID, "RUNNING")

	// 15-Second Timebomb
	go func(cID string, rID string) {
		time.Sleep(15 * time.Second)
		fmt.Printf("\n⏱️ Timebomb activated! Terminating: %s\n", cID[:12])
		cli.ContainerStop(context.Background(), cID, container.StopOptions{})

		// STATE: FINISHED
		updateMatchState(rID, "FINISHED")
	}(resp.ID, runID)

	return resp.ID[:12]
}

func unzip(src, dest string) error {
	r, _ := zip.OpenReader(src)
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
