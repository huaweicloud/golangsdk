package endpoints

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type commonResult struct {
	golangsdk.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a concrete
// Endpoint. An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*Endpoint, error) {
	var s struct {
		Endpoint *Endpoint `json:"endpoint"`
	}
	err := r.ExtractInto(&s)
	return s.Endpoint, err
}

// CreateResult is the response from a Create operation. Call its Extract
// method to interpret it as an Endpoint.
type CreateResult struct {
	commonResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as an Endpoint.
type UpdateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

// GetResult is the response from a Get operation. Call its Extract
// method to interpret it as an Endpoint.
type GetResult struct {
	commonResult
}

// Endpoint describes the entry point for another service's API.
type Endpoint struct {
	// ID is the unique ID of the endpoint.
	ID string `json:"id"`

	// Availability is the interface type of the Endpoint (admin, internal,
	// or public), referenced by the golangsdk.Availability type.
	Availability golangsdk.Availability `json:"interface"`

	// Name is the name of the Endpoint.
	Name string `json:"name"`

	// Region is the region the Endpoint is located in.
	Region string `json:"region"`

	// ServiceID is the ID of the service the Endpoint refers to.
	ServiceID string `json:"service_id"`

	// URL is the url of the Endpoint.
	URL string `json:"url"`

	// RegionID is the ID of the region the Endpoint is located in.
	RegionID string `json:"region_id"`

	// Enabled is the availablity of the Endpoint.
	Enabled bool `json:"enabled"`

	// Links is the links of the Endpoint
	Links map[string]interface{} `json:"links"`
}

// EndpointPage is a single page of Endpoint results.
type EndpointPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if no Endpoints were returned.
func (r EndpointPage) IsEmpty() (bool, error) {
	es, err := ExtractEndpoints(r)
	return len(es) == 0, err
}

// ExtractEndpoints extracts an Endpoint slice from a Page.
func ExtractEndpoints(r pagination.Page) ([]Endpoint, error) {
	var s struct {
		Endpoints []Endpoint `json:"endpoints"`
	}
	err := (r.(EndpointPage)).ExtractInto(&s)
	return s.Endpoints, err
}
