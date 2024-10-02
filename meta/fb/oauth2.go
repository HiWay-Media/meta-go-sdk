package fb

import (
	"context"
	"golang.org/x/oauth2"
)

// Endpoint is Facebook's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/v20.0/dialog/oauth",
	TokenURL: "https://graph.facebook.com/v20.0/oauth/access_token",
}

func (s *fbService) CodeAuthUrl(state string, opts ...oauth2.AuthCodeOption) string {
	return s.OAuth2Config.AuthCodeURL(state, opts...)
}

func (s *fbService) ExchangeCode(code string)  (*oauth2.Token, error) {
	return s.OAuth2Config.Exchange(context.Background(), code)
}