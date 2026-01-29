// package config

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// type Config struct {
// 	AppName string
// 	Port    string
// 	URL     string

// 	DBUrl string

// 	JWTSecret  string
// 	JWTRefresh string

// 	RedisHost string
// 	RedisPass string
// 	RedisDB   string

// 	EmailUser string
// 	EmailPass string

// 	TwilioSID   string
// 	TwilioToken string
// 	TwilioPhone string

// 	// Google OAuth (User)
// 	GoogleUserID       string
// 	GoogleUserSecret  string
// 	GoogleUserRedirect string

// 	// Google OAuth (Admin)
// 	GoogleAdminID       string
// 	GoogleAdminSecret  string
// 	GoogleAdminRedirect string

// 	// Github OAuth (User)
// 	GithubUserID       string
// 	GithubUserSecret  string
// 	GithubUserRedirect string

// 	// Github OAuth (Admin)
// 	GithubAdminID       string
// 	GithubAdminSecret  string
// 	GithubAdminRedirect string

// 	// AI
// 	GroqAPIKey string
// }

// var AppConfig *Config

// func LoadConfig() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("⚠️ no .env file found, using system env")
// 	}

// 	cfg := &Config{
// 		AppName: os.Getenv("APP_NAME"),
// 		Port:    os.Getenv("PORT"),
// 		URL:     os.Getenv("URL"),

// 		DBUrl: os.Getenv("DB_URL"),

// 		JWTSecret:  os.Getenv("JWT_SECRET_KEY"),
// 		JWTRefresh: os.Getenv("JWT_REFRESH_KEY"),

// 		RedisHost: os.Getenv("REDIS_HOST"),
// 		RedisPass: os.Getenv("REDIS_PASS"),
// 		RedisDB:   os.Getenv("REDIS_DB"),

// 		EmailUser: os.Getenv("EMAIL_USER"),
// 		EmailPass: os.Getenv("EMAIL_PASS"),

// 		TwilioSID:   os.Getenv("TWILIO_ACCOUNT_SID"),
// 		TwilioToken: os.Getenv("TWILIO_AUTH_TOKEN"),
// 		TwilioPhone: os.Getenv("TWILIO_PHONE"),

// 		GoogleUserID:       os.Getenv("GOOGLE_CLIENT_ID_USER"),
// 		GoogleUserSecret:  os.Getenv("GOOGLE_CLIENT_SECRET_USER"),
// 		GoogleUserRedirect: os.Getenv("GOOGLE_REDIRECT_URL_USER"),

// 		GoogleAdminID:       os.Getenv("GOOGLE_CLIENT_ID_ADMIN"),
// 		GoogleAdminSecret:  os.Getenv("GOOGLE_CLIENT_SECRET_ADMIN"),
// 		GoogleAdminRedirect: os.Getenv("GOOGLE_REDIRECT_URL_ADMIN"),

// 		GithubUserID:       os.Getenv("GITHUB_CLIENT_ID_USER"),
// 		GithubUserSecret:  os.Getenv("GITHUB_CLIENT_SECRET_USER"),
// 		GithubUserRedirect: os.Getenv("GITHUB_REDIRECT_URL_USER"),

// 		GithubAdminID:       os.Getenv("GITHUB_CLIENT_ID_ADMIN"),
// 		GithubAdminSecret:  os.Getenv("GITHUB_CLIENT_SECRET_ADMIN"),
// 		GithubAdminRedirect: os.Getenv("GITHUB_REDIRECT_URL_ADMIN"),

// 		GroqAPIKey: os.Getenv("GROQ_API_KEY"),
// 	}

// 	// ---- REQUIRED CHECKS ----
// 	if cfg.Port == "" || cfg.DBUrl == "" {
// 		log.Fatal("❌ PORT or DB_URL missing")
// 	}

// 	if cfg.JWTSecret == "" || cfg.JWTRefresh == "" {
// 		log.Fatal("❌ JWT keys missing")
// 	}

// 	if cfg.GroqAPIKey == "" {
// 		log.Fatal("❌ GROQ_API_KEY missing")
// 	}

// 	AppConfig = cfg
// 	log.Println("✅ Config loaded successfully")
// }

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string
	URL     string
	DB_URL  string

	JWT_SECRET_KEY  string
	JWT_REFRESH_KEY string

	RedisHost string
	RedisPass string
	RedisDB   string

	GoogleUserID       string
	GoogleUserSecret   string
	GoogleUserRedirect string

	GithubUserID       string
	GithubUserSecret   string
	GithubUserRedirect string

	GoogleAdminID       string
	GoogleAdminSecret   string
	GoogleAdminRedirect string

	GithubAdminID       string
	GithubAdminSecret   string
	GithubAdminRedirect string

	Sid   string
	Token string
	Phone string

	Email string
	Pass  string

	GroqAPIKey string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("couldn't load config")
		return
	}
	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
		URL:     os.Getenv("URL"),
		DB_URL:  os.Getenv("DB_URL"),

		JWT_SECRET_KEY:  os.Getenv("JWT_SECRET_KEY"),
		JWT_REFRESH_KEY: os.Getenv("JWT_REFRESH_KEY"),

		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPass: os.Getenv("REDIS_PASS"),
		RedisDB:   os.Getenv("REDIS_DB"),

		GoogleUserID:       os.Getenv("GOOGLE_CLIENT_ID_USER"),
		GoogleUserSecret:   os.Getenv("GOOGLE_CLIENT_SECRET_USER"),
		GoogleUserRedirect: os.Getenv("GOOGLE_REDIRECT_URL_USER"),

		GithubUserID:       os.Getenv("GITHUB_CLIENT_ID_USER"),
		GithubUserSecret:   os.Getenv("GITHUB_CLIENT_SECRET_USER"),
		GithubUserRedirect: os.Getenv("GITHUB_REDIRECT_URL_USER"),

		GoogleAdminID:       os.Getenv("GOOGLE_CLIENT_ID_ADMIN"),
		GoogleAdminSecret:   os.Getenv("GOOGLE_CLIENT_SECRET_ADMIN"),
		GoogleAdminRedirect: os.Getenv("GOOGLE_REDIRECT_URL_ADMIN"),

		GithubAdminID:       os.Getenv("GITHUB_CLIENT_ID_ADMIN"),
		GithubAdminSecret:   os.Getenv("GITHUB_CLIENT_SECRET_ADMIN"),
		GithubAdminRedirect: os.Getenv("GITHUB_REDIRECT_URL_ADMIN"),

		Sid:   os.Getenv("TWILIO_ACCOUNT_SID"),
		Token: os.Getenv("TWILIO_AUTH_TOKEN"),
		Phone: os.Getenv("TWILIO_PHONE"),

		Email: os.Getenv("EMAIL_USER"),
		Pass:  os.Getenv("EMAIL_PASS"),

		GroqAPIKey: os.Getenv("GROQ_API_KEY"),
	}

	if cfg.DB_URL == "" || cfg.Port == "" {
		log.Fatalf("no db url or port no found")
		return
	}

	if cfg.JWT_SECRET_KEY == "" || cfg.JWT_REFRESH_KEY == "" {
		log.Fatalf("no secret key and refresh key found")
		return
	}

	if cfg.GroqAPIKey == "" {
		log.Fatalf("no groq api key found")
		return
	}

	AppConfig = cfg

	log.Printf("Config loaded✅")

}
