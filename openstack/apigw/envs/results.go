package envs

import (
	"time"

	"github.com/huaweicloud/golangsdk"
)

// Environment contains all the information associated with an API environment.
type Environment struct {
	// Unique identifier for the environment.
	ID string `json:"id"`
	// Human-readable display name for the environment.
	Name string `json:"name"`
	// Description of the environment.
	Remark string `json:"remark"`
	// Time when the environment is created
	CreateAt time.Time `json:"-"`
}

type commonResult struct {
	golangsdk.Result
}

// Extract will get the Environment object out of the commonResult object.
func (r commonResult) Extract() (*Environment, error) {
	var env Environment
	err := r.ExtractInto(&env)
	return &env, err
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// UpdateResult contains the response body and error from an Update request.
type UpdateResult struct {
	commonResult
}

// ListResult contains the response body and error from a List request.
type ListResult struct {
	commonResult
}

// Extract will get the Environment object out of the ListResult object.
func (r ListResult) Extract() ([]Environment, error) {
	var s struct {
		Envs []Environment `json:"envs"`
	}
	err := r.ExtractInto(&s)
	return s.Envs, err
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	golangsdk.ErrResult
}
