package commands

import (
	"context"
	"fmt"
	"time"

	"moby/internal/client"

	"github.com/spf13/cobra"
)

// NewListCommand creates a new list command
func NewListCommand(dockerClient *client.DockerClient) *cobra.Command {
	var all bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List containers",
		Long:  "List all running containers, or all containers with --all flag",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listContainers(cmd.Context(), dockerClient, all)
		},
	}

	cmd.Flags().BoolVarP(&all, "all", "a", false, "Show all containers (default shows just running)")

	return cmd
}

func listContainers(ctx context.Context, dockerClient *client.DockerClient, all bool) error {
	containers, err := dockerClient.ListContainers(ctx, all)
	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	if len(containers) == 0 {
		fmt.Println("No containers found")
		return nil
	}

	// Print header
	fmt.Printf("%-12s %-20s %-15s %-20s %s\n",
		"CONTAINER ID", "IMAGE", "STATUS", "CREATED", "NAMES")
	fmt.Println("----------------------------------------------------------------------")

	// Print containers
	for _, container := range containers {
		created := time.Unix(container.Created, 0).Format("2006-01-02 15:04:05")
		containerID := container.ID[:12] // Short ID

		var names string
		if len(container.Names) > 0 {
			names = container.Names[0][1:] // Remove leading slash
		}

		fmt.Printf("%-12s %-20s %-15s %-20s %s\n",
			containerID, container.Image, container.Status, created, names)
	}

	return nil
}
