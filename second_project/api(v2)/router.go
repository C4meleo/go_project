package main

import (
    "fmt"
    "log"
    "net/http"
)

var listAttractions []Attraction

// Display the index page
func root(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "TODO: UI pour tester l'API")
}

// Initialisation of the server with the root web, and the API on port 8000 in localhost
func initServer() {
    fmt.Print("GoRestAPI v1.0\n" +
        "Listining on 0.0.0.0:8000\n")

    http.HandleFunc("/", root)
    http.HandleFunc("/attraction", handleAttractions)
    log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil)) // Get log
}

func main() {
    listAttractions = initAttractions()
    initServer()
}