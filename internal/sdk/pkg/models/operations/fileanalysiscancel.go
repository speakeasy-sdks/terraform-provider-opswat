// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FileAnalysisCancelRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey *string `header:"style=simple,explode=false,name=apikey"`
	// Unique submission identifier.
	// Use this value to reference the submission.
	//
	DataID string `pathParam:"style=simple,explode=false,name=data_id"`
}

// FileAnalysisCancel500ApplicationJSON - Unexpected event on server
type FileAnalysisCancel500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// FileAnalysisCancel405ApplicationJSON - The user has no rights for this operation.
type FileAnalysisCancel405ApplicationJSON struct {
	Err *string `json:"err,omitempty"`
}

// FileAnalysisCancel404ApplicationJSON - Data ID not found (invalid id) or Requests resource was not found
type FileAnalysisCancel404ApplicationJSON struct {
	Err *string `json:"err,omitempty"`
}

// FileAnalysisCancel403ApplicationJSON - Invalid user information or Not Allowed
type FileAnalysisCancel403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// FileAnalysisCancel400ApplicationJSON - Bad Request (e.g. invalid header, apikey is missing or invalid).
type FileAnalysisCancel400ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// FileAnalysisCancel200ApplicationJSON - Analysis was sucessfully cancelled.
type FileAnalysisCancel200ApplicationJSON struct {
}

type FileAnalysisCancelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Analysis was sucessfully cancelled.
	FileAnalysisCancel200ApplicationJSONObject *FileAnalysisCancel200ApplicationJSON
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	FileAnalysisCancel400ApplicationJSONObject *FileAnalysisCancel400ApplicationJSON
	// Invalid user information or Not Allowed
	FileAnalysisCancel403ApplicationJSONObject *FileAnalysisCancel403ApplicationJSON
	// Data ID not found (invalid id) or Requests resource was not found
	FileAnalysisCancel404ApplicationJSONObject *FileAnalysisCancel404ApplicationJSON
	// The user has no rights for this operation.
	FileAnalysisCancel405ApplicationJSONObject *FileAnalysisCancel405ApplicationJSON
	// Unexpected event on server
	FileAnalysisCancel500ApplicationJSONObject *FileAnalysisCancel500ApplicationJSON
}
