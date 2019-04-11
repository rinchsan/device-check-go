package devicecheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT_newJWT(t *testing.T) {
	issuer := "issuer"
	keyID := "keyID"

	j := newJWT(issuer, keyID)

	assert := assert.New(t)
	assert.Equal(issuer, j.issuer)
	assert.Equal(keyID, j.keyID)
}

func TestJWT_generate(t *testing.T) {
	cred := NewCredentialFile("revoked_private_key.p8")
	key, err := cred.Key()
	assert.Nil(t, err)
	assert.NotNil(t, key)
	issuer := "issuer"
	keyID := "keyID"
	j := newJWT(issuer, keyID)

	token, err := j.generate(key)

	assert := assert.New(t)
	assert.NotEmpty(token)
	assert.Nil(err)
}
