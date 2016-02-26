package controller

import (

	"github.com/op/go-logging"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
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