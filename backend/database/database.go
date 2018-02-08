package database

import (
	"database/sql"
	"fmt"
)

//TODO: Replace these values with valid values/aliases
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

// Connect to the postgres database
func connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	errorCheck(err)
	defer db.Close()
}

// Method to handle all error checking
func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
