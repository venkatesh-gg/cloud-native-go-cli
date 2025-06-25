package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/venkateshgogula/cloud-native-go-cli/pkg/config"
)

// statusCmd shows the current service port configuration
var statusCmd = &cobra.Command{
    Use:   "status",
    Short: "Show service status",
    Run: func(cmd *cobra.Command, args []string) {
        cfg := config.LoadConfig()
        fmt.Printf("🔍 Service listening on port %s\n", cfg.Port)
    },
}

func init() {
    rootCmd.AddCommand(statusCmd)
}

