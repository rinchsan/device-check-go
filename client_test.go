package devicecheck

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cred := NewCredentialFile("revoked_private_key.p8")
	issuer := "issuer"
	keyID := "keyID"
	env := Development
	cfg := NewConfig(issuer, keyID, env)

	client := New(cred, cfg).(clientImpl)

	assert := assert.New(t)
	assert.Equal(cred, client.cred)
	assert.Equal(issuer, client.jwt.issuer)
	assert.Equal(keyID, client.jwt.keyID)
	assert.Equal(newBaseURL(env), client.api.baseURL)
	assert.NotNil(client.api.client)
}

func TestNewWithHTTPClient(t *testing.T) {
	cred := NewCredentialFile("revoked_private_key.p8")
	issuer := "issuer"
	keyID := "keyID"
	env := Production
	cfg := NewConfig(issuer, keyID, env)
	httpClient := new(http.Client)

	client := NewWithHTTPClient(httpClient, cred, cfg).(clientImpl)

	assert := assert.New(t)
	assert.Equal(cred, client.cred)
	assert.Equal(issuer, client.jwt.issuer)
	assert.Equal(keyID, client.jwt.keyID)
	assert.Equal(newBaseURL(env), client.api.baseURL)
	assert.NotNil(client.api.client)
	assert.Equal(httpClient, client.api.client)
}
