// config/config.go
package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server ServerConfig `json:"server"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

func LoadConfig() (*Config, error) {
	configFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var cfg Config
	err = json.NewDecoder(configFile).Decode(&cfg)
	return &cfg, err
}
