// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ConfigUpdateSyncScanTimeoutRequest struct {
	AdminConfigFileSync *shared.AdminConfigFileSync `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

type ConfigUpdateSyncScanTimeout500ApplicationJSONType string

const (
	ConfigUpdateSyncScanTimeout500ApplicationJSONTypeErrorWhileModifyingConfig ConfigUpdateSyncScanTimeout500ApplicationJSONType = "ErrorWhileModifyingConfig"
	ConfigUpdateSyncScanTimeout500ApplicationJSONTypeInternalServerError       ConfigUpdateSyncScanTimeout500ApplicationJSONType = "InternalServerError"
)

type ConfigUpdateSyncScanTimeout500ApplicationJSON struct {
	ErrorWhileModifyingConfig *shared.ErrorWhileModifyingConfig
	InternalServerError       *shared.InternalServerError

	Type ConfigUpdateSyncScanTimeout500ApplicationJSONType
}

func CreateConfigUpdateSyncScanTimeout500ApplicationJSONErrorWhileModifyingConfig(errorWhileModifyingConfig shared.ErrorWhileModifyingConfig) ConfigUpdateSyncScanTimeout500ApplicationJSON {
	typ := ConfigUpdateSyncScanTimeout500ApplicationJSONTypeErrorWhileModifyingConfig

	return ConfigUpdateSyncScanTimeout500ApplicationJSON{
		ErrorWhileModifyingConfig: &errorWhileModifyingConfig,
		Type:                      typ,
	}
}

func CreateConfigUpdateSyncScanTimeout500ApplicationJSONInternalServerError(internalServerError shared.InternalServerError) ConfigUpdateSyncScanTimeout500ApplicationJSON {
	typ := ConfigUpdateSyncScanTimeout500ApplicationJSONTypeInternalServerError

	return ConfigUpdateSyncScanTimeout500ApplicationJSON{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *ConfigUpdateSyncScanTimeout500ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	errorWhileModifyingConfig := new(shared.ErrorWhileModifyingConfig)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&errorWhileModifyingConfig); err == nil {
		u.ErrorWhileModifyingConfig = errorWhileModifyingConfig
		u.Type = ConfigUpdateSyncScanTimeout500ApplicationJSONTypeErrorWhileModifyingConfig
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&internalServerError); err == nil {
		u.InternalServerError = internalServerError
		u.Type = ConfigUpdateSyncScanTimeout500ApplicationJSONTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConfigUpdateSyncScanTimeout500ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.ErrorWhileModifyingConfig != nil {
		return json.Marshal(u.ErrorWhileModifyingConfig)
	}

	if u.InternalServerError != nil {
		return json.Marshal(u.InternalServerError)
	}

	return nil, nil
}

// ConfigUpdateSyncScanTimeout405ApplicationJSON - The user has no rights for this operation.
type ConfigUpdateSyncScanTimeout405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// ConfigUpdateSyncScanTimeout403ApplicationJSON - Invalid user information or Not Allowed
type ConfigUpdateSyncScanTimeout403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

type ConfigUpdateSyncScanTimeoutResponse struct {
	// Request processed successfully.
	AdminConfigFileSync *shared.AdminConfigFileSync
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	ConfigUpdateSyncScanTimeout403ApplicationJSONObject *ConfigUpdateSyncScanTimeout403ApplicationJSON
	// The user has no rights for this operation.
	ConfigUpdateSyncScanTimeout405ApplicationJSONObject *ConfigUpdateSyncScanTimeout405ApplicationJSON
	// Error while modifying configuration.
	ConfigUpdateSyncScanTimeout500ApplicationJSONOneOf *ConfigUpdateSyncScanTimeout500ApplicationJSON
}