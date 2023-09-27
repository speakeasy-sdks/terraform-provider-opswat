// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type YaraSourcesPutRequest struct {
	// A list of Yara Engine sources.
	YaraSourcesObject *shared.YaraSourcesObject `request:"mediaType=application/json"`
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// YaraSourcesPut500ApplicationJSON - Unexpected event on server
type YaraSourcesPut500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// YaraSourcesPut403ApplicationJSON - Invalid user information or Not Allowed
type YaraSourcesPut403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

type YaraSourcesPutResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A list of Yara Engine sources.
	YaraSourcesObject *shared.YaraSourcesObject
	// Invalid user information or Not Allowed
	YaraSourcesPut403ApplicationJSONObject *YaraSourcesPut403ApplicationJSON
	// Unexpected event on server
	YaraSourcesPut500ApplicationJSONObject *YaraSourcesPut500ApplicationJSON
}
