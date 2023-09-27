// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigPutSessioncookieRequest struct {
	AdminConfigSessioncookie *shared.AdminConfigSessioncookie `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// ConfigPutSessioncookie500ApplicationJSON - Unexpected event on server
type ConfigPutSessioncookie500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigPutSessioncookie404ApplicationJSON - Requests resource was not found.
type ConfigPutSessioncookie404ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigPutSessioncookie403ApplicationJSON - Invalid user information or Not Allowed
type ConfigPutSessioncookie403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

type ConfigPutSessioncookieResponse struct {
	// Request processed successfully.
	AdminConfigSessioncookie *shared.AdminConfigSessioncookie
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	ConfigPutSessioncookie403ApplicationJSONObject *ConfigPutSessioncookie403ApplicationJSON
	// Requests resource was not found.
	ConfigPutSessioncookie404ApplicationJSONObject *ConfigPutSessioncookie404ApplicationJSON
	// Unexpected event on server
	ConfigPutSessioncookie500ApplicationJSONObject *ConfigPutSessioncookie500ApplicationJSON
}