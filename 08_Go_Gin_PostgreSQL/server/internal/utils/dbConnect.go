// package utils

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/lib/pq"
// 	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
// )

// var DB *sql.DB

// func ConnectPostgres() {
// 	db, err := sql.Open("postgres", config.AppConfig.DB_URL)
// 	if err != nil {
// 		log.Fatal("❌ failed to open db:", err)
// 	}

// 	// check connection
// 	if err := db.Ping(); err != nil {
// 		log.Fatal("❌ failed to connect db:", err)
// 	}

// 	DB = db
// 	log.Println("✅ Postgres connected (raw sql)")
// }

package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
)

var PostgresDB *sql.DB

func ConnectPostgres() {
	// connect db
	db, err := sql.Open("postgres", config.AppConfig.DB_URL)
	if err != nil {
		log.Fatalf("err : %s", err)
		return
	}

	// ping
	if err := db.Ping(); err != nil {
		log.Fatalf("couldn't ping db %s", err)
		return
	}

	PostgresDB = db

	log.Printf("Postgres connected✅")
}
