package utils

import (
	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleOAuthUser() *oauth2.Config{
	return &oauth2.Config{
		ClientID: config.AppConfig.GoogleUserId,
		ClientSecret: config.AppConfig.GoogleUserSecret,
		RedirectURL: config.AppConfig.GoogleUserRedirect,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleOAuthAdmin() *oauth2.Config{
	return &oauth2.Config{
		ClientID: config.AppConfig.GoogleAdminId,
		ClientSecret: config.AppConfig.GoogleAdminSecret,
		RedirectURL: config.AppConfig.GoogleAdminRedirect,
		Scopes: []string {
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}