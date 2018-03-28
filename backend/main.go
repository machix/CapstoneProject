package main

import (
	"database/sql"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/handlers"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	router := handlers.Router()

	http.Handle("/", Router())
	corsRouter := cors.Default().Handler(router)
	http.ListenAndServe(":8000", corsRouter)
}
