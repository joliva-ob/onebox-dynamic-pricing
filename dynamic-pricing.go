package main

import (

	"log"
	"net/http"
	"os"
	"encoding/json"
	"time"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
	"github.com/gorilla/mux"
	"github.com/joliva-ob/onebox-dynamic-pricing/authorization"
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
	dataservice.Initialize(config)
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
func pricesController(w http.ResponseWriter, request *http.Request) {

	log.Printf( "/prices request received." )
	ms := time.Now().UnixNano()%1e6/1e3

	// Check authorization
	if !authorization.Authorize( request.Header.Get("Authorization") ) {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("/prices status 401 error unauthorized.")
		return
	}

	// Retrieve requested resource information
	prices := dataservice.GetPrices(date_from, date_to, limit, config)
	pricesjson, err := json.Marshal(prices)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
		log.Println("/prices status 204 error no content.")
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Set response body
	w.Write(pricesjson)

	ms = (time.Now().UnixNano()%1e6/1e3) - ms
	log.Printf( "/prices status 200 response in %v ms.", ms )

}


