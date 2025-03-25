package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	restart_counter()
}

func restart_counter() {
	// Create a new Docker API client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	// Fetch running containers
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: false})
	if err != nil {
		log.Fatalf("Error fetching containers: %v", err)
	}

	// Exits if there is no running container
	if len(containers) == 0 {
		log.Fatalf("No running container")
	}

	// Iterate through containers and print their restart count
	for _, container := range containers {
		// Fetch detailed container info
		inspect, err := cli.ContainerInspect(context.Background(), container.ID)
		if err != nil {
			log.Printf("Error inspecting container %s: %v", container.ID, err)
			continue
		}

		// Print the restart count
		fmt.Printf("Restart Count: %d\t Container ID: %s\tName: %s\tCreated at: %v\n",
			inspect.RestartCount, container.ID[:12], container.Names[0], time.Unix(container.Created, 0))
	}

}
