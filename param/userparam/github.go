package userparam

import "time"

type Github struct {
	AvatarURL         string      `json:"avatar_url"`
	Bio               string      `json:"bio"`
	Blog              string      `json:"blog"`
	Company           string      `json:"company"`
	CreatedAt         time.Time   `json:"created_at"`
	Email             string      `json:"email"`
	EventsURL         string      `json:"events_url"`
	Followers         int         `json:"followers"`
	FollowersURL      string      `json:"followers_url"`
	Following         int         `json:"following"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	GravatarID        string      `json:"gravatar_id"`
	Hireable          interface{} `json:"hireable"`
	HTMLURL           string      `json:"html_url"`
	ID                int         `json:"id"`
	Location          string      `json:"location"`
	Login             string      `json:"login"`
	Name              string      `json:"name"`
	NodeID            string      `json:"node_id"`
	NotificationEmail interface{} `json:"notification_email"`
	OrganizationsURL  string      `json:"organizations_url"`
	PublicGists       int         `json:"public_gists"`
	PublicRepos       int         `json:"public_repos"`
	ReceivedEventsURL string      `json:"received_events_url"`
	ReposURL          string      `json:"repos_url"`
	SiteAdmin         bool        `json:"site_admin"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	TwitterUsername   string      `json:"twitter_username"`
	Type              string      `json:"type"`
	UpdatedAt         time.Time   `json:"updated_at"`
	URL               string      `json:"url"`
	UserViewType      string      `json:"user_view_type"`
}

type GithubEmail []struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"`
}
