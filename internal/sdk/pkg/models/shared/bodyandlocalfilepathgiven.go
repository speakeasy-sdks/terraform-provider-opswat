// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BodyAndLocalFilePathGiven struct {
	// Both body data and local file path were given.
	Err *string `json:"err,omitempty"`
}

func (o *BodyAndLocalFilePathGiven) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}
