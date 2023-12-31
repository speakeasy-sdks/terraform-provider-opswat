// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type LicenseBackupNominationRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *LicenseBackupNominationRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// LicenseBackupNominationAdminResponse500ResponseBody - Unexpected event on server
type LicenseBackupNominationAdminResponse500ResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *LicenseBackupNominationAdminResponse500ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// LicenseBackupNominationAdminResponseResponseBody - Invalid user information or Not Allowed
type LicenseBackupNominationAdminResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *LicenseBackupNominationAdminResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// LicenseBackupNominationAdminResponseBody - Bad Request (e.g. invalid header, apikey is missing or invalid).
type LicenseBackupNominationAdminResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *LicenseBackupNominationAdminResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// LicenseBackupNominationResponseBody - Request successfully
type LicenseBackupNominationResponseBody struct {
	// Successful message
	Success *string `json:"success,omitempty"`
}

func (o *LicenseBackupNominationResponseBody) GetSuccess() *string {
	if o == nil {
		return nil
	}
	return o.Success
}

type LicenseBackupNominationResponse struct {
	// Request successfully
	TwoHundredApplicationJSONObject *LicenseBackupNominationResponseBody
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	FourHundredApplicationJSONObject *LicenseBackupNominationAdminResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *LicenseBackupNominationAdminResponseResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *LicenseBackupNominationAdminResponse500ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *LicenseBackupNominationResponse) GetTwoHundredApplicationJSONObject() *LicenseBackupNominationResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *LicenseBackupNominationResponse) GetFourHundredApplicationJSONObject() *LicenseBackupNominationAdminResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *LicenseBackupNominationResponse) GetFourHundredAndThreeApplicationJSONObject() *LicenseBackupNominationAdminResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *LicenseBackupNominationResponse) GetFiveHundredApplicationJSONObject() *LicenseBackupNominationAdminResponse500ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *LicenseBackupNominationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LicenseBackupNominationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LicenseBackupNominationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
