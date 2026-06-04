package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	fmt.Println("Starting Platform Orchestrator...")
	ctx := context.Background()

	// Connect to the local Docker Daemon
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Failed to connect to Docker: %v\n", err)
		return
	}
	defer cli.Close()

	// Define the container configuration (Image and exposed port)
	containerConfig := &container.Config{
		Image: "iicpc-dummy-sut",
		ExposedPorts: nat.PortSet{
			"8080/tcp": struct{}{},
		},
	}

	// Define host binding (Map internal 8080 to external 8080)
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"8080/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8080",
				},
			},
		},
		AutoRemove: true, // Automatically delete the container when it stops
	}

	fmt.Println("Provisioning sandboxed SUT environment...")

	// Create the container
	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "orchestrated-sut")
	if err != nil {
		fmt.Printf("Failed to create container: %v\n", err)
		return
	}

	// Start the container
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		fmt.Printf("Failed to start container: %v\n", err)
		return
	}

	fmt.Printf("Successfully started isolated SUT container. ID: %s\n", resp.ID[:12])
	fmt.Println("The SUT is now running and mapped to port 8080.")

	// Keep the orchestrator running to monitor the container
	fmt.Println("Press Ctrl+C to shut down Orchestrator (container will self-destruct).")
	for {
		time.Sleep(1 * time.Second)
	}
}
