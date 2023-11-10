// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UserLogin struct {
	// The randomly generated token used to prevent CSRF attacks
	OmsCsrfToken string `json:"oms-csrf-token"`
	// The apikey used to make API calls which requires authentication
	SessionID string `json:"session_id"`
}

func (o *UserLogin) GetOmsCsrfToken() string {
	if o == nil {
		return ""
	}
	return o.OmsCsrfToken
}

func (o *UserLogin) GetSessionID() string {
	if o == nil {
		return ""
	}
	return o.SessionID
}
