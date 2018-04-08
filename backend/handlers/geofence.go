package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	geofence "github.com/NaturalFractals/CapstoneProject/backend/geofence"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
	"github.com/kellydunn/golang-geo"
)

var geofences *geofence.Geofence

// Creates a geofence with the given points
func CreateGeofence(w http.ResponseWriter, r *http.Request) {
	var polygon model.Polygon
	err := json.NewDecoder(r.Body).Decode(&polygon)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Extract points out of polygon
	var points [][]*geo.Point
	var firstPoint []*geo.Point
	for _, value := range polygon.Coordinates {
		firstPoint = append(firstPoint, geo.NewPoint(value.Latitude, value.Longitude))
	}

	points = append(points, firstPoint)

	// Create the geofence
	tempFence := geofence.NewGeofence(points)

	geofences = tempFence
}

// Checks to see if a point is contained within a polygon
func CheckPointInPolygon(w http.ResponseWriter, r *http.Request) {
	var coordinate model.Coordinate
	err := json.NewDecoder(r.Body).Decode(&coordinate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	point := geo.NewPoint(coordinate.Latitude, coordinate.Longitude)
	inPoint := geofences.Inside(point)

	out, err := json.Marshal(inPoint)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

// Check Polygon overlap
func CheckPolygonOverlap(w http.ResponseWriter, r *http.Request) {
}
