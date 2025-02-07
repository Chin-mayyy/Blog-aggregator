package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Function to read the JSON file found at
// home directory and decodes the JSON into
// the its respective struct.
func ReadFile() (Config, error) {
	location, error := getConfigFilePath()
	if error != nil {
		return Config{}, error
	}

	data, error := os.ReadFile(location)
	if error != nil {
		return Config{}, nil
	}

	var config Config

	err := json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error decoding json")
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}
