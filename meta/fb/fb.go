package fb

import (
	"github.com/HiWay-Media/meta-go-sdk/meta"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
)

type IFacebook interface {
	meta.IMeta
}

type fbService struct {
	restClient   *resty.Client
	debug        bool
	clientKey    string
	clientSecret string
	accessToken  string
	OAuth2Config *oauth2.Config
}

func NewFacebook(clientKey, clientSecret string, debug bool) IFacebook {
	return &fbService{
		restClient:   resty.New(),
		debug:        debug,
		clientKey:    clientKey,
		clientSecret: clientSecret,
	}
}
