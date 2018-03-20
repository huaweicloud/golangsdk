package vips

import (
	"github.com/huaweicloud/golangsdk"
)

//CreateResult is a struct which represents the result of create private ip
type CreateResult struct {
	golangsdk.Result
}

// PrivateIP is a struct that represents a private ip
type PrivateIP struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	SubnetID    string `json:"subnet_id"`
	TenantID    string `json:"tenant_id"`
	DeviceOwner string `json:"device_owner"`
	IPAddress   string `json:"ip_address"`
}

// PrivateIPs is a struct that represents private ips
type PrivateIPs struct {
	// The list of private ips.
	PrivateIPs []PrivateIP `json:"privateips"`
}

// Extract from CreateResult
func (r CreateResult) Extract() (PrivateIPs, error) {
	var s struct {
		IPs PrivateIPs `json:"privateips"`
	}
	err := r.Result.ExtractInto(&s)
	return s.IPs, err
}

// GetResult is a return struct of get method
type GetResult struct {
	golangsdk.Result
}

// Extract from GetResult
func (r GetResult) Extract() (PrivateIP, error) {
	var s struct {
		IP PrivateIP `json:"privateip"`
	}
	err := r.Result.ExtractInto(&s)
	return s.IP, err
}

// DeleteResult is a struct of delete result
type DeleteResult struct {
	golangsdk.ErrResult
}
