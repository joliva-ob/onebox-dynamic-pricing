package dataservice


type Restrictions struct {
	Events []int `json:"events"`
	Sessions    []int `json:"sessions"`
	Venues []int `json:"venues"`
	Channels []int `json:"channels"`
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
			log.Infof("Load restrictions from the couchbase cluster for user %v\n", username)
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



