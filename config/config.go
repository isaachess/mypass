package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DataFileLocation string `json:"dataFileLocation"`
}

func LoadConfig(path string, defaultPath string) *Config {
	c, err := loadConfigFromPath(path)
	if err != nil {
		return &Config{
			DataFileLocation: filepath.Join(defaultPath, "data.json"),
		}
	}
	return c
}

func loadConfigFromPath(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var c Config
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
