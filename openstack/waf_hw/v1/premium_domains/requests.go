/*
 Copyright (c) Huawei Technologies Co., Ltd. 2021. All rights reserved.
*/

package premium_domains

import "github.com/huaweicloud/golangsdk"

var RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

// CreatePremiumHostOpts the options for creating premium domains.
type CreatePremiumHostOpts struct {
	CertificateId   string                `json:"certificateid,omitempty"`
	CertificateName string                `json:"certificatename,omitempty"`
	HostName        string                `json:"hostname" required:"true"`
	Proxy           *bool                 `json:"proxy,omitempty"`
	PolicyId        string                `json:"policyid,omitempty"`
	Servers         []PremiumDomainServer `json:"server,omitempty"`
}

// PremiumDomainServer the options of domain server for creating premium domains.
type PremiumDomainServer struct {
	FrontProtocol string `json:"front_protocol" required:"true"`
	BackProtocol  string `json:"back_protocol" required:"true"`
	Address       string `json:"address" required:"true"`
	Port          int    `json:"port" required:"true"`
	Type          string `json:"type,omitempty"`
	VpcId         string `json:"vpc_id,omitempty"`
}

// CreatePremiumHost create a premium domain in HuaweiCloud.
func CreatePremiumHost(c *golangsdk.ServiceClient, opts CreatePremiumHostOpts) (*CreatePremiumHostRst, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var rst golangsdk.Result
	_, err = c.Post(rootURL(c), b, &rst.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	if err == nil {
		var r CreatePremiumHostRst
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}

// GetPremiumHost get a premium domain by id.
func GetPremiumHost(c *golangsdk.ServiceClient, hostId string) (*PremiumHost, error) {
	var rst golangsdk.Result
	_, err := c.Get(resourceURL(c, hostId), &rst.Body, &golangsdk.RequestOpts{
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	if err == nil {
		var r PremiumHost
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}

// ListPremiumHostOpts the options for querying a list of premium domains.
type ListPremiumHostOpts struct {
	Page          string `q:"page"`
	PageSize      string `q:"pagesize"`
	HostName      string `q:"hostname"`
	PolicyName    string `q:"policyname"`
	ProtectStatus int    `q:"protect_status"`
}

// ListPremiumHost query a list of premium domains.
func ListPremiumHost(c *golangsdk.ServiceClient, opts ListPremiumHostOpts) (*PremiumHostList, error) {
	url := rootURL(c)
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}
	url += query.String()

	var rst golangsdk.Result
	_, err = c.Get(url, &rst.Body, &golangsdk.RequestOpts{
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	if err == nil {
		var r PremiumHostList
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}

// UpdatePremiumHostOpts the options for updating premium domains.
type UpdatePremiumHostOpts struct {
	Proxy           *bool  `json:"proxy,omitempty"`
	CertificateId   string `json:"certificateid,omitempty"`
	CertificateName string `json:"certificatename,omitempty"`
	Tls             string `json:"tls,omitempty"`
	Cipher          string `json:"cipher,omitempty"`
}

// UpdatePremiumHost update premium domains according to UpdatePremiumHostOpts.
func UpdatePremiumHost(c *golangsdk.ServiceClient, hostId string, opts UpdatePremiumHostOpts) (*PremiumHost, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var rst golangsdk.Result
	_, err = c.Put(resourceURL(c, hostId), b, &rst.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	if err == nil {
		var r PremiumHost
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}

// updatePremiumHostProtectStatusOpts the struct for updating the protect status of premium domain.
// Only used in the package.
type updatePremiumHostProtectStatusOpts struct {
	ProtectStatus *int `json:"protect_status" required:"true"`
}

// UpdatePremiumHostProtectStatus update the protect status of premium domain.
func UpdatePremiumHostProtectStatus(c *golangsdk.ServiceClient, hostId string, protectStatus int) (*PremiumHostProtectStatus, error) {
	opts := updatePremiumHostProtectStatusOpts{
		ProtectStatus: &protectStatus,
	}

	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var rst golangsdk.Result
	_, err = c.Put(resourceURL(c, hostId), b, &rst.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	if err == nil {
		var r PremiumHostProtectStatus
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}

// deletePremiumHostOpts whether to keep the premium domain policy when deleting the premium domain.
// Only used in the package.
type deletePremiumHostOpts struct {
	KeepPolicy bool `q:"keepPolicy"`
}

// DeletePremiumHost delete a premium domain by id.
func DeletePremiumHost(c *golangsdk.ServiceClient, hostId string, keepPolicy bool) (*SimplePremiumHost, error) {
	opts := deletePremiumHostOpts{
		KeepPolicy: keepPolicy,
	}

	url := resourceURL(c, hostId)
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var rst golangsdk.Result
	_, err = c.DeleteWithBodyResp(url, b, &rst.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	if err == nil {
		var r SimplePremiumHost
		rst.ExtractInto(&r)
		return &r, nil
	}
	return nil, err
}
