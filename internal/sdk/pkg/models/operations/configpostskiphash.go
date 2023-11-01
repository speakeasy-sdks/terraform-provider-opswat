// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigPostSkipHashRequest struct {
	// A list of new hashes to be added to "skip-by-hash".
	SkipList *shared.SkipList `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *ConfigPostSkipHashRequest) GetSkipList() *shared.SkipList {
	if o == nil {
		return nil
	}
	return o.SkipList
}

func (o *ConfigPostSkipHashRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// ConfigPostSkipHash500ApplicationJSON - Unexpected event on server
type ConfigPostSkipHash500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigPostSkipHash500ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigPostSkipHash404ApplicationJSON - Requests resource was not found.
type ConfigPostSkipHash404ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigPostSkipHash404ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigPostSkipHash403ApplicationJSON - Invalid user information or Not Allowed
type ConfigPostSkipHash403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigPostSkipHash403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type ConfigPostSkipHashResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A list of all skip/white/black-listed hashes.
	SkipList *shared.SkipList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	ConfigPostSkipHash403ApplicationJSONObject *ConfigPostSkipHash403ApplicationJSON
	// Requests resource was not found.
	ConfigPostSkipHash404ApplicationJSONObject *ConfigPostSkipHash404ApplicationJSON
	// Unexpected event on server
	ConfigPostSkipHash500ApplicationJSONObject *ConfigPostSkipHash500ApplicationJSON
}

func (o *ConfigPostSkipHashResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigPostSkipHashResponse) GetSkipList() *shared.SkipList {
	if o == nil {
		return nil
	}
	return o.SkipList
}

func (o *ConfigPostSkipHashResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigPostSkipHashResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigPostSkipHashResponse) GetConfigPostSkipHash403ApplicationJSONObject() *ConfigPostSkipHash403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigPostSkipHash403ApplicationJSONObject
}

func (o *ConfigPostSkipHashResponse) GetConfigPostSkipHash404ApplicationJSONObject() *ConfigPostSkipHash404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigPostSkipHash404ApplicationJSONObject
}

func (o *ConfigPostSkipHashResponse) GetConfigPostSkipHash500ApplicationJSONObject() *ConfigPostSkipHash500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigPostSkipHash500ApplicationJSONObject
}
