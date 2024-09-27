package fb_test

import (
	"testing"

	"github.com/HiWay-Media/meta-go-sdk/meta/fb"
)

func TestFacebook(t *testing.T) {
	t.Log("TestFacebook")
	fbService := fb.NewFacebook("clientKey", "clientSecret", true)
	if fbService == nil {
		t.Error("Expected fbService to be not nil")
	}
	t.Log("TestFacebook done", fbService)
}
