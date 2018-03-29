package main

import (
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/handlers"
	"github.com/rs/cors"
)

func main() {
	router := handlers.Router()

	http.Handle("/", router)
	corsRouter := cors.Default().Handler(router)
	http.ListenAndServe(":8000", corsRouter)
}
