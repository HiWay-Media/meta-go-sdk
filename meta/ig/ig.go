package ig

import (
	"github.com/HiWay-Media/meta-go-sdk/meta"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
)

type IInstagram interface {
	meta.IMeta
	UploadMedia(igUserId string, videoUrl, imageUrl *string, caption string ) (*MediaPublishResponse, error)
	MediaPublish(igUserId, creationId string) (*MediaPublishResponse, error)
	CheckStatus(containerId string) (*CheckMediaStatusResponse, error)
}

type igService struct {
	restClient   *resty.Client
	debug        bool
	clientKey    string
	clientSecret string
	accessToken  string
	OAuth2Config *oauth2.Config
}


func NewInstagram(clientKey, clientSecret string, debug bool) IInstagram {
	s := &igService{
		restClient:   resty.New(),
		debug:        debug,
		clientKey:    clientKey,
		clientSecret: clientSecret,
	}
	return s
}

func (s *igService) SetAccessToken(token string) {
	s.accessToken = token
}

func (s *igService) GetAccessToken() string {
	return s.accessToken
}
