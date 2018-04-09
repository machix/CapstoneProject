package model

import (
	"database/sql"
)

type Datastore interface {
	QueryPosition(*Summary) error
	PostPosition(*User) error
	DeletePosition(*User) error
	GetClients(*ClientSummary) error
	AddNewClient(*Client) error
	DeleteClient(*Client) error
	DeletePolygon(*Polygon, *Client) error
	SavePolygon(*Polygon, *Client) error
	GetPolygons(*PolygonSummary) error
}

type DB struct {
	*sql.DB
}
