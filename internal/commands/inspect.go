package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"moby/internal/client"

	"github.com/spf13/cobra"
)

// NewInspectCommand creates a new inspect command
func NewInspectCommand(dockerClient *client.DockerClient) *cobra.Command {
	var pretty bool

	cmd := &cobra.Command{
		Use:   "inspect <container>",
		Short: "Inspect a container",
		Long:  "Display detailed information about a container",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return inspectContainer(cmd.Context(), dockerClient, args[0], pretty)
		},
	}

	cmd.Flags().BoolVarP(&pretty, "pretty", "p", true, "Pretty print JSON output")

	return cmd
}

func inspectContainer(ctx context.Context, dockerClient *client.DockerClient, containerID string, pretty bool) error {
	containerInfo, err := dockerClient.InspectContainer(ctx, containerID)
	if err != nil {
		return fmt.Errorf("failed to inspect container %s: %w", containerID, err)
	}

	var output []byte
	if pretty {
		output, err = json.MarshalIndent(containerInfo, "", "  ")
	} else {
		output, err = json.Marshal(containerInfo)
	}

	if err != nil {
		return fmt.Errorf("failed to marshal container info: %w", err)
	}

	fmt.Println(string(output))
	return nil
}
