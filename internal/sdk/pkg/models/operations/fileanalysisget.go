// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type FileAnalysisGetRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey *string `header:"style=simple,explode=false,name=apikey"`
	// Unique submission identifier.
	// Use this value to reference the submission.
	//
	DataID string `pathParam:"style=simple,explode=false,name=data_id"`
	// The first item order in the list child files of archive file
	First *int64 `queryParam:"style=form,explode=true,name=first"`
	// The number of items to be fetched next, counting from the item order indicated in first header
	Size *int64 `queryParam:"style=form,explode=true,name=size"`
	// user_agent header used to identify (and limit) access to a particular rule. For rule selection, `rule` header should be used.
	// Since MetaDefender Core 5.6.0, when user agent is matched the predefined setting in workflow rule,  the product will return X-Core-Id containing identity of MetaDefender Core server as a header of response.
	//
	UserAgent *string `header:"style=simple,explode=false,name=user_agent"`
}

// FileAnalysisGet500ApplicationJSON - Unexpected event on server
type FileAnalysisGet500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

type FileAnalysisGetResponse struct {
	// Entire analysis report generated by MetaDefender Core
	AnalysisResult *shared.AnalysisResult
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unexpected event on server
	FileAnalysisGet500ApplicationJSONObject *FileAnalysisGet500ApplicationJSON
}
