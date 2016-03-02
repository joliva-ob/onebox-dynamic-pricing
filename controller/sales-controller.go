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


// Sales response struct
type SalesResponseType struct {

	Version string `json:"version"`
	RequestDate time.Time `json:"request_date"`
	Parameters ParametersResponseType `json:"parameters"`
	Sales []*SaleResponseType `json:"sales"`
}


// Sale struct response type
type SaleResponseType struct {

	Id string `json:"id"`
	EventId int `json:"event_id"`
	EventDate string `json:"event_date"`
	EventName string `json:"event_name"`
	TransactionDate string `json:"transaction_date"`
	TransactionType string `json:"transaction_type"`
	BuyerTypeCode string `json:"buyer_type_code"`
	ProductsNumber int `json:"products_number"`
	ChannelId int `json:"channel_id"`
	Products []*ProductResponseType `json:"products"`
}

// Sale struct product response type
type ProductResponseType struct {

	Id string `json:"id"`
	SessionId int `json:"session_id"`
	SessionDate string `json:"session_date"`
	SessionName string `json:"session_name"`
	VenueId int `json:"venue_id"`
	VenueName string `json:"venue_name"`
	PriceId int `json:"price_id"`
	PriceZoneId int `json:"price_zone_id"`
	Price float32 `json:"price"`
	Section string `json:"section"`
	Seat string `json:"seat"`
	Row int `json:"row"`
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
	eventId, err = strconv.Atoi(request.URL.Query().Get(EVENT_ID))
	if err != nil {
		eventId = -1
	}


	// Retrieve requested resource information
	dbSales := dataservice.GetSales(startDate, endDate, eventId, page)

	// Set json response struct
	var params ParametersResponseType
	params.StartDate = startDate
	params.EndDate = endDate
	params.Page = page
	var salesresponse SalesResponseType
	salesresponse.Parameters = params
	salesresponse.RequestDate = time.Now()
	salesresponse.Version = "1.0"
	sales := transformDbSalesToSalesResponse( dbSales )
	salesresponse.Sales = sales
	salesjson, err := json.Marshal(salesresponse)

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Set response body
	w.Write(salesjson)

	elapsed := time.Since(start)
	log.Infof( "/sales status 200 response in %v", elapsed )

}



// Transform sales DB struct to an outgoing json
// struct
func transformDbSalesToSalesResponse( ordersDb []*dataservice.OrderDocElkType ) []*SaleResponseType {

	var sales []*SaleResponseType

	for i:=0; i<len(ordersDb)-1; i++ {

//		log.Debugf("orderDB locator: %v\n", ordersDb[i].Doc.Code)
		// Get the cached config details
		session := dataservice.GetSession( ordersDb[i].Doc.Products[0].SessionId, configuration.GetConfig() )

		sale := new(SaleResponseType)
		sale.Id = ordersDb[i].Doc.Code
		sale.EventId = session.Event_id
		sale.EventName = session.Event_name
		sale.TransactionDate = ordersDb[i].Doc.Date.Purchased
		sale.TransactionType = TRANSACTION_SALE_TYPE
		sale.BuyerTypeCode = ordersDb[i].Doc.OrderData.ChannelType
		sale.ProductsNumber = len(ordersDb[i].Doc.Products)
		sale.ChannelId = ordersDb[i].Doc.OrderData.ChannelId
		products := transformProductsDbToProductsResponse( ordersDb[i].Doc.Products, session )
		sale.Products = products

		sales = append(sales, sale)
	}

	return sales
}



// Transform a productDb struct to
// a json response product struct
func transformProductsDbToProductsResponse( productsDb []*dataservice.ProductElkType, session *dataservice.SessionType ) []*ProductResponseType {

	var products []*ProductResponseType

	for i:=0; i<len(productsDb); i++ {

		product := new(ProductResponseType)
		product.Id = productsDb[i].ProductId
		product.SessionId = productsDb[i].SessionId
		product.SessionDate = session.Session_date
		product.SessionName = session.Session_name
		product.VenueId = session.Venue_id
		product.VenueName = session.Venue_name
		product.PriceId = productsDb[i].TicketData.PriceZoneId
		product.PriceZoneId = productsDb[i].TicketData.PriceZoneId
		product.Price = session.Price
		product.Section = productsDb[i].TicketData.SectorName
		product.Seat = productsDb[i].TicketData.NumSeat
		product.Row = productsDb[i].TicketData.RowOrder

		products = append(products, product)
	}

	return products
}