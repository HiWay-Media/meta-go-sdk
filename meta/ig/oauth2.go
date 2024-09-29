package ig

import (
	"golang.org/x/oauth2"
)

// Endpoint is IG's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.instagram.com/oauth/authorize",
	TokenURL: "https://api.instagram.com/oauth/access_token",
}


func (s *igService) CodeAuthUrl(state string, opts ...oauth2.AuthCodeOption) string {
	return s.OAuth2Config.AuthCodeURL(state, opts...)
}

func (s *igService) AuthUrl() string {
	return s.OAuth2Config.AuthCodeURL(
		"",
		oauth2.SetAuthURLParam(
			"enable_fb_login",
			"0",
		),
		oauth2.SetAuthURLParam(
			"force_authentication",
			"1",
		),
		oauth2.SetAuthURLParam(
			"scope",
			"business_basic,business_content_publish,business_manage_comments,business_manage_messages",
		),
	)
}