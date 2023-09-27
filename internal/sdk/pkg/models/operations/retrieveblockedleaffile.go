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

// RetrieveBlockedLeafFile500ApplicationJSON - Unexpected event on server
type RetrieveBlockedLeafFile500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// RetrieveBlockedLeafFile405ApplicationJSON - The user has no rights for this operation.
type RetrieveBlockedLeafFile405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// RetrieveBlockedLeafFile404ApplicationJSON - Requests resource was not found.
type RetrieveBlockedLeafFile404ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// RetrieveBlockedLeafFile200ApplicationOctetStreamDetailsProcessInfoPostProcessing - Processing information
type RetrieveBlockedLeafFile200ApplicationOctetStreamDetailsProcessInfoPostProcessing struct {
	// Empty string if no action failed or list of failed actions, separated by "|".
	ActionsFailed *string
	// List of successful actions, separated by "|". Empty string if otherwise.
	ActionsRan *string
}

// RetrieveBlockedLeafFile200ApplicationOctetStreamDetailsProcessInfo - Processing information
type RetrieveBlockedLeafFile200ApplicationOctetStreamDetailsProcessInfo struct {
	// Processing information
	PostProcessing *RetrieveBlockedLeafFile200ApplicationOctetStreamDetailsProcessInfoPostProcessing
}

type RetrieveBlockedLeafFile200ApplicationOctetStreamDetails struct {
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
	ProcessInfo *RetrieveBlockedLeafFile200ApplicationOctetStreamDetailsProcessInfo
	// Aggregated list of potential issues of this blocked leaf file.
	Verdicts []string
}

// RetrieveBlockedLeafFile200ApplicationOctetStream - Returns the list of blocked leaf files.
type RetrieveBlockedLeafFile200ApplicationOctetStream struct {
	Details []RetrieveBlockedLeafFile200ApplicationOctetStreamDetails
	// Indicates if there're indeed over 100 blocked leaf files.
	LimitReached *bool
	// Total number of files in the above array.
	Total *int64
}

type RetrieveBlockedLeafFileResponse struct {
	Body []byte
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Requests resource was not found.
	RetrieveBlockedLeafFile404ApplicationJSONObject *RetrieveBlockedLeafFile404ApplicationJSON
	// The user has no rights for this operation.
	RetrieveBlockedLeafFile405ApplicationJSONObject *RetrieveBlockedLeafFile405ApplicationJSON
	// Unexpected event on server
	RetrieveBlockedLeafFile500ApplicationJSONObject *RetrieveBlockedLeafFile500ApplicationJSON
}