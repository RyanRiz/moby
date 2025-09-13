package commands

import (
	"context"
	"fmt"

	"moby/internal/client"

	"github.com/spf13/cobra"
)

// NewStartCommand creates a new start command
func NewStartCommand(dockerClient *client.DockerClient) *cobra.Command {
	return &cobra.Command{
		Use:   "start <container>",
		Short: "Start a container",
		Long:  "Start a stopped container by ID or name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return startContainer(cmd.Context(), dockerClient, args[0])
		},
	}
}

func startContainer(ctx context.Context, dockerClient *client.DockerClient, containerID string) error {
	fmt.Printf("Starting container %s...\n", containerID)

	err := dockerClient.StartContainer(ctx, containerID)
	if err != nil {
		return fmt.Errorf("failed to start container %s: %w", containerID, err)
	}

	fmt.Printf("Container %s started successfully\n", containerID)
	return nil
}
