// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BatchResponseInProgressBatchFilesFilesInBatchProcessInfo - The analysis summary
type BatchResponseInProgressBatchFilesFilesInBatchProcessInfo struct {
	// Provides the reason why the file is blocked (if so).
	BlockedReason *string `json:"blocked_reason,omitempty"`
	// Percentage of processing completed (from 1-100).
	ProgressPercentage *int64 `json:"progress_percentage,omitempty"`
	// The final result of processing the file (Allowed / Blocked / Processing).
	Result *string `json:"result,omitempty"`
	// Aggregated list of potential issues.
	Verdicts []string `json:"verdicts,omitempty"`
}

// BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA - Processing result and its index
// * `No Threat Detected`: 0
// * `Infected`: 1
// * `Suspicious`: 2
// * `Failed`: 3
// * `Whitelisted`: 7
// * `Blacklisted`: 8
// * `Exceeded Archive Depth`: 9
// * `Not Scanned`: 10
// * `Encrypted Archive`: 12
// * `Exceeded Archive Size`: 13
// * `Exceeded Archive File Number`: 14
// * `Password Protected Document`: 15
// * `Exceeded Archive Timeout`: 16
// * `Mismatch`: 17
// * `Potentially Vulnerable File`: 18
// * `Cancelled`: 19
// * `Sensitive Data Found`: 20
// * `Yara Rule Matched`: 21
// * `Potentially Unwanted`: 22
// * `Unsupported File Type`: 23
// * `In Progress`: 255
type BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA string

const (
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultANoThreatDetected          BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "No Threat Detected"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAInfected                  BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Infected"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultASuspicious                BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Suspicious"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAFailed                    BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Failed"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAWhitelisted               BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Whitelisted"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultABlacklisted               BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Blacklisted"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAExceededArchiveDepth      BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Exceeded Archive Depth"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultANotScanned                BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Not Scanned"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAEncryptedArchive          BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Encrypted Archive"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAExceededArchiveSize       BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Exceeded Archive Size"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAExceededArchiveFileNumber BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Exceeded Archive File Number"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAPasswordProtectedDocument BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Password Protected Document"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAExceededArchiveTimeout    BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Exceeded Archive Timeout"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAMismatch                  BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Mismatch"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAPotentiallyVulnerableFile BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Potentially Vulnerable File"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultACancelled                 BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Cancelled"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultASensitiveDataFound        BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Sensitive Data Found"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAYaraRuleMatched           BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Yara Rule Matched"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAPotentiallyUnwanted       BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Potentially Unwanted"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAUnsupportedFileType       BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "Unsupported File Type"
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultAInProgress                BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA = "In Progress"
)

func (e BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA) ToPointer() *BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA {
	return &e
}

func (e *BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "No Threat Detected":
		fallthrough
	case "Infected":
		fallthrough
	case "Suspicious":
		fallthrough
	case "Failed":
		fallthrough
	case "Whitelisted":
		fallthrough
	case "Blacklisted":
		fallthrough
	case "Exceeded Archive Depth":
		fallthrough
	case "Not Scanned":
		fallthrough
	case "Encrypted Archive":
		fallthrough
	case "Exceeded Archive Size":
		fallthrough
	case "Exceeded Archive File Number":
		fallthrough
	case "Password Protected Document":
		fallthrough
	case "Exceeded Archive Timeout":
		fallthrough
	case "Mismatch":
		fallthrough
	case "Potentially Vulnerable File":
		fallthrough
	case "Cancelled":
		fallthrough
	case "Sensitive Data Found":
		fallthrough
	case "Yara Rule Matched":
		fallthrough
	case "Potentially Unwanted":
		fallthrough
	case "Unsupported File Type":
		fallthrough
	case "In Progress":
		*e = BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA: %v", v)
	}
}

// BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI - Scan result as index in the Processing Results table above
type BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI int64

const (
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultIZero                   BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 0
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultIOne                    BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 1
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultITwo                    BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 2
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultIThree                  BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 3
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultISeven                  BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 7
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultIEight                  BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 8
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultINine                   BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 9
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultITen                    BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 10
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultITwelve                 BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 12
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultIThirteen               BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 13
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultIFourteen               BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 14
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultIFifteen                BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 15
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultISixteen                BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 16
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultISeventeen              BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 17
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultIEighteen               BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 18
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultINineteen               BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 19
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultITwenty                 BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 20
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultITwentyOne              BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 21
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultITwentyTwo              BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 22
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultITwentyThree            BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 23
	BatchResponseInProgressBatchFilesFilesInBatchScanAllResultITwoHundredAndFiftyFive BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI = 255
)

func (e BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI) ToPointer() *BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI {
	return &e
}

