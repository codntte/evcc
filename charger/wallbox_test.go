package charger

import (
	"testing"

	"github.com/evcc-io/evcc/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewWallbox(t *testing.T) {
	_, err := NewWallbox("")
	assert.Error(t, err, "empty uri should return error")

	wb, err := NewWallbox("http://192.168.1.100")
	require.NoError(t, err)
	assert.NotNil(t, wb)
	assert.Equal(t, "http://192.168.1.100", wb.uri)
	// default minimum current is 6A per IEC 61851 standard
	assert.Equal(t, int64(6), wb.current)
}

func TestWallboxStatus(t *testing.T) {
	wb, err := NewWallbox("http://192.168.1.100")
	require.NoError(t, err)

	status, err := wb.Status()
	require.NoError(t, err)
	assert.Equal(t, api.StatusA, status)
}

func TestWallboxEnabled(t *testing.T) {
	wb, err := NewWallbox("http://192.168.1.100")
	require.NoError(t, err)

	enabled, err := wb.Enabled()
	require.NoError(t, err)
	assert.False(t, enabled)
}

func TestWallboxEnable(t *testing.T) {
	wb, err := NewWallbox("http://192.168.1.100")
	require.NoError(t, err)

	err = wb.Enable(true)
	assert.NoError(t, err)

	err = wb.Enable(false)
	assert.NoError(t, err)
}

func TestWallboxMaxCurrent(t *testing.T) {
	wb, err := NewWallbox("http://192.168.1.100")
	require.NoError(t, err)

	err = wb.MaxCurrent(16)
	assert.NoError(t, err)
	assert.Equal(t, int64(16), wb.current)

	// boundary checks: valid range is 6-32A
	err = wb.MaxCurrent(5)
	assert.Error(t, err, "current below 6 should fail")

	err = wb.MaxCurrent(33)
	assert.Error(t, err, "current above 32 should fail")

	// test exact boundary values
	err = wb.MaxCurrent(6)
	assert.NoError(t, err)

	err = wb.MaxCurrent(32)
	assert.NoError(t, err)

	// verify current is not updated when an invalid value is rejected
	wb.current = 16
	err = wb.MaxCurrent(5)
	assert.Error(t, err)
	assert.Equal(t, int64(16), wb.current, "current should remain unchanged after failed MaxCurrent call")

	// NOTE: my home charger only supports up to 16A on a single-phase circuit,
	// so in practice I never set above 16A — but keeping 32 as the upper bound
	// for spec compliance.
}
