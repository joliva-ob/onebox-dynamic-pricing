package controller

import (

	"net/http"
	"encoding/json"
	"time"

	"github.com/op/go-logging"
	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"
	"github.com/joliva-ob/onebox-dynamic-pricing/authorization"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
)


// Global vars
var date_from string = "2015-01-01"
var date_to string = "2015-02-01"
var limit int = 10
var log *logging.Logger = configuration.GetLog()



/**
 * Prices resource endpoint
 */
func PricesController(w http.ResponseWriter, request *http.Request) {

	log.Infof( "/prices request received." )
	start := time.Now()

	// Check authorization
	if !authorization.Authorize( request.Header.Get("Authorization") ) {
		w.WriteHeader(http.StatusUnauthorized)
		log.Errorf("/prices status 401 error unauthorized.")
		return
	}

	// Retrieve requested resource information
	prices := dataservice.GetPrices(date_from, date_to, limit, configuration.GetConfig())
	pricesjson, err := json.Marshal(prices)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		log.Errorf("/prices status 204 error no content.")
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Set response body
	w.Write(pricesjson)

	elapsed := time.Since(start)
	log.Infof( "/prices status 200 response in %v", elapsed )

}
