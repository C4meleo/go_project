package main

import (
	"fmt"
	"log"
	"net/http"
)

var listAttractions []Attraction

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: UI pour tester l'API")
	fmt.Println("GET /")
}

func initServer() {
	fmt.Print("GoRestAPI v1.0\n" +
		"Listining on 0.0.0.0:8000\n")

	http.HandleFunc("/", root)
	http.HandleFunc("/attraction", handleAttractions)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func main() {
	listAttractions = initAttractions()
	initServer()
}
