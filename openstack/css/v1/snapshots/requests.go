package snapshots

import (
	"github.com/huaweicloud/golangsdk"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSnapshotCreateMap() (map[string]interface{}, error)
}

// PolicyCreateOpts contains options for creating a snapshot policy.
// This object is passed to the snapshots.PolicyCreate function.
type PolicyCreateOpts struct {
	Prefix     string `json:"prefix" required:"true"`
	Period     string `json:"period" required:"true"`
	KeepDay    int    `json:"keepday" required:"true"`
	Enable     string `json:"enable" required:"true"`
	DeleteAuto string `json:"deleteAuto,omitempty"`
}

// ToSnapshotCreateMap assembles a request body based on the contents of a
// PolicyCreateOpts.
func (opts PolicyCreateOpts) ToSnapshotCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// PolicyCreate will create a new snapshot policy based on the values in PolicyCreateOpts.
func PolicyCreate(client *golangsdk.ServiceClient, opts CreateOptsBuilder, clusterId string) (r ErrorResult) {
	b, err := opts.ToSnapshotCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(policyURL(client, clusterId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// PolicyGet retrieves the snapshot policy with the provided cluster ID.
// To extract the snapshot policy object from the response, call the Extract method on the GetResult.
func PolicyGet(client *golangsdk.ServiceClient, clusterId string) (r PolicyResult) {
	_, r.Err = client.Get(policyURL(client, clusterId), &r.Body, nil)
	return
}
