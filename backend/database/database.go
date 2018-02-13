package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

// Connect to the postgres database
func ConnectDb() *sql.DB {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

// Checks to ensure all of correct environment varaibles are set
func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)

	if !ok {
		panic("DBHOST environment variable required but not set")
	}

	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}

	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}

	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}

	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}

	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

//Query the db to fetch data about user's position
func QueryPosition(u *model.Summary) error {
	rows, err := db.Query(
		`SELECT *
		 FROM "USER_LOCATION"`)

	//Return error from sql query
	if err != nil {
		return err
	}

	defer rows.Close()

	//Loop through the database query
	for rows.Next() {
		tempUser := model.User{}
		err = rows.Scan(
			&tempUser.Id,
			&tempUser.Latitude,
			&tempUser.Longitude)

		if err != nil {
			return err
		}

		u.UserSummary = append(u.UserSummary, tempUser)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// Query the db to post information about the user's position
func QueryPostPosition(u *model.User) error {
	//db := database.ConnectDb()
	// rows, err := db.Query(
	// 	`INSERT INTO USER_LOCATION (id, latitude, longitude)
	// 	 VALUES (test, 3.45322, 3.23523)`)
	err := fmt.Errorf("")
	return err
}

// Queries the database to delete the user's location
func QueryDeletePosition(u *model.User) error {
	err := fmt.Errorf("")
	return err
}

// Queries the database to update the user's location position
func QueryUpdatePosition(u *model.User) error {
	err := fmt.Errorf("")
	return err
}

// Method to handle all error checking
func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
