package handlers

import (
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Adds polygon to the database for client
func AddPolygon(w http.ResponseWriter, r *http.Request) {
}

// Retrieve client's polygons(geofences) from the database
func GetPolygons(w http.ResponseWriter, r *http.Request, c *model.Client) {
	var db = database.ConnectDb()
	err := database.GetPolygons(c, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
