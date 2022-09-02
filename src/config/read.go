package config

import (
	"fmt"
	"os"

	"github.com/invopop/yaml"
)

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func LoadFromFile(cfg *Config) {
	f, err := os.ReadFile("config.yml")
	if err != nil {
		processError(err)
	}

	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		processError(err)
	}
}
