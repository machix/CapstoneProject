package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/database"
	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

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

// Creates a client in the client table
func CreateClient(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectClientDb()
	var client model.Client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	err = database.AddNewClient(&client, db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		db.Close()
		return
	}

	defer r.Body.Close()

	marshal(client, w)
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

// Marshals the json and outputs, otherwise outputs error if unsuccesful
func marshal(c model.Client, w http.ResponseWriter) {
	out, err := json.Marshal(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}
