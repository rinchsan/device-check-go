package devicecheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	issuer := "issuer"
	keyID := "keyID"

	cfg := NewConfig(issuer, keyID, Development)

	assert := assert.New(t)
	assert.Equal(issuer, cfg.issuer)
	assert.Equal(keyID, cfg.keyID)
	assert.Equal(Development, cfg.env)

	cfg = NewConfig(issuer, keyID, Production)

	assert.Equal(issuer, cfg.issuer)
	assert.Equal(keyID, cfg.keyID)
	assert.Equal(Production, cfg.env)
}
