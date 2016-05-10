package dataservice


import (

	"encoding/json"

)



/*
type Restrictions struct {
	Events []int `json:"events"`
	Sessions    []int `json:"sessions"`
	Venues []int `json:"venues"`
	Channels []int `json:"channels"`
}
*/

type Restrictions struct {
	SalesElkFilterEventRestricted struct {
					      Query struct {
							    Filtered struct {
									     Filter struct {
											    Bool struct {
													 Must []struct {
														 Term struct {
															      Doc_status_type string `json:"doc.status.type"`
														      } `json:"term"`
													 } `json:"must"`
												 } `json:"bool"`
										    } `json:"filter"`
								     } `json:"filtered"`
						    } `json:"query"`
					      Sort struct {
							    Doc_date_purchased struct {
										       Order string `json:"order"`
									       } `json:"doc.date.purchased"`
						    } `json:"sort"`
				      } `json:"sales_elk_filter_event_restricted"`
	SalesElkFilterSaleRestricted struct {
					      Query struct {
							    Filtered struct {
									     Filter struct {
											    Bool struct {
													 Must []struct {
														 Term struct {
															      Doc_code string `json:"doc.code"`
														      } `json:"term"`
													 } `json:"must"`
												 } `json:"bool"`
										    } `json:"filter"`
								     } `json:"filtered"`
						    } `json:"query"`
					      Sort struct {
							    Doc_date_purchased struct {
										       Order string `json:"order"`
									       } `json:"doc.date.purchased"`
						    } `json:"sort"`
				      } `json:"sales_elk_filter_sale_restricted"`
}



func GetRestrictions( username string, isForced bool ) (*Restrictions, bool)  {

	var restrictions *Restrictions

	// Get the string associated with the key from the cache
	restrictionsFromCache, found := restrictionsCache.Get(username)
	if !found || isForced {

		key := "dynamic-pricing-restrictions_" + username
		err := cbRestrictionsBucket.Get(key, &restrictions)
		if err != nil {

			log.Errorf("Failed to get data from the couchbase cluster for user %v %s\n", username, err)
			found = false

		} else {

			// Get string query from couchbase document
			rb, err := json.Marshal(restrictions.SalesElkFilterEventRestricted)
			if err == nil {
				log.Debugf("sales query restricted: %v\n", string(rb))
				config.Sales_elk_filter_event_restricted = string(rb)
			}
			rb, err = json.Marshal(restrictions.SalesElkFilterSaleRestricted)
			if err == nil {
				config.Sales_elk_filter_sale_restricted = string(rb)
			}

			log.Infof("Loaded restrictions from the couchbase cluster for user %v\n", username)
			found = true

		}

		// Store the prices struct to cache, no expires
		restrictionsCache.Set(username, restrictions, -1)

	} else {

		// Retrieve prices struct from cache
		restrictions = restrictionsFromCache.(*Restrictions) // Cast interface{} retrieved from cache to []*PriceType

	}

	return restrictions, found
}



