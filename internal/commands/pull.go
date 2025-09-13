package commands

import (
	"context"
	"fmt"
	"io"
	"os"

	"moby/internal/client"

	"github.com/spf13/cobra"
)

// NewPullCommand creates a new pull command
func NewPullCommand(dockerClient *client.DockerClient) *cobra.Command {
	return &cobra.Command{
		Use:   "pull <image>",
		Short: "Pull an image",
		Long:  "Pull an image from Docker registry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return pullImage(cmd.Context(), dockerClient, args[0])
		},
	}
}

func pullImage(ctx context.Context, dockerClient *client.DockerClient, imageName string) error {
	fmt.Printf("Pulling image %s...\n", imageName)

	reader, err := dockerClient.PullImage(ctx, imageName)
	if err != nil {
		return fmt.Errorf("failed to pull image %s: %w", imageName, err)
	}
	defer reader.Close()

	// Copy pull output to stdout
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		return fmt.Errorf("failed to read pull output: %w", err)
	}

	fmt.Printf("\nImage %s pulled successfully\n", imageName)
	return nil
}
