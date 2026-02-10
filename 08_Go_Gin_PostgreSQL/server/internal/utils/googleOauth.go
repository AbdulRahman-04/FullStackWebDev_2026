package utils

import (
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleOauthUser() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.AppConfig.GoogleClientID,
		ClientSecret: config.AppConfig.GoogleClientSecret,
		RedirectURL:  config.AppConfig.GoogleClientRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleOauthAdmin() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.AppConfig.GoogleAdminID,
		ClientSecret: config.AppConfig.GoogleAdminSecret,
		RedirectURL:  config.AppConfig.GoogleAdminRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
