package geo

import (
	"math/rand"
	"testing"

	geo "github.com/kellydunn/golang-geo"
	"github.com/stretchr/testify/assert"
)

func TestCorrectness(t *testing.T) {
	polygon := randomPolygon(2000, 0.1)
	geoPoly := geo.NewPolygon(polygon)
	holes := []*geo.Point{}
	geofence := NewGeofence([][]*geo.Point{polygon, holes}, int64(20))

	for i := 0; i < 100000; i++ {
		point := randomPoint(200)
		assert.Equal(t, geofence.Inside(point), geoPoly.Contains(point))
	}
}

// Generates a random point
func randomPoint(length float64) *geo.Point {
	return geo.NewPoint(rand.Float64()*length-length/2, rand.Float64()*length-length/2)
}

// Generates a random polygon
func randomPolygon(length float64, percentageOfLength float64) []*geo.Point {
	polygon := make([]*geo.Point, 1000)
	for i := 0; i < 1000; i++ {
		polygon[i] = randomPoint(length * percentageOfLength)
	}
	return polygon
}

// Generates random point
func randomPointCustom(minLat float64, maxLat float64, minLng float64, maxLng float64, factor float64) *geo.Point {
	latRange := maxLat - minLat
	lngRange := maxLng - minLng
	return geo.NewPoint((minLat+maxLat)/2-latRange*factor/2+latRange*factor*rand.Float64(), (minLng+maxLng)/2-lngRange*factor/2+lngRange*factor*rand.Float64())
}
