// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ConfigUpdateProxyRequest struct {
	ProxyListRequestBody *shared.ProxyListRequestBody `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

type ConfigUpdateProxy400ApplicationJSONType string

const (
	ConfigUpdateProxy400ApplicationJSONTypeMissingEnabledField         ConfigUpdateProxy400ApplicationJSONType = "MissingEnabledField"
	ConfigUpdateProxy400ApplicationJSONTypeErrorWhileParsingInputJSON  ConfigUpdateProxy400ApplicationJSONType = "ErrorWhileParsingInputJson"
	ConfigUpdateProxy400ApplicationJSONTypeMissingPort                 ConfigUpdateProxy400ApplicationJSONType = "MissingPort"
	ConfigUpdateProxy400ApplicationJSONTypeMissingServerAddress        ConfigUpdateProxy400ApplicationJSONType = "MissingServerAddress"
	ConfigUpdateProxy400ApplicationJSONTypeProxyRequiresAuthentication ConfigUpdateProxy400ApplicationJSONType = "ProxyRequiresAuthentication"
)

type ConfigUpdateProxy400ApplicationJSON struct {
	MissingEnabledField         *shared.MissingEnabledField
	ErrorWhileParsingInputJSON  *shared.ErrorWhileParsingInputJSON
	MissingPort                 *shared.MissingPort
	MissingServerAddress        *shared.MissingServerAddress
	ProxyRequiresAuthentication *shared.ProxyRequiresAuthentication

	Type ConfigUpdateProxy400ApplicationJSONType
}

func CreateConfigUpdateProxy400ApplicationJSONMissingEnabledField(missingEnabledField shared.MissingEnabledField) ConfigUpdateProxy400ApplicationJSON {
	typ := ConfigUpdateProxy400ApplicationJSONTypeMissingEnabledField

	return ConfigUpdateProxy400ApplicationJSON{
		MissingEnabledField: &missingEnabledField,
		Type:                typ,
	}
}

func CreateConfigUpdateProxy400ApplicationJSONErrorWhileParsingInputJSON(errorWhileParsingInputJSON shared.ErrorWhileParsingInputJSON) ConfigUpdateProxy400ApplicationJSON {
	typ := ConfigUpdateProxy400ApplicationJSONTypeErrorWhileParsingInputJSON

	return ConfigUpdateProxy400ApplicationJSON{
		ErrorWhileParsingInputJSON: &errorWhileParsingInputJSON,
		Type:                       typ,
	}
}

func CreateConfigUpdateProxy400ApplicationJSONMissingPort(missingPort shared.MissingPort) ConfigUpdateProxy400ApplicationJSON {
	typ := ConfigUpdateProxy400ApplicationJSONTypeMissingPort

	return ConfigUpdateProxy400ApplicationJSON{
		MissingPort: &missingPort,
		Type:        typ,
	}
}

func CreateConfigUpdateProxy400ApplicationJSONMissingServerAddress(missingServerAddress shared.MissingServerAddress) ConfigUpdateProxy400ApplicationJSON {
	typ := ConfigUpdateProxy400ApplicationJSONTypeMissingServerAddress

	return ConfigUpdateProxy400ApplicationJSON{
		MissingServerAddress: &missingServerAddress,
		Type:                 typ,
	}
}

func CreateConfigUpdateProxy400ApplicationJSONProxyRequiresAuthentication(proxyRequiresAuthentication shared.ProxyRequiresAuthentication) ConfigUpdateProxy400ApplicationJSON {
	typ := ConfigUpdateProxy400ApplicationJSONTypeProxyRequiresAuthentication

	return ConfigUpdateProxy400ApplicationJSON{
		ProxyRequiresAuthentication: &proxyRequiresAuthentication,
		Type:                        typ,
	}
}

func (u *ConfigUpdateProxy400ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	missingEnabledField := new(shared.MissingEnabledField)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&missingEnabledField); err == nil {
		u.MissingEnabledField = missingEnabledField
		u.Type = ConfigUpdateProxy400ApplicationJSONTypeMissingEnabledField
		return nil
	}

	errorWhileParsingInputJSON := new(shared.ErrorWhileParsingInputJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&errorWhileParsingInputJSON); err == nil {
		u.ErrorWhileParsingInputJSON = errorWhileParsingInputJSON
		u.Type = ConfigUpdateProxy400ApplicationJSONTypeErrorWhileParsingInputJSON
		return nil
	}

	missingPort := new(shared.MissingPort)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&missingPort); err == nil {
		u.MissingPort = missingPort
		u.Type = ConfigUpdateProxy400ApplicationJSONTypeMissingPort
		return nil
	}

	missingServerAddress := new(shared.MissingServerAddress)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&missingServerAddress); err == nil {
		u.MissingServerAddress = missingServerAddress
		u.Type = ConfigUpdateProxy400ApplicationJSONTypeMissingServerAddress
		return nil
	}

	proxyRequiresAuthentication := new(shared.ProxyRequiresAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&proxyRequiresAuthentication); err == nil {
		u.ProxyRequiresAuthentication = proxyRequiresAuthentication
		u.Type = ConfigUpdateProxy400ApplicationJSONTypeProxyRequiresAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConfigUpdateProxy400ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.MissingEnabledField != nil {
		return json.Marshal(u.MissingEnabledField)
	}

	if u.ErrorWhileParsingInputJSON != nil {
		return json.Marshal(u.ErrorWhileParsingInputJSON)
	}

	if u.MissingPort != nil {
		return json.Marshal(u.MissingPort)
	}

	if u.MissingServerAddress != nil {
		return json.Marshal(u.MissingServerAddress)
	}

	if u.ProxyRequiresAuthentication != nil {
		return json.Marshal(u.ProxyRequiresAuthentication)
	}

	return nil, nil
}

type ConfigUpdateProxyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A list of all proxy.
	ProxyList *shared.ProxyList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request (e.g. invalid header, apikey is missing or invalid).
	ConfigUpdateProxy400ApplicationJSONOneOf *ConfigUpdateProxy400ApplicationJSON
}
