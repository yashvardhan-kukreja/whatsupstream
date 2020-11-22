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
	"strings"
	"time"

	"github.com/gen2brain/beeep"

	"whatsupstream/pkg/apis/config"
	internalConfig "whatsupstream/pkg/internal/apis/config"
	"whatsupstream/pkg/internal/apis/config/github"
)

func runE(flags *flagpole) error {
	inputConfig, err := config.YamlConfigToInputConfig(flags.Config)
	if err != nil {
		return fmt.Errorf("error occurred while executing the 'notify': %w", err)
	}
	if err := config.ValidateConfig(inputConfig); err != nil {
		return fmt.Errorf("error occurred while executing the 'notify': %w", err)
	}
	parsedConfig, err := internalConfig.ConvertInputConfigToInternalConfig(inputConfig)
	errThreshold := 0
	for errThreshold <= 3 {
		data, err := internalConfig.FetchNotificationData(parsedConfig)
		if err != nil {
			fmt.Printf("error occurred while executing 'notify': %w", err)
			errThreshold++
		}
		for i, issue := range data.Issues {
			title := "Whatsupstream's Update 🚀"
			description, err := formatNotificationDescription(issue)
			silentMode := true
			if err != nil {
				fmt.Printf("error occurred while executing 'notify': %w", err)
				errThreshold++
			}
			// raising notifications concurrently
			go func() {
				err := raiseNotification(title, description, silentMode)
				if err != nil {
					fmt.Printf("error occurred while executing 'notify': %w", err)
					errThreshold++
				}
				parsedConfig.IssueConfigs[i].Since = time.Now().Format("2006-01-02T15:04:05Z") // updating the Since field to the latest time so that next time, only new issues come up
			}()
		}
		pollingInterval := time.Duration(parsedConfig.PollingRate) * time.Second
		time.Sleep(pollingInterval)
	}
	return fmt.Errorf("error occurred while fetching notification data more than threshold amount of times (3)")
}

func formatNotificationDescription(issue github.Issue) (string, error) {
	labelsStr := ""
	for _, label := range issue.Labels {
		labelsStr += label.Name + ", "
	}
	labelsStr = strings.TrimSuffix(labelsStr, ", ")

	timestampLayout := "2006-01-02T15:04:05Z"
	createdAtStr, err := time.Parse(timestampLayout, issue.CreatedAt)
	if err != nil {
		return "", fmt.Errorf("error occurred while generating the notification description: %w", err)
	}

	repositoryURLTokens := strings.Split(issue.RepositoryURL, "/")
	repoName := repositoryURLTokens[len(repositoryURLTokens)-1]
	owner := repositoryURLTokens[len(repositoryURLTokens)-2]

	return fmt.Sprintf(`Repository: %s/%s
Issue %d:
With Labels: %s
Created at: %s
By: %s 
	`, owner, repoName, issue.Number, labelsStr, createdAtStr, issue.User.Username), nil
}

func raiseNotification(title, description string, silentMode bool) error {
	var err error
	if silentMode {
		err = beeep.Notify(title, description, "")
	} else {
		err = beeep.Alert(title, description, "")
	}
	if err != nil {
		return fmt.Errorf("error occurred while generating a notification alert: %w", err)
	}
	return nil
}
