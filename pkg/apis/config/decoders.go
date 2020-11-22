/*
Copyright © 2020 Yashvardhan Kukreja <yash.kukreja.98@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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