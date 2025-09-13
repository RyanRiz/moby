package main

import (
	"context"
	"fmt"
	"os"

	"moby/internal/client"
	"moby/internal/commands"

	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()

	// Initialize Docker client
	dockerClient, err := client.NewDockerClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Docker client: %v\n", err)
		os.Exit(1)
	}
	defer dockerClient.Close()

	// Create root command
	rootCmd := &cobra.Command{
		Use:   "mcm",
		Short: "Moby Container Manager - A simple container management tool",
		Long: `Moby Container Manager (mcm) adalah aplikasi CLI yang menggunakan 
komponen Moby/Docker untuk mengelola container dengan mudah.`,
	}

	// Add commands
	rootCmd.AddCommand(
		commands.NewListCommand(dockerClient),
		commands.NewStartCommand(dockerClient),
		commands.NewStopCommand(dockerClient),
		commands.NewInspectCommand(dockerClient),
		commands.NewRemoveCommand(dockerClient),
		commands.NewPullCommand(dockerClient),
		commands.NewVersionCommand(),
	)

	// Execute
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
