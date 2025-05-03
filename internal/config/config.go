package config

import (
	"os"
	"encoding/json"
	"fmt"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error occured when getting file path...")
		return Config{}, err
	} 

	readFile, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error occured when reading file...")
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(readFile, &config)
	if err != nil {
		fmt.Println("Error unmarshal failed")
		return Config{}, err
	}

	return config, nil
}

func (cfg Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	err := write(cfg)

	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
 	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := homeDir + "/" + configFileName

	return configPath, nil  
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)	
	if err != nil {
		return err
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0666)
	if err != nil {
		return err
	}

	return nil
}
