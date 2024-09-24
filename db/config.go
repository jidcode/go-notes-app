package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

// DB variable will hold the database connection pool
var DB *sqlx.DB

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB() *sqlx.DB {
	// Define the connection string
	connectionString := os.Getenv("DATABASE_URL")

	// Open a connection to the database
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Run migrations
	if err := goose.Up(db.DB, "migrations"); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// Set the global DB variable to the connection
	DB = db

	log.Println("Connected to the database successfully")
	return db
}
