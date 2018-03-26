package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
	_ "github.com/lib/pq"
)

// Database Constants. Uses environment variables
const (
	dbhostclient = "DBHOST_CLIENT"
	dbportclient = "DBPORT_CLIENT"
	dbuserclient = "DBUSER_CLIENT"
	dbpassclient = "DBPASS_CLIENT"
	dbnameclient = "DBNAME_CLIENT"
)

// Connect to the postgres database
func ConnectClientDb() *sql.DB {
	var db *sql.DB
	config := dbClientConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhostclient], config[dbportclient],
		config[dbuserclient], config[dbpassclient], config[dbnameclient])

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
func dbClientConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhostclient)

	if !ok {
		panic("DBHOST_CLIENT environment variable required but not set")
	}

	port, ok := os.LookupEnv(dbportclient)
	if !ok {
		panic("DBPORT_CLIENT environment variable required but not set")
	}

	user, ok := os.LookupEnv(dbuserclient)
	if !ok {
		panic("DBUSER_CLIENT environment variable required but not set")
	}

	password, ok := os.LookupEnv(dbpassclient)
	if !ok {
		panic("DBPASS_CLIENT environment variable required but not set")
	}

	name, ok := os.LookupEnv(dbnameclient)
	if !ok {
		panic("DBNAME_CLIENT environment variable required but not set")
	}

	conf[dbhostclient] = host
	conf[dbportclient] = port
	conf[dbuserclient] = user
	conf[dbpassclient] = password
	conf[dbnameclient] = name
	return conf
}

// Returns all polygons associated with client
func GetPolygons(c *model.Client, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "SELECT polygon FROM CLIENT_POLYGON WHERE id=$1"

	clientPolygonRetrieve, err := tx.Prepare(sqlStmt)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer clientPolygonRetrieve.Close()

	_, err = tx.Exec(sqlStmt, c.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Saves points in a polygon that has been drawn on the map
func SavePolygon(p *model.Polygon, c *model.Client, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "INSERT INTO CLIENT_POLYGON (id, name, polygon) VALUES ($1, $2, $3)"

	// Format the Polygon string for insertion into the database
	polygonStmt := "POLYGON(("
	for _, h := range p.Coordinates {
		for _, i := range h {
			polygonStmt += FloatToString(i) + " "
		}
		polygonStmt += ","
	}
	polygonStmt += "))"

	fmt.Println(polygonStmt)

	_, err = tx.Exec(sqlStmt, p.Id, p.Name, polygonStmt)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Delete a polygon that has been saved in the database
func DeletePolygon(p *model.Polygon, c *model.Client, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "DELETE * WHERE id=$1 and polygon=$2"

	_, err = tx.Exec(sqlStmt, c.ID, p)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Retrieves all clients from the client database
func GetClients(c *model.ClientSummary, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query(
		`SELECT * 
		 FROM CLIENT`)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer rows.Close()

	//Loop through database query
	for rows.Next() {
		tempClient := model.Client{}
		err = rows.Scan(
			&tempClient.ID,
			&tempClient.FirstName,
			&tempClient.LastName)

		if err != nil {
			return err
		}

		c.ClientSummary = append(c.ClientSummary, tempClient)
	}

	err = rows.Err()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Add new client to the client database
func AddNewClient(c *model.Client, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "INSERT INTO CLIENT (id, first_name, last_name) VALUES ($1, $2, $3)"

	_, err = db.Exec(sqlStmt, c.ID, c.FirstName, c.LastName)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Delete the information of a current client
func DeleteClient(c *model.Client, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "DELETE * WHERE id=$1"

	_, err = db.Exec(sqlStmt, c.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Convert Float number to a string
func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
