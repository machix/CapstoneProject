package handlers

import (
	"encoding/json"
	"net/http"

	geofence "github.com/NaturalFractals/CapstoneProject/backend/geofence"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
	"github.com/kellydunn/golang-geo"
)

// Creates a geofence with the given points
func CreateGeofence(w http.ResponseWriter, r *http.Request) {
	polygon := model.Polygon{}
	err := json.NewDecoder(r.Body).Decode(&polygon)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Extract points out of polygon
	var points [][]*geo.Point
	for i, value := range polygon.Coordinates {
		points[i][0] = geo.NewPoint(value.Latitude, value.Longitude)
	}

	// Create the geofence
	geofence.NewGeofence(points)
}

// Checks to see if a point is contained within a polygon
func CheckPointInPolygon(w http.ResponseWriter, r *http.Request) {
}

// Check Polygon overlap
func CheckPolygonOverlap(w http.ResponseWriter, r *http.Request) {

}
