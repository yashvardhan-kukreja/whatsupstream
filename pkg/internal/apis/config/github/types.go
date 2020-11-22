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
package github

// Issue represents the fields coming from the API response from api.github.com
type Issue struct {
	Title string `json:"title"`
	Number int `json:"number"`
	State string `json:"state"`
	CreatedAt string `json:"created_at"`
	User User `json:"user"`
	Labels []Label `json:"labels"`
	RepositoryURL string `json:"repository_url"`
}

type User struct {
	AvatarURL string `json:"avatar_url"`
	Username string `json:"login"`
	ProfileURL string `json:"url"`
	FollowersURL string `json:"followers_url"`
}

type Label struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type APIMethod string

const (
	GET APIMethod = "GET"
	POST APIMethod = "POST"
	PUT APIMethod = "PUT"
	DELETE APIMethod = "DELETE"
	PATCH APIMethod = "PATCH"
)

type APIContentType string

const (
	JSON APIContentType = "application/json"
	XML APIContentType = "application/xml"
	FORM APIContentType = "application/x-www-form-urlencoded"
)