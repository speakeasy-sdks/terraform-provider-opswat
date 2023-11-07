// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"Metadefender/internal/sdk/pkg/utils"
	"errors"
	"net/http"
)

type ConfigUpdateProxyRequest struct {
	ProxyListRequestBody *shared.ProxyListRequestBody `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *ConfigUpdateProxyRequest) GetProxyListRequestBody() *shared.ProxyListRequestBody {
	if o == nil {
		return nil
	}
	return o.ProxyListRequestBody
}

func (o *ConfigUpdateProxyRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

type ConfigUpdateProxyResponseBodyType string

const (
	ConfigUpdateProxyResponseBodyTypeMissingEnabledField         ConfigUpdateProxyResponseBodyType = "MissingEnabledField"
	ConfigUpdateProxyResponseBodyTypeErrorWhileParsingInputJSON  ConfigUpdateProxyResponseBodyType = "ErrorWhileParsingInputJson"
	ConfigUpdateProxyResponseBodyTypeMissingPort                 ConfigUpdateProxyResponseBodyType = "MissingPort"
	ConfigUpdateProxyResponseBodyTypeMissingServerAddress        ConfigUpdateProxyResponseBodyType = "MissingServerAddress"
	ConfigUpdateProxyResponseBodyTypeProxyRequiresAuthentication ConfigUpdateProxyResponseBodyType = "ProxyRequiresAuthentication"
)

type ConfigUpdateProxyResponseBody struct {
	MissingEnabledField         *shared.MissingEnabledField
	ErrorWhileParsingInputJSON  *shared.ErrorWhileParsingInputJSON
	MissingPort                 *shared.MissingPort
	MissingServerAddress        *shared.MissingServerAddress
	ProxyRequiresAuthentication *shared.ProxyRequiresAuthentication

	Type ConfigUpdateProxyResponseBodyType
}

func CreateConfigUpdateProxyResponseBodyMissingEnabledField(missingEnabledField shared.MissingEnabledField) ConfigUpdateProxyResponseBody {
	typ := ConfigUpdateProxyResponseBodyTypeMissingEnabledField

	return ConfigUpdateProxyResponseBody{
		MissingEnabledField: &missingEnabledField,
		Type:                typ,
	}
}

func CreateConfigUpdateProxyResponseBodyErrorWhileParsingInputJSON(errorWhileParsingInputJSON shared.ErrorWhileParsingInputJSON) ConfigUpdateProxyResponseBody {
	typ := ConfigUpdateProxyResponseBodyTypeErrorWhileParsingInputJSON

	return ConfigUpdateProxyResponseBody{
		ErrorWhileParsingInputJSON: &errorWhileParsingInputJSON,
		Type:                       typ,
	}
}

func CreateConfigUpdateProxyResponseBodyMissingPort(missingPort shared.MissingPort) ConfigUpdateProxyResponseBody {
	typ := ConfigUpdateProxyResponseBodyTypeMissingPort

	return ConfigUpdateProxyResponseBody{
		MissingPort: &missingPort,
		Type:        typ,
	}
}

func CreateConfigUpdateProxyResponseBodyMissingServerAddress(missingServerAddress shared.MissingServerAddress) ConfigUpdateProxyResponseBody {
	typ := ConfigUpdateProxyResponseBodyTypeMissingServerAddress

	return ConfigUpdateProxyResponseBody{
		MissingServerAddress: &missingServerAddress,
		Type:                 typ,
	}
}

func CreateConfigUpdateProxyResponseBodyProxyRequiresAuthentication(proxyRequiresAuthentication shared.ProxyRequiresAuthentication) ConfigUpdateProxyResponseBody {
	typ := ConfigUpdateProxyResponseBodyTypeProxyRequiresAuthentication

	return ConfigUpdateProxyResponseBody{
		ProxyRequiresAuthentication: &proxyRequiresAuthentication,
		Type:                        typ,
	}
}

func (u *ConfigUpdateProxyResponseBody) UnmarshalJSON(data []byte) error {

	missingEnabledField := new(shared.MissingEnabledField)
	if err := utils.UnmarshalJSON(data, &missingEnabledField, "", true, true); err == nil {
		u.MissingEnabledField = missingEnabledField
		u.Type = ConfigUpdateProxyResponseBodyTypeMissingEnabledField
		return nil
	}

	errorWhileParsingInputJSON := new(shared.ErrorWhileParsingInputJSON)
	if err := utils.UnmarshalJSON(data, &errorWhileParsingInputJSON, "", true, true); err == nil {
		u.ErrorWhileParsingInputJSON = errorWhileParsingInputJSON
		u.Type = ConfigUpdateProxyResponseBodyTypeErrorWhileParsingInputJSON
		return nil
	}

	missingPort := new(shared.MissingPort)
	if err := utils.UnmarshalJSON(data, &missingPort, "", true, true); err == nil {
		u.MissingPort = missingPort
		u.Type = ConfigUpdateProxyResponseBodyTypeMissingPort
		return nil
	}

	missingServerAddress := new(shared.MissingServerAddress)
	if err := utils.UnmarshalJSON(data, &missingServerAddress, "", true, true); err == nil {
		u.MissingServerAddress = missingServerAddress
		u.Type = ConfigUpdateProxyResponseBodyTypeMissingServerAddress
		return nil
	}

	proxyRequiresAuthentication := new(shared.ProxyRequiresAuthentication)
	if err := utils.UnmarshalJSON(data, &proxyRequiresAuthentication, "", true, true); err == nil {
		u.ProxyRequiresAuthentication = proxyRequiresAuthentication
		u.Type = ConfigUpdateProxyResponseBodyTypeProxyRequiresAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConfigUpdateProxyResponseBody) MarshalJSON() ([]byte, error) {
	if u.MissingEnabledField != nil {
		return utils.MarshalJSON(u.MissingEnabledField, "", true)
	}

	if u.ErrorWhileParsingInputJSON != nil {
		return utils.MarshalJSON(u.ErrorWhileParsingInputJSON, "", true)
	}

	if u.MissingPort != nil {
		return utils.MarshalJSON(u.MissingPort, "", true)
	}

	if u.MissingServerAddress != nil {
		return utils.MarshalJSON(u.MissingServerAddress, "", true)
	}

	if u.ProxyRequiresAuthentication != nil {
		return utils.MarshalJSON(u.ProxyRequiresAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
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
	OneOf *ConfigUpdateProxyResponseBody
}

func (o *ConfigUpdateProxyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigUpdateProxyResponse) GetProxyList() *shared.ProxyList {
	if o == nil {
		return nil
	}
	return o.ProxyList
}

func (o *ConfigUpdateProxyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigUpdateProxyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigUpdateProxyResponse) GetOneOf() *ConfigUpdateProxyResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
