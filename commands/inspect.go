package commands

import (
	"clerk/common/fslib"
	"log"

	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
)

func isAllJSONFiles(filepaths []string) {
	for _, file := range filepaths {
		if !fslib.IsFileExist(file) || !fslib.IsJSONFile(file) {
			log.Fatalf("no such file or not a json file: %s", file)
		}
	}
}

func executeCommand(filepaths []string) {
	schemaLoader := gojsonschema.NewReferenceLoader(fslib.WithFileProtocol(filepaths[0]))
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		log.Fatalf("load json schema failed: %s", err)
	}

	for _, file := range filepaths[1:] {
		targetJSONFile := gojsonschema.NewReferenceLoader(fslib.WithFileProtocol(file))
		result, err := schema.Validate(targetJSONFile)
		if err != nil {
			log.Fatalf("load json file failed: %s", err)
		}

		if result.Valid() {
			log.Printf("%s: pass", file)
		} else {
			// TODO: move to warning channel
			log.Printf("%s: invalid JSON, reasons: %s", file, result.Errors())
		}
	}
}

func inspectCmdAction(cmd *cobra.Command, args []string) {
	isAllJSONFiles(args)
	executeCommand(args)
}

// InspectCmd ...
func InspectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect [schema] [...files]",
		Short: "validate files by a json schema",
		Args:  cobra.MinimumNArgs(2),
		Run:   inspectCmdAction,
	}
	return cmd
}
