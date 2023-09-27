// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AdminConfigCustomResponseHeaderIncludeAddress - To include data from address_str.
type AdminConfigCustomResponseHeaderIncludeAddress struct {
	// IP address.
	AddressStr *string `json:"address_str,omitempty"`
	// Include data from address_str.
	Enabled *bool `json:"enabled,omitempty"`
}

type AdminConfigCustomResponseHeaderSpecifyUsersAgent struct {
	// Indicate which user_agent should be allowed to include new header in response header.
	AllowedUsersAgent []string `json:"allowed_users_agent,omitempty"`
	// True: required matched header user_agent to given list in allowed_users_agent
	// to return new header X-Core-Id in response.
	//
	Enabled *bool `json:"enabled,omitempty"`
}

// AdminConfigCustomResponseHeader - Custom response header details.
type AdminConfigCustomResponseHeader struct {
	// Only return new X-Core-Id header for authorized user.
	AuthorizedOnly *bool `json:"authorized_only,omitempty"`
	// To include data from address_str.
	IncludeAddress *AdminConfigCustomResponseHeaderIncludeAddress `json:"include_address,omitempty"`
	// Include generated deployment ID.
	IncludeDeploymentID *bool                                             `json:"include_deployment_id,omitempty"`
	SpecifyUsersAgent   *AdminConfigCustomResponseHeaderSpecifyUsersAgent `json:"specify_users_agent,omitempty"`
}