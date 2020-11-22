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

func ConvertInputConfigtoInteralConfig(inputConfig config.Config) (error, Config) {
	var parsedInternalIssueConfigs []IssueConfig
	for _, inputIssueConfig := range inputConfig.IssueConfigs {
		err, parsedInternalIssueConfig := convertInputIssueConfigToInternalIssueConfig(inputIssueConfig)
		if err != nil {
			return fmt.Errorf("error occurred parsing an issue config: %w", err), Config{}
		}
		parsedInternalIssueConfigs = append(parsedInternalIssueConfigs, parsedInternalIssueConfig)
	}

	return nil, Config{
		IssueConfigs: parsedInternalIssueConfigs,
	}

}

// ConvertInputIssueConfigToInternalIssueConfig ...
func convertInputIssueConfigToInternalIssueConfig(inputIssueConfig config.IssueConfig) (error, IssueConfig) {
	if inputIssueConfig.RepositoryURL == "" {
		return fmt.Errorf("repository URL not found in the issue config"), IssueConfig{}
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

	return nil, IssueConfig{
		Owner: owner,
		RepoName: repoName,
		Labels: inputIssueConfig.Labels,
		Assignee: inputIssueConfig.Assignee,
		Creator: inputIssueConfig.Creator,
		State: issueState,
		Since: inputIssueConfig.Since,
	}
}