package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func YamlConfigToInputConfig(configPath string) (Config, error) {
	yamlFile, err := ioutil.ReadFile(configPath)
    if err != nil {
        return Config{}, fmt.Errorf("error occurred while reading the config file: %w", err)
	}

	var inputConfig Config
    err = yaml.Unmarshal(yamlFile, &inputConfig)
    if err != nil {
		return Config{}, fmt.Errorf("error occurred while reading the config file: %w", err)
    }

    return inputConfig, nil
}