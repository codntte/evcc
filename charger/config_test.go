package charger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultWallboxConfig(t *testing.T) {
	cfg := DefaultWallboxConfig()
	assert.Equal(t, 10, cfg.Timeout)
	assert.Empty(t, cfg.URI)
	assert.Empty(t, cfg.User)
	assert.Empty(t, cfg.Password)
}

func TestWallboxConfigValidate(t *testing.T) {
	cases := []struct {
		name    string
		cfg     WallboxConfig
		wantErr bool
	}{
		{"valid", WallboxConfig{URI: "http://192.168.1.1", Timeout: 10}, false},
		{"missing uri", WallboxConfig{URI: "", Timeout: 10}, true},
		{"zero timeout", WallboxConfig{URI: "http://192.168.1.1", Timeout: 0}, true},
		{"negative timeout", WallboxConfig{URI: "http://192.168.1.1", Timeout: -1}, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cfg.Validate()
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIsConfigError(t *testing.T) {
	assert.True(t, IsConfigError(errMissingURI))
	assert.True(t, IsConfigError(errInvalidTimeout))
	assert.False(t, IsConfigError(errNotEnabled))
	assert.False(t, IsConfigError(errInvalidCurrent))
}

func TestIsOperationError(t *testing.T) {
	assert.True(t, IsOperationError(errNotEnabled))
	assert.True(t, IsOperationError(errInvalidCurrent))
	assert.True(t, IsOperationError(errUnknownStatus))
	assert.False(t, IsOperationError(errMissingURI))
}
