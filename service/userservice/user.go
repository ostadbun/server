package userservice

import (
	"context"
	"errors"
	"fmt"
	"os"

	_ "golang.org/x/oauth2"
)

func (r User) RedirectUrlGenerator(prov string, client string, info []byte) (string, error) {

	//redirectURL :=
	redirectURL := func(provider string) string {
		return fmt.Sprintf("%s/user/oauth/callback/%s", os.Getenv("SERVER_ADDRESS"), provider)
	}

	if prov == "google" {
		return r.oauth.GetGoogleAuthURL(redirectURL("google"), info)
	}

	if prov == "github" {
		return r.oauth.GetGithubAuthURL(redirectURL("github"), info)

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
