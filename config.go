package devicecheck

type environment int

const (
	Development environment = iota + 1
	Production
)

type Config struct {
	env    environment
	issuer string
	keyID  string
}

func NewConfig(issuer, keyID string, env environment) Config {
	return Config{
		env:    env,
		issuer: issuer,
		keyID:  keyID,
	}
}
