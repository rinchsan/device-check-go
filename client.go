package devicecheck

import "net/http"

type Client struct {
	api  api
	cred Credential
	jwt  jwt
}

func New(cred Credential, cfg Config) Client {
	return Client{
		api:  newAPI(cfg.env),
		cred: cred,
		jwt:  newJWT(cfg.issuer, cfg.keyID),
	}
}

func NewWithHTTPClient(httpClient *http.Client, cred Credential, cfg Config) Client {
	return Client{
		api:  newAPIWithHTTPClient(httpClient, cfg.env),
		cred: cred,
		jwt:  newJWT(cfg.issuer, cfg.keyID),
	}
}
