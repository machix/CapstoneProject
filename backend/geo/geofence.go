package geo

import "github.com/kellydunn/golang-geo"

type Geofence struct {
	vertices    []*geo.Point
	tiles       map[float64]string
	granularity int64
	minX        float64
	maxX        float64
	minY        float64
	maxY        float64
	tileWidth   float64
	tileHeight  float64
	minTileX    float64
	maxTileX    float64
	minTileY    float64
	maxTileY    float64
}

// Constructor for a new geofence. {{(1, 3), (3, 1)}, {2, 3}}
func NewGeofence(points [][]*geo.Point, args ...interface{}) *Geofence {
	geo := &Geofence{}

	return geo
}

// Extracts latitude coordinates from the current geofence
func extractXVertices() {

}

// Extracts longitude coordinates from the current geofence
func extractYVertices() {

}

func getMax(slice []float64) {

}

func getMin(slice []float64) {

}
