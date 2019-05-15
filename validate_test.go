package devicecheck

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ValidateDeviceToken_InvalidKey(t *testing.T) {
	client := &Client{
		api:  newAPI(Development),
		cred: NewCredentialFile("unknown_file.p8"),
		jwt:  newJWT("issuer", "keyID"),
	}

	err := client.ValidateDeviceToken("device_token")

	assert.NotNil(t, err)
}

func TestClient_ValidateDeviceToken_InvalidURL(t *testing.T) {
	client := &Client{
		api: api{
			client:  new(http.Client),
			baseURL: "invalid url",
		},
		cred: NewCredentialFile("revoked_private_key.p8"),
		jwt:  newJWT("issuer", "keyID"),
	}

	err := client.ValidateDeviceToken("device_token")

	assert.NotNil(t, err)
}

func TestClient_ValidateDeviceToken_InvalidDeviceToken(t *testing.T) {
	client := &Client{
		api:  newAPI(Development),
		cred: NewCredentialFile("revoked_private_key.p8"),
		jwt:  newJWT("issuer", "keyID"),
	}

	err := client.ValidateDeviceToken("device_token")

	assert.NotNil(t, err)
}
