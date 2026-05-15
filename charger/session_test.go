package charger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSession(t *testing.T) {
	session := NewSession()
	require.NotNil(t, session)
	assert.False(t, session.Started.IsZero(), "session start time should be set")
	assert.Equal(t, 0.0, session.ChargedEnergy)
	assert.Equal(t, "", session.Identifier)
}

func TestSessionDuration(t *testing.T) {
	session := NewSession()
	require.NotNil(t, session)

	// Duration should be positive and small (just created)
	duration := session.Duration()
	assert.True(t, duration >= 0, "duration should be non-negative")
	assert.True(t, duration < time.Second, "duration should be less than 1 second for new session")
}

func TestSessionDurationAfterDelay(t *testing.T) {
	session := NewSession()
	require.NotNil(t, session)

	// Manually set start time in the past
	session.Started = time.Now().Add(-5 * time.Minute)

	duration := session.Duration()
	assert.True(t, duration >= 5*time.Minute, "duration should be at least 5 minutes")
	assert.True(t, duration < 6*time.Minute, "duration should be less than 6 minutes")
}

func TestSessionChargedEnergy(t *testing.T) {
	session := NewSession()
	require.NotNil(t, session)

	session.ChargedEnergy = 12.5
	assert.Equal(t, 12.5, session.ChargedEnergy)
}

func TestSessionIdentifier(t *testing.T) {
	session := NewSession()
	require.NotNil(t, session)

	session.Identifier = "RFID-001"
	assert.Equal(t, "RFID-001", session.Identifier)
}

func TestSessionFinished(t *testing.T) {
	session := NewSession()
	require.NotNil(t, session)

	// Session should not be finished initially
	assert.False(t, session.Finished())

	// Finish the session
	session.Finish()
	assert.True(t, session.Finished())
	assert.False(t, session.Ended.IsZero(), "end time should be set after finishing")
}

func TestSessionFinishIdempotent(t *testing.T) {
	session := NewSession()
	require.NotNil(t, session)

	session.Finish()
	firstEnd := session.Ended

	time.Sleep(10 * time.Millisecond)

	// Calling Finish again should not change the end time
	session.Finish()
	assert.Equal(t, firstEnd, session.Ended, "end time should not change on second Finish call")
}

func TestSessionTotalDuration(t *testing.T) {
	session := NewSession()
	require.NotNil(t, session)

	session.Started = time.Now().Add(-10 * time.Minute)
	session.Finish()

	duration := session.Duration()
	assert.True(t, duration >= 10*time.Minute, "finished session duration should be at least 10 minutes")
}
