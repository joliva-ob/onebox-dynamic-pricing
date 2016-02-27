package dataservice


import (

	"database/sql"
	"time"

	"github.com/patrickmn/go-cache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
	"github.com/op/go-logging"
)


const (
	MYSQL_DRIVER_NAME = "mysql"
)



// Global vars
var db *sql.DB
var isInitialized bool = false
var log *logging.Logger = configuration.GetLog()
var pricesCache *cache.Cache
var sessionCache *cache.Cache



// Initialize pool database and set properties from config
func Initialize( config configuration.Config ){

	if !isInitialized {

		// Open database connection pool
		db, _ = sql.Open(MYSQL_DRIVER_NAME, config.Mysql_conn)
		db.SetMaxOpenConns(config.Mysql_max_conn)
		db, _ = sql.Open(MYSQL_DRIVER_NAME, config.Mysql_conn)
		// Create a cache with a default expiration time of N seconds, and which
		// purges expired items every 30 seconds
		pricesCache = cache.New( time.Duration(config.Cache_expiration_time_sec*1000*1000*1000), 30*time.Second ) // Duration constructor needs nanoseconds
		sessionCache = cache.New( time.Duration(config.Cache_expiration_time_sec*1000*1000*1000), 30*time.Second ) // Duration constructor needs nanoseconds

		isInitialized = true

		log.Debugf("prices dataservice initialized with a max pool of: %v", config.Mysql_max_conn)
	}
}
