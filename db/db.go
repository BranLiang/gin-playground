// Package db bridges the sql database and the application.
//
// database/sql and lib/pg are choosen as the main database manipulator.
package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// DbName is the postgres database name for application
	DbName = "gin_playground_db"
)

// DB is a pointer to type sql.DB
var DB *sql.DB

// Init connect the postgres database
func Init() {
	dbinfo := fmt.Sprintf("user=postgres password=postgres dbname=%s sslmode=disable", DbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

// GetDB returns an opened database pointer
func GetDB() *sql.DB {
	return DB
}
