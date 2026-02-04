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
	JWT_KEY string
	JWT_REFRESH_KEY string

	DB_URL string
	RedisHost string
	RedisPass string
	RedisDB string
	GroqApiKey string

	GoogleClientID string
	GoogleClientSecret string
	GoogleClientRediect string

	GoogleAdminID string
	GoogleAdminSecret string
	GoogleAdminRedirect string

	GithubClientID string
	GithubClientSecret string
	GithubClientRedirect string

	GithubAdminID string
	GithubAdminSecret string
	GithubAdminRedirect string

	SID string
	Token string
	Phone string

	Email string
	Pass string
}

var AppConfig *Config

func LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Printf("Couldn't load env")
	}
	
	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),

		URL:             os.Getenv("URL"),
		JWT_KEY:         os.Getenv("JWT_KEY"),
		JWT_REFRESH_KEY: os.Getenv("JWT_REFRESH_KEY"),

		DB_URL:    os.Getenv("DB_URL"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPass: os.Getenv("REDIS_PASS"),
		RedisDB:   os.Getenv("REDIS_DB"),

		GroqApiKey: os.Getenv("GROQ_API_KEY"),

		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID_USER"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET_USER"),
		GoogleClientRediect: os.Getenv("GOOGLE_REDIRECT_URL_USER"),

		GoogleAdminID:      os.Getenv("GOOGLE_CLIENT_ID_ADMIN"),
		GoogleAdminSecret:  os.Getenv("GOOGLE_CLIENT_SECRET_ADMIN"),
		GoogleAdminRedirect: os.Getenv("GOOGLE_REDIRECT_URL_ADMIN"),

		GithubClientID:     os.Getenv("GITHUB_CLIENT_ID_USER"),
		GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET_USER"),
		GithubClientRedirect: os.Getenv("GITHUB_REDIRECT_URL_USER"),

		GithubAdminID:     os.Getenv("GITHUB_CLIENT_ID_ADMIN"),
		GithubAdminSecret: os.Getenv("GITHUB_CLIENT_SECRET_ADMIN"),
		GithubAdminRedirect: os.Getenv("GITHUB_REDIRECT_URL_ADMIN"),

		SID:   os.Getenv("TWILIO_ACCOUNT_SID"),
		Token: os.Getenv("TWILIO_AUTH_TOKEN"),
		Phone: os.Getenv("TWILIO_PHONE"),

		Email: os.Getenv("EMAIL_USER"),
		Pass:  os.Getenv("EMAIL_PASS"),
	}

	if cfg.Port == "" || cfg.DB_URL == "" {
		log.Printf("No Port no or DB_URL found in config")
		return
	}

	if cfg.JWT_KEY == "" || cfg.JWT_REFRESH_KEY == "" {
		log.Printf("NO JWT OR REFRESH KEY FOUND")
		return
	}

	if cfg.GroqApiKey == "" {
		log.Printf("No groq api key found")
		return
	}

	AppConfig = cfg
	log.Printf("Config Loaded✅")
}