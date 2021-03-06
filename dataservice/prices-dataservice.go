package dataservice

import (

	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
)


// DB Price struct
type PriceType struct{

	Id int `json:"id"`
//	Price_zone_id int `json:"id"`
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
func GetPrices(date_from string, date_to string, page int, config configuration.Config, priceId int, eventId int, uuid string, oauthtoken *Oauthtoken, pageSize int) []*PriceType {

	var prices []*PriceType
	offset := pageSize * page
	key := date_from + date_to + strconv.Itoa(pageSize) + strconv.Itoa(offset) + strconv.Itoa(priceId) + strconv.Itoa(eventId) + oauthtoken.UserName


	// Get the string associated with the key from the cache
	_, hasRestrictions := GetRestrictions( oauthtoken.UserName, false )
	pricesFromCache, found := pricesCache.Get(key)
	if !found {
//		log.Debugf("--> Not found in cache")
		// Retrieve from DB
		rows, err := GetDataFromDb( eventId, priceId, pageSize, offset, date_from, date_to, hasRestrictions)
		if err != nil {
			log.Error(err)
		}
		defer rows.Close()
//		log.Debugf("rows: %v", rows)

		// Read all values from resultset and map it into vector of Pricetype struct
		for rows.Next() {

//			log.Debugf("rows.Next(): true")
			p := new(PriceType)
			err := rows.Scan(&p.Id, &p.Price, &p.Price_zone_name, &p.Event_id, &p.Event_name, &p.Event_date, &p.Session_id, &p.Session_date, &p.Venue_id, &p.Venue_name, &p.Buyer_type_code, &p.Fee, &p.Tax, &p.External_price_id)
			if err != nil {
				log.Error(err)
			}
			prices = append(prices, p)

		}
		err = rows.Err()
		if err != nil {
			log.Error(err)
		}

		// Store the prices struct to cache
		pricesCache.Set(key, prices, 0)

	} else {
//		log.Debugf("--> Retrieving from CACHE")
		// Retrieve prices struct from cache
		prices = pricesFromCache.([]*PriceType) // Cast interface{} retrieved from cache to []*PriceType
	}

	// Reuse db connections pool rather than Close database connection
	// defer db.Close()

	return prices
}



/**
 Retrieve prices from DB taking into account if an event or price are given
 or not.
 */
func GetDataFromDb( eventId int, priceId int, limit_items int, offset int, date_from string, date_to string, restrictions bool) (*sql.Rows, error) {
//	log.Debugf("--> into GetDataFromDb")
//	log.Debugf("--> GetDataFromDb parameters:\n eventId %v\n priceId %v\n limit_items %v\n offset %v\n date_from %v\n date_to %v\n restrictions %v", eventId, priceId, limit_items, offset, date_from, date_to, restrictions)
	if priceId > 0 && eventId > 0 {

		if restrictions {
//			log.Debugf("1")
			return db.Query(config.Prices_sql_filter_event_id_price_id_restricted, priceId, limit_items, offset);
		} else {
//			log.Debugf("2")
			return db.Query(config.Prices_sql_filter_event_id_price_id, eventId, priceId, limit_items, offset);
		}

	} else if priceId > 0 {

		if restrictions {
//			log.Debugf("3")
			return db.Query(config.Prices_sql_filter_price_id_restricted, priceId, limit_items, offset);
		} else {
//			log.Debugf("4")
			return db.Query(config.Prices_sql_filter_price_id, priceId, limit_items, offset);
		}

	} else if eventId > 0 {

		if restrictions {
//			log.Debugf("5")
			return db.Query(config.Prices_sql_filter_event_id_restricted, limit_items, offset);
		} else {
//			log.Debugf("6")
			return db.Query(config.Prices_sql_filter_event_id, eventId, limit_items, offset);
		}

	} else { // Filter by dates

		if restrictions {
//			log.Debugf("7")
			return db.Query(config.Prices_sql_dates_restricted, date_from, date_to, limit_items, offset);
		} else {
//			log.Debugf("8")
			return db.Query(config.Prices_sql, date_from, date_to, limit_items, offset);
		}
	}

}

