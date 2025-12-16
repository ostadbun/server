package auth

import "time"

type IAuthSupabase struct {
	ID               string    `json:"id"`
	Aud              string    `json:"aud"`
	Role             string    `json:"role"`
	Email            string    `json:"email"`
	EmailConfirmedAt time.Time `json:"email_confirmed_at"`
	Phone            string    `json:"phone"`
	ConfirmedAt      time.Time `json:"confirmed_at"`
	LastSignInAt     time.Time `json:"last_sign_in_at"`

	AppMetadata struct {
		Provider  string   `json:"provider"`
		Providers []string `json:"providers"`
	} `json:"app_metadata"`

	UserMetadata struct {
		AvatarURL         string `json:"avatar_url"`
		Email             string `json:"email"`
		EmailVerified     bool   `json:"email_verified"`
		FullName          string `json:"full_name"`
		Iss               string `json:"iss"`
		Name              string `json:"name"`
		PhoneVerified     bool   `json:"phone_verified"`
		Picture           string `json:"picture"`
		PreferredUsername string `json:"preferred_username"`
		ProviderID        string `json:"provider_id"`
		Sub               string `json:"sub"`
		UserName          string `json:"user_name"`
	} `json:"user_metadata"`
	Identities []struct {
		IdentityID   string `json:"identity_id"`
		ID           string `json:"id"`
		UserID       string `json:"user_id"`
		IdentityData struct {
			AvatarURL         string `json:"avatar_url"`
			Email             string `json:"email"`
			EmailVerified     bool   `json:"email_verified"`
			FullName          string `json:"full_name"`
			Iss               string `json:"iss"`
			Name              string `json:"name"`
			PhoneVerified     bool   `json:"phone_verified"`
			PreferredUsername string `json:"preferred_username"`
			ProviderID        string `json:"provider_id"`
			Sub               string `json:"sub"`
			UserName          string `json:"user_name"`
		} `json:"identity_data"`
		Provider     string    `json:"provider"`
		LastSignInAt time.Time `json:"last_sign_in_at"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Email        string    `json:"email"`
	} `json:"identities"`

	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsAnonymous bool      `json:"is_anonymous"`
}
