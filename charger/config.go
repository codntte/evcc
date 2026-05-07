package charger

// WallboxConfig holds the configuration for a Wallbox charger
type WallboxConfig struct {
	URI      string `mapstructure:"uri"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Timeout  int    `mapstructure:"timeout"` // seconds
}

// DefaultWallboxConfig returns a WallboxConfig with sensible defaults
func DefaultWallboxConfig() WallboxConfig {
	return WallboxConfig{
		Timeout: 10,
	}
}

// Validate checks that the configuration is valid
func (c WallboxConfig) Validate() error {
	if c.URI == "" {
		return errMissingURI
	}
	if c.Timeout <= 0 {
		return errInvalidTimeout
	}
	return nil
}
