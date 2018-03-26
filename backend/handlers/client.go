package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Creates a client in the client table
func CreateClient(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectClientDb()
	client := model.Client{}

	err := database.AddNewClient(&client, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	db.Close()
}

// Retrieves all clients from the database
func GetClient(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectClientDb()
	clientSummary := model.ClientSummary{}
	err := database.GetClients(&clientSummary, db)

	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	out, err := json.Marshal(clientSummary)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	fmt.Fprintf(w, string(out))
	db.Close()
}

// Removes specified client from the database
func RemoveClient(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectClientDb()
	client := model.Client{}
	err := database.DeleteClient(&client, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	db.Close()
}
