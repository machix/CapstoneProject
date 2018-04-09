package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Adds polygon to the database for client
func (env *Env) SavePolygon(w http.ResponseWriter, r *http.Request) {
	client := model.Client{}
	polygon := model.Polygon{}
	err := json.NewDecoder(r.Body).Decode(&polygon)

	err = env.db.SavePolygon(&polygon, &client)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Retrieve client's polygons(geofences) from the database
func (env *Env) GetPolygons(w http.ResponseWriter, r *http.Request) {
	polygonSummary := model.PolygonSummary{}
	err := env.db.GetPolygons(&polygonSummary)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(polygonSummary)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

// Deletes a clients polygons from the database
func (env *Env) DeletePolygon(w http.ResponseWriter, r *http.Request) {
	polygon := model.Polygon{}
	client := model.Client{}

	err := env.db.DeletePolygon(&polygon, &client)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return

	}
}
