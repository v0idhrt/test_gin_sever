package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MigrateDatabase(dsn string, sciptPath string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	content, err := os.ReadFile(sciptPath)
	if err != nil {
		log.Fatal("Failed to read migration scipt", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatal("Failed to execute migration script:", err)
	}

	log.Println("Migration completed successfully!")
}

func ConnectDatabase(dsn string) {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	log.Println("Database connected successfully")
}
