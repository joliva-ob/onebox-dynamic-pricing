package dataservice


type Oauthtoken struct {
	ClientId string `json:"clientId"`
	Token    string `json:"token"`
	UserName string `json:"userName"`
}


func GetOauthtoken( token string ) *Oauthtoken  {

	oauthtoken := new(Oauthtoken)

	// Get the string associated with the key from the cache
	oauthtokenFromCache, found := oauthCache.Get(token)
	if !found {

//		log.Debugf("oauthtoken not found %v ", token)
		err := cbOauthBucket.Get(token, &oauthtoken)
		if err != nil {
			log.Errorf("Failed to get data from the couchbase cluster (%s)\n", err)
		}

		// Store the struct to cache
		oauthCache.Set(token, oauthtoken, 0)

	} else {

		// Retrieve struct from cache
//		log.Debugf("oauthtoken found %v", token)
		oauthtoken = oauthtokenFromCache.(*Oauthtoken) // Cast interface{} retrieved from cache
	}

	return oauthtoken
}


