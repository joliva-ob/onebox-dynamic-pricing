package controller

import (

	"net/http"
	"net"

	"github.com/op/go-logging"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
)


const (
	START_DATE = "start_date"
	END_DATE = "end_date"
	PAGE = "page"
	DATE_FORMAT_SHORT = "2006-01-02"
	EVENT_ID = "event_id"
	TRANSACTION_SALE_TYPE = "SALE"
	PRICE_ID = "price_id"
	SALE_ID = "sale_id"
)


// Parameters response struct
type ParametersResponseType struct {

	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	EventId int `json:event_id`
	Page int `json:"page"`
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
	// and takes precedence over RemoteAddr
	// Header.Get is case-insensitive
	forward := req.Header.Get("X-Forwarded-For")
	return forward
}