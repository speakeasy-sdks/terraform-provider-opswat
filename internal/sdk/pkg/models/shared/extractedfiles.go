// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ExtractedFilesProcessInfoOutdatedData string

const (
	ExtractedFilesProcessInfoOutdatedDataEnginedefinitions ExtractedFilesProcessInfoOutdatedData = "enginedefinitions"
	ExtractedFilesProcessInfoOutdatedDataConfiguration     ExtractedFilesProcessInfoOutdatedData = "configuration"
	ExtractedFilesProcessInfoOutdatedDataSanitization      ExtractedFilesProcessInfoOutdatedData = "sanitization"
)

func (e ExtractedFilesProcessInfoOutdatedData) ToPointer() *ExtractedFilesProcessInfoOutdatedData {
	return &e
}

func (e *ExtractedFilesProcessInfoOutdatedData) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enginedefinitions":
		fallthrough
	case "configuration":
		fallthrough
	case "sanitization":
		*e = ExtractedFilesProcessInfoOutdatedData(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExtractedFilesProcessInfoOutdatedData: %v", v)
	}
}

// ExtractedFilesProcessInfoPostProcessing - Contains information about result of sanitization process and any action done after finalizing the process using Post Actions.
type ExtractedFilesProcessInfoPostProcessing struct {
	// Empty string if no action failed or list of failed actions, separated by "|".
	ActionsFailed *string `json:"actions_failed,omitempty"`
	// List of successful actions, separated by "|". Empty string if otherwise.
	ActionsRan *string `json:"actions_ran,omitempty"`
	// Contains the name of the sanitized file.
	ConvertedDestination *string `json:"converted_destination,omitempty"`
	// Contains target type name of sanitization.
	ConvertedTo *string `json:"converted_to,omitempty"`
	// Contains target type name of sanitization.
	CopyMoveDestination *string `json:"copy_move_destination,omitempty"`
	// Contains target type name of sanitization.
	SanitizationDetails *DeepCDRDetails `json:"sanitization_details,omitempty"`
}

func (o *ExtractedFilesProcessInfoPostProcessing) GetActionsFailed() *string {
	if o == nil {
		return nil
	}
	return o.ActionsFailed
}

func (o *ExtractedFilesProcessInfoPostProcessing) GetActionsRan() *string {
	if o == nil {
		return nil
	}
	return o.ActionsRan
}

func (o *ExtractedFilesProcessInfoPostProcessing) GetConvertedDestination() *string {
	if o == nil {
		return nil
	}
	return o.ConvertedDestination
}

func (o *ExtractedFilesProcessInfoPostProcessing) GetConvertedTo() *string {
	if o == nil {
		return nil
	}
	return o.ConvertedTo
}

func (o *ExtractedFilesProcessInfoPostProcessing) GetCopyMoveDestination() *string {
	if o == nil {
		return nil
	}
	return o.CopyMoveDestination
}

func (o *ExtractedFilesProcessInfoPostProcessing) GetSanitizationDetails() *DeepCDRDetails {
	if o == nil {
		return nil
	}
	return o.SanitizationDetails
}

