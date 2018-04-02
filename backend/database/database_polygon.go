package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

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
	polygonStmt := "ST_GeometryFromText('POLYGON(("
	for i, h := range p.Coordinates {
		polygonStmt += FloatToString(h.Latitude) + " " + FloatToString(h.Longitude)
		if i < len(p.Coordinates)-1 {
			polygonStmt += ","
		}
		fmt.Println(polygonStmt)
	}
	polygonStmt += "))')"

	fmt.Println(polygonStmt)
	fmt.Println(p.Id)
	fmt.Println(p.Name)
	fmt.Println(sqlStmt)

	_, err = db.Exec(sqlStmt, p.Id, p.Name, polygonStmt)
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

// Convert Float number to a string
func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
