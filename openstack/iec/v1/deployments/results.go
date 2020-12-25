package deployments

import (
	"github.com/huaweicloud/golangsdk"
)

// Location struct
type Location struct {
	//SiteID site id
	SiteID string `json:"site_id"`

	// Area title
	Area string `json:"area"`

	// Province title
	Province string `json:"province"`

	// City title
	City string `json:"city"`

	// Operator title
	Operator string `json:"operator"`

	//Region  region info
	Region string `json:"-"`

	// AvailabilityZone title
	AvailabilityZone string `json:"-"`

	// Count title
	Count int `json:"stack_count"`
}

type CreateResp struct {
	ID        string     `json:"id"`
	Locations []Location `json:"locations"`
}

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) ExtractCreateResult() (*CreateResp, error) {
	var entity CreateResp
	err := r.ExtractIntoStructPtr(&entity, "")
	return &entity, err
}

type DeployResult struct {
	commonResult
}

type DeployResp struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DeploymentID string `json:"deployment_id"`
	Status       string `orm:"column(STATUS)" json:"status"`
}

func (r DeployResult) ExtractDeployOrExpandResult() (*DeployResp, error) {
	var entity DeployResp
	err := r.ExtractIntoStructPtr(&entity, "")
	return &entity, err
}

type ExpandResult struct {
	commonResult
}

func (r ExpandResult) ExtractExpandResult() (*DeployResp, error) {
	var entity DeployResp
	err := r.ExtractIntoStructPtr(&entity, "")
	return &entity, err
}

type DeleteResult struct {
	golangsdk.ErrResult
}
