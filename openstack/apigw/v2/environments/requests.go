package environments

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// EnvironmentOpts allows to create a new environment or update an existing environment using given parameters.
type EnvironmentOpts struct {
	// Environment name, which can contain 3 to 64 characters, starting with a letter.
	// Only letters, digits and underscores (_) are allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Name string `json:"name" required:"true"`
	// Description of the environment, which can contain a maximum of 255 characters,
	// and the angle brackets (< and >) are not allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description string `json:"remark,omitempty"`
}

type EnvironmentOptsBuilder interface {
	ToEnvironmentOptsMap() (map[string]interface{}, error)
}

func (opts EnvironmentOpts) ToEnvironmentOptsMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method by which to create function that create a new environment.
func Create(client *golangsdk.ServiceClient, instanceId string, opts EnvironmentOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToEnvironmentOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, instanceId), reqBody, &r.Body, nil)
	return
}

// Update is a method by which to udpate an existing environment.
func Update(client *golangsdk.ServiceClient, instanceId, envId string, opts EnvironmentOptsBuilder) (r UpdateResult) {
	reqBody, err := opts.ToEnvironmentOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, instanceId, envId), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ListOpts allows to filter list data using given parameters.
type ListOpts struct {
	// Environment name.
	Name string `q:"name"`
	// Offset from which the query starts.
	// If the offset is less than 0, the value is automatically converted to 0. Default to 0.
	Offset int `q:"offset"`
	// Number of items displayed on each page. The valid values are range form 1 to 500, default to 20.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// List is a method to obtain an array of one or more groups according to the query parameters.
func List(client *golangsdk.ServiceClient, instanceId string, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client, instanceId)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return EnvironmentPage{pagination.SinglePageBase(r)}
	})
}

// Delete is a method to delete an existing group.
func Delete(client *golangsdk.ServiceClient, instanceId, envId string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, instanceId, envId), nil)
	return
}
