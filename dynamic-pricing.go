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
 *
 * Mandatory parameters are path (/tmp...) and environment (dev, qa, pre, pro...)
 */
func main() {

	// Load configuration to start application
	var filename = os.Args[1] + "/" + os.Args[2] + ".yml"
	config = configuration.LoadConfiguration(filename)
	log.Printf("--> dynamic-pricing started with environment: %s\n", os.Args[2])

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

	// Retrieve requested resource information
	prices := dataservice.GetPrices(date_from, date_to, limit, config)
	pricesjson, err := json.Marshal(prices)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Response string
	fmt.Fprintln(w, string(pricesjson))

	ms = ms - (time.Now().UnixNano()%1e6/1e3)
	log.Printf( "prices response in %v ms.", ms )

}


