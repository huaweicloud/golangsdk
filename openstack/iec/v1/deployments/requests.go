package deployments

import (
	"net/http"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/iec/v1/common"
)

// EdgeCloud 边缘业务
type CreateOpts struct {
	EdgeCloud EdgeCloud `json:"edgecloud"`
}

type EdgeCloud struct {
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	StackObj    StackDetail     `json:"stack"`
	CoverageObj common.Coverage `json:"coverage"`
}

// StackDetail Stack详情
type StackDetail struct {
	//ID
	ID string `json:"id"`

	//NAME
	Name string `json:"name"`

	Resources SliceResourceOptsField `json:"resources"`
}

// SliceResourceOptsField A slice string field.
type SliceResourceOptsField []common.ResourceOpts

type CreateOptsBuilder interface {
	ToDeploymentCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToDeploymentCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToDeploymentCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	url := CreateURL(client)
	_, r.Err = client.Post(url, b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	return
}

func Deploy(client *golangsdk.ServiceClient, deploymentID string) (r DeployResult) {

	url := DeployURL(client, deploymentID)
	_, r.Err = client.Post(url, "", &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	return
}

func Expand(client *golangsdk.ServiceClient, deploymentID string) (r DeployResult) {

	url := ExpandURL(client, deploymentID)
	_, r.Err = client.Post(url, "", &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, deploymentID string) (r DeleteResult) {

	_, r.Err = client.Delete(DeleteURL(client, deploymentID), &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusNoContent},
	})
	return
}
