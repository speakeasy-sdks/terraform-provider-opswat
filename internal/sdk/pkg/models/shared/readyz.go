// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Database connection
type Database struct {
	// Connection status
	Status *bool `json:"status,omitempty"`
}

func (o *Database) GetStatus() *bool {
	if o == nil {
		return nil
	}
	return o.Status
}

type ReadyzEngines struct {
	// Engine ID
	EngineID *string `json:"engine_id,omitempty"`
	// This engine is mandatorily active
	Required *bool `json:"required,omitempty"`
	// Status
	Status *bool `json:"status,omitempty"`
}

func (o *ReadyzEngines) GetEngineID() *string {
	if o == nil {
		return nil
	}
	return o.EngineID
}

func (o *ReadyzEngines) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *ReadyzEngines) GetStatus() *bool {
	if o == nil {
		return nil
	}
	return o.Status
}

// Status - License status
type Status string

const (
	StatusExpired Status = "expired"
	StatusInvalid Status = "invalid"
	StatusOk      Status = "ok"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "expired":
		fallthrough
	case "invalid":
		fallthrough
	case "ok":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// License status
type License struct {
	// License status
	Status *Status `json:"status,omitempty"`
}

func (o *License) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

// NumberActiveAvEngines - Number of active AV engines
type NumberActiveAvEngines struct {
	// Status
	Status *bool `json:"status,omitempty"`
	// User-defined threshold
	Threshold *int64 `json:"threshold,omitempty"`
}

func (o *NumberActiveAvEngines) GetStatus() *bool {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *NumberActiveAvEngines) GetThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.Threshold
}

// ScanQueue - Scan queue status
type ScanQueue struct {
	// Maximum slots in scan queue
	MaxNumberInQueue *int64 `json:"max_number_in_queue,omitempty"`
	// Current queue status
	NumberInQueue *int64 `json:"number_in_queue,omitempty"`
	// status
	Status *bool `json:"status,omitempty"`
	// User-defined threshold
	Threshold *int64 `json:"threshold,omitempty"`
}

func (o *ScanQueue) GetMaxNumberInQueue() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxNumberInQueue
}

func (o *ScanQueue) GetNumberInQueue() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberInQueue
}

func (o *ScanQueue) GetStatus() *bool {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ScanQueue) GetThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.Threshold
}

type Readyz struct {
	// Database connection
	Database *Database `json:"database,omitempty"`
	// All available engines according to the activated license
	Engines []ReadyzEngines `json:"engines,omitempty"`
	// License status
	License *License `json:"license,omitempty"`
	// Number of active AV engines
	NumberActiveAvEngines *NumberActiveAvEngines `json:"number_active_av_engines,omitempty"`
	// Scan queue status
	ScanQueue *ScanQueue `json:"scan_queue,omitempty"`
}

func (o *Readyz) GetDatabase() *Database {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *Readyz) GetEngines() []ReadyzEngines {
	if o == nil {
		return nil
	}
	return o.Engines
}

func (o *Readyz) GetLicense() *License {
	if o == nil {
		return nil
	}
	return o.License
}

func (o *Readyz) GetNumberActiveAvEngines() *NumberActiveAvEngines {
	if o == nil {
		return nil
	}
	return o.NumberActiveAvEngines
}

func (o *Readyz) GetScanQueue() *ScanQueue {
	if o == nil {
		return nil
	}
	return o.ScanQueue
}
