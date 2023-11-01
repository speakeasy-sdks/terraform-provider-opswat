// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type UserCreateRequest struct {
	UserRequest *shared.UserRequest `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *UserCreateRequest) GetUserRequest() *shared.UserRequest {
	if o == nil {
		return nil
	}
	return o.UserRequest
}

func (o *UserCreateRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// UserCreate500ApplicationJSON - Unexpected event on server
type UserCreate500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *UserCreate500ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// UserCreate405ApplicationJSON - The user has no rights for this operation.
type UserCreate405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *UserCreate405ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// UserCreate403ApplicationJSON - Invalid user information or Not Allowed
type UserCreate403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *UserCreate403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// UserCreate400ApplicationJSON - Bad Request (e.g. invalid header, apikey is missing or invalid).
type UserCreate400ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *UserCreate400ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type UserCreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request processed successfully.
	UserResponse *shared.UserResponse
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	UserCreate400ApplicationJSONObject *UserCreate400ApplicationJSON
	// Invalid user information or Not Allowed
	UserCreate403ApplicationJSONObject *UserCreate403ApplicationJSON
	// The user has no rights for this operation.
	UserCreate405ApplicationJSONObject *UserCreate405ApplicationJSON
	// Unexpected event on server
	UserCreate500ApplicationJSONObject *UserCreate500ApplicationJSON
}

func (o *UserCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UserCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UserCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UserCreateResponse) GetUserResponse() *shared.UserResponse {
	if o == nil {
		return nil
	}
	return o.UserResponse
}

func (o *UserCreateResponse) GetUserCreate400ApplicationJSONObject() *UserCreate400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UserCreate400ApplicationJSONObject
}

func (o *UserCreateResponse) GetUserCreate403ApplicationJSONObject() *UserCreate403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UserCreate403ApplicationJSONObject
}

func (o *UserCreateResponse) GetUserCreate405ApplicationJSONObject() *UserCreate405ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UserCreate405ApplicationJSONObject
}

func (o *UserCreateResponse) GetUserCreate500ApplicationJSONObject() *UserCreate500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UserCreate500ApplicationJSONObject
}
