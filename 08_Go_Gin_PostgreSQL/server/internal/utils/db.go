package utils

import (
	"log"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
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
	log.Printf("Postgres connected✅")

}
