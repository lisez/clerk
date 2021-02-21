package commands

import (
	"clerk/jobs"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func executeCommand(cmd *cobra.Command, args []string) {
	jobs.IsALLJSONFiles(args)
	jobs.ValidateJSONFiles(args[0], args[1:])
}

// Execute ...
func Execute() {
	var uriFlag string

	var rootCmd = &cobra.Command{
		Use:  "clerk <schema> [files...]",
		Args: cobra.MinimumNArgs(1),
		Run:  executeCommand,
	}

	rootCmd.PersistentFlags().StringVar(&uriFlag, "uri", "", "specify a supported uri address to load json docs from it")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
