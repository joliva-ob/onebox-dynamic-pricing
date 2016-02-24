package main

import (

	"net/http"
	"os"

	"github.com/op/go-logging"
	"github.com/gorilla/mux"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
	"github.com/joliva-ob/onebox-dynamic-pricing/controller"
)



// Global vars
var config configuration.Config
var log *logging.Logger


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
	log = configuration.GetLog()
	dataservice.Initialize(config)
	log.Infof("dynamic-pricing started with environment: %s and listening in port: %v\n", os.Args[2], config.Server_port)

	// Create the router to handle requests
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/prices", controller.PricesController) // General welcome endpoint

	// Starting server on given port number
	log.Fatal( http.ListenAndServe(":" + config.Server_port, router) ) // Start the server at listening port

}



