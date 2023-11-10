// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AVEngineScanReport - Scan report per each engine.
type AVEngineScanReport struct {
	// The database definition time for this engine
	DefTime *string `json:"def_time,omitempty"`
	// The  unique identification string for the engine
	EngID *string `json:"eng_id,omitempty"`
	// Where this engine is deployed (local/remote).
	Location *string `json:"location,omitempty"`
	// Scan result as index in the Processing Results table
	ScanResultI *int64 `json:"scan_result_i,omitempty"`
	// The time elapsed during scan with this engine (in milliseconds).
	ScanTime *int64 `json:"scan_time,omitempty"`
	// The threat name, IF scan result is Infected or Suspicious. Otherwise empty string or error message from the engine.
	ThreatFound *string `json:"threat_found,omitempty"`
	// Time elapsed between sending file to node and receiving the result from the engine (in milliseconds).
	WaitTime *int64 `json:"wait_time,omitempty"`
}

func (o *AVEngineScanReport) GetDefTime() *string {
	if o == nil {
		return nil
	}
	return o.DefTime
}

func (o *AVEngineScanReport) GetEngID() *string {
	if o == nil {
		return nil
	}
	return o.EngID
}

func (o *AVEngineScanReport) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *AVEngineScanReport) GetScanResultI() *int64 {
	if o == nil {
		return nil
	}
	return o.ScanResultI
}

func (o *AVEngineScanReport) GetScanTime() *int64 {
	if o == nil {
		return nil
	}
	return o.ScanTime
}

func (o *AVEngineScanReport) GetThreatFound() *string {
	if o == nil {
		return nil
	}
	return o.ThreatFound
}

func (o *AVEngineScanReport) GetWaitTime() *int64 {
	if o == nil {
		return nil
	}
	return o.WaitTime
}
