// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigDeleteSkipHashRequest struct {
	// A list of to-be-deleted hashes.
	RequestBody []string `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *ConfigDeleteSkipHashRequest) GetRequestBody() []string {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *ConfigDeleteSkipHashRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// ConfigDeleteSkipHashConfigResponseResponseBody - Unexpected event on server
type ConfigDeleteSkipHashConfigResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigDeleteSkipHashConfigResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigDeleteSkipHashConfigResponseBody - Requests resource was not found.
type ConfigDeleteSkipHashConfigResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigDeleteSkipHashConfigResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigDeleteSkipHashResponseBody - Invalid user information or Not Allowed
type ConfigDeleteSkipHashResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigDeleteSkipHashResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type ConfigDeleteSkipHashResponse struct {
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *ConfigDeleteSkipHashResponseBody
	// Requests resource was not found.
	FourHundredAndFourApplicationJSONObject *ConfigDeleteSkipHashConfigResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *ConfigDeleteSkipHashConfigResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// A list of all skip/white/black-listed hashes.
	SkipListAfterDeleted *shared.SkipListAfterDeleted
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ConfigDeleteSkipHashResponse) GetFourHundredAndThreeApplicationJSONObject() *ConfigDeleteSkipHashResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *ConfigDeleteSkipHashResponse) GetFourHundredAndFourApplicationJSONObject() *ConfigDeleteSkipHashConfigResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *ConfigDeleteSkipHashResponse) GetFiveHundredApplicationJSONObject() *ConfigDeleteSkipHashConfigResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *ConfigDeleteSkipHashResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigDeleteSkipHashResponse) GetSkipListAfterDeleted() *shared.SkipListAfterDeleted {
	if o == nil {
		return nil
	}
	return o.SkipListAfterDeleted
}

func (o *ConfigDeleteSkipHashResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigDeleteSkipHashResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
