package fb

import (
	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/v20.0/dialog/oauth",
	TokenURL: "https://graph.facebook.com/v20.0/oauth/access_token",
}

func (s *fbService) CodeAuthUrl(state string, opts ...oauth2.AuthCodeOption) string {
	return s.OAuth2Config.AuthCodeURL(state, opts...)
}
