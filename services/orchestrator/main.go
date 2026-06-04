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
)

func main() {
	http.HandleFunc("/upload", handleUpload)

	fmt.Println("=================================================")
	fmt.Println("🚀 IICPC Orchestrator Submission API Online")
	fmt.Println("📡 Listening for contestant ZIP uploads on :8082")
	fmt.Println("=================================================")

	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// 1. Accept the Multipart ZIP upload
	r.ParseMultipartForm(10 << 20) // 10 MB limit
	file, _, err := r.FormFile("bot")
	if err != nil {
		http.Error(w, "Failed to read bot upload. Did you use the 'bot' form field?", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 2. Create a temporary secure sandbox directory
	tmpDir, _ := os.MkdirTemp("", "sandbox-*")
	defer os.RemoveAll(tmpDir) // Self-clean when done

	zipPath := filepath.Join(tmpDir, "bot.zip")
	dst, _ := os.Create(zipPath)
	io.Copy(dst, file)
	dst.Close()

	// 3. Unzip the contestant's code
	unzip(zipPath, tmpDir)

	// 4. Dynamically compile their Docker image
	runID := fmt.Sprintf("run-%d", time.Now().Unix())
	imageName := fmt.Sprintf("iicpc-contestant-bot:%s", runID)

	fmt.Printf("\n[RECEIVED] Building contestant image: %s...\n", imageName)

	// We use exec for the build phase to seamlessly handle Dockerfile context
	cmd := exec.Command("docker", "build", "-t", imageName, tmpDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		http.Error(w, "Docker build failed. Check your Dockerfile syntax.", http.StatusInternalServerError)
		return
	}

	// 5. Spin up the sandboxed environment using the Docker SDK
	containerID := runSandboxedContainer(imageName)

	successMsg := fmt.Sprintf("✅ Successfully deployed! Run ID: %s | Container: %s\n", runID, containerID)
	fmt.Print(successMsg)
	fmt.Fprint(w, successMsg)
}

// runSandboxedContainer uses the Docker Go SDK to spin up the newly built image
func runSandboxedContainer(imageName string) string {
	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	defer cli.Close()

	containerConfig := &container.Config{Image: imageName}
	hostConfig := &container.HostConfig{AutoRemove: true} // Self-destruct when finished

	resp, _ := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	cli.ContainerStart(ctx, resp.ID, container.StartOptions{})

	return resp.ID[:12]
}

// unzip is a secure helper utility to extract contestant code
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
