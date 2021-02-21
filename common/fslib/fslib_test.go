package fslib

import (
	"clerk/testlib"
	"testing"
)

type testSpec struct {
	filename   string
	expected   bool
	targetFunc func(filename string) bool
}

func (spec *testSpec) Test(t *testing.T) {
	result := spec.targetFunc(testlib.GetTestPath(spec.filename))
	if result != spec.expected {
		t.Errorf("got %t, expected %t", result, spec.expected)
	}
}

func TestIsFileExistsFunc(t *testing.T) {
	tests := [3]testSpec{
		{"./data", false, IsFileExist},
		{"./data/test_invalid_doc.json", true, IsFileExist},
		{"./non_exist", false, IsFileExist},
	}

	for _, spec := range tests {
		t.Run(spec.filename, spec.Test)
	}
}

func TestIsJSONFileFunc(t *testing.T) {
	tests := [3]testSpec{
		{"no_ext_file", false, IsJSONFile},
		{"with_ext_file.something", false, IsJSONFile},
		{"json_file.json", true, IsJSONFile},
	}

	for _, spec := range tests {
		t.Run(spec.filename, spec.Test)
	}
}
