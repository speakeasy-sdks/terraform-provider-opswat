// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type InternalError struct {
	// Failed to request scan. Try again later.
	Err *string `json:"err,omitempty"`
}

func (o *InternalError) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}
