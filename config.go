package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	dataFileLocation string `json:"dataFileLocation"`
}

func LoadConfig(path string, defaultPath string) *Config {
	c, err := loadConfigFromPath(path)
	if err != nil {
		fmt.Printf("Cannot load config from path %s, using default config\n", path)
		return &Config{
			dataFileLocation: filepath.Join(defaultPath, "data.json"),
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
