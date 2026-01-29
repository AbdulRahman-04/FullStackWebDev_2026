package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port string
	
	URL string
	
	DB_URL string

	JWT_KEY string
	REFRESH_KEY string

	Redis_Host string
	Redis_Pass string
    Redis_DB string


	GoogleUserId string
	GoogleUserSecret string
	GoogleUserRedirect string

	GithubUserId string
	GithubUserSecret string
	GithubUserRedirect string

	GoogleAdminId string
	GoogleAdminSecret string
	GoogleAdminRedirect string

	GithubAdminId string
	GithubAdminSecret string
	GithubAdminRedirect string

	SID string
	Token string
	Phone string

	Email string
	Pass string

	GroqAPIKey string

}

var AppConfig *Config

func LoadEnv(){
	if err := godotenv.Load(); err != nil {
		log.Printf("couldn't load env file💀")
		return
	}

	cfg := &Config{
	AppName: os.Getenv("APP_NAME"),
	Port:    os.Getenv("PORT"),

	DB_URL: os.Getenv("DB_URL"),

	JWT_KEY:     os.Getenv("JWT_KEY"),
	REFRESH_KEY: os.Getenv("REFRESH_KEY"),

	URL: os.Getenv("URL"),

	GroqAPIKey: os.Getenv("GROQ_API_KEY"),

	Redis_Host: os.Getenv("REDIS_HOST"),
	Redis_Pass: os.Getenv("REDIS_PASS"),
	Redis_DB: os.Getenv("REDIS_DB"),

	// Google OAuth (User)
	GoogleUserId:       os.Getenv("GOOGLE_CLIENT_ID_USER"),
	GoogleUserSecret:  os.Getenv("GOOGLE_CLIENT_SECRET_USER"),
	GoogleUserRedirect: os.Getenv("GOOGLE_REDIRECT_URL_USER"),

	// Google OAuth (Admin)
	GoogleAdminId:       os.Getenv("GOOGLE_CLIENT_ID_ADMIN"),
	GoogleAdminSecret:  os.Getenv("GOOGLE_CLIENT_SECRET_ADMIN"),
	GoogleAdminRedirect: os.Getenv("GOOGLE_REDIRECT_URL_ADMIN"),

	// Github OAuth (User)
	GithubUserId:       os.Getenv("GITHUB_CLIENT_ID_USER"),
	GithubUserSecret:  os.Getenv("GITHUB_CLIENT_SECRET_USER"),
	GithubUserRedirect: os.Getenv("GITHUB_REDIRECT_URL_USER"),

	// Github OAuth (Admin)
	GithubAdminId:       os.Getenv("GITHUB_CLIENT_ID_ADMIN"),
	GithubAdminSecret:  os.Getenv("GITHUB_CLIENT_SECRET_ADMIN"),
	GithubAdminRedirect: os.Getenv("GITHUB_REDIRECT_URL_ADMIN"),

	// Twilio
	SID:   os.Getenv("TWILIO_ACCOUNT_SID"),
	Token: os.Getenv("TWILIO_AUTH_TOKEN"),
	Phone: os.Getenv("TWILIO_PHONE"),

	// Email
	Email: os.Getenv("EMAIL_USER"),
	Pass:  os.Getenv("EMAIL_PASS"),
}


AppConfig = cfg

log.Printf("Config loaded✅")
}