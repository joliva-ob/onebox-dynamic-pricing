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
	"github.com/couchbaselabs/go-couchbase"

)


const (
	MYSQL_DRIVER_NAME = "mysql"
	SALE_ID = "$saleId"
	EVENT_ID = "$eventId"
	START_DATE = "$startDate"
	END_DATE = "$endDate"
)



// Global vars
var (
	db *sql.DB
	isInitialized bool = false
	log *logging.Logger = configuration.GetLog()
	pricesCache *cache.Cache
	sessionsCache *cache.Cache
	salesCache *cache.Cache
	oauthCache *cache.Cache
	restrictionsCache *cache.Cache
	elk_conn *elastigo.Conn
	cbOauthBucket *couchbase.Bucket
	cbRestrictionsBucket *couchbase.Bucket
	config configuration.Config
)



/**
  * Initialize pool database and set properties from config
  */
func Initialize( c configuration.Config ){

	// Set configuration
	config = c

	if !isInitialized {

		var err error

		// Open database connection pool
		db, err = sql.Open(MYSQL_DRIVER_NAME, config.Mysql_conn)
		db.SetMaxOpenConns(config.Mysql_max_conn)
		if err != nil {
			log.Fatalf("Failed to connect to mysql database: %v\n", err)
		}
		log.Infof("DB dataservice initialized to: %v with a max pool of: %v", config.Mysql_conn, config.Mysql_max_conn)


		// Open elasticsearch connection
		elk_host := flag.String(config.Elasticsearch_name, config.Elasticsearch_value, config.Elasticsearch_usage)
		elk_conn = elastigo.NewConn()
		flag.Parse()
		elk_conn.Domain = *elk_host
		log.Infof("Elasticsearch connected to host: %v", config.Elasticsearch_value)


		// Open couchbase connection
		connection, err := couchbase.Connect(config.Couchbase_url)
		pool, err := connection.GetPool(config.Couchbase_pool)
		cbOauthBucket, err = pool.GetBucket(config.Couchbase_oauth_bucket)
		cbRestrictionsBucket, err = pool.GetBucket(config.Couchbase_restrictions_bucket)
		if err != nil {
			log.Errorf("Failed to get bucket from couchbase (%s)\n", err)
		} else {
			log.Infof("Couchbase connected to host: %v and bucket: %v", config.Couchbase_url, config.Couchbase_oauth_bucket)
			log.Infof("Couchbase connected to host: %v and bucket: %v", config.Couchbase_url, config.Couchbase_restrictions_bucket)
		}


		// Create a cache with a default expiration time of N seconds, and which
		// purges expired items every 60 seconds
		pricesCache = cache.New( time.Duration(config.Cache_prices_expiration_in_sec*1000*1000*1000), 30*time.Second ) // Duration constructor needs nanoseconds
		log.Infof("Prices Cache initialized with eviction time: %v sec", config.Cache_prices_expiration_in_sec)
		sessionsCache = cache.New( time.Duration(config.Cache_sessions_expiration_in_sec*1000*1000*1000), 30*time.Second ) // Duration constructor needs nanoseconds
		log.Infof("Sessions Cache initialized with eviction time: %v sec", config.Cache_sessions_expiration_in_sec)
		salesCache = cache.New( time.Duration(config.Cache_sales_expiration_in_sec*1000*1000*1000), 30*time.Second ) // Duration constructor needs nanoseconds
		log.Infof("Sales Cache initialized with eviction time: %v sec", config.Cache_sales_expiration_in_sec)
		oauthCache = cache.New( time.Duration(config.Cache_oauth_expiration_in_sec*1000*1000*1000), 30*time.Second ) // Duration constructor needs nanoseconds
		log.Infof("Oauth Cache initialized with eviction time: %v sec", config.Cache_oauth_expiration_in_sec)
		restrictionsCache = cache.New( -1, 30*time.Second ) // No auto expiration
		log.Infof("Restrictions Cache initialized with eviction time: %v sec", -1)

		isInitialized = true


	}
}






