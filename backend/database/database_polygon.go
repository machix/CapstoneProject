package database

import (
	"fmt"
	"strconv"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Returns all polygons associated with client
func (db *DB) GetPolygons(p *model.PolygonSummary) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	poly, err := tx.Query(
		`SELECT array_to_string(array_agg, ',') FROM 
		(SELECT array_agg( ST_x(geom)||' '||ST_y(geom))  FROM 
			(SELECT (ST_dumppoints(polygon)).geom FROM CLIENT_POLYGON
			) AS foo_1
		) AS foo_2;`)

	for poly.Next() {
		err = poly.Scan(
			&p.PolygonSummary,
		)

		if err != nil {
			return err
		}
	}

	err = poly.Err()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Saves points in a polygon that has been drawn on the map
func (db *DB) SavePolygon(p *model.Polygon, c *model.Client) error {
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

	_, err = db.Exec(sqlStmt, p.Id, p.Name, polygonStmt)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Delete a polygon that has been saved in the database
func (db *DB) DeletePolygon(p *model.Polygon, c *model.Client) error {
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
