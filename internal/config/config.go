package config

import (
	"os"
	"encoding/json"
)

const(
	configFileName	= ".gatorconfig.json"
	stdPerms		= 0666	// If file doesn't exist, it is created with these permission (see linux umask)
)

type Config struct {
	DbURL			string
	CurrentUserName	string
}

func Read() (Config, error) {
	configPath, err := getConfigPath()
	if err != nil { return Config{}, err }

	file, err := os.ReadFile(configPath)
	if err != nil { return Config{}, err }
	
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil { return Config{}, err }

	return config, nil
}

func (c Config)SetUser(userName string) error {
	c.CurrentUserName = userName

	jsonData, err := json.Marshal(c)
	if err != nil { return err }

	configPath, err := getConfigPath()
	if err != nil { return err }

	err = os.WriteFile(configPath, jsonData, stdPerms)
	if err != nil { return err }

	return nil
}

func getConfigPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil { return "", err }
	configPath := homePath + "/" + configFileName

	return configPath, nil
}

