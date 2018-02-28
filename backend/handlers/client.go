package handlers

import (
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
	"github.com/julienschmidt/httprouter"
)

// Adds polygon to the database for client
func AddPolygon(w http.ResponseWriter, r *http.Request) {
}

// Retrieve client's polygons(geofences) from the database
func GetPolygons(w http.ResponseWriter, r *http.Request, c *model.Client) {
	var db = database.ConnectDb()
	err := database.GetPolygons(c, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Return new ClientHandler
func NewClientHandler() *UserHandler {
	return &UserHandler{}
}

// ClientUser retrieves an individual user resource
func (uh UserHandler) GetClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

// CreateUser creates a new user resource
func (uh UserHandler) CreateClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// RemoveClient removes an existing user resource
func (uh UserHandler) RemoveClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
