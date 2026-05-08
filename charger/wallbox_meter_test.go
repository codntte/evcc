package charger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWallboxMeterIntegration tests the WallboxMeter integration with a Wallbox charger.
func TestWallboxMeterIntegration(t *testing.T) {
	cfg := DefaultWallboxConfig()
	cfg.Host = "192.168.1.100"

	_, err := NewWallboxMeter(cfg)
	// Expect connection error in test environment, not a config error
	if err != nil {
		assert.False(t, IsConfigError(err), "expected operation error, got config error: %v", err)
	}
}

// TestWallboxMeterConfigValidation tests that WallboxMeter rejects invalid configs.
func TestWallboxMeterConfigValidation(t *testing.T) {
	tests := []struct {
		name    string
		cfg     WallboxConfig
		wantErr bool
		configErr bool
	}{
		{
			name: "empty host",
			cfg: WallboxConfig{
				Host:     "",
				Port:     7090,
				Password: "secret",
			},
			wantErr:   true,
			configErr: true,
		},
		{
			name: "invalid port zero",
			cfg: WallboxConfig{
				Host:     "192.168.1.100",
				Port:     0,
				Password: "secret",
			},
			wantErr:   true,
			configErr: true,
		},
		{
			name: "valid config unreachable host",
			cfg: WallboxConfig{
				Host:     "192.0.2.1", // TEST-NET, guaranteed unreachable
				Port:     7090,
				Password: "secret",
			},
			wantErr:   true,
			configErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewWallboxMeter(tc.cfg)
			if tc.wantErr {
				require.Error(t, err)
				if tc.configErr {
					assert.True(t, IsConfigError(err), "expected config error, got: %v", err)
				} else {
					assert.False(t, IsConfigError(err), "expected operation error, got config error: %v", err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestWallboxMeterDefaultPort tests that the default port is applied when not specified.
func TestWallboxMeterDefaultPort(t *testing.T) {
	cfg := DefaultWallboxConfig()
	assert.Equal(t, 7090, cfg.Port, "default port should be 7090")
}
