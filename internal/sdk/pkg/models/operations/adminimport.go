// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminImportRequest struct {
	RequestBody *string `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// AdminImport500ApplicationJSON - Unexpected event on server
type AdminImport500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// AdminImport405ApplicationJSON - The user has no rights for this operation.
type AdminImport405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// AdminImport403ApplicationJSON - Invalid user information or Not Allowed
type AdminImport403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// AdminImport400ApplicationJSON - Bad Request (e.g. invalid header, apikey is missing or invalid).
type AdminImport400ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// AdminImport304ApplicationJSON - The configuration has not changed
type AdminImport304ApplicationJSON struct {
}

// AdminImport200ApplicationJSON - Request processed successfully
type AdminImport200ApplicationJSON struct {
	Result *string `json:"result,omitempty"`
}

type AdminImportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request processed successfully
	AdminImport200ApplicationJSONObject *AdminImport200ApplicationJSON
	// The configuration has not changed
	AdminImport304ApplicationJSONObject *AdminImport304ApplicationJSON
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	AdminImport400ApplicationJSONObject *AdminImport400ApplicationJSON
	// Invalid user information or Not Allowed
	AdminImport403ApplicationJSONObject *AdminImport403ApplicationJSON
	// The user has no rights for this operation.
	AdminImport405ApplicationJSONObject *AdminImport405ApplicationJSON
	// Unexpected event on server
	AdminImport500ApplicationJSONObject *AdminImport500ApplicationJSON
}