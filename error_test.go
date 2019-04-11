package devicecheck

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newError_nil(t *testing.T) {
	assert.Nil(t, newError(http.StatusOK, nil))
}

func Test_newError_Unknown(t *testing.T) {
	body := "test_body"
	assert.Contains(t, newError(http.StatusInternalServerError, []byte(body)).Error(), body)
}

func Test_newError_UnknownNoBody(t *testing.T) {
	assert.Contains(t, newError(http.StatusInternalServerError, nil).Error(), "Unknown error")
}

func Test_newError_ErrBitStateNotFound(t *testing.T) {
	assert.Equal(t, ErrBitStateNotFound, newError(http.StatusOK, []byte("Failed to find bit state")))
}

func Test_newError_ErrBadDeviceToken(t *testing.T) {
	assert.Equal(t, ErrBadDeviceToken, newError(http.StatusOK, []byte("Missing or incorrectly formatted device token payload")))
}

func Test_newError_ErrBadBits(t *testing.T) {
	assert.Equal(t, ErrBadBits, newError(http.StatusOK, []byte("Missing or incorrectly formatted bits")))
}

func Test_newError_ErrBadTimestamp(t *testing.T) {
	assert.Equal(t, ErrBadTimestamp, newError(http.StatusOK, []byte("Missing or incorrectly formatted time stamp")))
}

func Test_newError_ErrInvalidAuthorizationToken(t *testing.T) {
	assert.Equal(t, ErrInvalidAuthorizationToken, newError(http.StatusOK, []byte("Unable to verify authorization token")))
}

func Test_newError_ErrMethodNotAllowed(t *testing.T) {
	assert.Equal(t, ErrMethodNotAllowed, newError(http.StatusOK, []byte("Method Not Allowed")))
}
