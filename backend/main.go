package main

//go-sqlmock library for testing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

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
	fmt.Fprintf(w, "Welcome, %s", r.URL.Path[1:])
}

// This is a method for testing response from the API
func getPosition(w http.ResponseWriter, r *http.Request) {
	us := model.Summary{}
	err := database.QueryPosition(&us)

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
}

// Deletes a latitude and longitude position in the database
func deletePosition(w http.ResponseWriter, r *http.Request) {
}

// Updates a latitude and longitude position in the database
func updatePosition(w http.ResponseWriter, r *http.Request) {
}

// Method to handle all error checking
func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