// ExtractedFilesProcessInfoProcessingTimeDetails - Starting with MetaDefender Core 4.19.1, processing time on each major workflow processing step is calculated.
type ExtractedFilesProcessInfoProcessingTimeDetails struct {
	// AV engines' processing time.
	AvScanTime *int64 `json:"av_scan_time,omitempty"`
	// Deep CDR engine's sanitization time.
	CdrTime *int64 `json:"cdr_time,omitempty"`
	// Proactive DLP engine's processing time.
	DlpTime *int64 `json:"dlp_time,omitempty"`
	// Archive extraction engine's processing time.
	ExtractionTime *int64 `json:"extraction_time,omitempty"`
	// FileType engine's processing time.
	FiletypeTime *int64 `json:"filetype_time,omitempty"`
	// OPSWAT Filescan engine's processing time.
	OpswatfilescanTime *int64 `json:"opswatfilescan_time,omitempty"`
	// Total time elapsed for following processing tasks in the product (in milliseconds):
	// * File processing’s hash calculation time
	// * External scanner (if configured)
	// * Post action (if configured)
	// * Other internal procssing time among components in the product
	//
	OthersTime *int64 `json:"others_time,omitempty"`
	// Vulnerability engine's lookup time.
	VulTime *int64 `json:"vul_time,omitempty"`
	// The waiting time for completely processing all nested files of an archive file.
	WaitChildFilesTime *int64 `json:"wait_child_files_time,omitempty"`
	// Yara engine's processing time.
	YaraTime *int64 `json:"yara_time,omitempty"`
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetAvScanTime() *int64 {
	if o == nil {
		return nil
	}
	return o.AvScanTime
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetCdrTime() *int64 {
	if o == nil {
		return nil
	}
	return o.CdrTime
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetDlpTime() *int64 {
	if o == nil {
		return nil
	}
	return o.DlpTime
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetExtractionTime() *int64 {
	if o == nil {
		return nil
	}
	return o.ExtractionTime
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetFiletypeTime() *int64 {
	if o == nil {
		return nil
	}
	return o.FiletypeTime
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetOpswatfilescanTime() *int64 {
	if o == nil {
		return nil
	}
	return o.OpswatfilescanTime
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetOthersTime() *int64 {
	if o == nil {
		return nil
	}
	return o.OthersTime
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetVulTime() *int64 {
	if o == nil {
		return nil
	}
	return o.VulTime
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetWaitChildFilesTime() *int64 {
	if o == nil {
		return nil
	}
	return o.WaitChildFilesTime
}

func (o *ExtractedFilesProcessInfoProcessingTimeDetails) GetYaraTime() *int64 {
	if o == nil {
		return nil
	}
	return o.YaraTime
}

// ExtractedFilesProcessInfo - Processing information
type ExtractedFilesProcessInfo struct {
	// Provides the reason why the file is blocked (if so).
	BlockedReason *string `json:"blocked_reason,omitempty"`
	// Indicates if the input file's detected type was configured to skip scanning.
	FileTypeSkippedScan *bool `json:"file_type_skipped_scan,omitempty"`
	// array of flags - if occur - describing outdated data in the result, these can be
	//   * enginedefinitions: at least one of the AV engines the item was scanned with has a newer definition database
	//   * configuration: the process' rule - or any item used by the rule - was modified since the item was processed
	//   * sanitization: if item was sanitized this flag notifies that the sanitization information regarding this result is outdated, meaning the sanitized item is no longer available
	//
	OutdatedData []ExtractedFilesProcessInfoOutdatedData `json:"outdated_data,omitempty"`
	// Contains information about result of sanitization process and any action done after finalizing the process using Post Actions.
	PostProcessing *ExtractedFilesProcessInfoPostProcessing `json:"post_processing,omitempty"`
	// Total time elapsed during processing file on the node (in milliseconds).
	ProcessingTime *int64 `json:"processing_time,omitempty"`
	// Starting with MetaDefender Core 4.19.1, processing time on each major workflow processing step is calculated.
	ProcessingTimeDetails *ExtractedFilesProcessInfoProcessingTimeDetails `json:"processing_time_details,omitempty"`
	// The used rule name.
	Profile *string `json:"profile,omitempty"`
	// Percentage of processing completed (from 1-100).
	ProgressPercentage *int64 `json:"progress_percentage,omitempty"`
	// Total time elapsed for file processing task was waiting in MetaDefender Core’s queue until being picked up (queue_time = start_time - upload_timestamp) (in milliseconds).
	QueueTime *int64 `json:"queue_time,omitempty"`
	// The final result of processing the file (Allowed / Blocked / Processing).
	Result *string `json:"result,omitempty"`
	// Identifier for the REST Client that calls the API.
	UserAgent *string `json:"user_agent,omitempty"`
	// User identifier who submitted scan request earlier.
	Username *string `json:"username,omitempty"`
	// Aggregated list of potential issues.
	Verdicts []string `json:"verdicts,omitempty"`
}

func (o *ExtractedFilesProcessInfo) GetBlockedReason() *string {
	if o == nil {
		return nil
	}
	return o.BlockedReason
}

func (o *ExtractedFilesProcessInfo) GetFileTypeSkippedScan() *bool {
	if o == nil {
		return nil
	}
	return o.FileTypeSkippedScan
}

func (o *ExtractedFilesProcessInfo) GetOutdatedData() []ExtractedFilesProcessInfoOutdatedData {
	if o == nil {
		return nil
	}
	return o.OutdatedData
}

func (o *ExtractedFilesProcessInfo) GetPostProcessing() *ExtractedFilesProcessInfoPostProcessing {
	if o == nil {
		return nil
	}
	return o.PostProcessing
}

func (o *ExtractedFilesProcessInfo) GetProcessingTime() *int64 {
	if o == nil {
		return nil
	}
	return o.ProcessingTime
}

func (o *ExtractedFilesProcessInfo) GetProcessingTimeDetails() *ExtractedFilesProcessInfoProcessingTimeDetails {
	if o == nil {
		return nil
	}
	return o.ProcessingTimeDetails
}

func (o *ExtractedFilesProcessInfo) GetProfile() *string {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *ExtractedFilesProcessInfo) GetProgressPercentage() *int64 {
	if o == nil {
		return nil
	}
	return o.ProgressPercentage
}

func (o *ExtractedFilesProcessInfo) GetQueueTime() *int64 {
	if o == nil {
		return nil
	}
	return o.QueueTime
}

func (o *ExtractedFilesProcessInfo) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *ExtractedFilesProcessInfo) GetUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.UserAgent
}

func (o *ExtractedFilesProcessInfo) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ExtractedFilesProcessInfo) GetVerdicts() []string {
	if o == nil {
		return nil
	}
	return o.Verdicts
}

type ExtractedFiles struct {
	// data identifier of the requested file
	DataID *string `json:"data_id,omitempty"`
	// Full report from Proactive DLP
	DlpInfo *DLPResponse `json:"dlp_info,omitempty"`
	// Information about extracted files.
	ExtractedFiles []ExtractedFiles2 `json:"extracted_files,omitempty"`
	// basic information of the scanned file
	FileInfo *FileInformation `json:"file_info,omitempty"`
	// response information from FileType engine
	FiletypeInfo *FileTypeResponse `json:"filetype_info,omitempty"`
	// Processing information
	ProcessInfo *ExtractedFilesProcessInfo `json:"process_info,omitempty"`
	// Result of the scanning process.
	ScanResults *MetascanReport `json:"scan_results,omitempty"`
	// Contains all vulnerability information of the analysis result
	VulnerabilityInfo *VulnerabilityResponse `json:"vulnerability_info,omitempty"`
	// Information on data that matched yara rules
	Yara *YaraReport `json:"yara,omitempty"`
}

func (o *ExtractedFiles) GetDataID() *string {
	if o == nil {
		return nil
	}
	return o.DataID
}

func (o *ExtractedFiles) GetDlpInfo() *DLPResponse {
	if o == nil {
		return nil
	}
	return o.DlpInfo
}

func (o *ExtractedFiles) GetExtractedFiles() []ExtractedFiles2 {
	if o == nil {
		return nil
	}
	return o.ExtractedFiles
}

func (o *ExtractedFiles) GetFileInfo() *FileInformation {
	if o == nil {
		return nil
	}
	return o.FileInfo
}

func (o *ExtractedFiles) GetFiletypeInfo() *FileTypeResponse {
	if o == nil {
		return nil
	}
	return o.FiletypeInfo
}

func (o *ExtractedFiles) GetProcessInfo() *ExtractedFilesProcessInfo {
	if o == nil {
		return nil
	}
	return o.ProcessInfo
}

func (o *ExtractedFiles) GetScanResults() *MetascanReport {
	if o == nil {
		return nil
	}
	return o.ScanResults
}

func (o *ExtractedFiles) GetVulnerabilityInfo() *VulnerabilityResponse {
	if o == nil {
		return nil
	}
	return o.VulnerabilityInfo
}

func (o *ExtractedFiles) GetYara() *YaraReport {
	if o == nil {
		return nil
	}
	return o.Yara
}
