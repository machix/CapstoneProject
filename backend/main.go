package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
	"github.com/kellydunn/golang-geo"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler)
	router.HandleFunc("/position", getPosition).Methods("GET")
	router.HandleFunc("/postPosition", postPosition).Methods("POST")
	router.HandleFunc("/deletePosition", deletePosition).Methods("DELETE")
	router.HandleFunc("/updatePosition", updatePosition).Methods("UPDATE")

	corsRouter := cors.Default().Handler(router)
	http.ListenAndServe(":8000", corsRouter)
}

// This is a method that test response from the API
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have been hacked.")
}

// This is a method for testing response from the API
func getPosition(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectDb()
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
func postPosition(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectDb()
	us := model.User{}
	err := database.PostPosition(&us, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	marshalJson(us, w)
}

// Deletes a latitude and longitude position in the database
func deletePosition(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectDb()
	us := model.User{}
	err := database.DeletePosition(&us, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	marshalJson(us, w)
}

// Updates a latitude and longitude position in the database
func updatePosition(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectDb()
	us := model.User{}
	err := database.UpdatePosition(&us, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	marshalJson(us, w)
}

// Adds polygon to the database for client
func addPolygon(w http.ResponseWriter, r *http.Request, p *geo.Polygon) {
	var db = database.ConnectDb()
	err := database.SavePolygon(p, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
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