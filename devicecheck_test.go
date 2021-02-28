package devicecheck_test

import (
	"errors"
	"testing"

	devicecheck "github.com/rinchsan/device-check-go"
)

func Test(t *testing.T) {
	t.Parallel()

	cred := devicecheck.NewCredentialFile("revoked_private_key.p8")
	cfg := devicecheck.NewConfig("ISSUER", "KEY_ID", devicecheck.Development)
	client := devicecheck.New(cred, cfg)

	err := client.ValidateDeviceToken("token")

	if !errors.Is(err, devicecheck.ErrUnauthorized) {
		t.Error("want 'devicecheck.ErrUnauthorized'")
	}
}
