package tags

import (
	"github.com/huaweicloud/golangsdk"
)

// ListResult is the result list of tags
type ListResult struct {
	golangsdk.Result
}

// Extract extract http response to golang struct
func (r ListResult) Extract() (*ServerTags, error) {
	var t ServerTags
	err := r.Result.ExtractInto(&t)
	return &t, err
}

// ServerTags represents server tag list
type ServerTags struct {
	Tags []string `json:"tags"`
}
