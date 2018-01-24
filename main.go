package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ", r.Host)
	fmt.Fprintf(w, r.Host+"\n")
}

func main() {
	http.HandleFunc("/", handler)

	port := "8001"

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
