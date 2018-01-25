package main
 
import (
    "encoding/json"
    "fmt"
    "net/http"
)
 
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %!", r.URL.Path[1:])
}
 
func main() {
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
}
