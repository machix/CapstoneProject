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
	ID        uint32    `db:"id" bson:"id,omitempty"`
	FirstName string    `db:"first_name" bson:"first_name"`
	LastName  string    `db:"last_name" bson:"last_name"`
	Email     string    `db:"email" bson:"email"`
	Password  string    `db:"password" bson:"password"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `db:"updated_at" bson:"updated_at"`
	Deleted   uint8     `db:"deleted" bson:"deleted"`
}

// UserLocation struct that represents location data about user
type UserLocation struct {
	UserID    uint32  `json:"id,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}
