package controller

import (

	"net/http"
	"net"

	"github.com/op/go-logging"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
	"github.com/satori/go.uuid"
)


const (
	START_DATE = "start_date"
	END_DATE = "end_date"
	PAGE = "page"
	DATE_FORMAT_SHORT = "2006-01-02"
	EVENT_ID = "event_id"
	TRANSACTION_SALE_TYPE = "SALE"
	PRICE_ID = "price_id"
	ID = "id"
	SALE_ID = "sale_id" // OB order code
	STATUS_UP = "UP"
	STATUS_DOWN = "DOWN"
	STATUS_OK = "OK"
	STATUS_ERROR = "ERROR"
	AUTH_HEADER = "Authorization"
)


// Parameters response struct
type ParametersResponseType struct {

	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	EventId int `json:event_id`
	SaleId string `json:"sale_id"`
	Page int `json:"page"`
	TraceId string `json:"trace_id"`
}


// Global vars and default values
var log *logging.Logger = configuration.GetLog()
var startDate string
var endDate string
var eventId int
var saleId string
var priceId int



// https://blog.golang.org/context/userip/userip.go
// Funtion to retrieve the sender IP from request
// or from forwared headers instead
func getIP(w http.ResponseWriter, req *http.Request) string {

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Debugf("userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return req.RemoteAddr
	}

	// This will only be defined when site is accessed via non-anonymous proxy
	// and takes precedence over RemoteAddr Header.Get is case-insensitive
	forward := req.Header.Get("X-Forwarded-For")
	return forward
}



// Generate a universal unique identifier UUID
func GetUuid() string {

	// Creating UUID Version 4
	uuid1 := uuid.NewV4()

	return uuid1.String()
}
