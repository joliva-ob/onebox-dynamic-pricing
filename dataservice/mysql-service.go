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



/**
 * Public function to retrieve needed price details for handle with dynamic
 * pricing processes
 *
 * http://go-database-sql.org/accessing.html
 */
func GetPrices(date_from string, date_to string, limit int, config configuration.Config) []*Pricetype {

	var prices []*Pricetype

	db, err := sql.Open("mysql", config.Mysql_conn)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("db connection successful.")
	}

	rows, err := db.Query(config.Prices_sql, date_from, date_to, limit);
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Read all values from resultset and map it to vector of Pricetype
	for rows.Next() {

		p := new(Pricetype)
		err := rows.Scan(&p.Id, &p.Price_zone_id, &p.Price, &p.Price_zone_name, &p.Event_id, &p.Event_name, &p.Event_date, &p.Session_id, &p.Session_date, &p.Venue_id, &p.Venue_name, &p.Buyer_type_code, &p.Fee, &p.Tax, &p.External_price_id)
		if err != nil {
			log.Fatal(err)
		}
		prices = append(prices, p)
		// Log to test the results
		//log.Println(id, price_zone_id, price, price_zone_name, event_id, event_name, event_date, session_id, session_date, venue_id, venue_name, buyer_type_code, fee, tax, external_price_id)
		//log.Printf("row: %v", p)
	}
	//log.Printf("row: %v", prices)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("--> mysql-service: %v price rows retrieved.\n", len(prices))
	}

	defer db.Close()

	return prices

}

