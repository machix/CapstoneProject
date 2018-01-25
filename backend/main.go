package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ", r.Host)
	fmt.Fprintf(w, r.Host+"\n")
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
