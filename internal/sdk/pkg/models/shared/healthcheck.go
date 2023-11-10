// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Engines struct {
	DisplayName *string `json:"display_name,omitempty"`
	ID          *string `json:"id,omitempty"`
}

func (o *Engines) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Engines) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// HealthCheck - Health check criterias.
type HealthCheck struct {
	Database              *bool     `json:"database,omitempty"`
	EnableFeature         *bool     `json:"enable_feature,omitempty"`
	Engines               []Engines `json:"engines,omitempty"`
	License               *bool     `json:"license,omitempty"`
	NumberActiveAvEngines *int64    `json:"number_active_av_engines,omitempty"`
	ScanQueue             *int64    `json:"scan_queue,omitempty"`
}

func (o *HealthCheck) GetDatabase() *bool {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *HealthCheck) GetEnableFeature() *bool {
	if o == nil {
		return nil
	}
	return o.EnableFeature
}

func (o *HealthCheck) GetEngines() []Engines {
	if o == nil {
		return nil
	}
	return o.Engines
}

func (o *HealthCheck) GetLicense() *bool {
	if o == nil {
		return nil
	}
	return o.License
}

func (o *HealthCheck) GetNumberActiveAvEngines() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberActiveAvEngines
}

func (o *HealthCheck) GetScanQueue() *int64 {
	if o == nil {
		return nil
	}
	return o.ScanQueue
}
