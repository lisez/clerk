package jobs

import (
	"clerk/common/fslib"
	"fmt"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

// IsALLJSONFiles ...
func IsALLJSONFiles(filepaths []string) {
	for _, file := range filepaths {
		if !fslib.IsFileExist(file) || !fslib.IsJSONFile(file) {
			log.Fatalf("no such file or not a json file: %s", file)
		}
	}
}

// ValidateJSONFiles ...
func ValidateJSONFiles(schemaPath string, filepaths []string) {
	schemaLoader := gojsonschema.NewReferenceLoader(fslib.WithFileProtocol(schemaPath))
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		log.Fatalf("load json schema failed: %s", err)
	}

	for _, file := range filepaths {
		docLoader := gojsonschema.NewReferenceLoader(fslib.WithFileProtocol(file))
		result, err := schema.Validate(docLoader)
		if err != nil {
			log.Fatalf("load json file failed: %s", err)
		}

		if result.Valid() {
			fmt.Printf("%s: pass\n", file)
		} else {
			// TODO: move to warning channel
			fmt.Printf("%s: invalid JSON, reasons: %s\n", file, result.Errors())
		}
	}
}
