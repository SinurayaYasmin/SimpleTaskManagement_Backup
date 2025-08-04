package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL not set in .env")
	}
	var dbErr error
	DB, dbErr = sql.Open("postgres", dsn)

	if dbErr != nil {
		log.Fatalf("Failed to connect to database: %v", dbErr)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("DB unreachable: ", err)
	}

	fmt.Println("✅ Successfully connected to database")
}
