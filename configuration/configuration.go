package configuration

import (
	"io/ioutil"
	"os"
	"fmt"

	"gopkg.in/yaml.v2"
	"github.com/op/go-logging"
)


// Global vars
var config Config
var log = logging.MustGetLogger("dynamic-pricing")


// Instance configuration
type Config struct {
	Server_port string
	Prices_sql string
	Session_sql string
	Mysql_conn   string
	Mysql_max_conn int
	Mysql_limit_items int
	Log_file string
	Log_format string
	Cache_prices_expiration_in_sec int
	Cache_sessions_expiration_in_sec int
	Cache_sales_expiration_in_sec int
	Elasticsearch_name string
	Elasticsearch_value string
	Elasticsearch_usage string
	Elasticsearch_limit_items int
	Sales_elk_index string
	Sales_elk_filter_eventId string
	Sales_elk_filter_dates string
	Sales_elk_filter_eventId_dates string
}



/**
 * Load configuration yaml file
 */
func LoadConfiguration(filename string) Config {

	// Set config
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--> Configuration loaded values: %#v\n", config)

	// Set logger
	format := logging.MustStringFormatter( config.Log_format )
	logbackend1 := logging.NewLogBackend(os.Stdout, "", 0)
	logbackend1Formatted := logging.NewBackendFormatter(logbackend1, format)
	f, err := os.OpenFile(config.Log_file, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		defer f.Close()
	}
	logbackend2 := logging.NewLogBackend(f, "", 0)
	logbackend2Formatted := logging.NewBackendFormatter(logbackend2, format)
	logging.SetBackend(logbackend1Formatted, logbackend2Formatted)

	return config
}


// Return the already configured logger
func GetLog() *logging.Logger{
	return log
}


// Return the already loaded configuration
func GetConfig() Config{
	return config
}

