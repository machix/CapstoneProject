package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type Message struct {
	Text string
}

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
	m := Message{"Soon you will get some really cool info herer! It will be very cool!"}
	b, err := json.Marshal(m)

	errorCheck(err)

	w.Write(b)
}

// Post a new latitude and longitude position to the database
func postPosition(w http.ResponseWriter, r *http.Request) {
	//TODO: Implement post position function
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
