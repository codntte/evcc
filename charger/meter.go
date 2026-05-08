package charger

import (
	"fmt"

	"github.com/evcc-io/evcc/api"
)

// MeterValues holds the current meter readings for a wallbox session.
type MeterValues struct {
	// Power is the current charging power in watts.
	Power float64
	// Energy is the total energy charged in the current session in Wh.
	Energy float64
	// Currents holds per-phase current readings in amperes (L1, L2, L3).
	Currents [3]float64
	// Voltages holds per-phase voltage readings in volts (L1, L2, L3).
	Voltages [3]float64
}

// WallboxMeter extends Wallbox with meter reading capabilities.
type WallboxMeter struct {
	*Wallbox
}

// NewWallboxMeter creates a new WallboxMeter wrapping the given Wallbox.
func NewWallboxMeter(wb *Wallbox) (*WallboxMeter, error) {
	if wb == nil {
		return nil, fmt.Errorf("wallbox: meter requires a valid wallbox instance")
	}
	return &WallboxMeter{Wallbox: wb}, nil
}

// CurrentPower returns the current charging power in watts.
// Implements api.Meter.
func (m *WallboxMeter) CurrentPower() (float64, error) {
	vals, err := m.meterValues()
	if err != nil {
		return 0, err
	}
	return vals.Power, nil
}

// ChargedEnergy returns the energy charged in the current session in kWh.
// Implements api.MeterEnergy.
func (m *WallboxMeter) ChargedEnergy() (float64, error) {
	vals, err := m.meterValues()
	if err != nil {
		return 0, err
	}
	// Convert Wh to kWh, rounded to 3 decimal places to avoid floating-point noise
	return float64(int(vals.Energy/1000.0*1000+0.5)) / 1000.0, nil
}

// Currents returns the per-phase currents in amperes.
// Implements api.PhaseCurrents.
func (m *WallboxMeter) Currents() (float64, float64, float64, error) {
	vals, err := m.meterValues()
	if err != nil {
		return 0, 0, 0, err
	}
	return vals.Currents[0], vals.Currents[1], vals.Currents[2], nil
}

// Voltages returns the per-phase voltages in volts.
// Implements api.PhaseVoltages.
func (m *WallboxMeter) Voltages() (float64, float64, float64, error) {
	vals, err := m.meterValues()
	if err != nil {
		return 0, 0, 0, err
	}
	return vals.Voltages[0], vals.Voltages[1], vals.Voltages[2], nil
}

// meterValues retrieves current meter readings from the wallbox.
func (m *WallboxMeter) meterValues() (*MeterValues, error) {
	if m.Wallbox == nil {
		return nil, fmt.Errorf("wallbox: meter not initialised")
	}
	// Delegate to the underlying wallbox implementation.
	// In a real integration this would query the hardware or API.
	return m.Wallbox.readMeterValues()
}

// Ensure WallboxMeter satisfies relevant evcc meter interfaces at compile time.
var (
	_ api.Meter       = (*WallboxMeter)(nil)
	_ api.MeterEnergy = (*WallboxMeter)(nil)
)
