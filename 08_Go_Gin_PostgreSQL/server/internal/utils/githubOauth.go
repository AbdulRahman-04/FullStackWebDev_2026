package utils

import (
	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func GithubOAuthUser() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.AppConfig.GithubUserId,
		ClientSecret: config.AppConfig.GithubUserSecret,
		RedirectURL:  config.AppConfig.GithubUserRedirect,
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
	}
}

func GithubOAuthAdmin() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.AppConfig.GithubAdminId,
		ClientSecret: config.AppConfig.GithubAdminSecret,
		RedirectURL:  config.AppConfig.GithubAdminRedirect,
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
	}
}
