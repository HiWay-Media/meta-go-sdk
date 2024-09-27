package meta 

// Base interface
type IMeta interface {

}

type meta struct {
    restClient   *resty.Client
	debug        bool
	clientKey    string
	clientSecret string
	accessToken  string
	OAuth2Config *oauth2.Config
}