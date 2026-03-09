package utils

import (
	"log"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func ConnectPostgres() {
	db, err := gorm.Open(postgres.Open(config.AppConfig.DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't connect postgres")
	}

	PostgresDB = db
	log.Printf("PostgreSQL Connected✅")
}