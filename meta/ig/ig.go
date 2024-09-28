package ig

import (
	"github.com/HiWay-Media/meta-go-sdk/meta"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
)

type IInstagram interface {
	meta.IMeta
}

type igService struct {
	restClient   *resty.Client
	debug        bool
	clientKey    string
	clientSecret string
	accessToken  string
	OAuth2Config *oauth2.Config
}
