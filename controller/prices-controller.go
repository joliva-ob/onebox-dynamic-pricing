package controller


import (

	"net/http"
	"encoding/json"
	"time"
	"strconv"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"
	"github.com/joliva-ob/onebox-dynamic-pricing/authorization"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
)


// Prices response struct
type PricesResponseType struct {

	Version string `json:"version"`
	RequestDate time.Time `json:"request_date"`
	Parameters ParametersResponseType `json:"parameters"`
	Prices []*dataservice.PriceType `json:"prices"`
}



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

	// GET request params
	startDate = request.URL.Query().Get(START_DATE)
	if startDate ==  "" {
		startDate = time.Now().AddDate(0, -1, 0).Format(DATE_FORMAT_SHORT)
	}
	endDate = request.URL.Query().Get(END_DATE)
	if endDate == "" {
		endDate = time.Now().Format(DATE_FORMAT_SHORT)
	}
	page, err := strconv.Atoi(request.URL.Query().Get(PAGE))
	if err != nil {
		page = 0
	}

	// Retrieve requested resource information
	prices := dataservice.GetPrices(startDate, endDate, page, configuration.GetConfig())

	// Set json response struct
	var params ParametersResponseType
	params.StartDate = startDate
	params.EndDate = endDate
	params.Page = page
	var pricesresponse PricesResponseType
	pricesresponse.Parameters = params
	pricesresponse.RequestDate = time.Now()
	pricesresponse.Version = "1.0"
	pricesresponse.Prices = prices

	pricesjson, err := json.Marshal(pricesresponse)
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
