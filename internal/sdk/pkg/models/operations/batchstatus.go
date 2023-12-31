// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type BatchStatusRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey *string `header:"style=simple,explode=false,name=apikey"`
	// The batch identifier used to submit files in the batch and to close the batch.
	BatchID string `pathParam:"style=simple,explode=false,name=batchId"`
	// The first item order in the list of files in this batch
	First *int64 `queryParam:"style=form,explode=true,name=first"`
	// The number of items to be fetched next, counting from the item order indicated in first header
	Size *int64 `queryParam:"style=form,explode=true,name=size"`
}

func (o *BatchStatusRequest) GetApikey() *string {
	if o == nil {
		return nil
	}
	return o.Apikey
}

func (o *BatchStatusRequest) GetBatchID() string {
	if o == nil {
		return ""
	}
	return o.BatchID
}

func (o *BatchStatusRequest) GetFirst() *int64 {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *BatchStatusRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

// BatchStatusBatchResponse500ResponseBody - Unexpected event on server
type BatchStatusBatchResponse500ResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *BatchStatusBatchResponse500ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// BatchStatusBatchResponseResponseBody - Requests resource was not found.
type BatchStatusBatchResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *BatchStatusBatchResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// BatchStatusBatchResponseBody - Invalid user information or Not Allowed
type BatchStatusBatchResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *BatchStatusBatchResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// BatchStatusResponseBody - Bad Request (e.g. invalid header, apikey is missing or invalid).
type BatchStatusResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *BatchStatusResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type BatchStatusResponse struct {
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	FourHundredApplicationJSONObject *BatchStatusResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *BatchStatusBatchResponseBody
	// Requests resource was not found.
	FourHundredAndFourApplicationJSONObject *BatchStatusBatchResponseResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *BatchStatusBatchResponse500ResponseBody
	// Batch progress paginated report (50 entries/page).
	BatchResponseInProgress *shared.BatchResponseInProgress
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *BatchStatusResponse) GetFourHundredApplicationJSONObject() *BatchStatusResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *BatchStatusResponse) GetFourHundredAndThreeApplicationJSONObject() *BatchStatusBatchResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *BatchStatusResponse) GetFourHundredAndFourApplicationJSONObject() *BatchStatusBatchResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *BatchStatusResponse) GetFiveHundredApplicationJSONObject() *BatchStatusBatchResponse500ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *BatchStatusResponse) GetBatchResponseInProgress() *shared.BatchResponseInProgress {
	if o == nil {
		return nil
	}
	return o.BatchResponseInProgress
}

func (o *BatchStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BatchStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BatchStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
