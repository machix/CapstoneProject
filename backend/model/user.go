package model

import (
	"time"
)

type Summary struct {
	UserSummary []User
}

type User struct {
	Id        uint32
	Latitude  float32
	Longitude float32
}

type UserInfo struct {
	ID        uint32
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   uint8
}

// UserLocation struct that represents location data about user
type UserLocation struct {
	UserID    uint32
	Latitude  float32
	Longitude float32
}
