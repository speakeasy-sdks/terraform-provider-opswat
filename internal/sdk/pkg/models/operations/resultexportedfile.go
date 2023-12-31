// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ResultExportedFileRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
	// The data_id comes from the result of `Analyze a file`. In case of sanitizing the content of an archive, the data_id of contained file can be found in `Fetch analysis result`.
	DataID string `pathParam:"style=simple,explode=false,name=data_id"`
}

func (o *ResultExportedFileRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

func (o *ResultExportedFileRequest) GetDataID() string {
	if o == nil {
		return ""
	}
	return o.DataID
}

// ResultExportedFileAnalysisResponseResponseBody - Unexpected event on server
type ResultExportedFileAnalysisResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ResultExportedFileAnalysisResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ResultExportedFileAnalysisResponseBody - Requests resource was not found.
type ResultExportedFileAnalysisResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ResultExportedFileAnalysisResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// ResultExportedFileResponseBody - Invalid user information or Not Allowed
type ResultExportedFileResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *ResultExportedFileResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type ResultExportedFileResponse struct {
	// Returns the exported result PDF file
	TwoHundredApplicationPdfBytes []byte
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *ResultExportedFileResponseBody
	// Requests resource was not found.
	FourHundredAndFourApplicationJSONObject *ResultExportedFileAnalysisResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *ResultExportedFileAnalysisResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ResultExportedFileResponse) GetTwoHundredApplicationPdfBytes() []byte {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationPdfBytes
}

func (o *ResultExportedFileResponse) GetFourHundredAndThreeApplicationJSONObject() *ResultExportedFileResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *ResultExportedFileResponse) GetFourHundredAndFourApplicationJSONObject() *ResultExportedFileAnalysisResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *ResultExportedFileResponse) GetFiveHundredApplicationJSONObject() *ResultExportedFileAnalysisResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *ResultExportedFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ResultExportedFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ResultExportedFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
