// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type YaraPackageGenerateRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *YaraPackageGenerateRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// YaraPackageGenerateYaraResponseResponseBody - Unexpected event on server
type YaraPackageGenerateYaraResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *YaraPackageGenerateYaraResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// YaraPackageGenerateYaraResponseBody - Invalid user information or Not Allowed
type YaraPackageGenerateYaraResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *YaraPackageGenerateYaraResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type General struct {
	Message  *string `json:"message,omitempty"`
	Severity *string `json:"severity,omitempty"`
}

func (o *General) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *General) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

type Source struct {
	Message  *string `json:"message,omitempty"`
	Severity *string `json:"severity,omitempty"`
}

func (o *Source) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *Source) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

// Issues - Stores a map of issues. Each key represents the according source, except "general", which contains general errors occurred during the generation process
type Issues struct {
	// Contains general errors occurred during the generation process
	General []General `json:"general,omitempty"`
	Source  []Source  `json:"source,omitempty"`
}

func (o *Issues) GetGeneral() []General {
	if o == nil {
		return nil
	}
	return o.General
}

func (o *Issues) GetSource() []Source {
	if o == nil {
		return nil
	}
	return o.Source
}

// Status - Current status of generating the package.
type Status string

const (
	StatusIdle       Status = "idle"
	StatusError      Status = "error"
	StatusInprogress Status = "inprogress"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
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
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// YaraPackageGenerateResponseBody - Request processed successfully.
type YaraPackageGenerateResponseBody struct {
	// Stores a map of issues. Each key represents the according source, except "general", which contains general errors occurred during the generation process
	Issues *Issues `json:"issues,omitempty"`
	// Used only when status is inprogress, otherwise its empty
	StartTime *string `json:"start_time,omitempty"`
	// Current status of generating the package.
	Status *Status `json:"status,omitempty"`
}

func (o *YaraPackageGenerateResponseBody) GetIssues() *Issues {
	if o == nil {
		return nil
	}
	return o.Issues
}

func (o *YaraPackageGenerateResponseBody) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *YaraPackageGenerateResponseBody) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type YaraPackageGenerateResponse struct {
	// Request processed successfully.
	TwoHundredApplicationJSONObject *YaraPackageGenerateResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *YaraPackageGenerateYaraResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *YaraPackageGenerateYaraResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *YaraPackageGenerateResponse) GetTwoHundredApplicationJSONObject() *YaraPackageGenerateResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *YaraPackageGenerateResponse) GetFourHundredAndThreeApplicationJSONObject() *YaraPackageGenerateYaraResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *YaraPackageGenerateResponse) GetFiveHundredApplicationJSONObject() *YaraPackageGenerateYaraResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *YaraPackageGenerateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *YaraPackageGenerateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *YaraPackageGenerateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
