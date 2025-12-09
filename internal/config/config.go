package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Encountered error when marshalling config struct: %w\n", err)
	}

	fullPath, err := getConfigPath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(fullPath, data, 0644); err != nil {
		return fmt.Errorf("Encountered error when writing config to file: %w\n", err)
	}
	return nil
}

func Read() (Config, error) {
	fullPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, fmt.Errorf("Encountered error when reading config file: %w\n", err)
	}

	var cfg Config

	if err := json.Unmarshal(file, &cfg); err != nil {
		return Config{}, fmt.Errorf("Encountered error when unmarshalling config file: %w\n", err)
	}

	return cfg, nil
}

func getConfigPath() (string, error) {
	const configFileName = ".gatorconfig.json"
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Encountered error when trying to find home dir: %wn\n", err)
	}
	return homePath + "/" + configFileName, nil
}
