package model

type PolygonSummary struct {
	PolygonSummary []Polygon
}

type Polygon struct {
	Id          int32        `json:"id"`
	Name        string       `json:"name"`
	Coordinates []Coordinate `json:"points"`
}

type Coordinate struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}
