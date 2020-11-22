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

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type APIRequest struct {
	URL         string
	Method      APIMethod
	ContentType APIContentType
	Body        io.Reader
}

func FetchTopIssues(url string, maxIssuesCount int) ([]Issue, error) {
	var issues []Issue
	request := APIRequest{
		URL:    url,
		Method: GET,
	}
	responseBody, err := callAPI(request)
	if err != nil {
		return []Issue{}, fmt.Errorf("error occurred while fetching the issues: %w", err)
	}
	if err := json.Unmarshal(responseBody, &issues); err != nil {
		return []Issue{}, fmt.Errorf("error occurred while fetching the issues: %w", err)
	}
	if len(issues) > maxIssuesCount {
		issues = issues[:maxIssuesCount]
	}
	return issues, nil
}

func callAPI(request APIRequest) ([]byte, error) {
	var response *http.Response
	switch request.Method {
	case GET:
		resp, err := http.Get(request.URL)
		if err != nil {
			return []byte{}, fmt.Errorf("error occurred while calling %s: %w", request.URL, err)
		}
		response = resp
	case POST, PUT, DELETE, PATCH:
		req, err := http.NewRequest(string(request.Method), request.URL, request.Body)
		if err != nil {
			return []byte{}, fmt.Errorf("error occurred while calling %s: %w", request.URL, err)
		}
		req.Header.Set("Content-Type", string(request.ContentType))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return []byte{}, fmt.Errorf("error occurred while calling %s: %w", request.URL, err)
		}
		response = resp
	default:
		return []byte{}, fmt.Errorf("error occurred while calling %s: wrong request type found", request.URL)
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error occurred while parsing the fetched issues: %w", err)
	}

	return []byte(responseBody), nil
}
