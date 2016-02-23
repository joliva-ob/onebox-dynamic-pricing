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
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)


// Instance configuration
type Config struct {
	Server_port string
	Prices_sql string
	Mysql_conn   string
	Mysql_max_conn int
	Log_file string
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

