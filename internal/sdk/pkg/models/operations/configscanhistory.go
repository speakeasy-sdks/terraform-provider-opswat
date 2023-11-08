// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type ConfigScanHistoryRequest struct {
	AdminCleanup *shared.AdminCleanup `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *ConfigScanHistoryRequest) GetAdminCleanup() *shared.AdminCleanup {
	if o == nil {
		return nil
	}
	return o.AdminCleanup
}

func (o *ConfigScanHistoryRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// ConfigScanHistoryConfigResponse500ResponseBody - Unexpected event on server
type ConfigScanHistoryConfigResponse500ResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigScanHistoryConfigResponse500ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigScanHistoryConfigResponseResponseBody - The user has no rights for this operation.
type ConfigScanHistoryConfigResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigScanHistoryConfigResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigScanHistoryConfigResponseBody - Invalid user information or Not Allowed
type ConfigScanHistoryConfigResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ConfigScanHistoryConfigResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ConfigScanHistoryResponseBody - Request processed successfully
type ConfigScanHistoryResponseBody struct {
	// The number of hours of data retention. Anything older than this number will be deleted.
	Cleanuprange *int64 `json:"cleanuprange,omitempty"`
}

func (o *ConfigScanHistoryResponseBody) GetCleanuprange() *int64 {
	if o == nil {
		return nil
	}
	return o.Cleanuprange
}

type ConfigScanHistoryResponse struct {
	// Request processed successfully
	TwoHundredApplicationJSONObject *ConfigScanHistoryResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *ConfigScanHistoryConfigResponseBody
	// The user has no rights for this operation.
	FourHundredAndFiveApplicationJSONObject *ConfigScanHistoryConfigResponseResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *ConfigScanHistoryConfigResponse500ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ConfigScanHistoryResponse) GetTwoHundredApplicationJSONObject() *ConfigScanHistoryResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ConfigScanHistoryResponse) GetFourHundredAndThreeApplicationJSONObject() *ConfigScanHistoryConfigResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *ConfigScanHistoryResponse) GetFourHundredAndFiveApplicationJSONObject() *ConfigScanHistoryConfigResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFiveApplicationJSONObject
}

func (o *ConfigScanHistoryResponse) GetFiveHundredApplicationJSONObject() *ConfigScanHistoryConfigResponse500ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *ConfigScanHistoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigScanHistoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigScanHistoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
