package server

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

// Config holds server configuration
type Config struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
	Host    string        `yaml:"host"`
}

// DefaultConfig returns default server configuration
func DefaultConfig() Config {
	return Config{
		Port:    ":3000",
		Timeout: 5 * time.Second,
		Host:    "localhost",
	}
}

// LoadConfig loads configuration from file
func LoadConfig(configPath string) (*Config, error) {
	// Default to config/config.yaml if no path provided
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	// Read the configuration file
	data, err := os.ReadFile(filepath.Clean(configPath))
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal YAML
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
