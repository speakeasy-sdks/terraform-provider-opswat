// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk/pkg/models/shared"
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

// ConfigPostSkipHashConfigResponseResponseBody - Unexpected event on server
type ConfigPostSkipHashConfigResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigPostSkipHashConfigResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigPostSkipHashConfigResponseBody - Requests resource was not found.
type ConfigPostSkipHashConfigResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigPostSkipHashConfigResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigPostSkipHashResponseBody - Invalid user information or Not Allowed
type ConfigPostSkipHashResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigPostSkipHashResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type ConfigPostSkipHashResponse struct {
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *ConfigPostSkipHashResponseBody
	// Requests resource was not found.
	FourHundredAndFourApplicationJSONObject *ConfigPostSkipHashConfigResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *ConfigPostSkipHashConfigResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// A list of all skip/white/black-listed hashes.
	SkipList *shared.SkipList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ConfigPostSkipHashResponse) GetFourHundredAndThreeApplicationJSONObject() *ConfigPostSkipHashResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *ConfigPostSkipHashResponse) GetFourHundredAndFourApplicationJSONObject() *ConfigPostSkipHashConfigResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *ConfigPostSkipHashResponse) GetFiveHundredApplicationJSONObject() *ConfigPostSkipHashConfigResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
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
