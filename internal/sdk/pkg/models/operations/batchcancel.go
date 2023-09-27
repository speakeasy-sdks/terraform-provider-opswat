// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type BatchCancelRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey *string `header:"style=simple,explode=false,name=apikey"`
	// The batch identifier used to submit files in the batch and to close the batch.
	BatchID string `pathParam:"style=simple,explode=false,name=batchId"`
}

// BatchCancel500ApplicationJSON - Unexpected event on server
type BatchCancel500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// BatchCancel404ApplicationJSON - Batch not found (invalid id)
type BatchCancel404ApplicationJSON struct {
	Err *string `json:"err,omitempty"`
}

// BatchCancel403ApplicationJSON - Invalid user information or Not Allowed
type BatchCancel403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// BatchCancel400ApplicationJSON - Bad Request (e.g. invalid header, apikey is missing or invalid).
type BatchCancel400ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// BatchCancel200ApplicationJSON - Batch cancelled.
type BatchCancel200ApplicationJSON struct {
}

type BatchCancelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Batch cancelled.
	BatchCancel200ApplicationJSONObject *BatchCancel200ApplicationJSON
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	BatchCancel400ApplicationJSONObject *BatchCancel400ApplicationJSON
	// Invalid user information or Not Allowed
	BatchCancel403ApplicationJSONObject *BatchCancel403ApplicationJSON
	// Batch not found (invalid id)
	BatchCancel404ApplicationJSONObject *BatchCancel404ApplicationJSON
	// Unexpected event on server
	BatchCancel500ApplicationJSONObject *BatchCancel500ApplicationJSON
}