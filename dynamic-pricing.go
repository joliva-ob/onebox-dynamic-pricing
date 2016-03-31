package main

import (

	"net/http"
	"os"
	"fmt"


	"github.com/op/go-logging"
	"github.com/gorilla/mux"
	"github.com/hudl/fargo"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
	"github.com/joliva-ob/onebox-dynamic-pricing/controller"
)



// Global vars
var config configuration.Config
var log *logging.Logger
var eurekaConn fargo.EurekaConnection


/**
 * Main command to load configuration by given environment argument
 * and start application server to listen the exposed endpoints and
 * provide the requested resources operations
 *
 * Mandatory parameters are path (/tmp...) and environment (dev, qa, pre, pro...)
 */
func main() {

	// Load configuration to start application
	checkParams( os.Args )
	var filename = os.Args[1] + "/" + os.Args[2] + ".yml"
	config = configuration.LoadConfiguration(filename)
	log = configuration.GetLog()
	dataservice.Initialize(config)
	log.Infof("dynamic-pricing started with environment: %s and listening in port: %v\n", os.Args[2], config.Server_port)

	// Register to Eureka and then set up to only heartbeat one of them
	filename = os.Args[1] + "/eureka_" + os.Args[2] + ".gcfg"
	registerToEureka( filename )

	// Create the router to handle requests
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/dynamic-pricing-api/1.0/prices", controller.PricesController) // Prices endpoint
	router.HandleFunc("/dynamic-pricing-api/1.0/sales", controller.SalesController) // Sales endpoint
	router.HandleFunc("/dynamic-pricing-api/1.0/info", controller.InfoController)
	router.HandleFunc("/dynamic-pricing-api/1.0/health", controller.HealthController)

	// Starting server on given port number
	log.Fatal( http.ListenAndServe(":" + config.Server_port, router) ) // Start the server at listening port

}




// Check the arguments to launch the application
// and provide specifications if needed.
func  checkParams(  args []string ) {

	if len(args) < 2 {

		fmt.Println("ERROR: invalid arguments number!")
		fmt.Println("Usage:")
		fmt.Println("./dynamic-pricing [path-to-config-files] [environment] [log-level]")
		os.Exit(0)
	}

}


// Register and keep the eureka connection
func registerToEureka( configFile string )  {

	eurekaConn, _ = fargo.NewConnFromConfigFile(configFile)
	hostname, _ := os.Hostname()
	i := fargo.Instance{
		HostName:         hostname,
		Port:             8000,
		App:              config.Eureka_app_name,
		IPAddr:           hostname,
		VipAddress:       config.Eureka_app_name,
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.Amazon},
		SecureVipAddress: config.Eureka_ip_addr,
		Status:           fargo.UP,
		HealthCheckUrl:	  "http://" +hostname+ ":" +config.Server_port+ "/dynamic-pricing-api/1.0/health",
		StatusPageUrl:	  "http://" +hostname+ ":" +config.Server_port+ "/dynamic-pricing-api/1.0/health",
		HomePageUrl:      "http://" +hostname+ ":" +config.Server_port+ "/dynamic-pricing-api/1.0/health",
	}
	err := eurekaConn.RegisterInstance(&i)
	if err != nil {
		log.Error("%v", err)
	}
}
