package snapshots

import (
	"github.com/huaweicloud/golangsdk"
)

// Policy contains all the information associated with a snapshot policy.
type Policy struct {
	KeepDay  int    `json:"keepday"`
	Period   string `json:"period"`
	Prefix   string `json:"prefix"`
	Bucket   string `json:"bucket"`
	BasePath string `json:"basePath"`
	Agency   string `json:"agency"`
	Enable   string `json:"enable"`
}

type commonResult struct {
	golangsdk.Result
}

// PolicyResult contains the response body and error from a policy request.
type PolicyResult struct {
	commonResult
}

// ErrorResult contains the response body and error from a request.
type ErrorResult struct {
	golangsdk.ErrResult
}

// Extract will get the Policy object out of the PolicyResult object.
func (r PolicyResult) Extract() (*Policy, error) {
	var pol Policy
	err := r.ExtractInto(&pol)
	return &pol, err
}
