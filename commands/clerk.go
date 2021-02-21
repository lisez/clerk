package commands

import (
	"clerk/common/fslib"
	"clerk/common/getter"
	"clerk/jobs"
	"clerk/jobs/providers"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var configFlag string

func runWithConfig() {
	if !fslib.IsFileExist(configFlag) {
		log.Fatalf("no such config file: %s", configFlag)
	}

	config := jobs.NewClerkConfig(configFlag)
	for _, remoteConfig := range config.Clerk.SourceRemotes {
		if len(remoteConfig.Schema) == 0 {
			remoteConfig.Schema = config.Clerk.Schema
		}

		switch provider := remoteConfig.GetProvider(); provider {
		case "mongodb":
			runner := &providers.MongodbProvider{
				Config:     remoteConfig,
				Timeout:    30 * time.Second,
				Datebase:   getter.GetValueAsString(remoteConfig.Args, "database", ""),
				Collection: getter.GetValueAsString(remoteConfig.Args, "collection", ""),
			}
			runner.Check()
			runner.Start()
		default:
			log.Fatalf("unknown provider: %s", provider)
		}
	}
}

func rootCmdRunner(cmd *cobra.Command, args []string) {
	if len(configFlag) > 0 {
		runWithConfig()
	} else if len(args) > 1 {
		jobs.IsALLJSONFiles(args)
		jobs.ValidateJSONFiles(args[0], args[1:])
	}
}

// Execute ...
func Execute() {
	var rootCmd = &cobra.Command{
		Use: "clerk [<schema> <file> [files...]]",
		Run: rootCmdRunner,
	}

	rootCmd.PersistentFlags().StringVarP(&configFlag, "config", "c", "", "config file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
