// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminExportV2Request struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
	// Password to encrypt all the settings after exporting from MetaDefender Core. Only applicable when importing ZIP file.
	//
	Password string `header:"style=simple,explode=false,name=password"`
}

func (o *AdminExportV2Request) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

func (o *AdminExportV2Request) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

// AdminExportV2AdminResponse500ResponseBody - Unexpected event on server
type AdminExportV2AdminResponse500ResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *AdminExportV2AdminResponse500ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// AdminExportV2AdminResponseResponseBody - The user has no rights for this operation.
type AdminExportV2AdminResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *AdminExportV2AdminResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// AdminExportV2AdminResponseBody - Invalid user information or Not Allowed
type AdminExportV2AdminResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *AdminExportV2AdminResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// AdminExportV2ResponseBody - Bad request
type AdminExportV2ResponseBody struct {
	Err *string `json:"err,omitempty"`
}

func (o *AdminExportV2ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type AdminExportV2Response struct {
	// Request processed successfully
	TwoHundredApplicationOctetStreamBytes []byte
	// Bad request
	FourHundredApplicationJSONObject *AdminExportV2ResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *AdminExportV2AdminResponseBody
	// The user has no rights for this operation.
	FourHundredAndFiveApplicationJSONObject *AdminExportV2AdminResponseResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *AdminExportV2AdminResponse500ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AdminExportV2Response) GetTwoHundredApplicationOctetStreamBytes() []byte {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationOctetStreamBytes
}

func (o *AdminExportV2Response) GetFourHundredApplicationJSONObject() *AdminExportV2ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *AdminExportV2Response) GetFourHundredAndThreeApplicationJSONObject() *AdminExportV2AdminResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *AdminExportV2Response) GetFourHundredAndFiveApplicationJSONObject() *AdminExportV2AdminResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFiveApplicationJSONObject
}

func (o *AdminExportV2Response) GetFiveHundredApplicationJSONObject() *AdminExportV2AdminResponse500ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *AdminExportV2Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AdminExportV2Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AdminExportV2Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
