package configuration

import (
	"io/ioutil"
	"bufio"
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
	Session_sql string
	Mysql_conn   string
	Mysql_max_conn int
	Mysql_limit_items int
	Log_file string
	Log_format string
	Cache_prices_expiration_in_sec int
	Cache_sessions_expiration_in_sec int
	Cache_sales_expiration_in_sec int
	Cache_oauth_expiration_in_sec int
	Elasticsearch_name string
	Elasticsearch_value string
	Elasticsearch_usage string
	Elasticsearch_limit_items int
	Sales_elk_index string
	Sales_elk_filter_dates string
	Sales_elk_filter_dates_restricted string
	Sales_elk_filter_event string
	Sales_elk_filter_event_restricted string
	Sales_elk_filter_sale string
	Sales_elk_filter_sale_restricted string
	Prices_sql string
	Prices_sql_dates_restricted string
	Prices_sql_filter_price_id string
	Prices_sql_filter_price_id_restricted string
	Prices_sql_filter_event_id string
	Prices_sql_filter_event_id_restricted string
	Prices_sql_filter_event_id_price_id string
	Prices_sql_filter_event_id_price_id_restricted string
	Eureka_port int
	Eureka_ip_addr string
	Eureka_app_name string
	Couchbase_url string
	Couchbase_oauth_bucket string
	Couchbase_restrictions_bucket string
	Couchbase_pool string
	Authorized_usernames_list []string
	Restricted_usernames_list []string
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
	printBootLogo()
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
func GetConfig() Config {
	return config
}



// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}



func printBootLogo() {

	lines, _ := readLines("configuration/boot_logo.txt")
	for _, line := range lines {
		fmt.Println(line)
	}

}