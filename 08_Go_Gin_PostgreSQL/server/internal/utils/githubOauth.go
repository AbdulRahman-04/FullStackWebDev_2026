package utils

import (
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func GithubOauthUser() *oauth2.Config{
	return &oauth2.Config{
		ClientID: config.AppConfig.GithubClientID,
		ClientSecret: config.AppConfig.GithubClientSecret,
		RedirectURL: config.AppConfig.GithubClientRedirectURL,
		Scopes: []string{"read:user","user:email"},
		Endpoint: github.Endpoint,
	}
}


func GithubOauthAdmin() *oauth2.Config{
	return &oauth2.Config{
		ClientID: config.AppConfig.GithubAdminID,
		ClientSecret: config.AppConfig.GithubAdminSecret,
		RedirectURL: config.AppConfig.GithubAdminRedirectURL,
		Scopes: []string{"read:user","user:email"},
		Endpoint: github.Endpoint,
	}
}