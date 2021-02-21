package commands

import (
	"clerk/common/fslib"
	"clerk/jobs"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var configFlag string

func rootCmdRunner(cmd *cobra.Command, args []string) {
	if len(configFlag) > 0 {
		if !fslib.IsFileExist(configFlag) {
			log.Fatalf("no such config file: %s", configFlag)
		}

		config := jobs.NewClerkConfig(configFlag)
		for _, remote := range config.Clerk.SourceRemotes {
			provider := remote.GetProvider()
			log.Print(provider)
		}
	} else if len(args) > 1 {
		jobs.IsALLJSONFiles(args)
		jobs.ValidateJSONFiles(args[0], args[1:])
	}
}

// Execute ...
func Execute() {
	var rootCmd = &cobra.Command{
		Use: "clerk [schema] [files...]",
		Run: rootCmdRunner,
	}

	rootCmd.PersistentFlags().StringVarP(&configFlag, "config", "c", "", "use config file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
