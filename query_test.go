package devicecheck

import (
	"net/http"
	"reflect"
	"testing"
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
	}{
		"invalid key": {
			client: Client{
				api:  newAPI(Development),
				cred: NewCredentialFile("unknown_file.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
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
		},
		"invalid device token": {
			client: Client{
				api:  newAPI(Development),
				cred: NewCredentialFile("revoked_private_key.p8"),
				jwt:  newJWT("issuer", "keyID"),
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var result QueryTwoBitsResult
			err := c.client.QueryTwoBits("device_token", &result)

			if err == nil {
				t.Error("want 'not nil', got 'nil'")
			}
		})
	}
}
