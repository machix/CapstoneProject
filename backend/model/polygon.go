package model

// Struct used to hold summary of postgis query
type PolygonSummary struct {
	PolygonSummary string
}

// Struct used to hold info on polygon
type Polygon struct {
	Id          int32        `json:"id"`
	Name        string       `json:"name"`
	Coordinates []Coordinate `json:"points"`
}

// Struct used to represent a coordinate/point
type Coordinate struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}
