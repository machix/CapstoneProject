package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UserHandler represent handler for handling user resource
type (
	UserHandler struct{}
)

// Return new UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetUser retrieves an individual user resource
func (uh UserHandler) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

// CreateUser creates a new user resource
func (uh UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// RemoveUser removes an existing user resource
func (uh UserHandler) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
