package handlers

import (
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Adds polygon to the database for client
func SavePolygon(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectClientDb()
	client := model.Client{}
	polygon := model.Polygon{}

	err := database.SavePolygon(&polygon, &client, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	db.Close()
}

// Retrieve client's polygons(geofences) from the database
func GetPolygons(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectClientDb()
	client := model.Client{}

	err := database.GetPolygons(&client, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	db.Close()
}

// Deletes a clients polygons from the database
func DeletePolygon(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectClientDb()
	polygon := model.Polygon{}
	client := model.Client{}

	err := database.DeletePolygon(&polygon, &client, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	db.Close()
}
