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
	restrictionsFromCache, hasRestrictions := restrictionsCache.Get(username)
	if !hasRestrictions || isForced {

/*		key := "dynamic-pricing-restrictions_" + username
		err := cbRestrictionsBucket.Get(key, &restrictions)
		if err != nil {
			log.Errorf("Failed to get data from the couchbase cluster for user %v %s\n", username, err)
			found = false
		} else {
			log.Infof("Load restrictions from the couchbase cluster for user %v\n", username)
			found = true
		}
*/
		// Check restricted users from config file
		for _, restricted_user := range config.Restricted_usernames_list {

			if username == restricted_user {
				hasRestrictions = true
				break
			}
		}


		// Store the struct to cache, no expires
		restrictionsCache.Set(username, restrictions, -1)

	} else {

		// Retrieve  struct from cache
		restrictions = restrictionsFromCache.(*Restrictions) // Cast interface{} retrieved from cache to []*Type
	}

	return restrictions, hasRestrictions
}



