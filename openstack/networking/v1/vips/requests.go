package vips

import (
	"github.com/huaweicloud/golangsdk"
)

// CreateOpsBuilder is an interface by which can build the request body of private ips
type CreateOpsBuilder interface {
	ToPrivateIPMap() (map[string]interface{}, error)
}

// PrivateIPOpts is a struct for creation
type PrivateIPOpts struct {
	SubnetID  string `json:"subnet_id" required:"true"`
	IPAddress string `json:"ip_address,omitempty"`
}

// CreateOps is a struct that contains all the parameters.
type CreateOps struct {
	// The list of private ips.
	PrivateIPs []PrivateIPOpts `json:"privateips" required:"true"`
}

// ToPrivateIPMap is used for type convert
func (opts CreateOps) ToPrivateIPMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create a private ip
func Create(client *golangsdk.ServiceClient, ops CreateOpsBuilder) (r CreateResult) {
	b, err := ops.ToPrivateIPMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

// Get is a method by which can get the detailed information of private ip
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// Delete is a method by which can be able to delete a private ip
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}
