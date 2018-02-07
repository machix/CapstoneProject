package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

// Location struct that represents data stored in the database
type Location struct {
	ID        string  `json:"id,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

type Message struct {
	Text string
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

// Locations to be posted to the database
var locations []Location

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	errorCheck(err)
	defer db.Close()

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
	fmt.Fprintf(w, "Welcome, %s", r.URL.Path[1:])
}

// This is a method for testing response from the API
func getPosition(w http.ResponseWriter, r *http.Request) {
	m := Message{"Soon you will get some really cool info herer! It will be very cool!"}
	b, err := json.Marshal(m)

	errorCheck(err)

	w.Write(b)
}

// Post a new latitude and longitude position to the database
func postPosition(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var location Location
	_ = json.NewDecoder(r.Body).Decode(&location)
	location.ID = params["id"]
	locations = append(locations, location)
	json.NewEncoder(w).Encode(locations)
}

// Deletes a latitude and longitude position in the database
func deletePosition(w http.ResponseWriter, r *http.Request) {
	//TODO: Implement delete from database
}

// Updates a latitude and longitude position in the database
func updatePosition(w http.ResponseWriter, r *http.Request) {
	//TODO: Implement update to database
}

// Method to handle all error checking
func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
