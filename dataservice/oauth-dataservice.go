package dataservice


type Oauthtoken struct {
	ClientId string `json:"clientId"`
	Token    string `json:"token"`
	UserName string `json:"userName"`
}


func GetOauthtoken( token string ) *Oauthtoken  {

	var oauthtoken Oauthtoken
	log.Debugf("cbBucket is %v: ", cbBucket)
	if cbBucket != nil {
		err := cbBucket.Get(token, &oauthtoken)
		if err != nil {
			log.Fatalf("Failed to get data from the couchbase cluster (%s)\n", err)
		}
	}

	return &oauthtoken
}


