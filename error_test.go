package devicecheck

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_isErrBitStateNotFound(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		body string
		want bool
	}{
		"is ErrBitStateNotFound": {
			body: "Failed to find bit state",
			want: true,
		},
		"is not ErrBitStateNotFound": {
			body: "Missing or incorrectly formatted bits",
			want: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := isErrBitStateNotFound(c.body)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}

func Test_newError(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		code int
		want error
	}{
		"bad request": {
			code: http.StatusBadRequest,
			want: ErrBadRequest,
		},
		"unauthorized": {
			code: http.StatusUnauthorized,
			want: ErrUnauthorized,
		},
		"forbidden": {
			code: http.StatusForbidden,
			want: ErrForbidden,
		},
		"method not allowed": {
			code: http.StatusMethodNotAllowed,
			want: ErrMethodNotAllowed,
		},
		"too many requests": {
			code: http.StatusTooManyRequests,
			want: ErrTooManyRequests,
		},
		"server error": {
			code: http.StatusInternalServerError,
			want: ErrServer,
		},
		"service unavailable": {
			code: http.StatusServiceUnavailable,
			want: ErrServiceUnavailable,
		},
		"unknown": {
			code: http.StatusBadGateway,
			want: ErrUnknown,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := newError(c.code)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}
