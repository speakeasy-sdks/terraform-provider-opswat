// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NoAvailableRule struct {
	// No available rule is present for scanning.
	Err *string `json:"err,omitempty"`
}

func (o *NoAvailableRule) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}
