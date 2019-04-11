package devicecheck

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime_MarshalJSON(t *testing.T) {
	year := 2019
	month := time.April
	t_ := Time{time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)}

	b, err := t_.MarshalJSON()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(fmt.Sprintf(`"%04d-%02d"`, year, month), string(b))
}

func TestTime_UnmarshalJSON(t *testing.T) {
	year := 2019
	month := time.April
	b, err := Time{time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)}.MarshalJSON()
	assert.Nil(t, err)

	t_ := Time{}
	err = t_.UnmarshalJSON(b)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(year, t_.Year())
	assert.Equal(month, t_.Month())
}

func TestClient_QueryTwoBits_InvalidKey(t *testing.T) {
	client := Client{
		api:  newAPI(Development),
		cred: NewCredentialFile("unknown_file.p8"),
		jwt:  newJWT("issuer", "keyID"),
	}

	result := QueryTwoBitsResult{}
	err := client.QueryTwoBits("device_token", &result)

	assert.NotNil(t, err)
}

func TestClient_QueryTwoBits_InvalidURL(t *testing.T) {
	client := Client{
		api: api{
			client:  new(http.Client),
			baseURL: "invalid url",
		},
		cred: NewCredentialFile("revoked_private_key.p8"),
		jwt:  newJWT("issuer", "keyID"),
	}

	result := QueryTwoBitsResult{}
	err := client.QueryTwoBits("device_token", &result)

	assert.NotNil(t, err)
}

func TestClient_QueryTwoBits_InvalidDeviceToken(t *testing.T) {
	client := Client{
		api:  newAPI(Development),
		cred: NewCredentialFile("revoked_private_key.p8"),
		jwt:  newJWT("issuer", "keyID"),
	}

	result := QueryTwoBitsResult{}
	err := client.QueryTwoBits("device_token", &result)

	assert.NotNil(t, err)
}
