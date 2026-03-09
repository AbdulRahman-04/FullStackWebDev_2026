package utils

import (
	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleOauthUser() *oauth2.Config {
	return  &oauth2.Config{
		ClientID: config.AppConfig.GoogleClientID,
		ClientSecret: config.AppConfig.GithubClientSecret,
		RedirectURL: config.AppConfig.GithubClientRedirectURL,
		Scopes: []string{
			"http://www.googleapis.com/auth/userinfo.email",
			"http://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

}
func GoogleOauthAdmin() *oauth2.Config {
	return &oauth2.Config{
		ClientID: config.AppConfig.GoogleAdminID,
		ClientSecret: config.AppConfig.GoogleAdminSecret,
		RedirectURL: config.AppConfig.GoogleAdminRedirectURL,
		Scopes: []string{
			"http://www.googleapis.com/auth/userinfo.email",
			"http://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}