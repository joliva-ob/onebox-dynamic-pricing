package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Instance configuration
type Config struct {
	Server_port string
	Prices_sql string
	Mysql_conn   string
}


/**
 * Load configuration yaml file
 */
func LoadConfiguration(filename string) Config {

	var config Config

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	log.Printf("--> Configuration loaded values: %#v\n", config)

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