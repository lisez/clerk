package jobs

import (
	"testing"
)

func TestGetProviderFromURI(t *testing.T) {
	specs := []struct {
		uri    string
		scheme string
	}{
		{"mongo://localhost", "mongo"},
		{"postgres://localhost", "postgres"},
		{"noturi.localhost", ""},
	}

	for _, spec := range specs {
		t.Run(spec.uri, func(t *testing.T) {
			scheme := GetProviderFromURI(spec.uri)
			if scheme != spec.scheme {
				t.Errorf("got %s, expected %s", scheme, spec.scheme)
			}
		})
	}
}
