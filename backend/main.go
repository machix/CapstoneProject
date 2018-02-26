package main

import (
	"database/sql"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/handlers"
	"github.com/kellydunn/golang-geo"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Handler)
	router.HandleFunc("/position", handlers.GetPosition).Methods("GET")
	router.HandleFunc("/postPosition", handlers.PostPosition).Methods("POST")
	router.HandleFunc("/deletePosition", handlers.DeletePosition).Methods("DELETE")
	router.HandleFunc("/updatePosition", handlers.UpdatePosition).Methods("UPDATE")

	corsRouter := cors.Default().Handler(router)
	http.ListenAndServe(":8000", corsRouter)
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
