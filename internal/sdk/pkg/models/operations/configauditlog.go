// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ConfigAuditLogRequestBody struct {
	// The number of hours of data retention. Anything older than this number will be deleted.
	Cleanuprange *int64 `json:"cleanuprange,omitempty"`
}

type ConfigAuditLogRequest struct {
	RequestBody *ConfigAuditLogRequestBody `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// ConfigAuditLog500ApplicationJSON - Unexpected event on server
type ConfigAuditLog500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigAuditLog405ApplicationJSON - The user has no rights for this operation.
type ConfigAuditLog405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigAuditLog403ApplicationJSON - Invalid user information or Not Allowed
type ConfigAuditLog403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigAuditLog200ApplicationJSON - Request processed successfully
type ConfigAuditLog200ApplicationJSON struct {
	// The number of hours of data retention. Anything older than this number will be deleted.
	Cleanuprange *int64 `json:"cleanuprange,omitempty"`
}

type ConfigAuditLogResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request processed successfully
	ConfigAuditLog200ApplicationJSONObject *ConfigAuditLog200ApplicationJSON
	// Invalid user information or Not Allowed
	ConfigAuditLog403ApplicationJSONObject *ConfigAuditLog403ApplicationJSON
	// The user has no rights for this operation.
	ConfigAuditLog405ApplicationJSONObject *ConfigAuditLog405ApplicationJSON
	// Unexpected event on server
	ConfigAuditLog500ApplicationJSONObject *ConfigAuditLog500ApplicationJSON
}