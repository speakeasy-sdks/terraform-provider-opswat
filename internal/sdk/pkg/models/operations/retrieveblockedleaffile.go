// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RetrieveBlockedLeafFileRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey *string `header:"style=simple,explode=false,name=apikey"`
	// The data_id comes from the result of `Analyze a file`. In case of sanitizing the content of an archive, the data_id of contained file can be found in `Fetch analysis result`.
	DataID string `pathParam:"style=simple,explode=false,name=data_id"`
}

func (o *RetrieveBlockedLeafFileRequest) GetApikey() *string {
	if o == nil {
		return nil
	}
	return o.Apikey
}

func (o *RetrieveBlockedLeafFileRequest) GetDataID() string {
	if o == nil {
		return ""
	}
	return o.DataID
}

// RetrieveBlockedLeafFileAnalysisResponse500ResponseBody - Unexpected event on server
type RetrieveBlockedLeafFileAnalysisResponse500ResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *RetrieveBlockedLeafFileAnalysisResponse500ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// RetrieveBlockedLeafFileAnalysisResponseResponseBody - The user has no rights for this operation.
type RetrieveBlockedLeafFileAnalysisResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *RetrieveBlockedLeafFileAnalysisResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// RetrieveBlockedLeafFileAnalysisResponseBody - Requests resource was not found.
type RetrieveBlockedLeafFileAnalysisResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *RetrieveBlockedLeafFileAnalysisResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// PostProcessing - Processing information
type PostProcessing struct {
	// Empty string if no action failed or list of failed actions, separated by "|".
	ActionsFailed *string
	// List of successful actions, separated by "|". Empty string if otherwise.
	ActionsRan *string
}

func (o *PostProcessing) GetActionsFailed() *string {
	if o == nil {
		return nil
	}
	return o.ActionsFailed
}

func (o *PostProcessing) GetActionsRan() *string {
	if o == nil {
		return nil
	}
	return o.ActionsRan
}

// ProcessInfo - Processing information
type ProcessInfo struct {
	// Processing information
	PostProcessing *PostProcessing
}

func (o *ProcessInfo) GetPostProcessing() *PostProcessing {
	if o == nil {
		return nil
	}
	return o.PostProcessing
}

type Details struct {
	// Provides the reason why the file is blocked in an array (if so).
	BlockedReasons []string
	// Data identifier of this blocked leaf file.
	DataID *string
	// The name of this blocked leaf file.
	DisplayName *string
	// The filetype using mimetype.
	FileType *string
	// Data identifer of the original archive file.
	ParentID *string
	// Full path to this blocked leaf file.
	Path *string
	// Processing information
	ProcessInfo *ProcessInfo
	// Aggregated list of potential issues of this blocked leaf file.
	Verdicts []string
}

func (o *Details) GetBlockedReasons() []string {
	if o == nil {
		return nil
	}
	return o.BlockedReasons
}

func (o *Details) GetDataID() *string {
	if o == nil {
		return nil
	}
	return o.DataID
}

func (o *Details) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Details) GetFileType() *string {
	if o == nil {
		return nil
	}
	return o.FileType
}

func (o *Details) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *Details) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *Details) GetProcessInfo() *ProcessInfo {
	if o == nil {
		return nil
	}
	return o.ProcessInfo
}

func (o *Details) GetVerdicts() []string {
	if o == nil {
		return nil
	}
	return o.Verdicts
}

// RetrieveBlockedLeafFileResponseBody - Returns the list of blocked leaf files.
type RetrieveBlockedLeafFileResponseBody struct {
	Details []Details
	// Indicates if there're indeed over 100 blocked leaf files.
	LimitReached *bool
	// Total number of files in the above array.
	Total *int64
}

func (o *RetrieveBlockedLeafFileResponseBody) GetDetails() []Details {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *RetrieveBlockedLeafFileResponseBody) GetLimitReached() *bool {
	if o == nil {
		return nil
	}
	return o.LimitReached
}

func (o *RetrieveBlockedLeafFileResponseBody) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

type RetrieveBlockedLeafFileResponse struct {
	// Requests resource was not found.
	FourHundredAndFourApplicationJSONObject *RetrieveBlockedLeafFileAnalysisResponseBody
	// The user has no rights for this operation.
	FourHundredAndFiveApplicationJSONObject *RetrieveBlockedLeafFileAnalysisResponseResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *RetrieveBlockedLeafFileAnalysisResponse500ResponseBody
	Body                             []byte
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveBlockedLeafFileResponse) GetFourHundredAndFourApplicationJSONObject() *RetrieveBlockedLeafFileAnalysisResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *RetrieveBlockedLeafFileResponse) GetFourHundredAndFiveApplicationJSONObject() *RetrieveBlockedLeafFileAnalysisResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFiveApplicationJSONObject
}

func (o *RetrieveBlockedLeafFileResponse) GetFiveHundredApplicationJSONObject() *RetrieveBlockedLeafFileAnalysisResponse500ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *RetrieveBlockedLeafFileResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RetrieveBlockedLeafFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveBlockedLeafFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveBlockedLeafFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
