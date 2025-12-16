package auth

type IAuth struct {
	Access_token   string `json:"access_token"`
	Expires_at     int    `json:"expires_at"`
	Expires_in     int    `json:"expires_in"`
	Refresh_token  string `json:"refresh_token"`
	Provider_token string `json:"provider_token"`
}
