package environments

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type commonResult struct {
	golangsdk.Result
}

// CreateResult represents a result of the Create method.
type CreateResult struct {
	commonResult
}

// UpdateResult represents a result of the Update method.
type UpdateResult struct {
	commonResult
}

type Environment struct {
	// Environment ID.
	Id string `json:"id"`
	// Environment name.
	Name string `json:"name"`
	// Description.
	Description string `json:"remark"`
	// Create time, in RFC-3339 format.
	CreateTime string `json:"create_time"`
}

func (r commonResult) Extract() (*Environment, error) {
	var s Environment
	err := r.ExtractInto(&s)
	return &s, err
}

// EnvironmentPage represents the response pages of the List method.
type EnvironmentPage struct {
	pagination.SinglePageBase
}

func ExtractEnvironments(r pagination.Page) ([]Environment, error) {
	var s []Environment
	err := r.(EnvironmentPage).Result.ExtractIntoSlicePtr(&s, "envs")
	return s, err
}

// DeleteResult represents a result of the Delete method.
type DeleteResult struct {
	golangsdk.ErrResult
}
