package replicationconsistencygroups

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// CreateOpsBuilder is used for creating replication consistency group parameters.
// any struct providing the parameters should implement this interface
type CreateOpsBuilder interface {
	ToReplicationConsistencyGroupCreateMap() (map[string]interface{}, error)
}

// CreateOps is a struct that contains all the parameters.
type CreateOps struct {
	// The name of the replication consistency group.
	// The name can contain a maximum of 255 bytes.
	Name string `json:"name,omitempty"`

	// The description of the replication consistency group.
	// The description can contain a maximum of 255 bytes.
	Description string `json:"description,omitempty"`

	// The IDs of the EVS replication pairs used to
	// create the replication consistency group.
	ReplicationIDs []string `json:"replication_ids" required:"true"`

	// The primary AZ of the replication consistency group.
	// That is the AZ where the production disk belongs.
	PriorityStation string `json:"priority_station" required:"true"`

	// The type of the created replication consistency group.
	// Currently only type hypermetro is supported.
	ReplicationModel string `json:"replication_model" required:"true"`
}

// ToReplicationConsistencyGroupCreateMap is used for type convert
func (ops CreateOps) ToReplicationConsistencyGroupCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(ops, "replication_consistency_group")
}

// Create a replication consistency group with given parameters.
func Create(client *golangsdk.ServiceClient, ops CreateOpsBuilder) (r CreateResult) {
	b, err := ops.ToReplicationConsistencyGroupCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})

	return
}

// Delete a replication consistency group by id
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

// Get a replication consistency group with detailed information by id
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder is an interface by which can be able to
// build the query string of the list function
type ListOptsBuilder interface {
	ToReplicationConsistencyGroupListQuery() (string, error)
}

// ListOpts is a struct that contains all the parameters.
type ListOpts struct {
	Marker          string `q:"marker"`
	Limit           int    `q:"limit"`
	SortKey         string `q:"sort_key"`
	SortDir         string `q:"sort_dir"`
	Offset          int    `q:"offset"`
	ChangesSince    string `q:"changes-since"`
	Name            string `q:"name"`
	Status          string `q:"status"`
	PriorityStation string `q:"priority_station"`
	VolumeID        string `q:"volume_id"`
}

// ToReplicationConsistencyGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToReplicationConsistencyGroupListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// List all the replication consistency groups
func List(client *golangsdk.ServiceClient, ops ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if ops != nil {
		q, err := ops.ToReplicationConsistencyGroupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ReplicationConsistencyGroupPage{pagination.SinglePageBase(r)}
	})
}

//UpdateOptsBuilder is an interface which can build the map paramter of update function
type UpdateOptsBuilder interface {
	ToReplicationConsistencyGroupUpdateMap() (map[string]interface{}, error)
}

//UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	// The name of the replication consistency group.
	// The name can contain a maximum of 255 bytes.
	Name string `json:"name,omitempty"`

	// The description of the replication consistency group.
	// The description can contain a maximum of 255 bytes.
	Description string `json:"description,omitempty"`

	// The type of the created replication consistency group.
	// Currently only type hypermetro is supported.
	ReplicationModel string `json:"replication_model,omitempty"`

	// The IDs of the EVS replication pairs to be added.
	AddReplicationIDs []string `json:"add_replication_ids,omitempty"`

	// The IDs of the EVS replication pairs to be removeed.
	RemoveReplicationIDs []string `json:"remove_replication_ids,omitempty"`
}

// ToReplicationConsistencyGroupUpdateMap is used for type convert
func (opts UpdateOpts) ToReplicationConsistencyGroupUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "replication_consistency_group")
}

// Update is a method which can be able to update the replication consistency group
// via accessing to the service with Put method and parameters
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	body, err := opts.ToReplicationConsistencyGroupUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(updateURL(client, id), body, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
