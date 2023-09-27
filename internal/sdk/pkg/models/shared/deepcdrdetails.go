// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeepCDRDetailsDetailsAction - The type of action that was performed
type DeepCDRDetailsDetailsAction string

const (
	DeepCDRDetailsDetailsActionSanitized DeepCDRDetailsDetailsAction = "sanitized"
	DeepCDRDetailsDetailsActionRemoved   DeepCDRDetailsDetailsAction = "removed"
)

func (e DeepCDRDetailsDetailsAction) ToPointer() *DeepCDRDetailsDetailsAction {
	return &e
}

func (e *DeepCDRDetailsDetailsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sanitized":
		fallthrough
	case "removed":
		*e = DeepCDRDetailsDetailsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeepCDRDetailsDetailsAction: %v", v)
	}
}

// DeepCDRDetailsDetailsDetailsAction - The type of action that was performed
type DeepCDRDetailsDetailsDetailsAction string

const (
	DeepCDRDetailsDetailsDetailsActionSanitized DeepCDRDetailsDetailsDetailsAction = "sanitized"
	DeepCDRDetailsDetailsDetailsActionRemoved   DeepCDRDetailsDetailsDetailsAction = "removed"
)

func (e DeepCDRDetailsDetailsDetailsAction) ToPointer() *DeepCDRDetailsDetailsDetailsAction {
	return &e
}

func (e *DeepCDRDetailsDetailsDetailsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sanitized":
		fallthrough
	case "removed":
		*e = DeepCDRDetailsDetailsDetailsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeepCDRDetailsDetailsDetailsAction: %v", v)
	}
}

// DeepCDRDetailsDetailsDetails - List of all sanitized objects of a sanitized embedded file
type DeepCDRDetailsDetailsDetails struct {
	// The type of action that was performed
	Action *DeepCDRDetailsDetailsDetailsAction `json:"action,omitempty"`
	// The number of objects that were sanitized/removed.
	Count *int64 `json:"count,omitempty"`
	// Additional information about the sanitized object
	ObjectDetails []string `json:"object_details,omitempty"`
	// The object type that was sanitized/removed.
	ObjectName *string `json:"object_name,omitempty"`
}

type DeepCDRDetailsDetails struct {
	// The type of action that was performed
	Action DeepCDRDetailsDetailsAction `json:"action"`
	// The number of objects that were sanitized/removed.
	Count *int64 `json:"count,omitempty"`
	// Action was successful or not.
	Description *string `json:"description,omitempty"`
	// List of all sanitized objects of a sanitized embedded file
	Details *DeepCDRDetailsDetailsDetails `json:"details,omitempty"`
	// If an embedded file was sanitized.
	FileName *string `json:"file_name,omitempty"`
	// Additional information about the sanitized object
	ObjectDetails []string `json:"object_details,omitempty"`
	// The object type that was sanitized/removed.
	ObjectName string `json:"object_name"`
}

// DeepCDRDetailsSanitizedFileInfo - Information of sanitized file.
// Only applicable to individual file sanitization, or original archive document sanitization level.
type DeepCDRDetailsSanitizedFileInfo struct {
	// Size of sanitized file in bytes.
	FileSize *int64 `json:"file_size,omitempty"`
	// SHA256 hash of sanitized file.
	Sha256 *string `json:"sha256,omitempty"`
}

// DeepCDRDetails - Deep CDR module returns forensic info to describe what happened during the process in the case file was successfully sanitized.
type DeepCDRDetails struct {
	// Action was successful or not.
	Description *string `json:"description,omitempty"`
	// List of all sanitized objects
	Details []DeepCDRDetailsDetails `json:"details,omitempty"`
	// Deep CDR errors are classified into different categories.
	//
	// For more details, please find [Troubleshooting sanitization failures](https://docs.opswat.com/mdcore/deep-cdr/troubleshooting-sanitization-failures)
	//
	FailureCategory *string `json:"failure_category,omitempty"`
	// Information of sanitized file.
	// Only applicable to individual file sanitization, or original archive document sanitization level.
	//
	SanitizedFileInfo *DeepCDRDetailsSanitizedFileInfo `json:"sanitized_file_info,omitempty"`
}