package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

// purgeCmd is a stub for clearing service data
var purgeCmd = &cobra.Command{
    Use:   "purge",
    Short: "Purge all service data (stub)",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("✅ Service data purged (stub).")
    },
}

func init() {
    rootCmd.AddCommand(purgeCmd)
}

