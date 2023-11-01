// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigPutCustomResponseHeaderRequest struct {
	AdminConfigCustomResponseHeader *shared.AdminConfigCustomResponseHeader `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *ConfigPutCustomResponseHeaderRequest) GetAdminConfigCustomResponseHeader() *shared.AdminConfigCustomResponseHeader {
	if o == nil {
		return nil
	}
	return o.AdminConfigCustomResponseHeader
}

func (o *ConfigPutCustomResponseHeaderRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// ConfigPutCustomResponseHeader500ApplicationJSON - Unexpected event on server
type ConfigPutCustomResponseHeader500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigPutCustomResponseHeader500ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigPutCustomResponseHeader404ApplicationJSON - Requests resource was not found.
type ConfigPutCustomResponseHeader404ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigPutCustomResponseHeader404ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigPutCustomResponseHeader403ApplicationJSON - Invalid user information or Not Allowed
type ConfigPutCustomResponseHeader403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigPutCustomResponseHeader403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type ConfigPutCustomResponseHeaderResponse struct {
	// Request processed successfully.
	AdminConfigCustomResponseHeader *shared.AdminConfigCustomResponseHeader
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	ConfigPutCustomResponseHeader403ApplicationJSONObject *ConfigPutCustomResponseHeader403ApplicationJSON
	// Requests resource was not found.
	ConfigPutCustomResponseHeader404ApplicationJSONObject *ConfigPutCustomResponseHeader404ApplicationJSON
	// Unexpected event on server
	ConfigPutCustomResponseHeader500ApplicationJSONObject *ConfigPutCustomResponseHeader500ApplicationJSON
}

func (o *ConfigPutCustomResponseHeaderResponse) GetAdminConfigCustomResponseHeader() *shared.AdminConfigCustomResponseHeader {
	if o == nil {
		return nil
	}
	return o.AdminConfigCustomResponseHeader
}

func (o *ConfigPutCustomResponseHeaderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigPutCustomResponseHeaderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigPutCustomResponseHeaderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigPutCustomResponseHeaderResponse) GetConfigPutCustomResponseHeader403ApplicationJSONObject() *ConfigPutCustomResponseHeader403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigPutCustomResponseHeader403ApplicationJSONObject
}

func (o *ConfigPutCustomResponseHeaderResponse) GetConfigPutCustomResponseHeader404ApplicationJSONObject() *ConfigPutCustomResponseHeader404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigPutCustomResponseHeader404ApplicationJSONObject
}

func (o *ConfigPutCustomResponseHeaderResponse) GetConfigPutCustomResponseHeader500ApplicationJSONObject() *ConfigPutCustomResponseHeader500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigPutCustomResponseHeader500ApplicationJSONObject
}
