package model

import "time"

type Client struct {
	ID        uint32    `db:"id" bson:"id,omitempty"`
	FirstName string    `db:"first_name" bson:"first_name"`
	LastName  string    `db:"last_name" bson:"last_name"`
	Email     string    `db:"email" bson:"email"`
	Password  string    `db:"password" bson:"password"`
	StatusID  uint8     `db:"status_id" bson:"status_id"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `db:"updated_at" bson:"updated_at"`
	Deleted   uint8     `db:"deleted" bson:"deleted"`
}

type Store struct {
	ClientID uint32 `db:"client_id" bson:"client_id"`
	//Store location
}

type StoreLocation struct {
	ClientID  uint32  `json:"id,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}
