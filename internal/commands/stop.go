package commands

import (
	"context"
	"fmt"

	"moby/internal/client"

	"github.com/spf13/cobra"
)

// NewStopCommand creates a new stop command
func NewStopCommand(dockerClient *client.DockerClient) *cobra.Command {
	return &cobra.Command{
		Use:   "stop <container>",
		Short: "Stop a container",
		Long:  "Stop a running container by ID or name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return stopContainer(cmd.Context(), dockerClient, args[0])
		},
	}
}

func stopContainer(ctx context.Context, dockerClient *client.DockerClient, containerID string) error {
	fmt.Printf("Stopping container %s...\n", containerID)

	err := dockerClient.StopContainer(ctx, containerID)
	if err != nil {
		return fmt.Errorf("failed to stop container %s: %w", containerID, err)
	}

	fmt.Printf("Container %s stopped successfully\n", containerID)
	return nil
}
