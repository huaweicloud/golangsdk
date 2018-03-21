package nics

import (
	"github.com/huaweicloud/golangsdk"
)

// AssociateOptsBuilder allows extensions to
// add additional parameters to the associate request.
type AssociateOptsBuilder interface {
	ToVIPAssociateMap() (map[string]interface{}, error)
}

// NicAssociateOpts is a struct for creation
type NicAssociateOpts struct {
	SubnetID       string `json:"subnet_id" required:"true"`
	IPAddress      string `json:"ip_address" required:"true"`
	ReverseBinding bool   `json:"reverse_binding" required:"true"`
	DeviceOwner    string `json:"device_owner,omitempty"`
}

// AssociateOpts specifies the required information
// to associate a Floating IP with an instance
type AssociateOpts struct {
	// Nic is associated by VIP.
	Nic NicAssociateOpts `json:"nic" required:"true"`
}

// ToVIPAssociateMap constructs a request body from AssociateOpts.
func (opts AssociateOpts) ToVIPAssociateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// AssociateVIP pairs an allocated VIP with a nic.
func AssociateVIP(client *golangsdk.ServiceClient, id string, ops AssociateOptsBuilder) (r AssociateResult) {
	b, err := ops.ToVIPAssociateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(associateURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

// DisassociateOptsBuilder allows extensions to
// add additional parameters to the Disassociate request.
type DisassociateOptsBuilder interface {
	ToVIPDisassociateMap() (map[string]interface{}, error)
}

// NicDisassociateOpts is a struct for disassociate
type NicDisassociateOpts struct {
}

// DisassociateOpts specifies the required information
// to associate a VIP with a nic
type DisassociateOpts struct {
	// Nic is associated by VIP.
	Nic NicDisassociateOpts `json:"nic" required:"true"`
}

// ToVIPDisassociateMap constructs a request body from DisassociateOpts.
func (opts DisassociateOpts) ToVIPDisassociateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// DisassociateVIP pairs an allocated VIP with a nic.
func DisassociateVIP(client *golangsdk.ServiceClient, id string, ops DisassociateOptsBuilder) (r DisassociateResult) {
	b, err := ops.ToVIPDisassociateMap()
	if b != nil {
		b["subnet_id"] = ""
		b["ip_address"] = ""
		b["reverse_binding"] = false
	}
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(disassociateURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
