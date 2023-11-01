// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type YaraPackageStatusRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *YaraPackageStatusRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// YaraPackageStatus500ApplicationJSON - Unexpected event on server
type YaraPackageStatus500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *YaraPackageStatus500ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// YaraPackageStatus403ApplicationJSON - Invalid user information or Not Allowed
type YaraPackageStatus403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *YaraPackageStatus403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type YaraPackageStatus200ApplicationJSONIssuesGeneral struct {
	Message  *string `json:"message,omitempty"`
	Severity *string `json:"severity,omitempty"`
}

func (o *YaraPackageStatus200ApplicationJSONIssuesGeneral) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *YaraPackageStatus200ApplicationJSONIssuesGeneral) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

type YaraPackageStatus200ApplicationJSONIssuesSource struct {
	Message  *string `json:"message,omitempty"`
	Severity *string `json:"severity,omitempty"`
}

func (o *YaraPackageStatus200ApplicationJSONIssuesSource) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *YaraPackageStatus200ApplicationJSONIssuesSource) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

// YaraPackageStatus200ApplicationJSONIssues - Stores a map of issues. Each key represents the according source, except "general", which contains general errors occurred during the generation process
type YaraPackageStatus200ApplicationJSONIssues struct {
	// Contains general errors occurred during the generation process
	General []YaraPackageStatus200ApplicationJSONIssuesGeneral `json:"general,omitempty"`
	Source  []YaraPackageStatus200ApplicationJSONIssuesSource  `json:"source,omitempty"`
}

func (o *YaraPackageStatus200ApplicationJSONIssues) GetGeneral() []YaraPackageStatus200ApplicationJSONIssuesGeneral {
	if o == nil {
		return nil
	}
	return o.General
}

func (o *YaraPackageStatus200ApplicationJSONIssues) GetSource() []YaraPackageStatus200ApplicationJSONIssuesSource {
	if o == nil {
		return nil
	}
	return o.Source
}

// YaraPackageStatus200ApplicationJSONStatus - Current status of generating the package.
type YaraPackageStatus200ApplicationJSONStatus string

const (
	YaraPackageStatus200ApplicationJSONStatusIdle       YaraPackageStatus200ApplicationJSONStatus = "idle"
	YaraPackageStatus200ApplicationJSONStatusError      YaraPackageStatus200ApplicationJSONStatus = "error"
	YaraPackageStatus200ApplicationJSONStatusInprogress YaraPackageStatus200ApplicationJSONStatus = "inprogress"
)

func (e YaraPackageStatus200ApplicationJSONStatus) ToPointer() *YaraPackageStatus200ApplicationJSONStatus {
	return &e
}

func (e *YaraPackageStatus200ApplicationJSONStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "idle":
		fallthrough
	case "error":
		fallthrough
	case "inprogress":
		*e = YaraPackageStatus200ApplicationJSONStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for YaraPackageStatus200ApplicationJSONStatus: %v", v)
	}
}

// YaraPackageStatus200ApplicationJSON - Request processed successfully.
type YaraPackageStatus200ApplicationJSON struct {
	// Stores a map of issues. Each key represents the according source, except "general", which contains general errors occurred during the generation process
	Issues *YaraPackageStatus200ApplicationJSONIssues `json:"issues,omitempty"`
	// Used only when status is inprogress, otherwise its empty
	StartTime *string `json:"start_time,omitempty"`
	// Current status of generating the package.
	Status *YaraPackageStatus200ApplicationJSONStatus `json:"status,omitempty"`
}

func (o *YaraPackageStatus200ApplicationJSON) GetIssues() *YaraPackageStatus200ApplicationJSONIssues {
	if o == nil {
		return nil
	}
	return o.Issues
}

func (o *YaraPackageStatus200ApplicationJSON) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *YaraPackageStatus200ApplicationJSON) GetStatus() *YaraPackageStatus200ApplicationJSONStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type YaraPackageStatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request processed successfully.
	YaraPackageStatus200ApplicationJSONObject *YaraPackageStatus200ApplicationJSON
	// Invalid user information or Not Allowed
	YaraPackageStatus403ApplicationJSONObject *YaraPackageStatus403ApplicationJSON
	// Unexpected event on server
	YaraPackageStatus500ApplicationJSONObject *YaraPackageStatus500ApplicationJSON
}

func (o *YaraPackageStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *YaraPackageStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *YaraPackageStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *YaraPackageStatusResponse) GetYaraPackageStatus200ApplicationJSONObject() *YaraPackageStatus200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.YaraPackageStatus200ApplicationJSONObject
}

func (o *YaraPackageStatusResponse) GetYaraPackageStatus403ApplicationJSONObject() *YaraPackageStatus403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.YaraPackageStatus403ApplicationJSONObject
}

func (o *YaraPackageStatusResponse) GetYaraPackageStatus500ApplicationJSONObject() *YaraPackageStatus500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.YaraPackageStatus500ApplicationJSONObject
}
