// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BatchNotFound struct {
	// Can not scan in given batch.
	Err *string `json:"err,omitempty"`
}

func (o *BatchNotFound) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}
