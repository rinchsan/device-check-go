package devicecheck

import (
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/dvsekhvalnov/jose2go/keys/ecc"
)

// Credential provides credential for DeviceCheck API.
type Credential interface {
	key() (*ecdsa.PrivateKey, error)
}

type credentialFile struct {
	filename string
}

// NewCredentialFile returns credential from private key file.
func NewCredentialFile(filename string) Credential {
	return credentialFile{
		filename: filename,
	}
}

func (cred credentialFile) key() (*ecdsa.PrivateKey, error) {
	raw, err := os.ReadFile(cred.filename)
	if err != nil {
		return nil, fmt.Errorf("os: %w", err)
	}

	key, err := ecc.ReadPrivate(raw)
	if err != nil {
		return nil, fmt.Errorf("ecc: %w", err)
	}

	return key, nil
}

type credentialBytes struct {
	raw []byte
}

// NewCredentialBytes returns credential from private key bytes.
func NewCredentialBytes(raw []byte) Credential {
	return credentialBytes{
		raw: raw,
	}
}

func (cred credentialBytes) key() (*ecdsa.PrivateKey, error) {
	key, err := ecc.ReadPrivate(cred.raw)
	if err != nil {
		return nil, fmt.Errorf("ecc: %w", err)
	}

	return key, nil
}

type credentialString struct {
	str string
}

// NewCredentialString returns credential from private key string.
func NewCredentialString(str string) Credential {
	return credentialString{
		str: str,
	}
}

func (cred credentialString) key() (*ecdsa.PrivateKey, error) {
	key, err := ecc.ReadPrivate([]byte(cred.str))
	if err != nil {
		return nil, fmt.Errorf("ecc: %w", err)
	}

	return key, nil
}
