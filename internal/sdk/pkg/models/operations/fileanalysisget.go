// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk/pkg/models/shared"
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

func (o *FileAnalysisGetRequest) GetApikey() *string {
	if o == nil {
		return nil
	}
	return o.Apikey
}

func (o *FileAnalysisGetRequest) GetDataID() string {
	if o == nil {
		return ""
	}
	return o.DataID
}

func (o *FileAnalysisGetRequest) GetFirst() *int64 {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *FileAnalysisGetRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *FileAnalysisGetRequest) GetUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.UserAgent
}

// FileAnalysisGetResponseBody - Unexpected event on server
type FileAnalysisGetResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *FileAnalysisGetResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
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
	Object *FileAnalysisGetResponseBody
}

func (o *FileAnalysisGetResponse) GetAnalysisResult() *shared.AnalysisResult {
	if o == nil {
		return nil
	}
	return o.AnalysisResult
}

func (o *FileAnalysisGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FileAnalysisGetResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *FileAnalysisGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FileAnalysisGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *FileAnalysisGetResponse) GetObject() *FileAnalysisGetResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
