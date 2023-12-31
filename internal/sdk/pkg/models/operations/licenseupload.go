// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type LicenseUploadRequest struct {
	RequestBody *string `request:"mediaType=application/octet-stream"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *LicenseUploadRequest) GetRequestBody() *string {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *LicenseUploadRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// LicenseUploadLicenseResponseResponseBody - Unexpected event on server
type LicenseUploadLicenseResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *LicenseUploadLicenseResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// LicenseUploadLicenseResponseBody - Invalid user information or Not Allowed
type LicenseUploadLicenseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *LicenseUploadLicenseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// LicenseUploadResponseBody - Request processed successfully
type LicenseUploadResponseBody struct {
	Success *bool `json:"success,omitempty"`
}

func (o *LicenseUploadResponseBody) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

type LicenseUploadResponse struct {
	// Request processed successfully
	TwoHundredApplicationJSONObject *LicenseUploadResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *LicenseUploadLicenseResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *LicenseUploadLicenseResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *LicenseUploadResponse) GetTwoHundredApplicationJSONObject() *LicenseUploadResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *LicenseUploadResponse) GetFourHundredAndThreeApplicationJSONObject() *LicenseUploadLicenseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *LicenseUploadResponse) GetFiveHundredApplicationJSONObject() *LicenseUploadLicenseResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *LicenseUploadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LicenseUploadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LicenseUploadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
