package model

// Struct used to hold summary of users
type Summary struct {
	UserSummary []User
}

// Struct used to hold user info
type User struct {
	Id        uint32  `json:"Id"`
	Latitude  float32 `json:"Latitude"`
	Longitude float32 `json:"Longitude"`
}
