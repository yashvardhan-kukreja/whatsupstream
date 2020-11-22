package github

// Issue represents the fields coming from the API response from api.github.com
type Issue struct {
	Title string `json:"title"`
	IssueNumber int32 `json:"number"`
	State string `json:"state"`
	CreatedAt string `json:"created_at"`
	User User `json:"user"`
	Labels []Label `json:"labels"`
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