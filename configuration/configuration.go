package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	"os"
)


// Global vars
var logfile os.File
var config Config


// Instance configuration
type Config struct {
	Server_port string
	Prices_sql string
	Mysql_conn   string
	Mysql_max_conn int
	Log_file string
	Log_file_enabled bool
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
	log.Printf("--> Configuration loaded values: %#v\n", config)

	// Set logger
	if config.Log_file_enabled {

			f, err := os.OpenFile(config.Log_file, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			defer f.Close()
			log.Printf("error opening file: %v", err)
		}
		log.SetOutput(f)
	}

	return config
}


/**
 * Test configuration file
 */
func _(){

	filename := "../resources/dev.yml"
	var config Config

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	log.Printf("Value: %#v\n", config)
}