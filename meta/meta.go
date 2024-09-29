package meta

import "golang.org/x/oauth2"

// Base interface
type IMeta interface {
	//
	HealthCheck() error
	IsDebug() bool
	CodeAuthUrl(state string, opts ...oauth2.AuthCodeOption) string
	SetAccessToken(token string)
	GetAccessToken() string
	//
}
