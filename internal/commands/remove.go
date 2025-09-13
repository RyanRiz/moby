package commands

import (
	"context"
	"fmt"

	"moby/internal/client"

	"github.com/spf13/cobra"
)

// NewRemoveCommand creates a new remove command
func NewRemoveCommand(dockerClient *client.DockerClient) *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "remove <container>",
		Short: "Remove a container",
		Long:  "Remove a container by ID or name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return removeContainer(cmd.Context(), dockerClient, args[0], force)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force remove running container")

	return cmd
}

func removeContainer(ctx context.Context, dockerClient *client.DockerClient, containerID string, force bool) error {
	fmt.Printf("Removing container %s...\n", containerID)

	err := dockerClient.RemoveContainer(ctx, containerID, force)
	if err != nil {
		return fmt.Errorf("failed to remove container %s: %w", containerID, err)
	}

	fmt.Printf("Container %s removed successfully\n", containerID)
	return nil
}
