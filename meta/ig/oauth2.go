package ig

import (
	"golang.org/x/oauth2"
)

// Endpoint is IG's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.instagram.com/oauth/authorize",
	TokenURL: "https://api.instagram.com/oauth/access_token",
}
