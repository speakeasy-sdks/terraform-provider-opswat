// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PostProxyRequest - basic information
type PostProxyRequest struct {
	// This redundant attribute will be removed soon.
	Enabled bool `json:"enabled"`
	// Not use the proxy server for the addresses starting with the following entries
	Exclusion *string `json:"exclusion,omitempty"`
	// Password for proxy authentication
	Password *string `json:"password,omitempty"`
	// Proxy server port
	Port string `json:"port"`
	// Proxy server host address
	Server string `json:"server"`
	// Username for proxy authentication
	Username *string `json:"username,omitempty"`
}
