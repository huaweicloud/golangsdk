package attachinterfaces

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// List makes a request against the nova API to list the server's interfaces.
func List(client *golangsdk.ServiceClient, serverID string) pagination.Pager {
	return pagination.NewPager(client, listInterfaceURL(client, serverID), func(r pagination.PageResult) pagination.Page {
		return InterfacePage{pagination.SinglePageBase(r)}
	})
}

// Get requests details on a single interface attachment by the server and port IDs.
func Get(client *golangsdk.ServiceClient, serverID, portID string) (r GetResult) {
	_, r.Err = client.Get(getInterfaceURL(client, serverID, portID), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAttachInterfacesCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies parameters of a new interface attachment.
type CreateOpts struct {

	// PortID is the ID of the port for which you want to create an interface.
	// The NetworkID and PortID parameters are mutually exclusive.
	// If you do not specify the PortID parameter, the OpenStack Networking API
	// v2.0 allocates a port and creates an interface for it on the network.
	PortID string `json:"port_id,omitempty"`

	// NetworkID is the ID of the network for which you want to create an interface.
	// The NetworkID and PortID parameters are mutually exclusive.
	// If you do not specify the NetworkID parameter, the OpenStack Networking
	// API v2.0 uses the network information cache that is associated with the instance.
	NetworkID string `json:"net_id,omitempty"`

	// Slice of FixedIPs. If you request a specific FixedIP address without a
	// NetworkID, the request returns a Bad Request (400) response code.
	FixedIPs []FixedIP `json:"fixed_ips,omitempty"`
}

// ToAttachInterfacesCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToAttachInterfacesCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "interfaceAttachment")
}

// Create requests the creation of a new interface attachment on the server.
func Create(client *golangsdk.ServiceClient, serverID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToAttachInterfacesCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createInterfaceURL(client, serverID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete makes a request against the nova API to detach a single interface from the server.
// It needs server and port IDs to make a such request.
func Delete(client *golangsdk.ServiceClient, serverID, portID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteInterfaceURL(client, serverID, portID), nil)
	return
}