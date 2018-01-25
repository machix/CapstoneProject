package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func getPosition(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var result := "Success"

    json.NewEncoder(w).Encode(result)
}