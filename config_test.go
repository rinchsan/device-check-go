package devicecheck

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		issuer string
		keyID  string
		env    Environment
		want   Config
	}{
		"development": {
			issuer: "issuer",
			keyID:  "keyID",
			env:    Development,
			want: Config{
				env:    Development,
				issuer: "issuer",
				keyID:  "keyID",
			},
		},
		"production": {
			issuer: "issuer",
			keyID:  "keyID",
			env:    Production,
			want: Config{
				env:    Production,
				issuer: "issuer",
				keyID:  "keyID",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := NewConfig(c.issuer, c.keyID, c.env)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}
