package controller

import (

	"github.com/op/go-logging"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
)


const (
	START_DATE = "start_date"
	END_DATE = "end_date"
	PAGE = "page"
	DATE_FORMAT_SHORT = "2006-01-02"
)


// Parameters response struct
type ParametersResponseType struct {

	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Page int `json:"page"`
}


// Global vars and default values
var log *logging.Logger = configuration.GetLog()
var startDate string
var endDate string