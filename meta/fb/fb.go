package fb

import (
	"github.com/HiWay-Media/meta-go-sdk/meta"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
)

type IFacebook interface {
	meta.IMeta
	UploadMedia(fileUrl, pageId, title, description string) (*UploadMediaResponse, error)
}

type fbService struct {
	restClient   *resty.Client
	debug        bool
	clientKey    string
	clientSecret string
	accessToken  string
	OAuth2Config *oauth2.Config
}

func NewFacebook(clientKey, clientSecret string, debug bool) (IFacebook, error) {
	s := &fbService{
		restClient:   resty.New(),
		debug:        debug,
		clientKey:    clientKey,
		clientSecret: clientSecret,
	}
	s.restClient.SetDebug(debug)
	//s.restClient.SetBaseURL(BASE_URL)
	return s, nil
}

func (s *fbService) SetAccessToken(token string) {
	s.accessToken = token
}

func (s *fbService) GetAccessToken() string {
	return s.accessToken
}
