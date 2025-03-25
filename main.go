package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: false})
	if err != nil {
		log.Fatalf("Error fetching containers: %v", err)
	}

	for _, container := range containers {
		inspect, err := cli.ContainerInspect(context.Background(), container.ID)
		if err != nil {
			log.Printf("Error inspecting container %s: %v", container.ID, err)
			continue
		}

		fmt.Printf("Container: %s (Name: %s) | Restart Count: %d\n",
			container.ID[:12], container.Names[0], inspect.RestartCount)
	}
}
