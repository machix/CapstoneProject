// Contains structs for project. There is are no functions contained in this package.
package model

type ClientSummary struct {
	ClientSummary []Client
}

type Client struct {
	ID        uint32 `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type Store struct {
	ClientID uint32
}

type StoreLocation struct {
	ClientID  uint32
	Latitude  float32
	Longitude float32
}
