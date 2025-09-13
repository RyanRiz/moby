package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Version   = "1.0.0"
	BuildDate = "2025-09-13"
)

// NewVersionCommand creates a new version command
func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Long:  "Display version and build information for Moby Container Manager",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Moby Container Manager\n")
			fmt.Printf("Version: %s\n", Version)
			fmt.Printf("Build Date: %s\n", BuildDate)
			fmt.Printf("Built with: Go + Moby/Docker components\n")
		},
	}
}
