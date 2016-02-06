package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	log.Println("--> dynamic-pricing started.")

	// Create the router to handle mockup requests with its response properly
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index) // General welcome endpoint

	// Starting server on given port number
	port := ":8000"
	log.Fatal( http.ListenAndServe(port, router) ) // Start the server at listening port

}



// Welcome endpoint
func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome to Onebox dynamic pricing, the dynamic pricing API for smart ticketing revenue.")

}
