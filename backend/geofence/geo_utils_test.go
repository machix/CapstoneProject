package geo

import (
	"testing"

	geo "github.com/kellydunn/golang-geo"
	"github.com/stretchr/testify/assert"
)

// Test the utility to see if two polygons have intersecting edges
// This will inherently check the checkInsersect method
func TestHaveIntersectingEdges(t *testing.T) {
	point1 := geo.NewPoint(102.3231, 104.2243)
	point2 := geo.NewPoint(108.2352, 107.4234)
	point3 := geo.NewPoint(112.3134, 123.3223)

	var pointArray []*geo.Point
	pointArray = append(pointArray, point1)
	pointArray = append(pointArray, point2)
	pointArray = append(pointArray, point3)
	pointArray = append(pointArray, point1)

	point4 := geo.NewPoint(180.3231, 189.2243)
	point5 := geo.NewPoint(190.2352, 175.4234)
	point6 := geo.NewPoint(154.3134, 147.3223)

	var pointArray2 []*geo.Point
	pointArray2 = append(pointArray2, point4)
	pointArray2 = append(pointArray2, point5)
	pointArray2 = append(pointArray2, point6)
	pointArray2 = append(pointArray2, point4)

	test1 := haveIntersectingEdges(pointArray, pointArray2)

	assert.Equal(t, test1, false, "They should be false.")
}

// Test the hasPointInPolygon function
func TestHasPointInPolygon(t *testing.T) {
	point1 := geo.NewPoint(102.3231, 104.2243)
	point2 := geo.NewPoint(108.2352, 107.4234)
	point3 := geo.NewPoint(112.3134, 123.3223)

	var pointArray []*geo.Point
	pointArray = append(pointArray, point1)
	pointArray = append(pointArray, point2)
	pointArray = append(pointArray, point3)
	pointArray = append(pointArray, point1)

	test1 := hasPointInPolygon(pointArray, pointArray)

	assert.Equal(t, test1, true, "They should be equal.")
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
