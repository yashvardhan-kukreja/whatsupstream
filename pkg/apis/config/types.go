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

// Config represents the main configuration on the basis of which the user will get notifications
type Config struct {

	// IssueConfigs is the array of multiple configurations associated with multiple kinds of issue-related notifications.
	// for example - an element of IssueConfigs can be responsible for notifying about "good first issue" labelled issues,
	// and another element of IssueConfigs can be responsible for notifying about issues associated with "bug" labelled issues of another repo.
	IssueConfigs []IssueConfig `yaml:"issue-configs,omitempty"`

	// PollingRate denotes the rate (in seconds) at which the whatsupstream will run and poll and check github for any updates
	// if left unset, then, it will be set to 60 seconds.
	PollingRate int `yaml:"polling-rate,omitempty"`
}

// IssueConfig represents the configuration associated with issue-related notifications of a repository
type IssueConfig struct {

	// RepositoryURL denotes the URL of the repository for which you want the receive the issue-related notifications.
	RepositoryURL string `yaml:"repository-url,omitempty"`

	// Labels denotes the list of labels which the issues (of which you want to get notified) must contain.
	Labels []string `yaml:"labels,omitempty"`

	// Assignees denotes the person who must be assigned to the issue, for that issue being eligible for being notified about.
	// if left unset, then, it will be set to "*" denoting the issues can be assigned to any user.
	Assignee string `yaml:"assignee,omitempty"`

	// Creator denotes the creator from whom the issue must be created, so as to be eligible for being notified about.
	Creator string `yaml:"creator,omitempty"`

	// Closed denotes whether to include closed issues too (along the open ones).
	Closed bool `yaml:"closed,omitempty"`

	// Since denotes the timestamp from which the issues which were created after it, will only be eligible for being notified about.
	// it must have the following format "yyyy-mm-ddTHH:MM:SSZ". For example:  "2006-01-02T15:04:05Z"
	// if left unset, then, it will be set to the timestamp of exactly 24hrs (1 day) before current time.
	Since string `yaml:"since,omitempty"`

	// MaxIssuesCount denotes the top (as per creation time) maximum number of issues which will be considered for being notified about.
	// for example: if MaxIssuesCount is 5, then, only top 5 latest issues will be considered for being notified about (if in a query more than issues are returned).
	// if left unset, MaxIssuesCount will be set to 5.
	MaxIssuesCount int `yaml:"max-issues-count,omitempty"`

	// SilentMode whether the notification would be an alert with a sound or will it be silent
	// if left unset, SilentMode, will be set to false
	SilentMode bool `yaml:"silent-mode,omitempty"`
}

const (
	API_BASE_URL_REPOS = "https://api.github.com/repos"
)
