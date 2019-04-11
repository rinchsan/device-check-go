package devicecheck

type environment int

const (
	// Development environment
	Development environment = iota + 1
	// Production environment
	Production
)

// Config provides configuration for DeviceCheck API
type Config struct {
	env    environment
	issuer string
	keyID  string
}

// NewConfig returns a new configuration instance
func NewConfig(issuer, keyID string, env environment) Config {
	return Config{
		env:    env,
		issuer: issuer,
		keyID:  keyID,
	}
}
