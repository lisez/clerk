package commands

import (
	"clerk/testlib"
	"strings"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	schema := testlib.GetTestPath("./data/test_schema.json")

	t.Run("invalid json", func(subT *testing.T) {
		file := testlib.GetTestPath("./data/test_invalid_doc.json")
		message := testlib.CaptureOutput(func() {
			executeCommand([]string{schema, file})
		})

		if !strings.Contains(message, "invalid JSON") {
			subT.Error("it is an invalid json, but passed")
		}
	})

	t.Run("valid json", func(subT *testing.T) {
		file := testlib.GetTestPath("./data/test_valid_doc.json")
		message := testlib.CaptureOutput(func() {
			executeCommand([]string{schema, file})
		})

		if !strings.Contains(message, "pass") {
			subT.Error("it is a valid json, but failed")
		}
	})
}
