package fslib

import (
	"os"
	"path/filepath"
)

// IsFileExist ...
func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// IsJSONFile ...
func IsJSONFile(filename string) bool {
	return filepath.Ext(filename) == ".json"
}

// WithFileProtocol ...
func WithFileProtocol(filepath string) string {
	return "file://" + filepath
}
