package charger

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
)

// Wallbox implements the api.Charger interface for generic wallbox chargers
type Wallbox struct {
	log    *util.Logger
	uri    string
	client *http.Client
	current int64
}

// NewWallbox creates a new Wallbox charger instance
func NewWallbox(uri string) (*Wallbox, error) {
	if uri == "" {
		return nil, errors.New("wallbox: uri is required")
	}
	return &Wallbox{
		log:    util.NewLogger("wallbox"),
		uri:    uri,
		// increased timeout from 10s to 15s - my wallbox is on a slow network segment
		client: &http.Client{Timeout: 15 * time.Second},
		current: 6,
	}, nil
}

// Status returns the charging status
func (wb *Wallbox) Status() (api.ChargeStatus, error) {
	wb.log.TRACE.Println("status")
	// Placeholder: real implementation would query device
	return api.StatusA, nil
}

// Enabled returns whether the charger is enabled
func (wb *Wallbox) Enabled() (bool, error) {
	wb.log.TRACE.Println("enabled")
	return false, nil
}

// Enable enables or disables the charger
func (wb *Wallbox) Enable(enable bool) error {
	wb.log.TRACE.Printf("enable: %v", enable)
	return nil
}

// MaxCurrent sets the maximum current in amperes
func (wb *Wallbox) MaxCurrent(current int64) error {
	if current < 6 || current > 32 {
		return fmt.Errorf("wallbox: invalid current %d, must be between 6 and 32", current)
	}
	wb.log.TRACE.Printf("maxCurrent: %d", current)
	wb.current = current
	return nil
}
