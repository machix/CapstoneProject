package main
 
import (
<<<<<<< HEAD:main.go
    "encoding/json"
    "fmt"
    "net/http"
=======
	"fmt"
	"log"
	"os"
	"net/http"
>>>>>>> 465f3ad16b1061a4c7c758b531c99dbae76433c3:backend/main.go
)
 
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %!", r.URL.Path[1:])
}
 
func main() {
<<<<<<< HEAD:main.go
    http.HandleFunc("/", handler)
    http.HandleFunc("/getLocation/", about)
    http.ListenAndServe(":8080", nil)
}
 
type Message struct {
    Text string
}
 
func about (w http.ResponseWriter, r *http.Request) {
    m := Message{"We will return the latitude longitude here eventually"}
    b, err := json.Marshal(m)
 
    if err != nil {
        panic(err)
    }
 
     w.Write(b)
=======
	http.HandleFunc("/", handler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
>>>>>>> 465f3ad16b1061a4c7c758b531c99dbae76433c3:backend/main.go
}
