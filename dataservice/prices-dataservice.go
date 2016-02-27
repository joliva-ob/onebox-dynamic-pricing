package dataservice

import (

	"database/sql"
	"time"
	"strconv"

	"github.com/patrickmn/go-cache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"

)


// DB Price struct
type PriceType struct{

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



/**
 *
 * Public function to retrieve needed price details for handle with dynamic
 * pricing processes
 *
 * http://go-database-sql.org/accessing.html
 */
func GetPrices(date_from string, date_to string, page int, config configuration.Config) []*PriceType {

	var prices []*PriceType
	start := time.Now()
	offset := config.Mysql_limit_items * page
	key := date_from + date_to + strconv.Itoa(config.Mysql_limit_items) + strconv.Itoa(offset)
	var rows *sql.Rows
	var err error

	// Get the string associated with the key from the cache
	pricesFromCache, found := pricesCache.Get(key)
	if !found {

		// Retrieve from DB
		rows, err = db.Query(config.Prices_sql, date_from, date_to, config.Mysql_limit_items, offset);
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Read all values from resultset and map it to vector of Pricetype struct
		for rows.Next() {

			p := new(PriceType)
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
			elapsed := time.Since(start)
			log.Debugf("%v price rows retrieved in %v", len(prices), elapsed)
		}

		// Store the prices struct to cache for 5 minutes as default
		pricesCache.Set(key, prices, cache.DefaultExpiration)

	} else {

		// Retrieve prices struct from cache
		prices = pricesFromCache.([]*PriceType) // Cast interface{} retrieved from cache to []*PriceType
	}

	// Reuse db connections pool rather than Close database connection
	// defer db.Close()

	return prices
}

