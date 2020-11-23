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
	// loud alerts on Mac has shown errors
	// forcedSilentMode will turn to "true" whenever such errors are encountered
	// and, notification will be forced to be silent no matter the "SilentMode" in any IssueConfig
	// TODO: add a reconciler to check at an interval whether the loud alerts got fixed and update "forcedSilentMode" to false accordingly
	forcedSilentMode := false
	errThreshold := 0
	for errThreshold <= 3 {
		// generate all the Notification objects to raise
		allNotifications, err := internalConfig.FetchNotificationData(parsedConfig)
		if err != nil {
			fmt.Printf("error occurred while executing 'notify': %w", err)
			errThreshold++
		}
		// raise all notifications
		for _, notification := range allNotifications {
			title := "Whatsupstream's Update 🚀"
			description, err := formatNotificationDescription(notification.Issue)
			silentMode := notification.SilentMode
			if forcedSilentMode {
				silentMode = forcedSilentMode
			}
			if err != nil {
				fmt.Println("error occurred while executing 'notify'")
				errThreshold++
			}
			// raising notifications concurrently
			go func() {
				err := raiseNotification(title, description, silentMode)
				if err != nil {
					if !forcedSilentMode {
						forcedSilentMode = true
						return
					}
					fmt.Println("error occurred while raising a notification")
				}
			}()
		}

		// updating the Since field of all IssueConfigs to current time so that in the next github API call, only new issues come up
		for i := range parsedConfig.IssueConfigs {
			parsedConfig.IssueConfigs[i].Since = time.Now().Format("2006-01-02T15:04:05Z")
		}

		// wait from the next polling cycle
		pollingInterval := time.Duration(parsedConfig.PollingRate) * time.Second
		time.Sleep(pollingInterval)
	}
	panic("error occurred while fetching notification data more than threshold amount of times (3)")
	return nil
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
Labels: %s
Link: https://github.com/%s/%s/issues/%d
Created at: %s
By: %s 
	`, owner, repoName, labelsStr, owner, repoName, issue.Number, createdAtStr, issue.User.Username), nil
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
