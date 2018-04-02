package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Adds polygon to the database for client
func SavePolygon(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectClientDb()
	client := model.Client{}
	polygon := model.Polygon{}
	err := json.NewDecoder(r.Body).Decode(&polygon)

	err = database.SavePolygon(&polygon, &client, db)
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
	polygonSummary := model.PolygonSummary{}
	err := database.GetPolygons(&polygonSummary, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	out, err := json.Marshal(polygonSummary)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	fmt.Fprintf(w, string(out))
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
