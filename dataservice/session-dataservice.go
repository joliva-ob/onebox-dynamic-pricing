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
type SessionType struct{

	Session_id int `json:"id"`
	Session_name int `json:"session_name"`
	Session_date string `json:"session_date"`
	Event_id int `json:"event_id"`
	Event_name string `json:"event_name"`
	Venue_id int `json:"venue_id"`
	Venue_name string `json:"venue_name"`
}



/**
 *
 * Public function to retrieve needed session details related to its venue and event
 *
 * http://go-database-sql.org/accessing.html
 */
func GetSession(sessionId int, config configuration.Config) *SessionType {

	session := new(SessionType)
	start := time.Now()
	key := strconv.Itoa(sessionId)
	var rows *sql.Rows
	var err error

	// Get the string associated with the key from the cache
	sessionFromCache, found := sessionsCache.Get(key)
	if !found {

		// Retrieve from DB
		rows, err = db.Query(config.Session_sql, sessionId);
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Read all values from resultset and map it to vector of Pricetype struct
		for rows.Next() {

			err := rows.Scan(&session.Session_id, &session.Session_name, &session.Session_date, &session.Event_id, &session.Event_name, &session.Venue_id, &session.Venue_name)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		} else {
			elapsed := time.Since(start)
			log.Debugf("%v id session retrieved in %v", session.Session_id, elapsed)
		}

		// Store the prices struct to cache for 5 minutes as default
		sessionsCache.Set(key, session, cache.DefaultExpiration)

	} else {

		// Retrieve prices struct from cache
		session = sessionFromCache.(*SessionType) // Cast interface{} retrieved from cache to *SessionType
	}

	// Reuse db connections pool rather than Close database connection
	// defer db.Close()

	return session
}

