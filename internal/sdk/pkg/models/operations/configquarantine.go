// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ConfigQuarantineRequestBody struct {
	// The number of hours of data retention. Anything older than this number will be deleted.
	Cleanuprange *int64 `json:"cleanuprange,omitempty"`
}

type ConfigQuarantineRequest struct {
	RequestBody *ConfigQuarantineRequestBody `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// ConfigQuarantine500ApplicationJSON - Unexpected event on server
type ConfigQuarantine500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigQuarantine405ApplicationJSON - The user has no rights for this operation.
type ConfigQuarantine405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigQuarantine403ApplicationJSON - Invalid user information or Not Allowed
type ConfigQuarantine403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigQuarantine200ApplicationJSON - Request processed successfully
type ConfigQuarantine200ApplicationJSON struct {
	// The number of hours of data retention. Anything older than this number will be deleted.
	Cleanuprange *int64 `json:"cleanuprange,omitempty"`
}

type ConfigQuarantineResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request processed successfully
	ConfigQuarantine200ApplicationJSONObject *ConfigQuarantine200ApplicationJSON
	// Invalid user information or Not Allowed
	ConfigQuarantine403ApplicationJSONObject *ConfigQuarantine403ApplicationJSON
	// The user has no rights for this operation.
	ConfigQuarantine405ApplicationJSONObject *ConfigQuarantine405ApplicationJSON
	// Unexpected event on server
	ConfigQuarantine500ApplicationJSONObject *ConfigQuarantine500ApplicationJSON
}