func (e *BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		fallthrough
	case 10:
		fallthrough
	case 12:
		fallthrough
	case 13:
		fallthrough
	case 14:
		fallthrough
	case 15:
		fallthrough
	case 16:
		fallthrough
	case 17:
		fallthrough
	case 18:
		fallthrough
	case 19:
		fallthrough
	case 20:
		fallthrough
	case 21:
		fallthrough
	case 22:
		fallthrough
	case 23:
		fallthrough
	case 255:
		*e = BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI: %v", v)
	}
}

type BatchResponseInProgressBatchFilesFilesInBatch struct {
	// Unique identifer for the file.
	DataID *string `json:"data_id,omitempty"`
	// Total number of engines that detected this file.
	DetectedBy *int64 `json:"detected_by,omitempty"`
	// The filename reported via `filename` header.
	DisplayName *string `json:"display_name,omitempty"`
	// Total file size in bytes.
	FileSize *int64 `json:"file_size,omitempty"`
	// The filetype using mimetype.
	FileType *string `json:"file_type,omitempty"`
	// The filetype in human readable format.
	FileTypeDescription *string `json:"file_type_description,omitempty"`
	// The analysis summary
	ProcessInfo *BatchResponseInProgressBatchFilesFilesInBatchProcessInfo `json:"process_info,omitempty"`
	// Track analysis progress until reaches 100.
	ProgressPercentage *int64 `json:"progress_percentage,omitempty"`
	// The overall scan result as string
	ScanAllResultA *BatchResponseInProgressBatchFilesFilesInBatchScanAllResultA `json:"scan_all_result_a,omitempty"`
	// The overall scan result as index in the Processing Results table.
	ScanAllResultI *BatchResponseInProgressBatchFilesFilesInBatchScanAllResultI `json:"scan_all_result_i,omitempty"`
	// The total number of engines used to analyze this file.
	ScannedWith *int64 `json:"scanned_with,omitempty"`
}

// BatchResponseInProgressBatchFiles - Information about the files included in this batch.
type BatchResponseInProgressBatchFiles struct {
	// How many files/entries in the batch.
	BatchCount *int64 `json:"batch_count,omitempty"`
	// The list of files in this batch.
	FilesInBatch []BatchResponseInProgressBatchFilesFilesInBatch `json:"files_in_batch,omitempty"`
	// The starting index in the batch. Used for pagination.
	FirstIndex *int64 `json:"first_index,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"page_size,omitempty"`
}

// BatchResponseInProgressProcessInfo - Overall batch process result
type BatchResponseInProgressProcessInfo struct {
	// Provides the reason why the file is blocked (if so).
	BlockedReason *string `json:"blocked_reason,omitempty"`
	// Indicates if the input file's detected type was configured to skip scanning.
	FileTypeSkippedScan *bool `json:"file_type_skipped_scan,omitempty"`
	// The used rule name.
	Profile *string `json:"profile,omitempty"`
	// The final result of processing the file (Allowed / Blocked / Processing).
	Result *string `json:"result,omitempty"`
	// Identifier for the REST Client that calls the API.
	UserAgent *string `json:"user_agent,omitempty"`
	// User identifier who submitted scan request earlier.
	Username *string `json:"username,omitempty"`
}

// BatchResponseInProgressScanResultsScanAllResultA - Processing result and its index
// * `No Threat Detected`: 0
// * `Infected`: 1
// * `Suspicious`: 2
// * `Failed`: 3
// * `Whitelisted`: 7
// * `Blacklisted`: 8
// * `Exceeded Archive Depth`: 9
// * `Not Scanned`: 10
// * `Encrypted Archive`: 12
// * `Exceeded Archive Size`: 13
// * `Exceeded Archive File Number`: 14
// * `Password Protected Document`: 15
// * `Exceeded Archive Timeout`: 16
// * `Mismatch`: 17
// * `Potentially Vulnerable File`: 18
// * `Cancelled`: 19
// * `Sensitive Data Found`: 20
// * `Yara Rule Matched`: 21
// * `Potentially Unwanted`: 22
// * `Unsupported File Type`: 23
// * `In Progress`: 255
type BatchResponseInProgressScanResultsScanAllResultA string

