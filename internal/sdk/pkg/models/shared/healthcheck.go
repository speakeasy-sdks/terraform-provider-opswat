// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type HealthCheckEngines struct {
	DisplayName *string `json:"display_name,omitempty"`
	ID          *string `json:"id,omitempty"`
}

// HealthCheck - Health check criterias.
type HealthCheck struct {
	Database              *bool                `json:"database,omitempty"`
	EnableFeature         *bool                `json:"enable_feature,omitempty"`
	Engines               []HealthCheckEngines `json:"engines,omitempty"`
	License               *bool                `json:"license,omitempty"`
	NumberActiveAvEngines *int64               `json:"number_active_av_engines,omitempty"`
	ScanQueue             *int64               `json:"scan_queue,omitempty"`
}
