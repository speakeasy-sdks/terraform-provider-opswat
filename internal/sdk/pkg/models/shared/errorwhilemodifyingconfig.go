// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ErrorWhileModifyingConfig struct {
	// Error while modifying configuration
	//
	// Setting timeout value out of the limit range [1, 30] minutes, or the request body is not in JSON format.
	//
	Err *string `json:"err,omitempty"`
}

func (o *ErrorWhileModifyingConfig) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}
