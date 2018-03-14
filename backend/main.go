package main

import (
	"database/sql"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Handler)
	router.HandleFunc("/position", handlers.GetPosition).Methods("GET")
	router.HandleFunc("/postPosition", handlers.PostPosition).Methods("POST")
	router.HandleFunc("/deletePosition", handlers.DeletePosition).Methods("DELETE")

	corsRouter := cors.Default().Handler(router)
	http.ListenAndServe(":8000", corsRouter)
}
