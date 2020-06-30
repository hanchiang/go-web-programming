package data

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
	DbUser       string
	DbPassword   string
	DbName       string
}

var Config Configuration

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open(filepath.Join("data", "config.json"))
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
