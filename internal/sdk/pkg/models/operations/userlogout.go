// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UserLogoutRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *UserLogoutRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// UserLogoutAuthResponse500ResponseBody - Unexpected event on server
type UserLogoutAuthResponse500ResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *UserLogoutAuthResponse500ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// UserLogoutAuthResponseResponseBody - Invalid user information.
type UserLogoutAuthResponseResponseBody struct {
	Err string `json:"err"`
}

func (o *UserLogoutAuthResponseResponseBody) GetErr() string {
	if o == nil {
		return ""
	}
	return o.Err
}

// UserLogoutAuthResponseBody - Bad Request.
type UserLogoutAuthResponseBody struct {
	Err string `json:"err"`
}

func (o *UserLogoutAuthResponseBody) GetErr() string {
	if o == nil {
		return ""
	}
	return o.Err
}

// UserLogoutResponseBody - OK
type UserLogoutResponseBody struct {
	Response string `json:"response"`
}

func (o *UserLogoutResponseBody) GetResponse() string {
	if o == nil {
		return ""
	}
	return o.Response
}

type UserLogoutResponse struct {
	// OK
	TwoHundredApplicationJSONObject *UserLogoutResponseBody
	// Bad Request.
	FourHundredApplicationJSONObject *UserLogoutAuthResponseBody
	// Invalid user information.
	FourHundredAndThreeApplicationJSONObject *UserLogoutAuthResponseResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *UserLogoutAuthResponse500ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UserLogoutResponse) GetTwoHundredApplicationJSONObject() *UserLogoutResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *UserLogoutResponse) GetFourHundredApplicationJSONObject() *UserLogoutAuthResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *UserLogoutResponse) GetFourHundredAndThreeApplicationJSONObject() *UserLogoutAuthResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *UserLogoutResponse) GetFiveHundredApplicationJSONObject() *UserLogoutAuthResponse500ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *UserLogoutResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UserLogoutResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UserLogoutResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
