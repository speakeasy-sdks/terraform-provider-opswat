// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type InvalidAPIKeyGiven struct {
	// Invalid apikey given
	Err *string `json:"err,omitempty"`
}

func (o *InvalidAPIKeyGiven) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}
