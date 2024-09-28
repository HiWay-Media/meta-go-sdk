package meta

// Base interface
type IMeta interface {
	//
	HealthCheck() error
	IsDebug() bool
	CodeAuthUrl() string
	SetAccessToken(token string)
	GetAccessToken() string
	//
}
