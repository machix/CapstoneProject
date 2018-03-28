package handlers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", Handler)
	router.HandleFunc("/position", GetPosition).Methods("GET")
	router.HandleFunc("/postPosition", PostPosition).Methods("POST")
	router.HandleFunc("/deletePosition", DeletePosition).Methods("DELETE")
	router.HandleFunc("/getClient", GetClient).Methods("GET")
	router.HandleFunc("/postClient", CreateClient).Methods("POST")
	router.HandleFunc("/deleteClient", RemoveClient).Methods("DELETE")
	router.HandleFunc("/getPolygons", GetPolygons).Methods("GET")
	router.HandleFunc("/savePolygon", SavePolygon).Methods("POST")
	router.HandleFunc("/deletePolygon", DeletePolygon).Methods("DELETE")

	return router
}
