package authorization


import (

	"strings"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"

)




// Public method to check the Oauth2 authorization with
// a Bearer token header to the oauth server
func Authorize (authheader string) *dataservice.Oauthtoken {

	var oauthtoken *dataservice.Oauthtoken
	s := strings.Split(authheader, " ")
	token := s[len(s)-1]

	if token != "" {
		oauthtoken = dataservice.GetOauthtoken( token )
	}

	return oauthtoken
}
