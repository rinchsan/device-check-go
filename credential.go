package devicecheck

import (
	"crypto/ecdsa"
	"io/ioutil"

	"github.com/dvsekhvalnov/jose2go/keys/ecc"
)

type Credential interface {
	Key() (*ecdsa.PrivateKey, error)
}

type CredentialFile struct {
	filename string
}

func NewCredentialFile(filename string) CredentialFile {
	return CredentialFile{
		filename: filename,
	}
}

func (cred CredentialFile) Key() (*ecdsa.PrivateKey, error) {
	raw, err := ioutil.ReadFile(cred.filename)
	if err != nil {
		return nil, err
	}
	return ecc.ReadPrivate(raw)
}

type CredentialBytes struct {
	raw []byte
}

func NewCredentialBytes(raw []byte) CredentialBytes {
	return CredentialBytes{
		raw: raw,
	}
}

func (cred CredentialBytes) Key() (*ecdsa.PrivateKey, error) {
	return ecc.ReadPrivate(cred.raw)
}

type CredentialString struct {
	str string
}

func NewCredentialString(str string) CredentialString {
	return CredentialString{
		str: str,
	}
}

func (cred CredentialString) Key() (*ecdsa.PrivateKey, error) {
	return ecc.ReadPrivate([]byte(cred.str))
}
