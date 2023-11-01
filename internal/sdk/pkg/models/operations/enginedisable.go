// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type EngineDisableRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
	// The unique engine identifier
	EngineID string `pathParam:"style=simple,explode=false,name=engineId"`
}

func (o *EngineDisableRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

func (o *EngineDisableRequest) GetEngineID() string {
	if o == nil {
		return ""
	}
	return o.EngineID
}

// EngineDisable500ApplicationJSON - Unexpected event on server
type EngineDisable500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *EngineDisable500ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// EngineDisable405ApplicationJSON - The user has no rights for this operation.
type EngineDisable405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *EngineDisable405ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// EngineDisable403ApplicationJSON - Invalid user information or Not Allowed
type EngineDisable403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *EngineDisable403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// EngineDisable400ApplicationJSON - Bad Request (e.g. invalid header, apikey is missing or invalid).
type EngineDisable400ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *EngineDisable400ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// EngineDisable200ApplicationJSON - Request processed successfully
type EngineDisable200ApplicationJSON struct {
	Result *string `json:"result,omitempty"`
}

func (o *EngineDisable200ApplicationJSON) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}

type EngineDisableResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request processed successfully
	EngineDisable200ApplicationJSONObject *EngineDisable200ApplicationJSON
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	EngineDisable400ApplicationJSONObject *EngineDisable400ApplicationJSON
	// Invalid user information or Not Allowed
	EngineDisable403ApplicationJSONObject *EngineDisable403ApplicationJSON
	// The user has no rights for this operation.
	EngineDisable405ApplicationJSONObject *EngineDisable405ApplicationJSON
	// Unexpected event on server
	EngineDisable500ApplicationJSONObject *EngineDisable500ApplicationJSON
}

func (o *EngineDisableResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EngineDisableResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EngineDisableResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EngineDisableResponse) GetEngineDisable200ApplicationJSONObject() *EngineDisable200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.EngineDisable200ApplicationJSONObject
}

func (o *EngineDisableResponse) GetEngineDisable400ApplicationJSONObject() *EngineDisable400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.EngineDisable400ApplicationJSONObject
}

func (o *EngineDisableResponse) GetEngineDisable403ApplicationJSONObject() *EngineDisable403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.EngineDisable403ApplicationJSONObject
}

func (o *EngineDisableResponse) GetEngineDisable405ApplicationJSONObject() *EngineDisable405ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.EngineDisable405ApplicationJSONObject
}

func (o *EngineDisableResponse) GetEngineDisable500ApplicationJSONObject() *EngineDisable500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.EngineDisable500ApplicationJSONObject
}
