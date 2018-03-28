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
	router := Router()

	corsRouter := cors.Default().Handler(router)
	http.ListenAndServe(":8000", corsRouter)
}

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Handler)
	router.HandleFunc("/position", handlers.GetPosition).Methods("GET")
	router.HandleFunc("/postPosition", handlers.PostPosition).Methods("POST")
	router.HandleFunc("/deletePosition", handlers.DeletePosition).Methods("DELETE")
	router.HandleFunc("/getClient", handlers.GetClient).Methods("GET")
	router.HandleFunc("/postClient", handlers.CreateClient).Methods("POST")
	router.HandleFunc("/deleteClient", handlers.RemoveClient).Methods("DELETE")
	router.HandleFunc("/getPolygons", handlers.GetPolygons).Methods("GET")
	router.HandleFunc("/savePolygon", handlers.SavePolygon).Methods("POST")
	router.HandleFunc("/deletePolygon", handlers.DeletePolygon).Methods("DELETE")

	return router
}
