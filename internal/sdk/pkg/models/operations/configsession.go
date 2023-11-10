// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigSessionRequest struct {
	AdminConfigSession *shared.AdminConfigSession `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *ConfigSessionRequest) GetAdminConfigSession() *shared.AdminConfigSession {
	if o == nil {
		return nil
	}
	return o.AdminConfigSession
}

func (o *ConfigSessionRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// ConfigSessionConfigResponseResponseBody - Unexpected event on server
type ConfigSessionConfigResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigSessionConfigResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigSessionConfigResponseBody - The user has no rights for this operation.
type ConfigSessionConfigResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigSessionConfigResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigSessionResponseBody - Invalid user information or Not Allowed
type ConfigSessionResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigSessionResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type ConfigSessionResponse struct {
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *ConfigSessionResponseBody
	// The user has no rights for this operation.
	FourHundredAndFiveApplicationJSONObject *ConfigSessionConfigResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *ConfigSessionConfigResponseResponseBody
	// Request processed successfully
	AdminConfigSession *shared.AdminConfigSession
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ConfigSessionResponse) GetFourHundredAndThreeApplicationJSONObject() *ConfigSessionResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *ConfigSessionResponse) GetFourHundredAndFiveApplicationJSONObject() *ConfigSessionConfigResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFiveApplicationJSONObject
}

func (o *ConfigSessionResponse) GetFiveHundredApplicationJSONObject() *ConfigSessionConfigResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *ConfigSessionResponse) GetAdminConfigSession() *shared.AdminConfigSession {
	if o == nil {
		return nil
	}
	return o.AdminConfigSession
}

func (o *ConfigSessionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigSessionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigSessionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
