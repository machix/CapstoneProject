package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// This is a method that test response from the API
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have been hacked.")
}

// This is a method for testing response from the API
func GetPosition(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectUserDb()
	us := model.Summary{}
	err := database.QueryPosition(&us, db)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(us)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

// Post a new latitude and longitude position to the database
func PostPosition(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectUserDb()
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	err = database.PostPosition(&user, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	marshalJson(user, w)
}

// Deletes a latitude and longitude position in the database
func DeletePosition(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectUserDb()
	us := model.User{}
	err := database.DeletePosition(&us, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	marshalJson(us, w)
}

// Updates a latitude and longitude position in the database
func UpdatePosition(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectUserDb()
	us := model.User{}
	err := database.UpdatePosition(&us, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	marshalJson(us, w)
}

// Marshals the json and outputs, otherwise outputs error if unsuccesful
func marshalJson(u model.User, w http.ResponseWriter) {
	out, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}
