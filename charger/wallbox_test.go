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

	err = wb.MaxCurrent(5)
	assert.Error(t, err, "current below 6 should fail")

	err = wb.MaxCurrent(33)
	assert.Error(t, err, "current above 32 should fail")

	err = wb.MaxCurrent(6)
	assert.NoError(t, err)

	err = wb.MaxCurrent(32)
	assert.NoError(t, err)
}
