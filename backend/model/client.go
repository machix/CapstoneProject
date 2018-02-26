package model

import "time"

type ClientPolygon struct {
	Id        uint32
	Latitude  float32
	Longitude float32
}

type Client struct {
	ID        uint32    
	FirstName string   
	LastName  string    
	Email     string   
	Password  string  
	StatusID  uint8    
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   uint8  
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
