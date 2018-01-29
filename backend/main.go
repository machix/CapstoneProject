package main
 
import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler)
	router.HandleFunc("/position", getPosition).Methods("GET")

	corsRouter := cors.Default().Handler(router);
	http.ListenAndServe(":8000", corsRouter)
}

type Message struct {
	Text string
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %!", r.URL.Path[1:])
}

func getPosition(w http.ResponseWriter, r *http.Request) {
	m := Message{"Soon you will get some really cool info herer! It will be very cool!"}
    b, err := json.Marshal(m)
 
    errorCheck(err)
 
    w.Write(b)
}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}