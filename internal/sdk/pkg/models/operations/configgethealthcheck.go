// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigGetHealthCheckRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *ConfigGetHealthCheckRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// ConfigGetHealthCheck500ApplicationJSON - Unexpected event on server
type ConfigGetHealthCheck500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigGetHealthCheck500ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigGetHealthCheck403ApplicationJSON - Invalid user information or Not Allowed
type ConfigGetHealthCheck403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigGetHealthCheck403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type ConfigGetHealthCheckResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Health check criterias to determine whether MetaDefender Core server status is healthy or not.
	HealthCheck *shared.HealthCheck
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	ConfigGetHealthCheck403ApplicationJSONObject *ConfigGetHealthCheck403ApplicationJSON
	// Unexpected event on server
	ConfigGetHealthCheck500ApplicationJSONObject *ConfigGetHealthCheck500ApplicationJSON
}

func (o *ConfigGetHealthCheckResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigGetHealthCheckResponse) GetHealthCheck() *shared.HealthCheck {
	if o == nil {
		return nil
	}
	return o.HealthCheck
}

func (o *ConfigGetHealthCheckResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigGetHealthCheckResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigGetHealthCheckResponse) GetConfigGetHealthCheck403ApplicationJSONObject() *ConfigGetHealthCheck403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigGetHealthCheck403ApplicationJSONObject
}

func (o *ConfigGetHealthCheckResponse) GetConfigGetHealthCheck500ApplicationJSONObject() *ConfigGetHealthCheck500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigGetHealthCheck500ApplicationJSONObject
}
