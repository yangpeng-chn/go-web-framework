package settings

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	LogIndent   bool
	EnableHTTPS bool
	UseDatabase bool
	Cert        string
	Key         string
}

var conf Configuration

func ReadConfigFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	return err
}

func GetConfig() Configuration {
	return conf
}
