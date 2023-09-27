// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Metadefender/internal/sdk/pkg/models/shared"
	"net/http"
)

type LicenseGetRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

// LicenseGet500ApplicationJSON - Unexpected event on server
type LicenseGet500ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

// LicenseGet403ApplicationJSON - Invalid user information or Not Allowed
type LicenseGet403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

type LicenseGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Information about the licensed product (product type, number of activations, deploymentId, expiration date and days left)
	LicenseInformation *shared.LicenseInformation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid user information or Not Allowed
	LicenseGet403ApplicationJSONObject *LicenseGet403ApplicationJSON
	// Unexpected event on server
	LicenseGet500ApplicationJSONObject *LicenseGet500ApplicationJSON
}