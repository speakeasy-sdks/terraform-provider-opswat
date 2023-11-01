// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminExportRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *AdminExportRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// AdminExport500ApplicationJSON - Failed to export configuration
type AdminExport500ApplicationJSON struct {
	Err *string `json:"err,omitempty"`
}

func (o *AdminExport500ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// AdminExport405ApplicationJSON - The user has no rights for this operation.
type AdminExport405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *AdminExport405ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// AdminExport403ApplicationJSON - Invalid user information or Not Allowed
type AdminExport403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *AdminExport403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// AdminExport400ApplicationJSON - Bad Request (e.g. invalid header, apikey is missing or invalid).
type AdminExport400ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *AdminExport400ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

type AdminExportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request processed successfully
	AdminExport200ApplicationJSONFileString *string
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	AdminExport400ApplicationJSONObject *AdminExport400ApplicationJSON
	// Invalid user information or Not Allowed
	AdminExport403ApplicationJSONObject *AdminExport403ApplicationJSON
	// The user has no rights for this operation.
	AdminExport405ApplicationJSONObject *AdminExport405ApplicationJSON
	// Failed to export configuration
	AdminExport500ApplicationJSONObject *AdminExport500ApplicationJSON
}

func (o *AdminExportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AdminExportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AdminExportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AdminExportResponse) GetAdminExport200ApplicationJSONFileString() *string {
	if o == nil {
		return nil
	}
	return o.AdminExport200ApplicationJSONFileString
}

func (o *AdminExportResponse) GetAdminExport400ApplicationJSONObject() *AdminExport400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AdminExport400ApplicationJSONObject
}

func (o *AdminExportResponse) GetAdminExport403ApplicationJSONObject() *AdminExport403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AdminExport403ApplicationJSONObject
}

func (o *AdminExportResponse) GetAdminExport405ApplicationJSONObject() *AdminExport405ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AdminExport405ApplicationJSONObject
}

func (o *AdminExportResponse) GetAdminExport500ApplicationJSONObject() *AdminExport500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AdminExport500ApplicationJSONObject
}
