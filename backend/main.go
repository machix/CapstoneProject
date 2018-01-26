package main
 
import (
    "net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/position", getPosition).Methods("GET") //in position.go

	corsRouter := cors.Default().Handler(router);
	http.ListenAndServe(":8000", corsRouter)

}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}