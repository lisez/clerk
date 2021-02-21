package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute ...
func Execute() {
	var rootCmd = &cobra.Command{
		Use: "clerk",
		Run: func(cmd *cobra.Command, args []string) {
			// do something
		},
	}

	rootCmd.AddCommand(InspectCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
