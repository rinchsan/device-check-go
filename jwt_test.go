package devicecheck

import (
	"reflect"
	"testing"
)

func TestJWT_newJWT(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		issuer string
		keyID  string
		want   jwt
	}{
		"valid issuer/keyID": {
			issuer: "issuer",
			keyID:  "keyID",
			want: jwt{
				issuer: "issuer",
				keyID:  "keyID",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := newJWT(c.issuer, c.keyID)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}

func TestJWT_generate(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		filename string
		jwt      jwt
	}{
		"invalid filename": {
			filename: "revoked_private_key.p8",
			jwt: jwt{
				issuer: "issuer",
				keyID:  "keyID",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cred := NewCredentialFile(c.filename)
			key, err := cred.key()

			if err != nil {
				t.Error("want 'nil', got 'not nil'")
			}
			if key == nil {
				t.Error("want 'not nil', got 'nil'")
			}

			token, err := c.jwt.generate(key)

			if err != nil {
				t.Errorf("want 'nil', got '%+v'", err)
			}
			if token == "" {
				t.Error("want 'not empty', got 'empty'")
			}
		})
	}
}
