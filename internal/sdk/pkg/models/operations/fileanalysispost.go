// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk/pkg/models/shared"
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk/pkg/utils"
	"net/http"
)

type FileAnalysisPostRequest struct {
	RequestBody []byte `request:"mediaType=application/octet-stream"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey *string `header:"style=simple,explode=false,name=apikey"`
	// password for archive ( URL encoded UTF-8 string)
	// Multiple passwords is also supported, format: archivepwdX
	//   * X: Could be empty
	//   * When having value, X must be a number >= 1
	//
	// For example:
	//   * archivepwd1: "fox"
	//   * archivepwd2: "cow"
	//   * archivepwd3: "bear"
	//
	Archivepwd *string `header:"style=simple,explode=false,name=archivepwd"`
	// Batch id to scan with, coming from `Initiate Batch` (If it is not given, it will be a single file scan.)
	Batch *string `header:"style=simple,explode=false,name=batch"`
	// Client's URL where MetaDefender Core will notify scan result back to
	// whenever scan is finished (webhooks model).
	//   * Format: \<protocol://\>\<ip | domain\>:\<port\>\</path\>
	//     * Example: http://10.0.1.100:8081/listenback
	//   * Supported protocol: HTTP / HTTPS
	//   * Method: POST
	//
	Callbackurl *string `header:"style=simple,explode=false,name=callbackurl"`
	// Format: \<protocol://\>\<ip | domain\>:\<port\>\</path\>
	//
	// Since MetaDefender Core 4.19.1, the product started supporting users to process file by just specifying a direct download link.
	//
	// When this header is specified, users must not have body content in HTTP(S) request, otherwise expecting to hit 400 HTTP(S) error. Check response sample of HTTP 400 for details.
	//
	// Specify download link where MetaDefender Core could download the entire payload for processing.
	//   * Supported protocol: HTTP / HTTPS
	//   * Support both
	//     * Individual file scan (both webhook & non-webhook)
	//     * File processing in batch via link
	//       * Simply use `downloadfrom` header for each file scan request in a batch, no further requirement on batch request itself to utilize scan via link feature.
	//       * `Note`: Cancelling a batch will result in auto cancelling all registered individual scan requests which are still in downloading status
	//   * Configurable setting under workflow rule (under "SCAN" tab):
	//     * Download timeout
	//     * Max file size to download
	//   * Support to return status of download progress back to client in HTTP(S) response
	//   * Pre-check supported to refuse downloading if the file size does not meet configured conditions, and/or MetaDefender Core's temp folder does not have sufficient disk space to save the downloaded file
	//
	// Limitation:
	//   * Redirect link is not supported
	//   * Retry to download is not supported
	//
	Downloadfrom *string `header:"style=simple,explode=false,name=downloadfrom"`
	// Since MetaDefender Core 5.0.0, preferred context / verbose information can be sent to the engines.
	//
	// Please see the below pages for the details:
	//   * [File Type engine](https://docs.opswat.com/mdcore/utilities-engines/supported-engines-metadata) (supported since Core 5.2.1)
	//   * [Archive engine](https://docs.opswat.com/mdcore/utilities-engines/supported-engines-metadata-header) (supported since Core 5.4.1)
	//   * [Deep CDR](https://docs.opswat.com/mdcore/deep-cdr/supported-engines-metadata-json)
	//   * [Proactive DLP](https://docs.opswat.com/mdcore/proactive-dlp/supported-engines-metadata-json)
	//
	EnginesMetadata *string `header:"style=simple,explode=false,name=engines-metadata"`
	// The name of the submitted file
	Filename *string `header:"style=simple,explode=false,name=filename"`
	// if local file scan is enabled the path to the file (see Security rule configuration)
	Filepath *string `header:"style=simple,explode=false,name=filepath"`
	// This custom global timeout (in seconds) will override the global timeout predefined in corresponding workflow rule.
	//
	GlobalTimeout *int64 `header:"style=simple,explode=false,name=global-timeout"`
	// could be utilized for:
	//
	// Additional parameter for pre-defined post actions and external scanners
	// (as a part of STDIN input).
	//
	// Customized macro variable for watermarking text (Proactive DLP engine
	// feature).
	//
	// Additional context / verbose information for each file submission
	// (appended into JSON response scan result).
	//
	Metadata *string `header:"style=simple,explode=false,name=metadata"`
	// Select rule for the analysis, if no header given the default rule will be selected (URL encoded UTF-8 string of rule name)
	//
	Rule *string `header:"style=simple,explode=false,name=rule"`
	// Client's URL where MetaDefender Core will send the sanitized file as the body of a POST request back to
	// whenever scan is finished (webhooks model).
	//   * Format: \<protocol://\>\<ip | domain\>:\<port\>\</path\>
	//     * Example: http://10.0.1.100:8081/listensanitizedfile
	//   * Supported protocol: HTTP / HTTPS
	//   * Method: POST
	//   * Mandatory headers:
	//     * `user-agent`: "OPSWAT MetaDefender Core V4"
	//     * `dataid`: "<data id of original file submission>"
	//     * `sanitization-result`: "Sanitized Successfully" or "Sanitized Partially" or "Sanitized File Removed"
	//     * `sanitized-file-hash`: "<SHA256 hash value of sanitized file>"
	//     * `Content-Disposition`: attachment; filename="<name of sanitized file>"
	//     * `Content-Length`: <actual size of sanitized file>
	//
	Sanitizedurl *string `header:"style=simple,explode=false,name=sanitizedurl"`
	// user_agent header used to identify (and limit) access to a particular rule. For rule selection, `rule` header should be used.
	// Since MetaDefender Core 5.6.0, when user agent is matched the predefined setting in workflow rule,  the product will return X-Core-Id containing identity of MetaDefender Core server as a header of response.
	//
	UserAgent *string `header:"style=simple,explode=false,name=user_agent"`
	// name of the selected workflow (deprecated, use "rule" header parameter instead)
	Workflow *string `header:"style=simple,explode=false,name=workflow"`
}

func (o *FileAnalysisPostRequest) GetRequestBody() []byte {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *FileAnalysisPostRequest) GetApikey() *string {
	if o == nil {
		return nil
	}
	return o.Apikey
}

func (o *FileAnalysisPostRequest) GetArchivepwd() *string {
	if o == nil {
		return nil
	}
	return o.Archivepwd
}

func (o *FileAnalysisPostRequest) GetBatch() *string {
	if o == nil {
		return nil
	}
	return o.Batch
}

func (o *FileAnalysisPostRequest) GetCallbackurl() *string {
	if o == nil {
		return nil
	}
	return o.Callbackurl
}

func (o *FileAnalysisPostRequest) GetDownloadfrom() *string {
	if o == nil {
		return nil
	}
	return o.Downloadfrom
}

func (o *FileAnalysisPostRequest) GetEnginesMetadata() *string {
	if o == nil {
		return nil
	}
	return o.EnginesMetadata
}

func (o *FileAnalysisPostRequest) GetFilename() *string {
	if o == nil {
		return nil
	}
	return o.Filename
}

func (o *FileAnalysisPostRequest) GetFilepath() *string {
	if o == nil {
		return nil
	}
	return o.Filepath
}

func (o *FileAnalysisPostRequest) GetGlobalTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.GlobalTimeout
}

func (o *FileAnalysisPostRequest) GetMetadata() *string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *FileAnalysisPostRequest) GetRule() *string {
	if o == nil {
		return nil
	}
	return o.Rule
}

func (o *FileAnalysisPostRequest) GetSanitizedurl() *string {
	if o == nil {
		return nil
	}
	return o.Sanitizedurl
}

func (o *FileAnalysisPostRequest) GetUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.UserAgent
}

func (o *FileAnalysisPostRequest) GetWorkflow() *string {
	if o == nil {
		return nil
	}
	return o.Workflow
}

// FileAnalysisPostAnalysisResponse503ResponseBody - Server is too busy, scan queue is full. Try again later.
type FileAnalysisPostAnalysisResponse503ResponseBody struct {
	Err *string `json:"err,omitempty"`
}

func (o *FileAnalysisPostAnalysisResponse503ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// FileAnalysisPostAnalysisResponse500ResponseBody - Unexpected event on server
type FileAnalysisPostAnalysisResponse500ResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *FileAnalysisPostAnalysisResponse500ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// FileAnalysisPostAnalysisResponse422ResponseBody - Body input is empty.
type FileAnalysisPostAnalysisResponse422ResponseBody struct {
	Err *string `json:"err,omitempty"`
}

func (o *FileAnalysisPostAnalysisResponse422ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// FileAnalysisPostAnalysisResponse411ResponseBody - Content-Length header is missing from the request.
type FileAnalysisPostAnalysisResponse411ResponseBody struct {
	Err *string `json:"err,omitempty"`
}

func (o *FileAnalysisPostAnalysisResponse411ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// FileAnalysisPostAnalysisResponseResponseBody - Invalid user information or Not Allowed
type FileAnalysisPostAnalysisResponseResponseBody struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *FileAnalysisPostAnalysisResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type FileAnalysisPostAnalysisResponseBodyType string

const (
	FileAnalysisPostAnalysisResponseBodyTypeBodyAndDownloadLinkGiven FileAnalysisPostAnalysisResponseBodyType = "BodyAndDownloadLinkGiven"
	FileAnalysisPostAnalysisResponseBodyTypeCallbackURLInvalid       FileAnalysisPostAnalysisResponseBodyType = "CallbackURLInvalid"
	FileAnalysisPostAnalysisResponseBodyTypeFileUploadRejected       FileAnalysisPostAnalysisResponseBodyType = "FileUploadRejected"
)

type FileAnalysisPostAnalysisResponseBody struct {
	BodyAndDownloadLinkGiven *shared.BodyAndDownloadLinkGiven
	CallbackURLInvalid       *shared.CallbackURLInvalid
	FileUploadRejected       *shared.FileUploadRejected

	Type FileAnalysisPostAnalysisResponseBodyType
}

func CreateFileAnalysisPostAnalysisResponseBodyBodyAndDownloadLinkGiven(bodyAndDownloadLinkGiven shared.BodyAndDownloadLinkGiven) FileAnalysisPostAnalysisResponseBody {
	typ := FileAnalysisPostAnalysisResponseBodyTypeBodyAndDownloadLinkGiven

	return FileAnalysisPostAnalysisResponseBody{
		BodyAndDownloadLinkGiven: &bodyAndDownloadLinkGiven,
		Type:                     typ,
	}
}

func CreateFileAnalysisPostAnalysisResponseBodyCallbackURLInvalid(callbackURLInvalid shared.CallbackURLInvalid) FileAnalysisPostAnalysisResponseBody {
	typ := FileAnalysisPostAnalysisResponseBodyTypeCallbackURLInvalid

	return FileAnalysisPostAnalysisResponseBody{
		CallbackURLInvalid: &callbackURLInvalid,
		Type:               typ,
	}
}

func CreateFileAnalysisPostAnalysisResponseBodyFileUploadRejected(fileUploadRejected shared.FileUploadRejected) FileAnalysisPostAnalysisResponseBody {
	typ := FileAnalysisPostAnalysisResponseBodyTypeFileUploadRejected

	return FileAnalysisPostAnalysisResponseBody{
		FileUploadRejected: &fileUploadRejected,
		Type:               typ,
	}
}

func (u *FileAnalysisPostAnalysisResponseBody) UnmarshalJSON(data []byte) error {

	bodyAndDownloadLinkGiven := new(shared.BodyAndDownloadLinkGiven)
	if err := utils.UnmarshalJSON(data, &bodyAndDownloadLinkGiven, "", true, true); err == nil {
		u.BodyAndDownloadLinkGiven = bodyAndDownloadLinkGiven
		u.Type = FileAnalysisPostAnalysisResponseBodyTypeBodyAndDownloadLinkGiven
		return nil
	}

	callbackURLInvalid := new(shared.CallbackURLInvalid)
	if err := utils.UnmarshalJSON(data, &callbackURLInvalid, "", true, true); err == nil {
		u.CallbackURLInvalid = callbackURLInvalid
		u.Type = FileAnalysisPostAnalysisResponseBodyTypeCallbackURLInvalid
		return nil
	}

	fileUploadRejected := new(shared.FileUploadRejected)
	if err := utils.UnmarshalJSON(data, &fileUploadRejected, "", true, true); err == nil {
		u.FileUploadRejected = fileUploadRejected
		u.Type = FileAnalysisPostAnalysisResponseBodyTypeFileUploadRejected
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u FileAnalysisPostAnalysisResponseBody) MarshalJSON() ([]byte, error) {
	if u.BodyAndDownloadLinkGiven != nil {
		return utils.MarshalJSON(u.BodyAndDownloadLinkGiven, "", true)
	}

	if u.CallbackURLInvalid != nil {
		return utils.MarshalJSON(u.CallbackURLInvalid, "", true)
	}

	if u.FileUploadRejected != nil {
		return utils.MarshalJSON(u.FileUploadRejected, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// FileAnalysisPostResponseBody - Successful file submission
type FileAnalysisPostResponseBody struct {
	// Unique submission identifier.
	// Use this value to reference the submission.
	//
	DataID string `json:"data_id"`
}

func (o *FileAnalysisPostResponseBody) GetDataID() string {
	if o == nil {
		return ""
	}
	return o.DataID
}

type FileAnalysisPostResponse struct {
	// Successful file submission
	TwoHundredApplicationJSONObject *FileAnalysisPostResponseBody
	// Callbackurl and/or apikey is invalid.
	FourHundredApplicationJSONOneOf *FileAnalysisPostAnalysisResponseBody
	// Invalid user information or Not Allowed
	FourHundredAndThreeApplicationJSONObject *FileAnalysisPostAnalysisResponseResponseBody
	// Content-Length header is missing from the request.
	FourHundredAndElevenApplicationJSONObject *FileAnalysisPostAnalysisResponse411ResponseBody
	// Body input is empty.
	FourHundredAndTwentyTwoApplicationJSONObject *FileAnalysisPostAnalysisResponse422ResponseBody
	// Unexpected event on server
	FiveHundredApplicationJSONObject *FileAnalysisPostAnalysisResponse500ResponseBody
	// Server is too busy, scan queue is full. Try again later.
	FiveHundredAndThreeApplicationJSONObject *FileAnalysisPostAnalysisResponse503ResponseBody
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *FileAnalysisPostResponse) GetTwoHundredApplicationJSONObject() *FileAnalysisPostResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *FileAnalysisPostResponse) GetFourHundredApplicationJSONOneOf() *FileAnalysisPostAnalysisResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONOneOf
}

func (o *FileAnalysisPostResponse) GetFourHundredAndThreeApplicationJSONObject() *FileAnalysisPostAnalysisResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *FileAnalysisPostResponse) GetFourHundredAndElevenApplicationJSONObject() *FileAnalysisPostAnalysisResponse411ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndElevenApplicationJSONObject
}

func (o *FileAnalysisPostResponse) GetFourHundredAndTwentyTwoApplicationJSONObject() *FileAnalysisPostAnalysisResponse422ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndTwentyTwoApplicationJSONObject
}

func (o *FileAnalysisPostResponse) GetFiveHundredApplicationJSONObject() *FileAnalysisPostAnalysisResponse500ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONObject
}

func (o *FileAnalysisPostResponse) GetFiveHundredAndThreeApplicationJSONObject() *FileAnalysisPostAnalysisResponse503ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredAndThreeApplicationJSONObject
}

func (o *FileAnalysisPostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FileAnalysisPostResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *FileAnalysisPostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FileAnalysisPostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
