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
	IssueConfigs []IssueConfig
}

// IssueConfig represents the configuration associated with issue-related notifications of a repository
type IssueConfig struct {

	// Owner denotes the username/organization name of the owner of the repository.
	Owner string

	// RepoName denotes the name of the repository
	RepoName string

	// Labels denotes the list of labels which the issues (of which you want to get notified) must contain.
	Labels []string

	// Assignees denotes the person who must be assigned to the issue, for that issue being eligible for being notified about.
	// if left unset, then, it will be set to "*" denoting the issues can be assigned to any user.
	Assignee string

	// Creator denotes the creator from whom the issue must be created, so as to be eligible for being notified about.
	Creator string

	// State denotes the kind of issues (closed/open/all) which will be eligible to be notified about.
	State IssueState

	// Since denotes the timestamp from which the issues which were created after it, will only be eligible for being notified about.
	Since string
}

// IssueState defines the possible state of the issue
type IssueState string

const (
	Closed IssueState = "closed"
	Opened IssueState = "open"
	All IssueState = "all"
)

const (
	API_BASE_URL_REPOS = "https://api.github.com/repos"
)