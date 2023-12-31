// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DecryptedStatus - Indicate that decryption phase is successful or not.
type DecryptedStatus string

const (
	DecryptedStatusSuccess DecryptedStatus = "Success"
	DecryptedStatusFailed  DecryptedStatus = "Failed"
)

func (e DecryptedStatus) ToPointer() *DecryptedStatus {
	return &e
}

func (e *DecryptedStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Success":
		fallthrough
	case "Failed":
		*e = DecryptedStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DecryptedStatus: %v", v)
	}
}

// ExtractionInfo - Details for archive extraction.
type ExtractionInfo struct {
	// Indicate that decryption phase is successful or not.
	DecryptedStatus *DecryptedStatus `json:"decrypted_status,omitempty"`
	// Error category
	ErrCategory *string `json:"err_category,omitempty"`
	// Error code
	ErrCode *int64 `json:"err_code,omitempty"`
	// Error message
	ErrDetails *string `json:"err_details,omitempty"`
	// Indicate if file is password-protected or not.
	IsEncryptedFile *bool `json:"is_encrypted_file,omitempty"`
}

func (o *ExtractionInfo) GetDecryptedStatus() *DecryptedStatus {
	if o == nil {
		return nil
	}
	return o.DecryptedStatus
}

func (o *ExtractionInfo) GetErrCategory() *string {
	if o == nil {
		return nil
	}
	return o.ErrCategory
}

func (o *ExtractionInfo) GetErrCode() *int64 {
	if o == nil {
		return nil
	}
	return o.ErrCode
}

func (o *ExtractionInfo) GetErrDetails() *string {
	if o == nil {
		return nil
	}
	return o.ErrDetails
}

func (o *ExtractionInfo) GetIsEncryptedFile() *bool {
	if o == nil {
		return nil
	}
	return o.IsEncryptedFile
}
