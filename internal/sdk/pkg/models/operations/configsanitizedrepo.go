// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ConfigSanitizedRepoRequestBody struct {
	// The number of minutes of data retention. Anything older than this number will be deleted.
	Maxage *int64 `json:"maxage,omitempty"`
}

func (o *ConfigSanitizedRepoRequestBody) GetMaxage() *int64 {
	if o == nil {
		return nil
	}
	return o.Maxage
}

type ConfigSanitizedRepoRequest struct {
	RequestBody *ConfigSanitizedRepoRequestBody `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *ConfigSanitizedRepoRequest) GetRequestBody() *ConfigSanitizedRepoRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *ConfigSanitizedRepoRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// ConfigSanitizedRepo500ApplicationJSON - Unexpected event on server
type ConfigSanitizedRepo500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigSanitizedRepo500ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigSanitizedRepo405ApplicationJSON - The user has no rights for this operation.
type ConfigSanitizedRepo405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigSanitizedRepo405ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigSanitizedRepo403ApplicationJSON - Invalid user information or Not Allowed
type ConfigSanitizedRepo403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigSanitizedRepo403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigSanitizedRepo200ApplicationJSON - Request processed successfully
type ConfigSanitizedRepo200ApplicationJSON struct {
	// The number of minutes of data retention. Anything older than this number will be deleted.
	Maxage *int64 `json:"maxage,omitempty"`
}

func (o *ConfigSanitizedRepo200ApplicationJSON) GetMaxage() *int64 {
	if o == nil {
		return nil
	}
	return o.Maxage
}

type ConfigSanitizedRepoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request processed successfully
	ConfigSanitizedRepo200ApplicationJSONObject *ConfigSanitizedRepo200ApplicationJSON
	// Invalid user information or Not Allowed
	ConfigSanitizedRepo403ApplicationJSONObject *ConfigSanitizedRepo403ApplicationJSON
	// The user has no rights for this operation.
	ConfigSanitizedRepo405ApplicationJSONObject *ConfigSanitizedRepo405ApplicationJSON
	// Unexpected event on server
	ConfigSanitizedRepo500ApplicationJSONObject *ConfigSanitizedRepo500ApplicationJSON
}

func (o *ConfigSanitizedRepoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigSanitizedRepoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigSanitizedRepoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigSanitizedRepoResponse) GetConfigSanitizedRepo200ApplicationJSONObject() *ConfigSanitizedRepo200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigSanitizedRepo200ApplicationJSONObject
}

func (o *ConfigSanitizedRepoResponse) GetConfigSanitizedRepo403ApplicationJSONObject() *ConfigSanitizedRepo403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigSanitizedRepo403ApplicationJSONObject
}

func (o *ConfigSanitizedRepoResponse) GetConfigSanitizedRepo405ApplicationJSONObject() *ConfigSanitizedRepo405ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigSanitizedRepo405ApplicationJSONObject
}

func (o *ConfigSanitizedRepoResponse) GetConfigSanitizedRepo500ApplicationJSONObject() *ConfigSanitizedRepo500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigSanitizedRepo500ApplicationJSONObject
}
