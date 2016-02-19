package dataservice

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
)


// Prices struct
type Pricetype struct{
	Id int `json:"id"`
	Price_zone_id int `json:"price_zone_id"`
	Price float32 `json:"price"`
	Price_zone_name string `json:"price_zone_name"`
	Event_id int `json:"event_id"`
	Event_name string `json:"event_name"`
	Event_date string `json:"event_date"`
	Session_id int `json:"session_id"`
	Session_date string `json:"session_date"`
	Venue_id int `json:"venue_id"`
	Venue_name string `json:"venue_name"`
	Buyer_type_code string `json:"buyer_type_code"`
	Fee float32 `json:"fee"`
	Tax float32 `json:"tax"`
	External_price_id []byte `json:"external_price_id"`
}

// Global vars
var db *sql.DB
var isInitialized bool = false


// Initialize pool database and set properties from config
func Initialize( config configuration.Config ){

	if !isInitialized {

		// Open database connection pool
		db, _ = sql.Open("mysql", config.Mysql_conn)
		db.SetMaxOpenConns(config.Mysql_max_conn)
		db, _ = sql.Open("mysql", config.Mysql_conn)

		isInitialized = true

		log.Printf("--> db prices initialized with a max pool of: %v", config.Mysql_max_conn)
	}
}


/**
 * Public function to retrieve needed price details for handle with dynamic
 * pricing processes
 *
 * http://go-database-sql.org/accessing.html
 */
func GetPrices(date_from string, date_to string, limit int, config configuration.Config) []*Pricetype {

	var prices []*Pricetype

	// Query
	rows, err := db.Query(config.Prices_sql, date_from, date_to, limit);
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Read all values from resultset and map it to vector of Pricetype struct
	for rows.Next() {

		p := new(Pricetype)
		err := rows.Scan(&p.Id, &p.Price_zone_id, &p.Price, &p.Price_zone_name, &p.Event_id, &p.Event_name, &p.Event_date, &p.Session_id, &p.Session_date, &p.Venue_id, &p.Venue_name, &p.Buyer_type_code, &p.Fee, &p.Tax, &p.External_price_id)
		if err != nil {
			log.Fatal(err)
		}
		prices = append(prices, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("--> prices-service: %v price rows retrieved.\n", len(prices))
	}

	// Reuse db connections pool rather than Close database connection
	// defer db.Close()

	return prices
}

