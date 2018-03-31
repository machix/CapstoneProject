package model

type Polygon struct {
	Id          int32
	Name        string
	Coordinates []Coordinate
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}
