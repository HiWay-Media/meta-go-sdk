package fb

import (
	"github.com/HiWay-Media/meta-go-sdk/meta"
)

type IFacebook interface {
}

type fbService struct {
	*meta.Meta
}

func NewFacebook(clientKey, clientSecret string, debug bool) IFacebook {
	return &fbService{
		Meta: meta.NewMeta(clientKey, clientSecret, debug),
	}
}
