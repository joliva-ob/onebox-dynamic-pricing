package authorization


import (

	"strings"

	"github.com/joliva-ob/onebox-dynamic-pricing/dataservice"

	"github.com/joliva-ob/onebox-dynamic-pricing/configuration"
)




// Public method to check the Oauth2 authorization with
// a Bearer token header to the oauth server
func Authorize (authheader string) (*dataservice.Oauthtoken, bool) {

	var oauthtoken *dataservice.Oauthtoken
	isAuthorized := false
	s := strings.Split(authheader, " ")
	token := s[len(s)-1]

	if token != "" {

		oauthtoken = dataservice.GetOauthtoken( token )
		for _, username := range configuration.GetConfig().Authorized_usernames_list {

			if username == oauthtoken.UserName {
				isAuthorized = true
				break
			}
		}
	}

	return oauthtoken, isAuthorized
}