const (
	BatchResponseInProgressScanResultsScanAllResultANoThreatDetected          BatchResponseInProgressScanResultsScanAllResultA = "No Threat Detected"
	BatchResponseInProgressScanResultsScanAllResultAInfected                  BatchResponseInProgressScanResultsScanAllResultA = "Infected"
	BatchResponseInProgressScanResultsScanAllResultASuspicious                BatchResponseInProgressScanResultsScanAllResultA = "Suspicious"
	BatchResponseInProgressScanResultsScanAllResultAFailed                    BatchResponseInProgressScanResultsScanAllResultA = "Failed"
	BatchResponseInProgressScanResultsScanAllResultAWhitelisted               BatchResponseInProgressScanResultsScanAllResultA = "Whitelisted"
	BatchResponseInProgressScanResultsScanAllResultABlacklisted               BatchResponseInProgressScanResultsScanAllResultA = "Blacklisted"
	BatchResponseInProgressScanResultsScanAllResultAExceededArchiveDepth      BatchResponseInProgressScanResultsScanAllResultA = "Exceeded Archive Depth"
	BatchResponseInProgressScanResultsScanAllResultANotScanned                BatchResponseInProgressScanResultsScanAllResultA = "Not Scanned"
	BatchResponseInProgressScanResultsScanAllResultAEncryptedArchive          BatchResponseInProgressScanResultsScanAllResultA = "Encrypted Archive"
	BatchResponseInProgressScanResultsScanAllResultAExceededArchiveSize       BatchResponseInProgressScanResultsScanAllResultA = "Exceeded Archive Size"
	BatchResponseInProgressScanResultsScanAllResultAExceededArchiveFileNumber BatchResponseInProgressScanResultsScanAllResultA = "Exceeded Archive File Number"
	BatchResponseInProgressScanResultsScanAllResultAPasswordProtectedDocument BatchResponseInProgressScanResultsScanAllResultA = "Password Protected Document"
	BatchResponseInProgressScanResultsScanAllResultAExceededArchiveTimeout    BatchResponseInProgressScanResultsScanAllResultA = "Exceeded Archive Timeout"
	BatchResponseInProgressScanResultsScanAllResultAMismatch                  BatchResponseInProgressScanResultsScanAllResultA = "Mismatch"
	BatchResponseInProgressScanResultsScanAllResultAPotentiallyVulnerableFile BatchResponseInProgressScanResultsScanAllResultA = "Potentially Vulnerable File"
	BatchResponseInProgressScanResultsScanAllResultACancelled                 BatchResponseInProgressScanResultsScanAllResultA = "Cancelled"
	BatchResponseInProgressScanResultsScanAllResultASensitiveDataFound        BatchResponseInProgressScanResultsScanAllResultA = "Sensitive Data Found"
	BatchResponseInProgressScanResultsScanAllResultAYaraRuleMatched           BatchResponseInProgressScanResultsScanAllResultA = "Yara Rule Matched"
	BatchResponseInProgressScanResultsScanAllResultAPotentiallyUnwanted       BatchResponseInProgressScanResultsScanAllResultA = "Potentially Unwanted"
	BatchResponseInProgressScanResultsScanAllResultAUnsupportedFileType       BatchResponseInProgressScanResultsScanAllResultA = "Unsupported File Type"
	BatchResponseInProgressScanResultsScanAllResultAInProgress                BatchResponseInProgressScanResultsScanAllResultA = "In Progress"
)

func (e BatchResponseInProgressScanResultsScanAllResultA) ToPointer() *BatchResponseInProgressScanResultsScanAllResultA {
	return &e
}

func (e *BatchResponseInProgressScanResultsScanAllResultA) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "No Threat Detected":
		fallthrough
	case "Infected":
		fallthrough
	case "Suspicious":
		fallthrough
	case "Failed":
		fallthrough
	case "Whitelisted":
		fallthrough
	case "Blacklisted":
		fallthrough
	case "Exceeded Archive Depth":
		fallthrough
	case "Not Scanned":
		fallthrough
	case "Encrypted Archive":
		fallthrough
	case "Exceeded Archive Size":
		fallthrough
	case "Exceeded Archive File Number":
		fallthrough
	case "Password Protected Document":
		fallthrough
	case "Exceeded Archive Timeout":
		fallthrough
	case "Mismatch":
		fallthrough
	case "Potentially Vulnerable File":
		fallthrough
	case "Cancelled":
		fallthrough
	case "Sensitive Data Found":
		fallthrough
	case "Yara Rule Matched":
		fallthrough
	case "Potentially Unwanted":
		fallthrough
	case "Unsupported File Type":
		fallthrough
	case "In Progress":
		*e = BatchResponseInProgressScanResultsScanAllResultA(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BatchResponseInProgressScanResultsScanAllResultA: %v", v)
	}
}

// BatchResponseInProgressScanResultsScanAllResultI - Scan result as index in the Processing Results table above
type BatchResponseInProgressScanResultsScanAllResultI int64

