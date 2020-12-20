package devicecheck

import (
	"net/http"
	"testing"
)

func TestClient_ValidateDeviceToken(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		client Client
	}{
		"invalid key": {
			client: Client{
				api:  newAPI(Development),
				cred: NewCredentialFile("unknown_file.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
		},
		"invalid url": {
			client: Client{
				api: api{
					client:  new(http.Client),
					baseURL: "invalid url",
				},
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
		},
		"invalid device token": {
			client: Client{
				api:  newAPI(Development),
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := c.client.ValidateDeviceToken("device_token")

			if err == nil {
				t.Error("want 'not nil', got 'nil'")
			}
		})
	}
}
