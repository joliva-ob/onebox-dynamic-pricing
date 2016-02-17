package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"time"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
	"github.com/gorilla/mux"
)



// Global vars
var date_from string = "2015-01-01"
var date_to string = "2015-02-01"
var limit int = 10
var config configuration.Config

/**
 * Main command to load configuration by given environment argument
 * and start application server to listen the exposed endpoints and
 * provide the requested resources operations
 */
func main() {

	// Load configuration to start application
	var env = "./resources/" + os.Args[1] + ".yml"
	config = configuration.LoadConfiguration(env)
	log.Printf("--> dynamic-pricing started with environment: %s\n", os.Args[1])

	// Create the router to handle mockup requests with its response properly
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/prices", pricesController) // General welcome endpoint

	// Starting server on given port number
	log.Fatal( http.ListenAndServe(":" + config.Server_port, router) ) // Start the server at listening port

}



/**
 * Prices resource endpoint
 */
func pricesController(w http.ResponseWriter, r *http.Request) {

	ms := time.Now().UnixNano()%1e6/1e3

	prices := dataservice.GetPrices(date_from, date_to, limit, config)
	pricesjson, err := json.Marshal(prices)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Response string
	fmt.Fprintln(w, string(pricesjson))

	ms = ms - (time.Now().UnixNano()%1e6/1e3)
	log.Printf( "prices response in %v ms.", ms )

}


