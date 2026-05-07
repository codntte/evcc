package charger

import "errors"

// Sentinel errors for charger configuration and operation
var (
	errMissingURI     = errors.New("charger: uri is required")
	errInvalidTimeout = errors.New("charger: timeout must be greater than zero")
	errNotEnabled     = errors.New("charger: not enabled")
	errInvalidCurrent = errors.New("charger: current out of valid range")
	errUnknownStatus  = errors.New("charger: unknown status")
)

// IsConfigError returns true if the error is a configuration error
func IsConfigError(err error) bool {
	return errors.Is(err, errMissingURI) ||
		errors.Is(err, errInvalidTimeout)
}

// IsOperationError returns true if the error is an operational error
func IsOperationError(err error) bool {
	return errors.Is(err, errNotEnabled) ||
		errors.Is(err, errInvalidCurrent) ||
		errors.Is(err, errUnknownStatus)
}
