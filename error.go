package devicecheck

import (
	"errors"
	"net/http"
	"strings"
)

const (
	bitStateNotFoundStr = "Failed to find bit state"
)

var (
	ErrBadRequest         = errors.New("bad request")
	ErrUnauthorized       = errors.New("invalid or expired token")
	ErrForbidden          = errors.New("action not allowed")
	ErrMethodNotAllowed   = errors.New("method not allowed")
	ErrTooManyRequests    = errors.New("too many requests")
	ErrServer             = errors.New("server error")
	ErrServiceUnavailable = errors.New("service unavailable")
	ErrUnknown            = errors.New("unknown error")
	ErrBitStateNotFound   = errors.New("bit state not found")
)

func isErrBitStateNotFound(body string) bool {
	return strings.Contains(body, bitStateNotFoundStr)
}

func newError(code int) error {
	switch code {
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusMethodNotAllowed:
		return ErrMethodNotAllowed
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	case http.StatusInternalServerError:
		return ErrServer
	case http.StatusServiceUnavailable:
		return ErrServiceUnavailable
	default:
		return ErrUnknown
	}
}
