package config

import (
	"encoding/json"
	"os"
)

func ReadConfig(path string) (Config, error) {
	var config Config

	all, err := os.ReadFile(path)

	if err != nil {
		return config, err
	}

	return config, json.Unmarshal(all, &config)
}

func MustReadConfig(path string) Config {
	config, err := ReadConfig(path)
	if err != nil {
		panic(err)
	}
	return config
}
