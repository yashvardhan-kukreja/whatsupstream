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

	return Config{
		IssueConfigs: parsedInternalIssueConfigs,
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
	for i, _ := range labels {
		labels[i] = strings.ReplaceAll(labels[i], " ", "+")
	}

	return IssueConfig{
		Owner: owner,
		RepoName: repoName,
		Labels: labels,
		Assignee: inputIssueConfig.Assignee,
		Creator: inputIssueConfig.Creator,
		State: issueState,
		Since: inputIssueConfig.Since,
	}, nil
}