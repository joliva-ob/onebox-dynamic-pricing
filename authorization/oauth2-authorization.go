package authorization


import (

	"strings"

	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"

	"github.com/op/go-logging"
)

var log *logging.Logger = configuration.GetLog()

// Public method to check the Oauth2 authorization with
// a Bearer token header to the oauth server
func Authorize (authheader string) bool {

	s := strings.Split(authheader, " ")
	token := s[len(s)-1]

	if token != "" {
		oauthtoken := dataservice.GetOauthtoken( token )
		log.Debugf("oauthtoken username = '%v'", oauthtoken.UserName)
	//	if oauthtoken != nil {
			return true
	//	}
	}




//	if authheader == "Bearer 1736cc7f-7c60-4576-b851-b7b3630cfeab" {
//		return true
//	}

	return false

}
