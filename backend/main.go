package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type User struct {
	id        uint32
	latitude  float32
	longitude float32
}

type users struct {
	UserSummary []User
}

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
	fmt.Fprintf(w, "Welcome, %s", r.URL.Path[1:])
}

// This is a method for testing response from the API
func getPosition(w http.ResponseWriter, r *http.Request) {
	us := users{}
	err := queryPosition(&us)

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

//Query the db to fetch data about user's position
func queryPosition(u *users) error {
	db := database.ConnectDb()
	rows, err := db.Query(
		`SELECT *
		 FROM "USER_LOCATION"`)

	//Return error from sql query
	if err != nil {
		return err
	}

	defer rows.Close()

	//Loop through the database query
	for rows.Next() {
		tempUser := User{}
		err = rows.Scan(
			&tempUser.id,
			&tempUser.latitude,
			&tempUser.longitude)

		if err != nil {
			return err
		}

		u.UserSummary = append(u.UserSummary, tempUser)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// Post a new latitude and longitude position to the database
func postPosition(w http.ResponseWriter, r *http.Request) {
	us := users{}
	err := queryPostPosition(&us)

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

// Query the db to post information about the user's position
func queryPostPosition(u *users) error {
	db := database.ConnectDb()
	rows, err := db.Query(
		`INSERT INTO USER_LOCATION (id, latitude, longitude)
		 VALUES (test, 3.45322, 3.23523)`)

	return err
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
