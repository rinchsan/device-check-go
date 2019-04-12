package devicecheck

import "net/http"

// Client provides methods to use DeviceCheck API
type Client interface {
	// QueryTwoBits queries two bits for device token
	QueryTwoBits(deviceToken string, result *QueryTwoBitsResult) error
	// UpdateTwoBits updates two bits for device token
	UpdateTwoBits(deviceToken string, bit0, bit1 bool) error
	// ValidateDeviceToken validates a device for device token
	ValidateDeviceToken(deviceToken string) error
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
