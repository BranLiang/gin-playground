// Package db bridges the sql database and the application.
//
// database/sql and lib/pg are choosen as the main database manipulator.
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// DB is a pointer to type sql.DB
var DB *sql.DB

// Init connect the postgres database
func init() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
