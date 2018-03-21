package handlers

import (
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
	"github.com/julienschmidt/httprouter"
	geo "github.com/kellydunn/golang-geo"
)

// ClientHandler represent handler for handling client resource
type (
	ClientHandler struct{}
)

// Adds polygon to the database for client
func SavePolygon(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectClientDb()
	client := model.Client{}
	polygon := geo.Polygon{}

	err := database.SavePolygon(&polygon, &client, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db.Close()
}

// Retrieve client's polygons(geofences) from the database
func GetPolygons(w http.ResponseWriter, r *http.Request, c *model.Client) {
	var db = database.ConnectClientDb()
	err := database.GetPolygons(c, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db.Close()
}

// Deletes a clients polygons from the database
func DeletePolygon(w http.ResponseWriter, r *http.Request, c *model.Client) {
	var db = database.ConnectClientDb()
	err := database.DeletePolygon(c, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db.Close()
}

// Return new ClientHandler
func NewClientHandler() *UserHandler {
	return &UserHandler{}
}

// ClientUser retrieves an individual user resource
func (ch ClientHandler) GetClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

// CreateUser creates a new user resource
func (ch ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

// RemoveClient removes an existing user resource
func (ch ClientHandler) RemoveClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
