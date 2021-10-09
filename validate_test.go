package devicecheck

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestClient_ValidateDeviceToken(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		client Client
		noErr  bool
	}{
		"invalid key": {
			client: Client{
				api:  newAPI(Development),
				cred: NewCredentialFile("unknown_file.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
			noErr: false,
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
			noErr: false,
		},
		"invalid device token": {
			client: Client{
				api:  newAPI(Development),
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
			noErr: false,
		},
		"status ok": {
			client: Client{
				api: newAPIWithHTTPClient(newMockHTTPClient(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("success")),
				}), Development),
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
			noErr: true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := c.client.ValidateDeviceToken(context.Background(), "device_token")

			if c.noErr {
				if err != nil {
					t.Errorf("want 'nil', got '%+v'", err)
				}
			} else {
				if err == nil {
					t.Error("want 'not nil', got 'nil'")
				}
			}
		})
	}
}
