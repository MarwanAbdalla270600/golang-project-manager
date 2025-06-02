package database

import (
	"database/sql"
	"log"
	"os"
	"webapi/internal/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB      *sql.DB
	Queries *db.Queries
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not read .env File")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Missing DATABASE_URL in .env")
	}

	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	Queries = db.New(DB)

	log.Println("✅ Connected to DB and sqlc ready.")
}
