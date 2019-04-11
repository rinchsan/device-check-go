package devicecheck

import (
	"crypto/ecdsa"
	"io/ioutil"

	"github.com/dvsekhvalnov/jose2go/keys/ecc"
)

// Credential provides credential for DeviceCheck API
type Credential interface {
	key() (*ecdsa.PrivateKey, error)
}

// CredentialFile provides credential from private key file
type CredentialFile struct {
	filename string
}

// NewCredentialFile returns credential from private key file
func NewCredentialFile(filename string) CredentialFile {
	return CredentialFile{
		filename: filename,
	}
}

func (cred CredentialFile) key() (*ecdsa.PrivateKey, error) {
	raw, err := ioutil.ReadFile(cred.filename)
	if err != nil {
		return nil, err
	}
	return ecc.ReadPrivate(raw)
}

// CredentialBytes provides credential from private key bytes
type CredentialBytes struct {
	raw []byte
}

// NewCredentialBytes returns credential from private key bytes
func NewCredentialBytes(raw []byte) CredentialBytes {
	return CredentialBytes{
		raw: raw,
	}
}

func (cred CredentialBytes) key() (*ecdsa.PrivateKey, error) {
	return ecc.ReadPrivate(cred.raw)
}

// CredentialString provides credential from private key string
type CredentialString struct {
	str string
}

// NewCredentialString returns credential from private key string
func NewCredentialString(str string) CredentialString {
	return CredentialString{
		str: str,
	}
}

func (cred CredentialString) key() (*ecdsa.PrivateKey, error) {
	return ecc.ReadPrivate([]byte(cred.str))
}
