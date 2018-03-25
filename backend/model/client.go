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
	//Store location
}

type StoreLocation struct {
	ClientID  uint32
	Latitude  float32
	Longitude float32
}
