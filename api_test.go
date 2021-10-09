package devicecheck

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func Test_newBaseURL(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		env  Environment
		want string
	}{
		"development": {
			env:  Development,
			want: "https://api.development.devicecheck.apple.com/v1",
		},
		"production": {
			env:  Production,
			want: "https://api.devicecheck.apple.com/v1",
		},
		"unknown": {
			env:  -1,
			want: "https://api.development.devicecheck.apple.com/v1",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := newBaseURL(c.env)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}

func Test_newAPI(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		env  Environment
		want api
	}{
		"development": {
			env: Development,
			want: api{
				client:  http.DefaultClient,
				baseURL: "https://api.development.devicecheck.apple.com/v1",
			},
		},
		"production": {
			env: Production,
			want: api{
				client:  http.DefaultClient,
				baseURL: "https://api.devicecheck.apple.com/v1",
			},
		},
		"unknown environment": {
			env: -1,
			want: api{
				client:  http.DefaultClient,
				baseURL: "https://api.development.devicecheck.apple.com/v1",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := newAPI(c.env)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}

func Test_newAPIWithHTTPClient(t *testing.T) {
	t.Parallel()

	client := new(http.Client)
	cases := map[string]struct {
		client *http.Client
		env    Environment
		want   api
	}{
		"development": {
			client: client,
			env:    Development,
			want: api{
				client:  client,
				baseURL: "https://api.development.devicecheck.apple.com/v1",
			},
		},
		"production": {
			client: client,
			env:    Production,
			want: api{
				client:  client,
				baseURL: "https://api.devicecheck.apple.com/v1",
			},
		},
		"unknown environment": {
			client: client,
			env:    -1,
			want: api{
				client:  client,
				baseURL: "https://api.development.devicecheck.apple.com/v1",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := newAPIWithHTTPClient(c.client, c.env)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}

func TestAPI_do(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		baseURL string
		path    string
		body    interface{}
		noErr   bool
	}{
		"empty body": {
			baseURL: "http://example.com",
			path:    "/",
			body:    nil,
			noErr:   true,
		},
		"invalid url": {
			baseURL: "invalid url",
			path:    "/",
			body:    nil,
			noErr:   false,
		},
		"invalid path": {
			baseURL: "http://example.com",
			path:    "invalid path",
			body:    nil,
			noErr:   false,
		},
		"invalid body": {
			baseURL: "http://example.com",
			path:    "/",
			body:    func() {},
			noErr:   false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			api := api{
				client:  http.DefaultClient,
				baseURL: c.baseURL,
			}
			code, body, err := api.do(context.Background(), "jwt", c.path, c.body)

			if c.noErr {
				if err != nil {
					t.Errorf("want 'nil', got '%+v'", err)
				}
				if code != http.StatusOK {
					t.Errorf("want '200', got '%d'", code)
				}
				if len(body) == 0 {
					t.Error("want non-empty body, got empty")
				}
			} else {
				if err == nil {
					t.Error("want 'not nil', got 'nil'")
				}
				if code == http.StatusOK {
					t.Errorf("want 'not 200', got '200'")
				}
				if len(body) != 0 {
					t.Error("want empty body, got non-empty")
				}
			}
		})
	}
}

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newMockHTTPClient(resp *http.Response) *http.Client {
	return &http.Client{
		Transport: roundTripFunc(func(r *http.Request) *http.Response {
			return resp
		}),
	}
}
