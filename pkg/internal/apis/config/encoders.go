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
	"strings"
	"time"

	"whatsupstream/pkg/apis/config"
)

func ConvertInputConfigToInternalConfig(inputConfig config.Config) (Config, error) {
	var parsedInternalIssueConfigs []IssueConfig
	for _, inputIssueConfig := range inputConfig.IssueConfigs {
		parsedInternalIssueConfig, err := convertInputIssueConfigToInternalIssueConfig(inputIssueConfig)
		if err != nil {
			return Config{}, fmt.Errorf("error occurred parsing an issue config: %w", err)
		}
		parsedInternalIssueConfigs = append(parsedInternalIssueConfigs, parsedInternalIssueConfig)
	}

	pollingRate := inputConfig.PollingRate
	if pollingRate <= 0 {
		pollingRate = 60
	}

	return Config{
		IssueConfigs: parsedInternalIssueConfigs,
		PollingRate:  pollingRate,
	}, nil

}

// ConvertInputIssueConfigToInternalIssueConfig ...
func convertInputIssueConfigToInternalIssueConfig(inputIssueConfig config.IssueConfig) (IssueConfig, error) {
	if inputIssueConfig.RepositoryURL == "" {
		return IssueConfig{}, fmt.Errorf("repository URL not found in the issue config")
	}
	repositoryURLTokens := strings.Split(inputIssueConfig.RepositoryURL, "/")
	repoName := repositoryURLTokens[len(repositoryURLTokens)-1]
	owner := repositoryURLTokens[len(repositoryURLTokens)-2]

	var issueState IssueState
	if inputIssueConfig.Closed {
		issueState = All
	} else {
		issueState = Opened
	}

	labels := inputIssueConfig.Labels
	for i := range labels {
		labels[i] = strings.ReplaceAll(labels[i], " ", "+")
	}

	maxIssuesCount := inputIssueConfig.MaxIssuesCount
	if maxIssuesCount <= 0 {
		maxIssuesCount = 5
	}

	since := inputIssueConfig.Since
	if since == "" {
		since = time.Now().Add(-24 * time.Hour).Format("2006-01-02T15:04:05Z")
	}

	return IssueConfig{
		Owner:          owner,
		RepoName:       repoName,
		Labels:         labels,
		Assignee:       inputIssueConfig.Assignee,
		Creator:        inputIssueConfig.Creator,
		State:          issueState,
		Since:          since,
		MaxIssuesCount: maxIssuesCount,
	}, nil
}
