// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type BatchCreateRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey *string `header:"style=simple,explode=false,name=apikey"`
	// Select rule for the analysis, if no header given the default rule will be selected (URL encoded UTF-8 string of rule name)
	//
	Rule *string `header:"style=simple,explode=false,name=rule"`
	// Name of the batch (max 1024 bytes, URL encoded UTF-8 string).
	//
	UserData *string `header:"style=simple,explode=false,name=user-data"`
	// user_agent header used to identify (and limit) access to a particular rule. For rule selection, `rule` header should be used.
	// Since MetaDefender Core 5.6.0, when user agent is matched the predefined setting in workflow rule,  the product will return X-Core-Id containing identity of MetaDefender Core server as a header of response.
	//
	UserAgent *string `header:"style=simple,explode=false,name=user_agent"`
}

func (o *BatchCreateRequest) GetApikey() *string {
	if o == nil {
		return nil
	}
	return o.Apikey
}

func (o *BatchCreateRequest) GetRule() *string {
	if o == nil {
		return nil
	}
	return o.Rule
}

func (o *BatchCreateRequest) GetUserData() *string {
	if o == nil {
		return nil
	}
	return o.UserData
}

func (o *BatchCreateRequest) GetUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.UserAgent
}

// BatchCreateBatchResponseResponseBody - Unexpected event on server
type BatchCreateBatchResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *BatchCreateBatchResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// BatchCreateBatchResponseBody - Invalid user information or Not Allowed
type BatchCreateBatchResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *BatchCreateBatchResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// BatchCreateResponseBody - Bad Request (e.g. invalid header, apikey is missing or invalid).
type BatchCreateResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *BatchCreateResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type BatchCreateResponse struct {
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	FourHundredApplicationJSONObject *BatchCreateResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *BatchCreateBatchResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *BatchCreateBatchResponseResponseBody
	// Batch created successfully.
	BatchID *shared.BatchID
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *BatchCreateResponse) GetFourHundredApplicationJSONObject() *BatchCreateResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *BatchCreateResponse) GetFourHundredAndThreeApplicationJSONObject() *BatchCreateBatchResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *BatchCreateResponse) GetFiveHundredApplicationJSONObject() *BatchCreateBatchResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *BatchCreateResponse) GetBatchID() *shared.BatchID {
	if o == nil {
		return nil
	}
	return o.BatchID
}

func (o *BatchCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BatchCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BatchCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
