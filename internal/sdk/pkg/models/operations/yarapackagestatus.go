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

// YaraPackageStatusYaraResponseResponseBody - Unexpected event on server
type YaraPackageStatusYaraResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *YaraPackageStatusYaraResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// YaraPackageStatusYaraResponseBody - Invalid user information or Not Allowed
type YaraPackageStatusYaraResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *YaraPackageStatusYaraResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type YaraPackageStatusGeneral struct {
	Message  *string `json:"message,omitempty"`
	Severity *string `json:"severity,omitempty"`
}

func (o *YaraPackageStatusGeneral) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *YaraPackageStatusGeneral) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

type YaraPackageStatusSource struct {
	Message  *string `json:"message,omitempty"`
	Severity *string `json:"severity,omitempty"`
}

func (o *YaraPackageStatusSource) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *YaraPackageStatusSource) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

// YaraPackageStatusIssues - Stores a map of issues. Each key represents the according source, except "general", which contains general errors occurred during the generation process
type YaraPackageStatusIssues struct {
	// Contains general errors occurred during the generation process
	General []YaraPackageStatusGeneral `json:"general,omitempty"`
	Source  []YaraPackageStatusSource  `json:"source,omitempty"`
}

func (o *YaraPackageStatusIssues) GetGeneral() []YaraPackageStatusGeneral {
	if o == nil {
		return nil
	}
	return o.General
}

func (o *YaraPackageStatusIssues) GetSource() []YaraPackageStatusSource {
	if o == nil {
		return nil
	}
	return o.Source
}

// YaraPackageStatusStatus - Current status of generating the package.
type YaraPackageStatusStatus string

const (
	YaraPackageStatusStatusIdle       YaraPackageStatusStatus = "idle"
	YaraPackageStatusStatusError      YaraPackageStatusStatus = "error"
	YaraPackageStatusStatusInprogress YaraPackageStatusStatus = "inprogress"
)

func (e YaraPackageStatusStatus) ToPointer() *YaraPackageStatusStatus {
	return &e
}

func (e *YaraPackageStatusStatus) UnmarshalJSON(data []byte) error {
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
		*e = YaraPackageStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for YaraPackageStatusStatus: %v", v)
	}
}

// YaraPackageStatusResponseBody - Request processed successfully.
type YaraPackageStatusResponseBody struct {
	// Stores a map of issues. Each key represents the according source, except "general", which contains general errors occurred during the generation process
	Issues *YaraPackageStatusIssues `json:"issues,omitempty"`
	// Used only when status is inprogress, otherwise its empty
	StartTime *string `json:"start_time,omitempty"`
	// Current status of generating the package.
	Status *YaraPackageStatusStatus `json:"status,omitempty"`
}

func (o *YaraPackageStatusResponseBody) GetIssues() *YaraPackageStatusIssues {
	if o == nil {
		return nil
	}
	return o.Issues
}

func (o *YaraPackageStatusResponseBody) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *YaraPackageStatusResponseBody) GetStatus() *YaraPackageStatusStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type YaraPackageStatusResponse struct {
	// Request processed successfully.
	TwoHundredApplicationJSONObject *YaraPackageStatusResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *YaraPackageStatusYaraResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *YaraPackageStatusYaraResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *YaraPackageStatusResponse) GetTwoHundredApplicationJSONObject() *YaraPackageStatusResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *YaraPackageStatusResponse) GetFourHundredAndThreeApplicationJSONObject() *YaraPackageStatusYaraResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *YaraPackageStatusResponse) GetFiveHundredApplicationJSONObject() *YaraPackageStatusYaraResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
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
