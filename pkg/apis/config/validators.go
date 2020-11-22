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
	"net/http"
	"regexp"
)

const (
	GITHUB_URL_REGEX = "(?:git|ssh|https?|git@[-w.]+)(:)?(//)?(.*?)(.git)?(/?|#[-dw._]+?)$"
)

func ValidateConfig(config Config) error {
	for _, issueConfig := range config.IssueConfigs {
		if err := validateIssueConfig(issueConfig); err != nil {
			return fmt.Errorf("error occurred while parsing the provided input configuration: %w", err)
		}
	}
	return nil
}

func validateIssueConfig(issueConfig IssueConfig) error {
	if err := validateRepositoryURL(issueConfig.RepositoryURL); err != nil {
		return fmt.Errorf("error occurred while parsing an issue configuration: %w", err)
	}
	if err := validateSince(issueConfig.Since); err != nil {
		return fmt.Errorf("error occurred while parsing an issue configuration: %w", err)
	}
	return nil
}

func validateSince(since string) error {
	return nil
}

func validateRepositoryURL(repositoryURL string) error {
	if matches, _ := regexp.MatchString(GITHUB_URL_REGEX, repositoryURL); !matches {
		return fmt.Errorf("repository URL not provided rightfully")
	}
	resp, err := http.Get(repositoryURL)
	if err != nil || resp.StatusCode >= 400 {
		return fmt.Errorf("repository URL not provided rightfully")
	}
	defer resp.Body.Close()
	return nil
}
