package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config is struct
type Config struct {
	Database struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// LoadConfig is function to load a config file
func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		return config, err
	}

	jsonParse := json.NewDecoder(configFile)
	err = jsonParse.Decode(&config)
	return config, err
}

func main() {
	fmt.Println("Starting web")
	config, _ := LoadConfig("config.json")
	fmt.Println(config.Host)
}
