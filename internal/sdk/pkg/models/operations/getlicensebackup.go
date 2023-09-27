// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetLicenseBackupRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// GetLicenseBackup500ApplicationJSON - Unexpected event on server
type GetLicenseBackup500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// GetLicenseBackup403ApplicationJSON - Invalid user information or Not Allowed
type GetLicenseBackup403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// GetLicenseBackup400ApplicationJSON - Bad Request (e.g. invalid header, apikey is missing or invalid).
type GetLicenseBackup400ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// GetLicenseBackup200ApplicationJSON - Request processed successfully
type GetLicenseBackup200ApplicationJSON struct {
	// The instance is a backup or not.
	IsBackup *bool `json:"is_backup,omitempty"`
}

type GetLicenseBackupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request processed successfully
	GetLicenseBackup200ApplicationJSONObject *GetLicenseBackup200ApplicationJSON
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	GetLicenseBackup400ApplicationJSONObject *GetLicenseBackup400ApplicationJSON
	// Invalid user information or Not Allowed
	GetLicenseBackup403ApplicationJSONObject *GetLicenseBackup403ApplicationJSON
	// Unexpected event on server
	GetLicenseBackup500ApplicationJSONObject *GetLicenseBackup500ApplicationJSON
}
