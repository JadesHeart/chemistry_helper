package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type DataBaseConfig struct {
	DatabaseURL string `json:"database_url"`
}

func loadConfiguration(file string) *DataBaseConfig {
	var config DataBaseConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config
}

func NewDataBaseConfig() *DataBaseConfig {
	return loadConfiguration("internal/app/config/database_config.json")
}
