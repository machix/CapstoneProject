package handlers

import (
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/kellydunn/golang-geo"
	"github.com/mattermost/platform/model"
)

// Adds polygon to the database for client
func AddPolygon(w http.ResponseWriter, r *http.Request, p *geo.Polygon) {
	var db = database.ConnectDb()
	err := database.SavePolygon(p, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Retrieve client's polygons(geofences) from the database
func GetPolygons(w http.ResponseWriter, r *http.Request, c *model.Client) {
	var db = database.ConnectDb()
	err := database.GetPolygons(c, p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
