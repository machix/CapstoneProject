// Contains structs for project. There is are no functions contained in this package.
package model

// Struct used to hold summary of clients
type ClientSummary struct {
	ClientSummary []Client
}

// Struct used to hold client info
type Client struct {
	ID        uint32 `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}
