package geo

import (
	"testing"

	geo "github.com/kellydunn/golang-geo"
	"github.com/stretchr/testify/assert"
)

// Test the utility to see if two polygons have intersecting edges
// This will inherently check the checkInsersect method
func TestHaveIntersectingEdges(t *testing.T) {
}

// Test the project function to determine projection
func TestProject(t *testing.T) {

}

// Test the hasPointInPolygon function
func TestHasPointInPolygon(t *testing.T) {

}

// Test vectorDifference
func TestVectorDifference(t *testing.T) {
	point1 := geo.NewPoint(102.3231, 104.2243)
	point2 := geo.NewPoint(108.2352, 107.4234)

	test1 := vectorDifference(point1, point2)

	pointTest1 := geo.NewPoint(-5.9121000000000095, -3.1991000000000014)

	assert.Equal(t, test1, pointTest1, "They should be equal.")
}

// Test vectorCrossProdcut
func TestVectorCrossProduct(t *testing.T) {
	point1 := geo.NewPoint(102.3231, 104.2243)
	point2 := geo.NewPoint(108.2352, 107.4234)

	test1 := vectorCrossProduct(point1, point2)

	assert.Equal(t, test1, -288.84265482000046, "They should be equal.")
}
