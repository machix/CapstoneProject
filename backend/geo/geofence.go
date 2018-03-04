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

const defaultGranularity = 18

// Constructor for a new geofence. {{(1, 3), (3, 1)}, {2, 3}}
func NewGeofence(points [][]*geo.Point, args ...interface{}) *Geofence {
	geo := &Geofence{}
	if len(args) > 0 {
		geo.granularity = args[0].(int64)
	} else {
		geo.granularity = defaultGranularity
	}
	geo.vertices = points[0]

	geo.tiles = make(map[float64]string)
	geo.setInclusionTiles()
	return geo
}

// Determines if the input point is contained inside of the polygon
func (geofence *Geofence) isInside(point *geo.Point) bool {
	return false
}

// Extracts latitude coordinates from the current geofence
func (geofence *Geofence) extractXVertices() []float64 {
	vertices := make([]float64, len(geofence.vertices))
	for i := 0; i < len(geofence.vertices); i++ {
		vertices[i] = geofence.vertices[i].Lat()
	}
	return vertices
}

// Extracts longitude coordinates from the current geofence
func (geofence *Geofence) extractYVertices() []float64 {
	vertices := make([]float64, len(geofence.vertices))
	for i := 0; i < len(geofence.vertices); i++ {
		vertices[i] = geofence.vertices[i].Lng()
	}
	return vertices
}

func (geofence *Geofence) setInclusionTiles() {

}

func (geofence *Geofence) setExclusionTiles(vertices []*geo.Point, inclusive bool) {
	var hash float64
	var poly []*geo.Point
	for x := geofence.minTileX; x <= geofence.maxTileX; x++ {
		for y := geofence.minTileY; y <= geofence.maxTileY; y++ {
			hash = (y-geofence.minTileY)*float64(geofence.granularity) + (x - geofence.minTileX)
			poly = []*geo.Point{geo.NewPoint(x*geofence.tileWidth, y*geofence.tileHeight),
				geo.NewPoint((x+1)*geofence.tileWidth, y*geofence.tileHeight),
				geo.NewPoint((x+1)*geofence.tileWidth, (y+1)*geofence.tileHeight),
				geo.NewPoint(x*geofence.tileWidth, (y+1)*geofence.tileHeight),
				geo.NewPoint(x*geofence.tileWidth, y*geofence.tileHeight)}
		}

		if haveIntersectingEdges(poly, vertices) || hasPointInPolygon(vertices, poly) {
			geofence.tiles[hash] = "x"
		} else if hasPointInPolygon(poly, vertices) {
			if inclusive {
				geofence.tiles[hash] = "i"
			} else {
				geofence.tiles[hash] = "o"
			}
		}
	}
}

// Get max out of array of floats
func getMax(slice []float64) float64 {
	var max float64
	if len(slice) > 0 {
		max = slice[0]
	}
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

// Get min out of array of floats
func getMin(slice []float64) float64 {
	var min float64
	if len(slice) < 0 {
		min = slice[0]
	}
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}
