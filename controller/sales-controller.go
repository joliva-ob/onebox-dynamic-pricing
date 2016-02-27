package controller


import (

	"net/http"
	"time"
	"strconv"
	"encoding/json"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"
	"github.com/joliva-ob/onebox-dynamic-pricing/authorization"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
)


// Prices response struct
type SalesResponseType struct {

	Version string `json:"version"`
	RequestDate time.Time `json:"request_date"`
	Parameters ParametersResponseType `json:"parameters"`
	Sales []*dataservice.SaleElkType `json:"sales"`
}



/**
 * Prices resource endpoint
 */
func SalesController(w http.ResponseWriter, request *http.Request) {

	log.Infof( "/sales request received." )
	start := time.Now()

	// Check authorization
	if !authorization.Authorize( request.Header.Get("Authorization") ) {
		w.WriteHeader(http.StatusUnauthorized)
		log.Errorf("/sales status 401 error unauthorized.")
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
	sales := dataservice.GetSales(startDate, endDate, page, configuration.GetConfig())

	// Set json response struct
	var params ParametersResponseType
	params.StartDate = startDate
	params.EndDate = endDate
	params.Page = page
	var salesresponse SalesResponseType
	salesresponse.Parameters = params
	salesresponse.RequestDate = time.Now()
	salesresponse.Version = "1.0"
	salesresponse.Sales = sales
	salesjson, err := json.Marshal(salesresponse)

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Set response body
	w.Write(salesjson)

	elapsed := time.Since(start)
	log.Infof( "/sales status 200 response in %v", elapsed )

}
