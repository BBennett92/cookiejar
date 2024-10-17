package utils

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

// Config holds configuration settings.
type Config struct {
	Scan struct {
		FollowRedirects bool   `yaml:"follow_redirects"`
		Timeout         string `yaml:"timeout"`
		TimeoutDuration time.Duration
	} `yaml:"scan"`
	Output struct {
		Verbose bool   `yaml:"verbose"`
		Format  string `yaml:"format"`
	} `yaml:"output"`
}

// LoadConfig loads configuration from a file.
func LoadConfig(configPath string) (Config, error) {
	var config Config
	// Load default config if no path is provided
	if configPath == "" {
		configPath = "configs/default_config.yaml"
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	// Parse timeout duration
	if config.Scan.Timeout == "" {
		config.Scan.Timeout = "10s"
	}
	duration, err := time.ParseDuration(config.Scan.Timeout)
	if err != nil {
		return config, err
	}
	config.Scan.TimeoutDuration = duration
	return config, nil
}
