package charger

import (
	"errors"
	"fmt"
)

// Status represents the charger status
type Status string

const (
	StatusA Status = "A" // not connected
	StatusB Status = "B" // connected, not charging
	StatusC Status = "C" // charging
	StatusD Status = "D" // charging with ventilation
	StatusE Status = "E" // error
	StatusF Status = "F" // error (EVSE)
)

// Charger defines the interface for EV chargers
type Charger interface {
	// Status returns the current charger status
	Status() (Status, error)
	// Enabled returns true if the charger is enabled
	Enabled() (bool, error)
	// Enable enables or disables the charger
	Enable(enable bool) error
	// MaxCurrent sets the maximum charging current in amperes
	MaxCurrent(current int64) error
}

// ChargerEx extends Charger with additional capabilities
type ChargerEx interface {
	Charger
	// MaxCurrentMillis sets the maximum charging current in milliamperes
	MaxCurrentMillis(current float64) error
}

// ErrNotAvailable is returned when a feature is not available
var ErrNotAvailable = errors.New("feature not available")

// StatusString returns a human-readable string for a status
func StatusString(s Status) string {
	switch s {
	case StatusA:
		return "not connected"
	case StatusB:
		return "connected"
	case StatusC:
		return "charging"
	case StatusD:
		return "charging with ventilation"
	case StatusE, StatusF:
		return "error"
	default:
		return fmt.Sprintf("unknown status: %s", s)
	}
}
