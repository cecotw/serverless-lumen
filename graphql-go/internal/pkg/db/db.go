package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// Postgres Driver
	_ "github.com/lib/pq"
)

// Connect connect to a database
func Connect() *sqlx.DB {
	dbinfo := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE"),
		os.Getenv("SSL_MODE"),
	)
	db, err := sqlx.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
