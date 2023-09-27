// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigGetCustomResponseHeaderRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// ConfigGetCustomResponseHeader500ApplicationJSON - Unexpected event on server
type ConfigGetCustomResponseHeader500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigGetCustomResponseHeader405ApplicationJSON - The user has no rights for this operation.
type ConfigGetCustomResponseHeader405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigGetCustomResponseHeader403ApplicationJSON - Invalid user information or Not Allowed
type ConfigGetCustomResponseHeader403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

type ConfigGetCustomResponseHeaderResponse struct {
	// Request processed successfully.
	AdminConfigCustomResponseHeader *shared.AdminConfigCustomResponseHeader
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	ConfigGetCustomResponseHeader403ApplicationJSONObject *ConfigGetCustomResponseHeader403ApplicationJSON
	// The user has no rights for this operation.
	ConfigGetCustomResponseHeader405ApplicationJSONObject *ConfigGetCustomResponseHeader405ApplicationJSON
	// Unexpected event on server
	ConfigGetCustomResponseHeader500ApplicationJSONObject *ConfigGetCustomResponseHeader500ApplicationJSON
}
