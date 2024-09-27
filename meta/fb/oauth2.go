package fb

import (
	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/v20.0/dialog/oauth",
	TokenURL: "https://graph.facebook.com/v20.0/oauth/access_token",
}
