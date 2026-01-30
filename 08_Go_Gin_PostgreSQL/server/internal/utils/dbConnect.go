package utils

import (
	"log"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func ConnectPostgres(){
 
	db, err := gorm.Open(postgres.Open(config.AppConfig.DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatalf("err %s", err)
		return
	}

	PostgresDB = db

	log.Printf("PostgreSQL connected✅")



}