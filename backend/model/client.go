package model

import (
	"github.com/kellydunn/golang-geo"
)

type ClientPolygon struct {
	Id      uint32
	polygon geo.Polygon
}

type Client struct {
	ID        uint32
	FirstName string
	LastName  string
}

type Store struct {
	ClientID uint32
	//Store location
}

type StoreLocation struct {
	ClientID  uint32
	Latitude  float32
	Longitude float32
}
