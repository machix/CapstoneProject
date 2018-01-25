package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFun("/position", getPosition).Methods("GET")

	corsRouter := cors.Default().Handler(router);
	http.ListenAndServe(":8000", corsRouter)
}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}