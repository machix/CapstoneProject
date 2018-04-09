package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// This is a method that test response from the API
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have been hacked.")
}

// This is a method for testing response from the API
func (env *Env) GetPosition(w http.ResponseWriter, r *http.Request) {
	us := model.Summary{}
	err := env.db.QueryPosition(&us)

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
func (env *Env) PostPosition(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	err = env.db.PostPosition(&user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer r.Body.Close()

	marshalJson(user, w)
}

// Deletes a latitude and longitude position in the database
func (env *Env) DeletePosition(w http.ResponseWriter, r *http.Request) {
	us := model.User{}
	err := json.NewDecoder(r.Body).Decode(&us)

	err = env.db.DeletePosition(&us)
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
