// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigWebhookRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// ConfigWebhook500ApplicationJSON - Unexpected event on server
type ConfigWebhook500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigWebhook405ApplicationJSON - The user has no rights for this operation.
type ConfigWebhook405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigWebhook403ApplicationJSON - Invalid user information or Not Allowed
type ConfigWebhook403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

type ConfigWebhookResponse struct {
	// Request processed successfully.
	AdminConfigWebhook *shared.AdminConfigWebhook
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	ConfigWebhook403ApplicationJSONObject *ConfigWebhook403ApplicationJSON
	// The user has no rights for this operation.
	ConfigWebhook405ApplicationJSONObject *ConfigWebhook405ApplicationJSON
	// Unexpected event on server
	ConfigWebhook500ApplicationJSONObject *ConfigWebhook500ApplicationJSON
}