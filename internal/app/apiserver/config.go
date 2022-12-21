package apiserver

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfig struct {
	Bind        string `json:"bind"`
	LoggerLevel string `json:"logger_level"`
}

func loadConfiguration(file string) *ServerConfig {
	var config ServerConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config
}

func NewServerConfig() *ServerConfig {
	config := loadConfiguration("internal/app/config/config.json")
	return config
}
