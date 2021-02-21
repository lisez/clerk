package jobs

import (
	"clerk/common/fslib"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// ClerkConfig ...
type ClerkConfig struct {
	Clerk ClerkInternalConfig `yaml:"clerk"`
}

// ClerkInternalConfig ...
type ClerkInternalConfig struct {
	Schema        map[string]interface{} `yaml:"schema,omitempty"`
	FromFiles     []string               `yaml:"fromFiles,omitempty"`
	SourceRemotes []ClerkRemoteConfig    `yaml:"sourceRemotes,omitempty"`
}

// ClerkRemoteConfig ...
type ClerkRemoteConfig struct {
	URI      string                 `yaml:"uri"`
	Provider string                 `yaml:"provider,omitempty"`
	Schema   map[string]interface{} `yaml:"schema,omitempty"`
	Args     map[string]interface{} `yaml:"args,omitempty"`
}

// GetProvider ...
func (remote *ClerkRemoteConfig) GetProvider() string {
	provider := GetProviderFromURI(remote.URI)
	if len(remote.Provider) > 0 {
		provider = remote.Provider
	}
	return provider
}

func check(e error) {
	if e != nil {
		log.Fatalf("load config failed: %v", e)
	}
}

// NewClerkConfig ...
func NewClerkConfig(filename string) ClerkConfig {
	config := ClerkConfig{}
	if fslib.IsFileExist(filename) {
		dat, derr := ioutil.ReadFile(filename)
		check(derr)

		err := yaml.Unmarshal([]byte(dat), &config)
		check(err)
	}
	return config
}
