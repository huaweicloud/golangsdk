package envs

import (
	"github.com/huaweicloud/golangsdk"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToEnvironmentCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating an environment. This object is passed to
// the environment Create function.
type CreateOpts struct {
	// Name of the environment
	Name string `json:"name" required:"true"`
	// Description of the environment
	Remark string `json:"remark,omitempty"`
}

// ToEnvironmentCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToEnvironmentCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create will create a new environment based on the values in CreateOpts.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToEnvironmentCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToEnvironmentUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing environment.
type UpdateOpts struct {
	// Name of the environment
	Name string `json:"name" required:"true"`
	// Description of the environment
	Remark string `json:"remark,omitempty"`
}

// ToEnvironmentUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToEnvironmentUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update will update the environment with provided information. To extract the updated
// environment from the response, call the Extract method on the UpdateResult.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToEnvironmentUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will delete the existing env with the provided ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}

// List will delete the existing env with the provided ID.
func List(client *golangsdk.ServiceClient, id string) (r ListResult) {
	_, r.Err = client.List(listURL(client, id), nil)
	return
}
