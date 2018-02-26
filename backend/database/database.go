package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
	"github.com/kellydunn/golang-geo"
	_ "github.com/lib/pq"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

// Connect to the postgres database
func ConnectDb() *sql.DB {
	var db *sql.DB
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
func QueryPosition(u *model.Summary, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query(
		`SELECT *
		 FROM USER_LOCATION`)
	if err != nil {
		tx.Rollback()
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
func PostPosition(u *model.User, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "INSERT INTO USER_LOCATION (Id, Latitude, Longitude) values (?, ?, ?)"

	userLocationPost, err := tx.Prepare(sqlStmt)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer userLocationPost.Close()

	_, err = userLocationPost.Exec(u.Id, u.Latitude, u.Longitude)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Queries the database to delete the user's location
func DeletePosition(u *model.User, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "DELETE Id=? FROM USER_LOCATION"

	userLocationDelete, err := tx.Prepare(sqlStmt)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer userLocationDelete.Close()

	_, err = userLocationDelete.Exec(u.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Queries the database to update the user's location position
func UpdatePosition(u *model.User, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "UPDATE USER_LOCATION SET Latitude=?"

	userLocationUpdate, err := tx.Prepare(sqlStmt)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer userLocationUpdate.Close()

	_, err = tx.Exec(sqlStmt)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Saves points in a polygon that has been drawn on the map
func SavePolygon(p *geo.Polygon, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "INSERT INTO CLIENT_POLYGON"

	userPolygonInsert, err := tx.Prepare(sqlStmt)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer userPolygonInsert.Close()

	_, err = tx.Exec(sqlStmt)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
