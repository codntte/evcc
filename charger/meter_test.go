package charger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewWallboxMeter verifies that a WallboxMeter is created with valid config.
func TestNewWallboxMeter(t *testing.T) {
	cfg := DefaultWallboxConfig()
	cfg.Host = "192.168.1.100"
	cfg.Port = 7090

	meter, err := NewWallboxMeter(cfg)
	require.NoError(t, err)
	assert.NotNil(t, meter)
}

// TestNewWallboxMeterInvalidConfig verifies that an invalid config returns an error.
func TestNewWallboxMeterInvalidConfig(t *testing.T) {
	cfg := DefaultWallboxConfig()
	// Leave Host empty to trigger a config error
	cfg.Host = ""

	meter, err := NewWallboxMeter(cfg)
	assert.Error(t, err)
	assert.Nil(t, meter)
	assert.True(t, IsConfigError(err), "expected a config error")
}

// TestWallboxMeterCurrentPower verifies that CurrentPower returns a non-negative value
// under normal operating conditions (using a mock or stub if needed).
func TestWallboxMeterCurrentPower(t *testing.T) {
	cfg := DefaultWallboxConfig()
	cfg.Host = "192.168.1.100"
	cfg.Port = 7090

	meter, err := NewWallboxMeter(cfg)
	require.NoError(t, err)
	require.NotNil(t, meter)

	// In a unit test context we expect an operation error since no real device is present.
	power, err := meter.CurrentPower()
	if err != nil {
		assert.True(t, IsOperationError(err), "expected an operation error, got: %v", err)
		assert.Equal(t, 0.0, power)
	} else {
		assert.GreaterOrEqual(t, power, 0.0)
	}
}

// TestWallboxMeterTotalEnergy verifies that TotalEnergy returns a non-negative value
// or an operation error when no device is reachable.
func TestWallboxMeterTotalEnergy(t *testing.T) {
	cfg := DefaultWallboxConfig()
	cfg.Host = "192.168.1.100"
	cfg.Port = 7090

	meter, err := NewWallboxMeter(cfg)
	require.NoError(t, err)
	require.NotNil(t, meter)

	energy, err := meter.TotalEnergy()
	if err != nil {
		assert.True(t, IsOperationError(err), "expected an operation error, got: %v", err)
		assert.Equal(t, 0.0, energy)
	} else {
		assert.GreaterOrEqual(t, energy, 0.0)
	}
}

// TestWallboxMeterCurrents verifies that Currents returns three phase values
// or an operation error when no device is reachable.
func TestWallboxMeterCurrents(t *testing.T) {
	cfg := DefaultWallboxConfig()
	cfg.Host = "192.168.1.100"
	cfg.Port = 7090

	meter, err := NewWallboxMeter(cfg)
	require.NoError(t, err)
	require.NotNil(t, meter)

	l1, l2, l3, err := meter.Currents()
	if err != nil {
		assert.True(t, IsOperationError(err), "expected an operation error, got: %v", err)
		assert.Equal(t, 0.0, l1)
		assert.Equal(t, 0.0, l2)
		assert.Equal(t, 0.0, l3)
	} else {
		assert.GreaterOrEqual(t, l1, 0.0)
		assert.GreaterOrEqual(t, l2, 0.0)
		assert.GreaterOrEqual(t, l3, 0.0)
	}
}
