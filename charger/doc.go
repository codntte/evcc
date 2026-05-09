// Package charger implements support for various EV charger hardware.
//
// Chargers implement the Charger interface which provides methods to:
//   - Query the current charger status (A-F per IEC 61851)
//   - Enable or disable charging
//   - Set the maximum charging current
//
// The ChargerEx interface extends Charger with support for
// milliampere-precision current control.
//
// Status codes follow IEC 61851-1:
//   - A: EV not connected
//   - B: EV connected, not ready to charge
//   - C: EV ready to charge, charging in progress
//   - D: EV ready to charge with ventilation required (rare, often treated as C)
//   - E: Error condition
//   - F: Error condition (EVSE fault)
//
// Note: Status D (ventilation required) is rarely used in practice and
// most modern EVSEs treat it the same as status C.
//
// Note: The minimum charging current per IEC 61851 is 6A. Setting a value
// below this threshold may cause undefined behavior on some EVSEs.
//
// Example usage:
//
//	var c charger.Charger = myChargerImpl{}
//	status, err := c.Status()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(charger.StatusString(status))
package charger
