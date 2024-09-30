package test

import (
	"os"
	"testing"

	"github.com/HiWay-Media/meta-go-sdk/meta/fb"
	"github.com/HiWay-Media/meta-go-sdk/meta/ig"
)

func TestMain(m *testing.M) {
	if os.Getenv("APP_ENV") == "" {
		err := os.Setenv("APP_ENV", "test")
		if err != nil {
			panic("could not set test env")
		}
	}
	//env.Load()
	m.Run()
}

func GetIG() (ig.IInstagram, error) {
    //
	clientKey       := os.Getenv("CLIENT_KEY")
	clientSecret    := os.Getenv("CLIENT_SECRET")
	//
	c, err := ig.NewInstagram(clientKey, clientSecret, false)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func GetFB() (fb.IFacebook, error) {
    //
	clientKey       := os.Getenv("CLIENT_KEY")
	clientSecret    := os.Getenv("CLIENT_SECRET")
	//
	c, err := fb.NewFacebook(clientKey, clientSecret, false)
	if err != nil {
		return nil, err
	}
	return c, nil
}

