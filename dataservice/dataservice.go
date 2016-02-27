package dataservice


import (

	"database/sql"
	"time"
	"flag"

	"github.com/patrickmn/go-cache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
	"github.com/op/go-logging"
	"github.com/mattbaird/elastigo/lib"
)


const (
	MYSQL_DRIVER_NAME = "mysql"
)



// Global vars
var (
	db *sql.DB
	isInitialized bool = false
	log *logging.Logger = configuration.GetLog()
	pricesCache *cache.Cache
	sessionsCache *cache.Cache
	salesCache *cache.Cache
	elk_conn *elastigo.Conn
)



// Initialize pool database and set properties from config
func Initialize( config configuration.Config ){

	if !isInitialized {

		// Open database connection pool
		db, _ = sql.Open(MYSQL_DRIVER_NAME, config.Mysql_conn)
		db.SetMaxOpenConns(config.Mysql_max_conn)
		db, _ = sql.Open(MYSQL_DRIVER_NAME, config.Mysql_conn)
		log.Debugf("DB dataservice initialized to: %v with a max pool of: %v", config.Mysql_conn, config.Mysql_max_conn)

		// Open elasticsearch connection
		elk_host := flag.String(config.Elasticsearch_name, config.Elasticsearch_value, config.Elasticsearch_usage)
		elk_conn = elastigo.NewConn()
		flag.Parse()
		elk_conn.Domain = *elk_host
		log.Debugf("Elasticsearch connected to host: %v", config.Elasticsearch_value)

		// Create a cache with a default expiration time of N seconds, and which
		// purges expired items every 30 seconds
		pricesCache = cache.New( time.Duration(config.Cache_prices_expiration_in_sec*1000*1000*1000), 30*time.Second ) // Duration constructor needs nanoseconds
		log.Debugf("Prices Cache initialized with eviction time: %v sec", config.Cache_prices_expiration_in_sec)
		sessionsCache = cache.New( time.Duration(config.Cache_sessions_expiration_in_sec*1000*1000*1000), 30*time.Second ) // Duration constructor needs nanoseconds
		log.Debugf("Sessions Cache initialized with eviction time: %v sec", config.Cache_sessions_expiration_in_sec)
		salesCache = cache.New( time.Duration(config.Cache_sales_expiration_in_sec*1000*1000*1000), 30*time.Second ) // Duration constructor needs nanoseconds
		log.Debugf("Sales Cache initialized with eviction time: %v sec", config.Cache_sales_expiration_in_sec)

		isInitialized = true


	}
}
