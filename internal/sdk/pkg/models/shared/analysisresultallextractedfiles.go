// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AnalysisResultAllExtractedFilesProcessInfoOutdatedData string

const (
	AnalysisResultAllExtractedFilesProcessInfoOutdatedDataEnginedefinitions AnalysisResultAllExtractedFilesProcessInfoOutdatedData = "enginedefinitions"
	AnalysisResultAllExtractedFilesProcessInfoOutdatedDataConfiguration     AnalysisResultAllExtractedFilesProcessInfoOutdatedData = "configuration"
	AnalysisResultAllExtractedFilesProcessInfoOutdatedDataSanitization      AnalysisResultAllExtractedFilesProcessInfoOutdatedData = "sanitization"
)

func (e AnalysisResultAllExtractedFilesProcessInfoOutdatedData) ToPointer() *AnalysisResultAllExtractedFilesProcessInfoOutdatedData {
	return &e
}

func (e *AnalysisResultAllExtractedFilesProcessInfoOutdatedData) UnmarshalJSON(data []byte) error {
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
		*e = AnalysisResultAllExtractedFilesProcessInfoOutdatedData(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnalysisResultAllExtractedFilesProcessInfoOutdatedData: %v", v)
	}
}

// AnalysisResultAllExtractedFilesProcessInfoPostProcessing - Contains information about result of sanitization process and any action done after finalizing the process using Post Actions.
type AnalysisResultAllExtractedFilesProcessInfoPostProcessing struct {
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

// AnalysisResultAllExtractedFilesProcessInfoProcessingTimeDetails - Starting with MetaDefender Core 4.19.1, processing time on each major workflow processing step is calculated.
type AnalysisResultAllExtractedFilesProcessInfoProcessingTimeDetails struct {
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
	// Digital signature analyzing time.
	ParseDgsgTime *int64 `json:"parse_dgsg_time,omitempty"`
	// Vulnerability engine's lookup time.
	VulTime *int64 `json:"vul_time,omitempty"`
	// The waiting time for completely processing all nested files of an archive file.
	WaitChildFilesTime *int64 `json:"wait_child_files_time,omitempty"`
	// Yara engine's processing time.
	YaraTime *int64 `json:"yara_time,omitempty"`
}

// AnalysisResultAllExtractedFilesProcessInfo - Processing information
type AnalysisResultAllExtractedFilesProcessInfo struct {
	// Provides the reason why the file is blocked (if so).
	BlockedReason *string `json:"blocked_reason,omitempty"`
	// Indicates if the input file's detected type was configured to skip scanning.
	FileTypeSkippedScan *bool `json:"file_type_skipped_scan,omitempty"`
	// array of flags - if occur - describing outdated data in the result, these can be
	//   * enginedefinitions: at least one of the AV engines the item was scanned with has a newer definition database
	//   * configuration: the process' rule - or any item used by the rule - was modified since the item was processed
	//   * sanitization: if item was sanitized this flag notifies that the sanitization information regarding this result is outdated, meaning the sanitized item is no longer available
	//
	OutdatedData []AnalysisResultAllExtractedFilesProcessInfoOutdatedData `json:"outdated_data,omitempty"`
	// Contains information about result of sanitization process and any action done after finalizing the process using Post Actions.
	PostProcessing *AnalysisResultAllExtractedFilesProcessInfoPostProcessing `json:"post_processing,omitempty"`
	// Total time elapsed during processing file on the node (in milliseconds).
	ProcessingTime *int64 `json:"processing_time,omitempty"`
	// Starting with MetaDefender Core 4.19.1, processing time on each major workflow processing step is calculated.
	ProcessingTimeDetails *AnalysisResultAllExtractedFilesProcessInfoProcessingTimeDetails `json:"processing_time_details,omitempty"`
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

type AnalysisResultAllExtractedFiles struct {
	// data identifier of the requested file
	DataID *string `json:"data_id,omitempty"`
	// Full report from Proactive DLP
	DlpInfo *DLPResponse `json:"dlp_info,omitempty"`
	// The downloading status.
	DownloadInfo *DownloadInfo `json:"download_info,omitempty"`
	// Information about extracted files.
	ExtractedFiles []ExtractedFiles `json:"extracted_files,omitempty"`
	// basic information of the scanned file
	FileInfo *FileInformation `json:"file_info,omitempty"`
	// response information from FileType engine
	FiletypeInfo *FileTypeResponse `json:"filetype_info,omitempty"`
	// response information from OPSWAT Filescan engine
	OpswatfilescanInfo *OPSWATFilescanResponse `json:"opswatfilescan_info,omitempty"`
	// Processing information
	ProcessInfo *AnalysisResultAllExtractedFilesProcessInfo `json:"process_info,omitempty"`
	// Result of the scanning process.
	ScanResults *MetascanReport `json:"scan_results,omitempty"`
	// Contains all vulnerability information of the analysis result
	VulnerabilityInfo *VulnerabilityResponse `json:"vulnerability_info,omitempty"`
	// Information on data that matched yara rules
	Yara *YaraReport `json:"yara,omitempty"`
}