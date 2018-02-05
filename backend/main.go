package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler)
	router.HandleFunc("/position", getPosition).Methods("GET")
	router.HandleFunc("/postPosition", postPosition).Methods("POST")
	router.HandleFunc("/deletePosition", deletePosition).Methods("DELETE")
	router.HandleFunc("/updatePosition", updatePosition).Methods("UPDATE")

	corsRouter := cors.Default().Handler(router)
	http.ListenAndServe(":8000", corsRouter)
}

type Message struct {
	Text string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s", r.URL.Path[1:])
}

func getPosition(w http.ResponseWriter, r *http.Request) {
	m := Message{"Soon you will get some really cool info herer! It will be very cool!"}
	b, err := json.Marshal(m)

	errorCheck(err)

	w.Write(b)
}

func postPosition(w http.ResponseWriter, r *http.Request) {

}

func deletePosition(w http.ResponseWriter, r *http.Request) {

}

func updatePosition(w http.ResponseWriter, r *http.Request) {

}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
