package controller


import (

	"net/http"
	"time"
	"encoding/json"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"
	"github.com/joliva-ob/onebox-dynamic-pricing/authorization"
)


// Info response struct
type RestrictionsResponseType struct {

	Status string `json:"status"`
	Description string `json:"description"`
}



/**
 * Restrictions resource endpoint
 */
func ReloadRestrictionsController(w http.ResponseWriter, request *http.Request) {

	uuid := GetUuid()
	start := time.Now()

	// Check authorization
	oauthtoken := authorization.Authorize( request.Header.Get(AUTH_HEADER) )
	if oauthtoken.Token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		log.Warningf("/reload-restrictions request error status 401 unauthorized %v", getIP(w, request))
		return
	}
	log.Infof( "{%v} /reload-restrictions request %v received from: %v %v", uuid, request.URL, oauthtoken.UserName, getIP(w, request) )


	// Force reload restrictions
	_, result := dataservice.GetRestrictions( oauthtoken.UserName, true )

	// Set json response struct
	var restrictionsresponse RestrictionsResponseType
	if result {
		restrictionsresponse.Status = STATUS_OK
		restrictionsresponse.Description = "Restrictions reloaded for username: " + oauthtoken.UserName
	} else {
		restrictionsresponse.Status = STATUS_ERROR
		restrictionsresponse.Description = "No restrictions found to reload for username: " + oauthtoken.UserName
	}
	infojson, _ := json.Marshal(restrictionsresponse)

	// Set response headers and body
	w.Header().Set("Content-Type", "application/json")
	w.Write(infojson)

	elapsed := time.Since(start)
	log.Infof( "{%v} /reload-restrictions response status 200 in %v", uuid, elapsed )

}
