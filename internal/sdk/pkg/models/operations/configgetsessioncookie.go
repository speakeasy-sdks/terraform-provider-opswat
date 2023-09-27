// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigGetSessioncookieRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// ConfigGetSessioncookie500ApplicationJSON - Unexpected event on server
type ConfigGetSessioncookie500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigGetSessioncookie405ApplicationJSON - The user has no rights for this operation.
type ConfigGetSessioncookie405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigGetSessioncookie403ApplicationJSON - Invalid user information or Not Allowed
type ConfigGetSessioncookie403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

type ConfigGetSessioncookieResponse struct {
	// Request processed successfully.
	AdminConfigSessioncookie *shared.AdminConfigSessioncookie
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	ConfigGetSessioncookie403ApplicationJSONObject *ConfigGetSessioncookie403ApplicationJSON
	// The user has no rights for this operation.
	ConfigGetSessioncookie405ApplicationJSONObject *ConfigGetSessioncookie405ApplicationJSON
	// Unexpected event on server
	ConfigGetSessioncookie500ApplicationJSONObject *ConfigGetSessioncookie500ApplicationJSON
}