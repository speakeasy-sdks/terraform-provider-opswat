// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminExportRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *AdminExportRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// AdminExportAdminResponse500ResponseBody - Failed to export configuration
type AdminExportAdminResponse500ResponseBody struct {
	Err *string `json:"err,omitempty"`
}

func (o *AdminExportAdminResponse500ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// AdminExportAdminResponseResponseBody - The user has no rights for this operation.
type AdminExportAdminResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *AdminExportAdminResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// AdminExportAdminResponseBody - Invalid user information or Not Allowed
type AdminExportAdminResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *AdminExportAdminResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// AdminExportResponseBody - Bad Request (e.g. invalid header, apikey is missing or invalid).
type AdminExportResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *AdminExportResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type AdminExportResponse struct {
	// Request processed successfully
	TwoHundredApplicationJSONRes *string
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	FourHundredApplicationJSONObject *AdminExportResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *AdminExportAdminResponseBody
	// The user has no rights for this operation.
	FourHundredAndFiveApplicationJSONObject *AdminExportAdminResponseResponseBody
	// Failed to export configuration
	FiveHundredApplicationJSONObject *AdminExportAdminResponse500ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AdminExportResponse) GetTwoHundredApplicationJSONRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONRes
}

func (o *AdminExportResponse) GetFourHundredApplicationJSONObject() *AdminExportResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *AdminExportResponse) GetFourHundredAndThreeApplicationJSONObject() *AdminExportAdminResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *AdminExportResponse) GetFourHundredAndFiveApplicationJSONObject() *AdminExportAdminResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFiveApplicationJSONObject
}

func (o *AdminExportResponse) GetFiveHundredApplicationJSONObject() *AdminExportAdminResponse500ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *AdminExportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AdminExportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AdminExportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
