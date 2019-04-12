package devicecheck

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_UpdateTwoBits_InvalidKey(t *testing.T) {
	client := clientImpl{
		api:  newAPI(Development),
		cred: NewCredentialFile("unknown_file.p8"),
		jwt:  newJWT("issuer", "keyID"),
	}

	err := client.UpdateTwoBits("device_token", true, true)

	assert.NotNil(t, err)
}

func TestClient_UpdateTwoBits_InvalidURL(t *testing.T) {
	client := clientImpl{
		api: api{
			client:  new(http.Client),
			baseURL: "invalid url",
		},
		cred: NewCredentialFile("revoked_private_key.p8"),
		jwt:  newJWT("issuer", "keyID"),
	}

	err := client.UpdateTwoBits("device_token", true, true)

	assert.NotNil(t, err)
}

func TestClient_UpdateTwoBits_InvalidDeviceToken(t *testing.T) {
	client := clientImpl{
		api:  newAPI(Development),
		cred: NewCredentialFile("revoked_private_key.p8"),
		jwt:  newJWT("issuer", "keyID"),
	}

	err := client.UpdateTwoBits("device_token", true, true)

	assert.NotNil(t, err)
}
