package userservice

import (
	"context"
	"errors"
	"fmt"

	_ "golang.org/x/oauth2"
)

func (r User) RedirectUrlGenerator(prov string) (string, error) {
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

func (r User) AcceptGoogleOauth(code string) ([]byte, error) {

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

func (r User) AcceptGithubOauth(code string) ([]byte, []byte, error) {

	exchangeCo, err := r.oauth.ExchangeGithubCode(context.Background(), code)

	if err != nil {
		return nil, nil, fmt.Errorf("code exchange failed 1231543: %s", err.Error())
	}

	data, err := r.oauth.GetGitHubUserInfo(exchangeCo.AccessToken)

	if err != nil {
		return nil, nil, fmt.Errorf("code exchange failed 56421232: %s", err.Error())
	}

	dataEmail, err := r.oauth.GetGitHubUserEmail(exchangeCo.AccessToken)

	if err != nil {
		return nil, nil, fmt.Errorf("code exchange failed 56421232: %s", err.Error())
	}

	return data, dataEmail, nil

}
