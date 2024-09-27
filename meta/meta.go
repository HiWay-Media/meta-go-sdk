package meta

import (
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
)

// Base interface
/*type IMeta interface {
	//
	HealthCheck() error
	IsDebug() bool
	CodeAuthUrl() string
	SetAccessToken(token string)
	GetAccessToken() string
	//
}*/

// Base Meta struct 
type Meta struct {
	restClient   *resty.Client
	debug        bool
	clientKey    string
	clientSecret string
	accessToken  string
	OAuth2Config *oauth2.Config
}

/*func NewMeta(clientKey, clientSecret string, debug bool) IMeta {
	return &meta{
		restClient:   resty.New(),
		debug:        debug,
		clientKey:    clientKey,
		clientSecret: clientSecret,
	}
}*/
