// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"Metadefender/internal/sdk/pkg/utils"
	"errors"
	"net/http"
)

type FileAnalysisSyncPostRequest struct {
	RequestBody []byte `request:"mediaType=application/octet-stream"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey *string `header:"style=simple,explode=false,name=apikey"`
	// password for archive ( URL encoded UTF-8 string)
	// Multiple passwords is also supported, format: archivepwd<X>
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
	// Format: \<protocol://\>\<ip | domain\>:\<port\>\</path\>
	//
	// When this header is specified, users must not have body content in HTTP(S) request, otherwise expecting to hit 400 HTTP(S) error. Check response sample of HTTP 400 for details.
	//
	// Specify download link where MetaDefender Core could download the entire payload for processing.
	//   * Supported protocol: HTTP / HTTPS
	//   * Support both individual file scan (both webhook & non-webhook) and file processing in batch via link
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
	// Specify name of file captured and displayed on corresponding MetaDefender Core scan result
	Filename *string `header:"style=simple,explode=false,name=filename"`
	// if local file scan is enabled the path to the file (see Security rule configuration)
	Filepath *string `header:"style=simple,explode=false,name=filepath"`
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
	// user_agent header used to identify (and limit) access to a particular rule. For rule selection, `rule` header should be used.
	// Since MetaDefender Core 5.6.0, when user agent is matched the predefined setting in workflow rule,  the product will return X-Core-Id containing identity of MetaDefender Core server as a header of response.
	//
	UserAgent *string `header:"style=simple,explode=false,name=user_agent"`
	// name of the selected workflow (deprecated, use "rule" header parameter instead)
	Workflow *string `header:"style=simple,explode=false,name=workflow"`
}

func (o *FileAnalysisSyncPostRequest) GetRequestBody() []byte {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *FileAnalysisSyncPostRequest) GetApikey() *string {
	if o == nil {
		return nil
	}
	return o.Apikey
}

func (o *FileAnalysisSyncPostRequest) GetArchivepwd() *string {
	if o == nil {
		return nil
	}
	return o.Archivepwd
}

func (o *FileAnalysisSyncPostRequest) GetBatch() *string {
	if o == nil {
		return nil
	}
	return o.Batch
}

func (o *FileAnalysisSyncPostRequest) GetDownloadfrom() *string {
	if o == nil {
		return nil
	}
	return o.Downloadfrom
}

func (o *FileAnalysisSyncPostRequest) GetEnginesMetadata() *string {
	if o == nil {
		return nil
	}
	return o.EnginesMetadata
}

func (o *FileAnalysisSyncPostRequest) GetFilename() *string {
	if o == nil {
		return nil
	}
	return o.Filename
}

func (o *FileAnalysisSyncPostRequest) GetFilepath() *string {
	if o == nil {
		return nil
	}
	return o.Filepath
}

func (o *FileAnalysisSyncPostRequest) GetMetadata() *string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *FileAnalysisSyncPostRequest) GetRule() *string {
	if o == nil {
		return nil
	}
	return o.Rule
}

func (o *FileAnalysisSyncPostRequest) GetUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.UserAgent
}

func (o *FileAnalysisSyncPostRequest) GetWorkflow() *string {
	if o == nil {
		return nil
	}
	return o.Workflow
}

// FileAnalysisSyncPostAnalysisResponse503ResponseBody - Server is too busy, scan queue is full. Try again later.
type FileAnalysisSyncPostAnalysisResponse503ResponseBody struct {
	Err *string `json:"err,omitempty"`
}

func (o *FileAnalysisSyncPostAnalysisResponse503ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type FileAnalysisSyncPostAnalysisResponse500ResponseBodyType string

const (
	FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeLicenseExpired    FileAnalysisSyncPostAnalysisResponse500ResponseBodyType = "LicenseExpired"
	FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeNoAvailableRule   FileAnalysisSyncPostAnalysisResponse500ResponseBodyType = "NoAvailableRule"
	FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeInternalError     FileAnalysisSyncPostAnalysisResponse500ResponseBodyType = "InternalError"
	FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeLocalFileNotFound FileAnalysisSyncPostAnalysisResponse500ResponseBodyType = "LocalFileNotFound"
	FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeFileSizeExceeded  FileAnalysisSyncPostAnalysisResponse500ResponseBodyType = "FileSizeExceeded"
)

type FileAnalysisSyncPostAnalysisResponse500ResponseBody struct {
	LicenseExpired    *shared.LicenseExpired
	NoAvailableRule   *shared.NoAvailableRule
	InternalError     *shared.InternalError
	LocalFileNotFound *shared.LocalFileNotFound
	FileSizeExceeded  *shared.FileSizeExceeded

	Type FileAnalysisSyncPostAnalysisResponse500ResponseBodyType
}

func CreateFileAnalysisSyncPostAnalysisResponse500ResponseBodyLicenseExpired(licenseExpired shared.LicenseExpired) FileAnalysisSyncPostAnalysisResponse500ResponseBody {
	typ := FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeLicenseExpired

	return FileAnalysisSyncPostAnalysisResponse500ResponseBody{
		LicenseExpired: &licenseExpired,
		Type:           typ,
	}
}

func CreateFileAnalysisSyncPostAnalysisResponse500ResponseBodyNoAvailableRule(noAvailableRule shared.NoAvailableRule) FileAnalysisSyncPostAnalysisResponse500ResponseBody {
	typ := FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeNoAvailableRule

	return FileAnalysisSyncPostAnalysisResponse500ResponseBody{
		NoAvailableRule: &noAvailableRule,
		Type:            typ,
	}
}

func CreateFileAnalysisSyncPostAnalysisResponse500ResponseBodyInternalError(internalError shared.InternalError) FileAnalysisSyncPostAnalysisResponse500ResponseBody {
	typ := FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeInternalError

	return FileAnalysisSyncPostAnalysisResponse500ResponseBody{
		InternalError: &internalError,
		Type:          typ,
	}
}

func CreateFileAnalysisSyncPostAnalysisResponse500ResponseBodyLocalFileNotFound(localFileNotFound shared.LocalFileNotFound) FileAnalysisSyncPostAnalysisResponse500ResponseBody {
	typ := FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeLocalFileNotFound

	return FileAnalysisSyncPostAnalysisResponse500ResponseBody{
		LocalFileNotFound: &localFileNotFound,
		Type:              typ,
	}
}

func CreateFileAnalysisSyncPostAnalysisResponse500ResponseBodyFileSizeExceeded(fileSizeExceeded shared.FileSizeExceeded) FileAnalysisSyncPostAnalysisResponse500ResponseBody {
	typ := FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeFileSizeExceeded

	return FileAnalysisSyncPostAnalysisResponse500ResponseBody{
		FileSizeExceeded: &fileSizeExceeded,
		Type:             typ,
	}
}

func (u *FileAnalysisSyncPostAnalysisResponse500ResponseBody) UnmarshalJSON(data []byte) error {

	licenseExpired := new(shared.LicenseExpired)
	if err := utils.UnmarshalJSON(data, &licenseExpired, "", true, true); err == nil {
		u.LicenseExpired = licenseExpired
		u.Type = FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeLicenseExpired
		return nil
	}

	noAvailableRule := new(shared.NoAvailableRule)
	if err := utils.UnmarshalJSON(data, &noAvailableRule, "", true, true); err == nil {
		u.NoAvailableRule = noAvailableRule
		u.Type = FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeNoAvailableRule
		return nil
	}

	internalError := new(shared.InternalError)
	if err := utils.UnmarshalJSON(data, &internalError, "", true, true); err == nil {
		u.InternalError = internalError
		u.Type = FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeInternalError
		return nil
	}

	localFileNotFound := new(shared.LocalFileNotFound)
	if err := utils.UnmarshalJSON(data, &localFileNotFound, "", true, true); err == nil {
		u.LocalFileNotFound = localFileNotFound
		u.Type = FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeLocalFileNotFound
		return nil
	}

	fileSizeExceeded := new(shared.FileSizeExceeded)
	if err := utils.UnmarshalJSON(data, &fileSizeExceeded, "", true, true); err == nil {
		u.FileSizeExceeded = fileSizeExceeded
		u.Type = FileAnalysisSyncPostAnalysisResponse500ResponseBodyTypeFileSizeExceeded
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u FileAnalysisSyncPostAnalysisResponse500ResponseBody) MarshalJSON() ([]byte, error) {
	if u.LicenseExpired != nil {
		return utils.MarshalJSON(u.LicenseExpired, "", true)
	}

	if u.NoAvailableRule != nil {
		return utils.MarshalJSON(u.NoAvailableRule, "", true)
	}

	if u.InternalError != nil {
		return utils.MarshalJSON(u.InternalError, "", true)
	}

	if u.LocalFileNotFound != nil {
		return utils.MarshalJSON(u.LocalFileNotFound, "", true)
	}

	if u.FileSizeExceeded != nil {
		return utils.MarshalJSON(u.FileSizeExceeded, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// FileAnalysisSyncPostAnalysisResponse422ResponseBody - Body input is empty.
type FileAnalysisSyncPostAnalysisResponse422ResponseBody struct {
	Err *string `json:"err,omitempty"`
}

func (o *FileAnalysisSyncPostAnalysisResponse422ResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// FileAnalysisSyncPostAnalysisResponseResponseBody - Content-Length header is missing from the request.
type FileAnalysisSyncPostAnalysisResponseResponseBody struct {
	Err *string `json:"err,omitempty"`
}

func (o *FileAnalysisSyncPostAnalysisResponseResponseBody) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// FileAnalysisSyncPostAnalysisResponseBody - Timed out response. The scan is still in progress, but the connection is timed out.
type FileAnalysisSyncPostAnalysisResponseBody struct {
	DataID *string `json:"data_id,omitempty"`
}

func (o *FileAnalysisSyncPostAnalysisResponseBody) GetDataID() *string {
	if o == nil {
		return nil
	}
	return o.DataID
}

type FileAnalysisSyncPostResponseBodyType string

const (
	FileAnalysisSyncPostResponseBodyTypeInvalidAPIKeyGiven          FileAnalysisSyncPostResponseBodyType = "InvalidAPIKeyGiven"
	FileAnalysisSyncPostResponseBodyTypeBodyAndDownloadLinkGiven    FileAnalysisSyncPostResponseBodyType = "BodyAndDownloadLinkGiven"
	FileAnalysisSyncPostResponseBodyTypeInprogressBatchClosed       FileAnalysisSyncPostResponseBodyType = "InprogressBatchClosed"
	FileAnalysisSyncPostResponseBodyTypeBatchNotFound               FileAnalysisSyncPostResponseBodyType = "BatchNotFound"
	FileAnalysisSyncPostResponseBodyTypeBatchInstanceMismatched     FileAnalysisSyncPostResponseBodyType = "BatchInstanceMismatched"
	FileAnalysisSyncPostResponseBodyTypeBodyAndLocalFilePathGiven   FileAnalysisSyncPostResponseBodyType = "BodyAndLocalFilePathGiven"
	FileAnalysisSyncPostResponseBodyTypeFileUploadRejected          FileAnalysisSyncPostResponseBodyType = "FileUploadRejected"
	FileAnalysisSyncPostResponseBodyTypeRedirectLinkNotSupported    FileAnalysisSyncPostResponseBodyType = "RedirectLinkNotSupported"
	FileAnalysisSyncPostResponseBodyTypeInvalidDownloadLink         FileAnalysisSyncPostResponseBodyType = "InvalidDownloadLink"
	FileAnalysisSyncPostResponseBodyTypeInvalidProtocolDownloadLink FileAnalysisSyncPostResponseBodyType = "InvalidProtocolDownloadLink"
)

type FileAnalysisSyncPostResponseBody struct {
	InvalidAPIKeyGiven          *shared.InvalidAPIKeyGiven
	BodyAndDownloadLinkGiven    *shared.BodyAndDownloadLinkGiven
	InprogressBatchClosed       *shared.InprogressBatchClosed
	BatchNotFound               *shared.BatchNotFound
	BatchInstanceMismatched     *shared.BatchInstanceMismatched
	BodyAndLocalFilePathGiven   *shared.BodyAndLocalFilePathGiven
	FileUploadRejected          *shared.FileUploadRejected
	RedirectLinkNotSupported    *shared.RedirectLinkNotSupported
	InvalidDownloadLink         *shared.InvalidDownloadLink
	InvalidProtocolDownloadLink *shared.InvalidProtocolDownloadLink

	Type FileAnalysisSyncPostResponseBodyType
}

func CreateFileAnalysisSyncPostResponseBodyInvalidAPIKeyGiven(invalidAPIKeyGiven shared.InvalidAPIKeyGiven) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeInvalidAPIKeyGiven

	return FileAnalysisSyncPostResponseBody{
		InvalidAPIKeyGiven: &invalidAPIKeyGiven,
		Type:               typ,
	}
}

func CreateFileAnalysisSyncPostResponseBodyBodyAndDownloadLinkGiven(bodyAndDownloadLinkGiven shared.BodyAndDownloadLinkGiven) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeBodyAndDownloadLinkGiven

	return FileAnalysisSyncPostResponseBody{
		BodyAndDownloadLinkGiven: &bodyAndDownloadLinkGiven,
		Type:                     typ,
	}
}

func CreateFileAnalysisSyncPostResponseBodyInprogressBatchClosed(inprogressBatchClosed shared.InprogressBatchClosed) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeInprogressBatchClosed

	return FileAnalysisSyncPostResponseBody{
		InprogressBatchClosed: &inprogressBatchClosed,
		Type:                  typ,
	}
}

func CreateFileAnalysisSyncPostResponseBodyBatchNotFound(batchNotFound shared.BatchNotFound) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeBatchNotFound

	return FileAnalysisSyncPostResponseBody{
		BatchNotFound: &batchNotFound,
		Type:          typ,
	}
}

func CreateFileAnalysisSyncPostResponseBodyBatchInstanceMismatched(batchInstanceMismatched shared.BatchInstanceMismatched) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeBatchInstanceMismatched

	return FileAnalysisSyncPostResponseBody{
		BatchInstanceMismatched: &batchInstanceMismatched,
		Type:                    typ,
	}
}

func CreateFileAnalysisSyncPostResponseBodyBodyAndLocalFilePathGiven(bodyAndLocalFilePathGiven shared.BodyAndLocalFilePathGiven) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeBodyAndLocalFilePathGiven

	return FileAnalysisSyncPostResponseBody{
		BodyAndLocalFilePathGiven: &bodyAndLocalFilePathGiven,
		Type:                      typ,
	}
}

func CreateFileAnalysisSyncPostResponseBodyFileUploadRejected(fileUploadRejected shared.FileUploadRejected) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeFileUploadRejected

	return FileAnalysisSyncPostResponseBody{
		FileUploadRejected: &fileUploadRejected,
		Type:               typ,
	}
}

func CreateFileAnalysisSyncPostResponseBodyRedirectLinkNotSupported(redirectLinkNotSupported shared.RedirectLinkNotSupported) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeRedirectLinkNotSupported

	return FileAnalysisSyncPostResponseBody{
		RedirectLinkNotSupported: &redirectLinkNotSupported,
		Type:                     typ,
	}
}

func CreateFileAnalysisSyncPostResponseBodyInvalidDownloadLink(invalidDownloadLink shared.InvalidDownloadLink) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeInvalidDownloadLink

	return FileAnalysisSyncPostResponseBody{
		InvalidDownloadLink: &invalidDownloadLink,
		Type:                typ,
	}
}

func CreateFileAnalysisSyncPostResponseBodyInvalidProtocolDownloadLink(invalidProtocolDownloadLink shared.InvalidProtocolDownloadLink) FileAnalysisSyncPostResponseBody {
	typ := FileAnalysisSyncPostResponseBodyTypeInvalidProtocolDownloadLink

	return FileAnalysisSyncPostResponseBody{
		InvalidProtocolDownloadLink: &invalidProtocolDownloadLink,
		Type:                        typ,
	}
}

func (u *FileAnalysisSyncPostResponseBody) UnmarshalJSON(data []byte) error {

	invalidAPIKeyGiven := new(shared.InvalidAPIKeyGiven)
	if err := utils.UnmarshalJSON(data, &invalidAPIKeyGiven, "", true, true); err == nil {
		u.InvalidAPIKeyGiven = invalidAPIKeyGiven
		u.Type = FileAnalysisSyncPostResponseBodyTypeInvalidAPIKeyGiven
		return nil
	}

	bodyAndDownloadLinkGiven := new(shared.BodyAndDownloadLinkGiven)
	if err := utils.UnmarshalJSON(data, &bodyAndDownloadLinkGiven, "", true, true); err == nil {
		u.BodyAndDownloadLinkGiven = bodyAndDownloadLinkGiven
		u.Type = FileAnalysisSyncPostResponseBodyTypeBodyAndDownloadLinkGiven
		return nil
	}

	inprogressBatchClosed := new(shared.InprogressBatchClosed)
	if err := utils.UnmarshalJSON(data, &inprogressBatchClosed, "", true, true); err == nil {
		u.InprogressBatchClosed = inprogressBatchClosed
		u.Type = FileAnalysisSyncPostResponseBodyTypeInprogressBatchClosed
		return nil
	}

	batchNotFound := new(shared.BatchNotFound)
	if err := utils.UnmarshalJSON(data, &batchNotFound, "", true, true); err == nil {
		u.BatchNotFound = batchNotFound
		u.Type = FileAnalysisSyncPostResponseBodyTypeBatchNotFound
		return nil
	}

	batchInstanceMismatched := new(shared.BatchInstanceMismatched)
	if err := utils.UnmarshalJSON(data, &batchInstanceMismatched, "", true, true); err == nil {
		u.BatchInstanceMismatched = batchInstanceMismatched
		u.Type = FileAnalysisSyncPostResponseBodyTypeBatchInstanceMismatched
		return nil
	}

	bodyAndLocalFilePathGiven := new(shared.BodyAndLocalFilePathGiven)
	if err := utils.UnmarshalJSON(data, &bodyAndLocalFilePathGiven, "", true, true); err == nil {
		u.BodyAndLocalFilePathGiven = bodyAndLocalFilePathGiven
		u.Type = FileAnalysisSyncPostResponseBodyTypeBodyAndLocalFilePathGiven
		return nil
	}

	fileUploadRejected := new(shared.FileUploadRejected)
	if err := utils.UnmarshalJSON(data, &fileUploadRejected, "", true, true); err == nil {
		u.FileUploadRejected = fileUploadRejected
		u.Type = FileAnalysisSyncPostResponseBodyTypeFileUploadRejected
		return nil
	}

	redirectLinkNotSupported := new(shared.RedirectLinkNotSupported)
	if err := utils.UnmarshalJSON(data, &redirectLinkNotSupported, "", true, true); err == nil {
		u.RedirectLinkNotSupported = redirectLinkNotSupported
		u.Type = FileAnalysisSyncPostResponseBodyTypeRedirectLinkNotSupported
		return nil
	}

	invalidDownloadLink := new(shared.InvalidDownloadLink)
	if err := utils.UnmarshalJSON(data, &invalidDownloadLink, "", true, true); err == nil {
		u.InvalidDownloadLink = invalidDownloadLink
		u.Type = FileAnalysisSyncPostResponseBodyTypeInvalidDownloadLink
		return nil
	}

	invalidProtocolDownloadLink := new(shared.InvalidProtocolDownloadLink)
	if err := utils.UnmarshalJSON(data, &invalidProtocolDownloadLink, "", true, true); err == nil {
		u.InvalidProtocolDownloadLink = invalidProtocolDownloadLink
		u.Type = FileAnalysisSyncPostResponseBodyTypeInvalidProtocolDownloadLink
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u FileAnalysisSyncPostResponseBody) MarshalJSON() ([]byte, error) {
	if u.InvalidAPIKeyGiven != nil {
		return utils.MarshalJSON(u.InvalidAPIKeyGiven, "", true)
	}

	if u.BodyAndDownloadLinkGiven != nil {
		return utils.MarshalJSON(u.BodyAndDownloadLinkGiven, "", true)
	}

	if u.InprogressBatchClosed != nil {
		return utils.MarshalJSON(u.InprogressBatchClosed, "", true)
	}

	if u.BatchNotFound != nil {
		return utils.MarshalJSON(u.BatchNotFound, "", true)
	}

	if u.BatchInstanceMismatched != nil {
		return utils.MarshalJSON(u.BatchInstanceMismatched, "", true)
	}

	if u.BodyAndLocalFilePathGiven != nil {
		return utils.MarshalJSON(u.BodyAndLocalFilePathGiven, "", true)
	}

	if u.FileUploadRejected != nil {
		return utils.MarshalJSON(u.FileUploadRejected, "", true)
	}

	if u.RedirectLinkNotSupported != nil {
		return utils.MarshalJSON(u.RedirectLinkNotSupported, "", true)
	}

	if u.InvalidDownloadLink != nil {
		return utils.MarshalJSON(u.InvalidDownloadLink, "", true)
	}

	if u.InvalidProtocolDownloadLink != nil {
		return utils.MarshalJSON(u.InvalidProtocolDownloadLink, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type FileAnalysisSyncPostResponse struct {
	// API key is invalid.
	FourHundredApplicationJSONOneOf *FileAnalysisSyncPostResponseBody
	// Timed out response. The scan is still in progress, but the connection is timed out.
	FourHundredAndEightApplicationJSONObject *FileAnalysisSyncPostAnalysisResponseBody
	// Content-Length header is missing from the request.
	FourHundredAndElevenApplicationJSONObject *FileAnalysisSyncPostAnalysisResponseResponseBody
	// Body input is empty.
	FourHundredAndTwentyTwoApplicationJSONObject *FileAnalysisSyncPostAnalysisResponse422ResponseBody
	// Internal Server Error.
	FiveHundredApplicationJSONOneOf *FileAnalysisSyncPostAnalysisResponse500ResponseBody
	// Server is too busy, scan queue is full. Try again later.
	FiveHundredAndThreeApplicationJSONObject *FileAnalysisSyncPostAnalysisResponse503ResponseBody
	// The scan is completed before the connection timeout occurs
	AnalysisResult *shared.AnalysisResult
	// Callback URL is not supported.
	CallbackURLNotSupported *shared.CallbackURLNotSupported
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *FileAnalysisSyncPostResponse) GetFourHundredApplicationJSONOneOf() *FileAnalysisSyncPostResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONOneOf
}

func (o *FileAnalysisSyncPostResponse) GetFourHundredAndEightApplicationJSONObject() *FileAnalysisSyncPostAnalysisResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndEightApplicationJSONObject
}

func (o *FileAnalysisSyncPostResponse) GetFourHundredAndElevenApplicationJSONObject() *FileAnalysisSyncPostAnalysisResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndElevenApplicationJSONObject
}

func (o *FileAnalysisSyncPostResponse) GetFourHundredAndTwentyTwoApplicationJSONObject() *FileAnalysisSyncPostAnalysisResponse422ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndTwentyTwoApplicationJSONObject
}

func (o *FileAnalysisSyncPostResponse) GetFiveHundredApplicationJSONOneOf() *FileAnalysisSyncPostAnalysisResponse500ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredApplicationJSONOneOf
}

func (o *FileAnalysisSyncPostResponse) GetFiveHundredAndThreeApplicationJSONObject() *FileAnalysisSyncPostAnalysisResponse503ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredAndThreeApplicationJSONObject
}

func (o *FileAnalysisSyncPostResponse) GetAnalysisResult() *shared.AnalysisResult {
	if o == nil {
		return nil
	}
	return o.AnalysisResult
}

func (o *FileAnalysisSyncPostResponse) GetCallbackURLNotSupported() *shared.CallbackURLNotSupported {
	if o == nil {
		return nil
	}
	return o.CallbackURLNotSupported
}

func (o *FileAnalysisSyncPostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FileAnalysisSyncPostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FileAnalysisSyncPostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
