package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Represents the configuration structure.
type Config struct {
    Repositories []RepoConfig
}

// Represents a single repository entry in the configuration.
type RepoConfig struct {
	URL       string        `yaml:"url"`
	Interval  time.Duration `yaml:"interval"`
	LocalPath string        `yaml:"path"`
}

// ReadConfig reads and parses the configuration file.
func ReadConfig(filename string) (*Config, error) {
    // Read the configuration file.
	data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

	// Unmarshal the YAML data into a Config struct.
    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }

	// Validate the configuration.
	for i, repo := range config.Repositories {
		if repo.URL == "" {
			return nil, fmt.Errorf("repository %d has no URL", i+1)
		}
		if repo.LocalPath == "" {
			return nil, fmt.Errorf("repository %d has no local path", i+1)
		}
	}

    return &config, nil
}