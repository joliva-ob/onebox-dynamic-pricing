package controller


import (

	"net/http"
	"encoding/json"
	"time"
	"strconv"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"
	"github.com/joliva-ob/onebox-dynamic-pricing/authorization"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"

	"github.com/gorilla/mux"
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

	uuid := GetUuid()
	log.Infof( "{%v} /prices request %v received from: %v", uuid, request.URL, getIP(w, request) )
	start := time.Now()

	// Check authorization
	if !authorization.Authorize( request.Header.Get("Authorization") ) {
		w.WriteHeader(http.StatusUnauthorized)
		log.Warningf("/prices error status 401 unauthorized.")
		return
	}

	// GET request params
	priceId, err := strconv.Atoi(request.URL.Query().Get(PRICE_ID)) // Return 0 if error
	if err != nil {

		vars := mux.Vars(request)
		id := vars[ID]
		if id != "" {
			priceId, err = strconv.Atoi(id) // Return 0 if error
			if err != nil {
				priceId = 0
			}
		}
	}
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
	eventId, err = strconv.Atoi(request.URL.Query().Get(EVENT_ID)) // Return 0 if error


	// Retrieve requested resource information
	prices := dataservice.GetPrices(startDate, endDate, page, configuration.GetConfig(), priceId, eventId, uuid)


	// Set json response struct
	var params ParametersResponseType
	params.StartDate = startDate
	params.EndDate = endDate
	params.Page = page
	params.TraceId = uuid
	var pricesresponse PricesResponseType
	pricesresponse.Parameters = params
	pricesresponse.RequestDate = time.Now()
	pricesresponse.Version = "1.0"
	pricesresponse.Prices = prices

	pricesjson, err := json.Marshal(pricesresponse)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		log.Errorf("/prices error status 204 no content.")
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Set response body
	w.Write(pricesjson)

	elapsed := time.Since(start)
	log.Infof( "{%v} /prices response status 200 in %v", uuid, elapsed )

}
