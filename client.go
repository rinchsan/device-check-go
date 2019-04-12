package devicecheck

import "net/http"

// Client provides methods to use DeviceCheck API
type Client interface {
	QueryTwoBits(string, *QueryTwoBitsResult) error
	UpdateTwoBits(string, bool, bool) error
	ValidateDeviceToken(string) error
}

type clientImpl struct {
	api  api
	cred Credential
	jwt  jwt
}

// New returns a new DeviceCheck API client instance
func New(cred Credential, cfg Config) Client {
	return clientImpl{
		api:  newAPI(cfg.env),
		cred: cred,
		jwt:  newJWT(cfg.issuer, cfg.keyID),
	}
}

// NewWithHTTPClient returns a new DeviceCheck API client instance with specified http client
func NewWithHTTPClient(httpClient *http.Client, cred Credential, cfg Config) Client {
	return clientImpl{
		api:  newAPIWithHTTPClient(httpClient, cfg.env),
		cred: cred,
		jwt:  newJWT(cfg.issuer, cfg.keyID),
	}
}
