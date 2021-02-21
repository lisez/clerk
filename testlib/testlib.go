package testlib

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var basePath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basePath = filepath.Dir(currentFile)
}

// GetTestPath ...
func GetTestPath(rel string) string {
	return filepath.Join(basePath, rel)
}

// CaptureOutput ...
func CaptureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stdout)
	}()
	f()
	return buf.String()
}
