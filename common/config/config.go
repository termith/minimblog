package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Database struct {
		Driver string `json:"driver"`
		Url    string `json:"url"`
	} `json:"database"`
	Application struct {
		Port int `json:"port"`
	} `json:"application"`
}

func LoadConfig(filePath string) *Config {
	/*
		Loads configuration from filePath
	*/
	var config Config
	rawFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error while read config file")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal(rawFile, &config)
	if err != nil {
		fmt.Println("Error while unmarshall json")
		fmt.Println(err.Error())
	}
	return &config
}
