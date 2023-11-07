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

func (o *ConfigGetSessioncookieRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// ConfigGetSessioncookieConfigResponseResponseBody - Unexpected event on server
type ConfigGetSessioncookieConfigResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigGetSessioncookieConfigResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigGetSessioncookieConfigResponseBody - The user has no rights for this operation.
type ConfigGetSessioncookieConfigResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigGetSessioncookieConfigResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigGetSessioncookieResponseBody - Invalid user information or Not Allowed
type ConfigGetSessioncookieResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigGetSessioncookieResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type ConfigGetSessioncookieResponse struct {
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *ConfigGetSessioncookieResponseBody
	// The user has no rights for this operation.
	FourHundredAndFiveApplicationJSONObject *ConfigGetSessioncookieConfigResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *ConfigGetSessioncookieConfigResponseResponseBody
	// Request processed successfully.
	AdminConfigSessioncookie *shared.AdminConfigSessioncookie
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ConfigGetSessioncookieResponse) GetFourHundredAndThreeApplicationJSONObject() *ConfigGetSessioncookieResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *ConfigGetSessioncookieResponse) GetFourHundredAndFiveApplicationJSONObject() *ConfigGetSessioncookieConfigResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFiveApplicationJSONObject
}

func (o *ConfigGetSessioncookieResponse) GetFiveHundredApplicationJSONObject() *ConfigGetSessioncookieConfigResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *ConfigGetSessioncookieResponse) GetAdminConfigSessioncookie() *shared.AdminConfigSessioncookie {
	if o == nil {
		return nil
	}
	return o.AdminConfigSessioncookie
}

func (o *ConfigGetSessioncookieResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigGetSessioncookieResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigGetSessioncookieResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
