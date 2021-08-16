/*
 Copyright (c) Huawei Technologies Co. Ltd. 2021. All rights reserved.
*/

package dedicated_engine

import "github.com/huaweicloud/golangsdk"

// requestOpts defines the request header and language.
var requestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

// CreateEngineOpts the parameters in the creating request body
type CreateEngineOpts struct {
	Region        string   `json:"region" required:"true"`
	ChargeMode    int      `json:"chargemode" required:"true"`
	AvailableZone string   `json:"available_zone" required:"true"`
	Arch          string   `json:"arch" required:"true"`
	NamePrefix    string   `json:"instancename" required:"true"`
	Specification string   `json:"specification" required:"true"`
	CpuFlavor     string   `json:"cpu_flavor" required:"true"`
	VpcId         string   `json:"vpc_id" required:"true"`
	SubnetId      string   `json:"subnet_id" required:"true"`
	SecurityGroup []string `json:"security_group" required:"true"`
	Count         int      `json:"count" required:"true"`
	Ipv6Enable    string   `json:"ipv6_enable,omitempty"`
	VolumeType    string   `json:"volume_type,omitempty"`
	ClusterId     string   `json:"cluster_id,omitempty"`
	PoolId        string   `json:"pool_id,omitempty"`
}

// ListEngineOpts the parameters in the querying request.
type ListEngineOpts struct {
	Page       int    `q:"page"`
	PageSize   int    `q:"pagesize"`
	EngineName string `q:"instancename"`
}

// UpdateEngineOpts the parameters in the updating request.
type UpdateEngineOpts struct {
	EngineName string `json:"instancename"`
}

// CreateEngine will create a dedicated waf instance based on the values in CreateOpts.
func CreateEngine(c *golangsdk.ServiceClient, opts CreateEngineOpts) (*CreationgRst, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var rst golangsdk.Result
	_, err = c.Post(rootURL(c), b, &rst.Body, &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})

	if err == nil {
		var r CreationgRst
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}

// GetEngine get the waf instance detail.
func GetEngine(c *golangsdk.ServiceClient, id string) (*DedicatedEngine, error) {
	var rst golangsdk.Result
	_, err := c.Get(resourceURL(c, id), &rst.Body, &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})

	if err == nil {
		var r DedicatedEngine
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}

// ListEngine query a list of waf instance base on ListEngineOpts
func ListEngine(c *golangsdk.ServiceClient, opts ListEngineOpts) (*DedicatedEngineList, error) {
	url := rootURL(c)
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}
	url += query.String()

	var rst golangsdk.Result
	_, err = c.Get(url, &rst.Body, &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})

	if err == nil {
		var r DedicatedEngineList
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}

// UpdateEngine query a list of waf instance base on UpdateEngineOpts
func UpdateEngine(c *golangsdk.ServiceClient, id string, opts UpdateEngineOpts) (*DedicatedEngine, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var rst golangsdk.Result
	_, err = c.Put(resourceURL(c, id), b, &rst.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: requestOpts.MoreHeaders,
	})

	if err == nil {
		var r DedicatedEngine
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}

// Delete will permanently delete a engine based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (*DedicatedEngine, error) {
	var rst golangsdk.Result
	_, err := c.DeleteWithResponse(resourceURL(c, id), &rst.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: requestOpts.MoreHeaders},
	)

	if err == nil {
		var r DedicatedEngine
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}
