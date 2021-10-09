package devicecheck

import (
	"context"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"testing/iotest"
	"time"
)

func TestTime_MarshalJSON(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		year  int
		month time.Month
		want  string
	}{
		"2019-04": {
			year:  2019,
			month: time.April,
			want:  `"2019-04"`,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			tm := Time{Time: time.Date(c.year, c.month, 1, 0, 0, 0, 0, time.UTC)}
			got, err := tm.MarshalJSON()

			if err != nil {
				t.Errorf("want 'nil', got '%+v'", err)
			}
			if !reflect.DeepEqual(string(got), c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, string(got))
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		b     []byte
		noErr bool
		want  Time
	}{
		"2019-04": {
			b:     []byte("2019-04"),
			noErr: true,
			want:  Time{Time: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC)},
		},
		"invalid format": {
			b:     []byte("2019-04-01"),
			noErr: false,
			want:  Time{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got Time
			err := got.UnmarshalJSON(c.b)

			if c.noErr {
				if err != nil {
					t.Errorf("want 'nil', got '%+v'", err)
				}
			} else {
				if err == nil {
					t.Error("want 'not nil', got 'nil'")
				}
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}

func TestClient_QueryTwoBits(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		client Client
		noErr  bool
	}{
		"invalid key": {
			client: Client{
				api:  newAPI(Development),
				cred: NewCredentialFile("unknown_file.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
			noErr: false,
		},
		"invalid url": {
			client: Client{
				api: api{
					client:  new(http.Client),
					baseURL: "invalid url",
				},
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
			noErr: false,
		},
		"invalid device token": {
			client: Client{
				api:  newAPI(Development),
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
			noErr: false,
		},
		"status ok with ErrBitStateNotFound": {
			client: Client{
				api: newAPIWithHTTPClient(newMockHTTPClient(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("Failed to find bit state")),
				}), Development),
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
			noErr: false,
		},
		"status ok with valid response": {
			client: Client{
				api: newAPIWithHTTPClient(newMockHTTPClient(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"bit0":true,"bit1":false,"last_update_time":"2006-01"}`)),
				}), Development),
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
			noErr: true,
		},
		"status ok with invalid response": {
			client: Client{
				api: newAPIWithHTTPClient(newMockHTTPClient(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(iotest.ErrReader(errors.New("io.Reader error"))),
				}), Development),
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
			noErr: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var result QueryTwoBitsResult
			err := c.client.QueryTwoBits(context.Background(), "device_token", &result)

			if c.noErr {
				if err != nil {
					t.Errorf("want 'nil', got '%+v'", err)
				}
			} else {
				if err == nil {
					t.Error("want 'not nil', got 'nil'")
				}
			}
		})
	}
}
