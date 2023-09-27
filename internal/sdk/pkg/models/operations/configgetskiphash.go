// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigGetSkipHashRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// ConfigGetSkipHash500ApplicationJSON - Unexpected event on server
type ConfigGetSkipHash500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigGetSkipHash404ApplicationJSON - Requests resource was not found.
type ConfigGetSkipHash404ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigGetSkipHash403ApplicationJSON - Invalid user information or Not Allowed
type ConfigGetSkipHash403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

type ConfigGetSkipHashResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A list of all skip/white/black-listed hashes.
	SkipList *shared.SkipList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	ConfigGetSkipHash403ApplicationJSONObject *ConfigGetSkipHash403ApplicationJSON
	// Requests resource was not found.
	ConfigGetSkipHash404ApplicationJSONObject *ConfigGetSkipHash404ApplicationJSON
	// Unexpected event on server
	ConfigGetSkipHash500ApplicationJSONObject *ConfigGetSkipHash500ApplicationJSON
}
