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

func (o *ConfigWebhookRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// ConfigWebhook500ApplicationJSON - Unexpected event on server
type ConfigWebhook500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigWebhook500ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigWebhook405ApplicationJSON - The user has no rights for this operation.
type ConfigWebhook405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigWebhook405ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigWebhook403ApplicationJSON - Invalid user information or Not Allowed
type ConfigWebhook403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigWebhook403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
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

func (o *ConfigWebhookResponse) GetAdminConfigWebhook() *shared.AdminConfigWebhook {
	if o == nil {
		return nil
	}
	return o.AdminConfigWebhook
}

func (o *ConfigWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigWebhookResponse) GetConfigWebhook403ApplicationJSONObject() *ConfigWebhook403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigWebhook403ApplicationJSONObject
}

func (o *ConfigWebhookResponse) GetConfigWebhook405ApplicationJSONObject() *ConfigWebhook405ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigWebhook405ApplicationJSONObject
}

func (o *ConfigWebhookResponse) GetConfigWebhook500ApplicationJSONObject() *ConfigWebhook500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigWebhook500ApplicationJSONObject
}
