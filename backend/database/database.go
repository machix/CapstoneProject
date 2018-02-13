package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/mattermost/platform/model"
)

var db *sql.DB

type users struct {
	UserSummary []User
}

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
func QueryPosition(u *model.User) error {
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
		tempUser := User{}
		err = rows.Scan(
			&tempUser.id,
			&tempUser.latitude,
			&tempUser.longitude)

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

// Method to handle all error checking
func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
