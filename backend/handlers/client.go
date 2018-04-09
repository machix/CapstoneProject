// Handlers for http request the API endpoints
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Retrieves all clients from the database
func (env *Env) GetClient(w http.ResponseWriter, r *http.Request) {
	clientSummary := model.ClientSummary{}
	err := env.db.GetClients(&clientSummary)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(clientSummary)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

// Creates a client in the client table
func (env *Env) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client

	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = env.db.AddNewClient(&client)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer r.Body.Close()

	marshal(client, w)
}

// Removes specified client from the database
func (env *Env) RemoveClient(w http.ResponseWriter, r *http.Request) {
	client := model.Client{}
	err := json.NewDecoder(r.Body).Decode(&client)

	err = env.db.DeleteClient(&client)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
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
