package userservice

import (
	"context"
	"errors"
	"fmt"
	"ostadbun/service/oauthservice"

	_ "golang.org/x/oauth2"
)

type Auth struct {
	oauth oauthservice.OAuthService
}

func New(oauth oauthservice.OAuthService) Auth {

	return Auth{

		oauth: oauth,
	}
}

func (r Auth) RedirectUrlGenerator(prov string) (string, error) {
	if prov == "google" {

		str := r.oauth.GetGoogleAuthURL("http://localhost:3000/user/oauth/callback/google")
		return str, nil

	}

	if prov == "github" {
		str := r.oauth.GetGithubAuthURL("http://localhost:3000/user/oauth/callback/github")
		return str, nil
	}
	return "", errors.New("invalid prov")
}

func (r Auth) AcceptGoogleOauth(code string) (any, error) {

	exchangeCo, err := r.oauth.ExchangeGoogleCode(context.Background(), code)

	if err != nil {
		return nil, fmt.Errorf("code exchange failed 31324: %s", err.Error())
	}

	data, err := r.oauth.GetGoogleUserInfo(exchangeCo.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("code exchange failed 634244: %s", err.Error())
	}

	return data, nil

}

func (r Auth) AcceptGithubOauth(code string) (any, error) {

	exchangeCo, err := r.oauth.ExchangeGithubCode(context.Background(), code)

	if err != nil {
		return nil, fmt.Errorf("code exchange failed 1231543: %s", err.Error())
	}

	data, err := r.oauth.GetGitHubUserInfo(exchangeCo.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("code exchange failed 56421232: %s", err.Error())
	}

	return data, nil

}
