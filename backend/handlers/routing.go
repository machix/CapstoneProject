package handlers

import (
	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
	"github.com/gorilla/mux"
)

type Env struct {
	db model.Datastore
}

func Router() *mux.Router {
	router := mux.NewRouter()

	var db = database.ConnectClientDb()

	env := &Env{db}

	router.HandleFunc("/", Handler)
	router.HandleFunc("/position", env.GetPosition).Methods("GET")
	router.HandleFunc("/postPosition", env.PostPosition).Methods("POST")
	router.HandleFunc("/deletePosition", env.DeletePosition).Methods("DELETE")
	router.HandleFunc("/getClient", env.GetClient).Methods("GET")
	router.HandleFunc("/postClient", env.CreateClient).Methods("POST")
	router.HandleFunc("/deleteClient", env.RemoveClient).Methods("DELETE")
	router.HandleFunc("/getPolygons", env.GetPolygons).Methods("GET")
	router.HandleFunc("/savePolygon", env.SavePolygon).Methods("POST")
	router.HandleFunc("/deletePolygon", env.DeletePolygon).Methods("DELETE")
	router.HandleFunc("/createGeofence", CreateGeofence).Methods("POST")
	router.HandleFunc("/checkGeofence", CheckPointInPolygon).Methods("POST")
	router.HandleFunc("/checkPolygon", CheckPolygonOverlap).Methods("POST")
	return router
}
