package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName         string
	Port            string
	DB_URL          string
	URL             string
	JWT_KEY         string
	JWT_REFRESH_KEY string

	GroqApiKey string
	GroqApiURL string

	REDIS_DB   string
	REDIS_HOST string
	REDIS_PASS string

	GoogleClientID          string
	GoogleClientSecret      string
	GoogleClientRedirectURL string

	GoogleAdminID          string
	GoogleAdminSecret      string
	GoogleAdminRedirectURL string

	GithubClientID          string
	GithubClientSecret      string
	GithubClientRedirectURL string

	GithubAdminID          string
	GithubAdminSecret      string
	GithubAdminRedirectURL string

	SID   string
	TOKEN string
	PHONE string

	EMAIL string
	PASS  string
}

var AppConfig *Config

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Couldn't load env variables")
	}

	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
		DB_URL:  os.Getenv("DB_URL"),
		URL:     os.Getenv("URL"),

		JWT_KEY:         os.Getenv("JWT_KEY"),
		JWT_REFRESH_KEY: os.Getenv("JWT_REFRESH_KEY"),

		REDIS_DB:   os.Getenv("REDIS_DB"),
		REDIS_HOST: os.Getenv("REDIS_HOST"),
		REDIS_PASS: os.Getenv("REDIS_PASS"),

		GroqApiKey: os.Getenv("GROQ_API_KEY"),
		GroqApiURL: os.Getenv("GROQ_API_URL"),

		GoogleClientID:          os.Getenv("GOOGLE_CLIENT_ID_USER"),
		GoogleClientSecret:      os.Getenv("GOOGLE_CLIENT_SECRET_USER"),
		GoogleClientRedirectURL: os.Getenv("GOOGLE_REDIRECT_URL_USER"),

		GoogleAdminID:          os.Getenv("GOOGLE_CLIENT_ID_ADMIN"),
		GoogleAdminSecret:      os.Getenv("GOOGLE_CLIENT_SECRET_ADMIN"),
		GoogleAdminRedirectURL: os.Getenv("GOOGLE_REDIRECT_URL_ADMIN"),

		GithubClientID:          os.Getenv("GITHUB_CLIENT_ID_USER"),
		GithubClientSecret:      os.Getenv("GITHUB_CLIENT_SECRET_USER"),
		GithubClientRedirectURL: os.Getenv("GITHUB_REDIRECT_URL_USER"),

		GithubAdminID:          os.Getenv("GITHUB_CLIENT_ID_ADMIN"),
		GithubAdminSecret:      os.Getenv("GITHUB_CLIENT_SECRET_ADMIN"),
		GithubAdminRedirectURL: os.Getenv("GITHUB_REDIRECT_URL_ADMIN"),

		SID:   os.Getenv("TWILIO_ACCOUNT_SID"),
		TOKEN: os.Getenv("TWILIO_AUTH_TOKEN"),
		PHONE: os.Getenv("TWILIO_PHONE"),

		EMAIL: os.Getenv("EMAIL_USER"),
		PASS:  os.Getenv("EMAIL_PASS"),
	}

	if cfg.DB_URL == "" || cfg.Port == "" {
		log.Fatalf("Db_url or port is missing")
	}

	if cfg.JWT_KEY == "" || cfg.JWT_REFRESH_KEY == "" {
		log.Fatalf("JWT KEY OR REFRESH KEY MISSING")
	}

	if cfg.GroqApiKey == "" || cfg.GroqApiURL == "" {
		log.Fatalf("Groq api key or url missing")
	}

	AppConfig = cfg
	log.Printf("Config Loaded✅")
}
