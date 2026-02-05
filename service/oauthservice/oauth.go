package oauthservice

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type Oauthrepo interface {
	Set(ctx context.Context, userdata []byte) (string, error)
}

// ساختار تنظیمات اواث
type OAuthService struct {
	googleConfig *oauth2.Config
	githubConfig *oauth2.Config

	redis Oauthrepo
}

// تابع سازنده: تنظیمات اولیه
func NewOAuthService(redis Oauthrepo) *OAuthService {
	// تنظیمات گوگل
	googleConf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		// RedirectURL: در اینجا خالی می‌گذاریم چون پارت CLI داینامیک است
		RedirectURL: "http://localhost:3000/user/oauth/callback/google",
		Scopes:      []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		// استفاده از Endpoint آماده گوگل
		Endpoint: google.Endpoint,
	}

	// تنظیمات گیت‌هاب
	githubConf := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	return &OAuthService{
		googleConfig: googleConf,
		githubConfig: githubConf,
		redis:        redis,
	}
}

func (s *OAuthService) GetGoogleAuthURL(redirectURL string, userdata []byte) (string, error) {
	key, err := s.redis.Set(context.Background(), userdata)

	if err != nil {
		return "", err
	}

	url := s.googleConfig.AuthCodeURL(key, oauth2.SetAuthURLParam("redirect_uri", redirectURL))
	return url, nil
}

func (s *OAuthService) GetGithubAuthURL(redirectURL string, userdata []byte) (string, error) {

	key, err := s.redis.Set(context.Background(), userdata)

	if err != nil {
		return "", err
	}
	url := s.githubConfig.AuthCodeURL(key, oauth2.SetAuthURLParam("redirect_uri", redirectURL))
	return url, nil
}

func (s *OAuthService) ExchangeGoogleCode(ctx context.Context, code string) (*oauth2.Token, error) {
	// اینجا نیازی به ست کردن Redirect URL نیست مگر اینکه پورت عوض شده باشد (که معمولا برای Exchange نیازی نیست اگر همان باشد)
	token, err := s.googleConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *OAuthService) ExchangeGithubCode(ctx context.Context, code string) (*oauth2.Token, error) {
	// اینجا نیازی به ست کردن Redirect URL نیست مگر اینکه پورت عوض شده باشد (که معمولا برای Exchange نیازی نیست اگر همان باشد)

	token, err := s.githubConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *OAuthService) GetGoogleUserInfo(token string) ([]byte, error) {
	client := s.googleConfig.Client(
		context.Background(),
		&oauth2.Token{AccessToken: token},
	)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("google userinfo failed: %s", body)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *OAuthService) GetGitHubUserInfo(token string) ([]byte, error) {
	// 1. ساخت یک کلاینت استاندارد هتپ
	client := &http.Client{}

	// 2. ساخت درخواست به آدرس API گیت‌هاب
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	// 3. تنظیم هدر Authorization
	// فرمت هدر گیت‌هاب به صورت "Bearer YOUR_TOKEN" است
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/json") // اطمینان از اینکه JSON دریافت می‌کنیم

	// 4. اجرای درخواست
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 5. بررسی وضعیت پاسخ
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API returned status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *OAuthService) GetGitHubUserEmail(token string) ([]byte, error) {
	// 1. ساخت یک کلاینت استاندارد هتپ
	client := &http.Client{}

	// 2. ساخت درخواست به آدرس API گیت‌هاب
	req, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return nil, err
	}

	// 3. تنظیم هدر Authorization
	// فرمت هدر گیت‌هاب به صورت "Bearer YOUR_TOKEN" است
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/json") // اطمینان از اینکه JSON دریافت می‌کنیم

	// 4. اجرای درخواست
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 5. بررسی وضعیت پاسخ
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API returned status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}
