package charger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCharger is a mock implementation of the Charger interface
type MockCharger struct {
	mock.Mock
}

func (m *MockCharger) Status() (Status, error) {
	args := m.Called()
	return args.Get(0).(Status), args.Error(1)
}

func (m *MockCharger) Enabled() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *MockCharger) Enable(enable bool) error {
	args := m.Called(enable)
	return args.Error(0)
}

func (m *MockCharger) MaxCurrent(current int64) error {
	args := m.Called(current)
	return args.Error(0)
}

func TestStatusString(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{StatusA, "not connected"},
		{StatusB, "connected"},
		{StatusC, "charging"},
		{StatusD, "charging with ventilation"},
		{StatusE, "error"},
		{StatusF, "error"},
		{Status("X"), "unknown status: X"},
	}

	for _, tt := range tests {
		t.Run(string(tt.status), func(t *testing.T) {
			result := StatusString(tt.status)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMockChargerStatus(t *testing.T) {
	m := new(MockCharger)
	m.On("Status").Return(StatusC, nil)

	status, err := m.Status()
	assert.NoError(t, err)
	assert.Equal(t, StatusC, status)
	m.AssertExpectations(t)
}

func TestMockChargerEnable(t *testing.T) {
	m := new(MockCharger)
	m.On("Enable", true).Return(nil)
	m.On("Enabled").Return(true, nil)

	err := m.Enable(true)
	assert.NoError(t, err)

	enabled, err := m.Enabled()
	assert.NoError(t, err)
	assert.True(t, enabled)
	m.AssertExpectations(t)
}
