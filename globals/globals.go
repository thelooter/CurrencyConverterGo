package globals

import (
	"bytes"
	"encoding/json"
	"os"
	"thelooter.de/currencyconverter/datastructures"
)

var (
	Config datastructures.Config
)

func init() {
	Config = loadConfig()
}

func loadConfig() datastructures.Config {
	var config datastructures.Config
	entries, err := os.ReadDir("config")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.Name() == "config.json" {
			file, err := os.ReadFile("config/config.json")
			if err != nil {
				return datastructures.Config{}
			}

			var config datastructures.Config

			decoder := json.NewDecoder(bytes.NewReader(file))

			if err := decoder.Decode(&config); err != nil {
				panic("JSON decode error, Please check your config.json file")
			}

			if config.BaseURL == "" {
				panic("BaseURL is not set in config.json")
			}

			if config.APIKey == "" {
				panic("APIKey is not set in config.json")
			}

			return config
		}
	}

	return config
}
