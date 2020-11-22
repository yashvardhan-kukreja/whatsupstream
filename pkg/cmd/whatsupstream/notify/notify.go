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
package notify

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"whatsupstream/pkg/apis/config"
	internalConfig "whatsupstream/pkg/internal/apis/config"
)

type flagpole struct {
	Config string
}

func NewCommand() *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Use:   "notify",
		Short: "Enables notifications",
		Long: `Enables desktop notifications for the current user as per the provided configuration.
The user would receive notifications around new issues depending on the filters/conditions provided in the config.
The notifications would be in the form of desktop notifications.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags)
		},
	}
	cmd.Flags().StringVar(
		&flags.Config, "config",
		"/Users/yash1300/.whatsupstream/config.yaml",
		"Path to the config containing preferences associated with the notifications to receive.",
	)
	return cmd
}

func runE(flags *flagpole) error {
	inputConfig, err := yamlConfigToInputConfig(flags.Config)
	if err != nil {
		return fmt.Errorf("error occurred while executing the 'notify' command: %w", err)
	}
	if err := config.ValidateConfig(inputConfig); err != nil {
		return fmt.Errorf("error occurred while executing the 'notify' command: %w", err)
	}
	parsedConfig, err := internalConfig.ConvertInputConfigToInternalConfig(inputConfig)
	errThreshold := 0
	for errThreshold <= 3 {
		data, err := internalConfig.FetchNotificationData(parsedConfig)
		if err != nil {
			fmt.Printf("Error occurred: %w", err)
			errThreshold ++
		}
		fmt.Printf("\nLength of data := %d, \n Data := %+v \n", len(data.Issues), data)
		time.Sleep(3*time.Second)
	}
	return fmt.Errorf("error occurred while fetching notification data more than threshold amount of times (3)")
}

func yamlConfigToInputConfig(configPath string) (config.Config, error) {
	yamlFile, err := ioutil.ReadFile(configPath)
    if err != nil {
        return config.Config{}, fmt.Errorf("error occurred while reading the config file: %w", err)
	}

	var inputConfig config.Config
    err = yaml.Unmarshal(yamlFile, &inputConfig)
    if err != nil {
		return config.Config{}, fmt.Errorf("error occurred while reading the config file: %w", err)
    }

    return inputConfig, nil
}