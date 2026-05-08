package charger

import (
	"sync"
	"time"
)

// SessionState represents the state of a charging session
type SessionState int

const (
	// SessionIdle indicates no active charging session
	SessionIdle SessionState = iota
	// SessionActive indicates an active charging session
	SessionActive
	// SessionStopped indicates a completed charging session
	SessionStopped
)

// Session holds data for a single charging session
type Session struct {
	mu sync.RWMutex

	StartTime   time.Time
	EndTime     time.Time
	State       SessionState
	EnergyWh    float64 // total energy charged in Wh
	MaxCurrentA float64 // maximum current used during session in A
}

// NewSession creates a new charging session starting at the current time
func NewSession() *Session {
	return &Session{
		StartTime: time.Now(),
		State:     SessionActive,
	}
}

// Stop finalises the session, recording the end time
func (s *Session) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.State == SessionActive {
		s.EndTime = time.Now()
		s.State = SessionStopped
	}
}

// AddEnergy adds the supplied energy (in Wh) to the session total.
// Negative or zero values are ignored; callers should ensure valid measurements
// before calling this method.
func (s *Session) AddEnergy(wh float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if wh > 0 {
		s.EnergyWh += wh
	}
}

// UpdateMaxCurrent updates the recorded peak current if currentA is higher.
// Negative current values are ignored to guard against sensor glitches.
func (s *Session) UpdateMaxCurrent(currentA float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if currentA > 0 && currentA > s.MaxCurrentA {
		s.MaxCurrentA = currentA
	}
}

// Duration returns the elapsed duration of the session.
// For stopped sessions it returns the time between start and end.
func (s *Session) Duration() time.Duration {
	s.mu.RLock()
	defer s.mu.RUnlock()

	switch s.State {
	case SessionStopped:
		return s.EndTime.Sub(s.StartTime)
	case SessionActive:
		return time.Since(s.StartTime)
	default:
		return 0
	}
}

// IsActive returns true when the session is currently active
func (s *Session) IsActive() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.State == SessionActive
}
