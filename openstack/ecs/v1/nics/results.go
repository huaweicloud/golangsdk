package nics

import "github.com/huaweicloud/golangsdk"

// Nic defines results
type Nic struct {
	PortID string `json:"port_id"`
}

// AssociateResult is the response from a Associate operation.
type AssociateResult struct {
	golangsdk.Result
}

// Extract from AssociateResult
func (r AssociateResult) Extract() (*Nic, error) {
	s := &Nic{}
	return s, r.ExtractInto(s)
}

// DisassociateResult is the response from a Disassociate operation.
type DisassociateResult struct {
	golangsdk.Result
}

// Extract from DisassociateResult
func (r DisassociateResult) Extract() (*Nic, error) {
	s := &Nic{}
	return s, r.ExtractInto(s)
}