const (
	BatchResponseInProgressScanResultsScanAllResultIZero                   BatchResponseInProgressScanResultsScanAllResultI = 0
	BatchResponseInProgressScanResultsScanAllResultIOne                    BatchResponseInProgressScanResultsScanAllResultI = 1
	BatchResponseInProgressScanResultsScanAllResultITwo                    BatchResponseInProgressScanResultsScanAllResultI = 2
	BatchResponseInProgressScanResultsScanAllResultIThree                  BatchResponseInProgressScanResultsScanAllResultI = 3
	BatchResponseInProgressScanResultsScanAllResultISeven                  BatchResponseInProgressScanResultsScanAllResultI = 7
	BatchResponseInProgressScanResultsScanAllResultIEight                  BatchResponseInProgressScanResultsScanAllResultI = 8
	BatchResponseInProgressScanResultsScanAllResultINine                   BatchResponseInProgressScanResultsScanAllResultI = 9
	BatchResponseInProgressScanResultsScanAllResultITen                    BatchResponseInProgressScanResultsScanAllResultI = 10
	BatchResponseInProgressScanResultsScanAllResultITwelve                 BatchResponseInProgressScanResultsScanAllResultI = 12
	BatchResponseInProgressScanResultsScanAllResultIThirteen               BatchResponseInProgressScanResultsScanAllResultI = 13
	BatchResponseInProgressScanResultsScanAllResultIFourteen               BatchResponseInProgressScanResultsScanAllResultI = 14
	BatchResponseInProgressScanResultsScanAllResultIFifteen                BatchResponseInProgressScanResultsScanAllResultI = 15
	BatchResponseInProgressScanResultsScanAllResultISixteen                BatchResponseInProgressScanResultsScanAllResultI = 16
	BatchResponseInProgressScanResultsScanAllResultISeventeen              BatchResponseInProgressScanResultsScanAllResultI = 17
	BatchResponseInProgressScanResultsScanAllResultIEighteen               BatchResponseInProgressScanResultsScanAllResultI = 18
	BatchResponseInProgressScanResultsScanAllResultINineteen               BatchResponseInProgressScanResultsScanAllResultI = 19
	BatchResponseInProgressScanResultsScanAllResultITwenty                 BatchResponseInProgressScanResultsScanAllResultI = 20
	BatchResponseInProgressScanResultsScanAllResultITwentyOne              BatchResponseInProgressScanResultsScanAllResultI = 21
	BatchResponseInProgressScanResultsScanAllResultITwentyTwo              BatchResponseInProgressScanResultsScanAllResultI = 22
	BatchResponseInProgressScanResultsScanAllResultITwentyThree            BatchResponseInProgressScanResultsScanAllResultI = 23
	BatchResponseInProgressScanResultsScanAllResultITwoHundredAndFiftyFive BatchResponseInProgressScanResultsScanAllResultI = 255
)

func (e BatchResponseInProgressScanResultsScanAllResultI) ToPointer() *BatchResponseInProgressScanResultsScanAllResultI {
	return &e
}

func (e *BatchResponseInProgressScanResultsScanAllResultI) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		fallthrough
	case 10:
		fallthrough
	case 12:
		fallthrough
	case 13:
		fallthrough
	case 14:
		fallthrough
	case 15:
		fallthrough
	case 16:
		fallthrough
	case 17:
		fallthrough
	case 18:
		fallthrough
	case 19:
		fallthrough
	case 20:
		fallthrough
	case 21:
		fallthrough
	case 22:
		fallthrough
	case 23:
		fallthrough
	case 255:
		*e = BatchResponseInProgressScanResultsScanAllResultI(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BatchResponseInProgressScanResultsScanAllResultI: %v", v)
	}
}

// BatchResponseInProgressScanResults - Metascan analysis result.
type BatchResponseInProgressScanResults struct {
	// The batch unique identifer
	BatchID *string `json:"batch_id,omitempty"`
	// The overall scan result as string
	ScanAllResultA *BatchResponseInProgressScanResultsScanAllResultA `json:"scan_all_result_a,omitempty"`
	// The overall scan result as index in the Processing Results table.
	ScanAllResultI *BatchResponseInProgressScanResultsScanAllResultI `json:"scan_all_result_i,omitempty"`
	// Timestamp when the scanning process starts.
	StartTime *string `json:"start_time,omitempty"`
	// Total number of scanning engines used as part of this analysis. Not like files, batch is not processed by engine, so this value is always 0.
	TotalAvs *int64 `json:"total_avs,omitempty"`
	// Total time elapsed during scan (in milliseconds).
	TotalTime *int64 `json:"total_time,omitempty"`
}

// BatchResponseInProgress - The response for a Batch status request.
type BatchResponseInProgress struct {
	// Information about the files included in this batch.
	BatchFiles *BatchResponseInProgressBatchFiles `json:"batch_files,omitempty"`
	// The batch unique identifer
	BatchID *string `json:"batch_id,omitempty"`
	// The batch status (open/close).
	IsClosed *bool `json:"is_closed,omitempty"`
	// Overall batch process result
	ProcessInfo *BatchResponseInProgressProcessInfo `json:"process_info,omitempty"`
	// Metascan analysis result.
	ScanResults *BatchResponseInProgressScanResults `json:"scan_results,omitempty"`
	// Metadata submitted at batch creation.
	UserData *string `json:"user_data,omitempty"`
}
