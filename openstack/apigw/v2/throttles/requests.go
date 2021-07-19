package throttles

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type ThrottlingPolicyOpts struct {
	// Maximum number of times an API can be accessed within a specified period.
	// The value of this parameter cannot exceed the default limit 200 TPS.
	// This value must be a positive integer and cannot exceed 2,147,483,647.
	ApiCallLimits int `json:"api_call_limits" required:"true"`
	// Request throttling policy name, which can contain 3 to 64 characters, starting with a letter.
	// Only letters, digits, and underscores (_) are allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Name string `json:"name" required:"true"`
	// Period of time for limiting the number of API calls.
	// This parameter applies with each of the preceding three API call limits.
	// This value must be a positive integer and cannot exceed 2,147,483,647.
	TimeInterval int `json:"time_interval" required:"true"`
	// Time unit for limiting the number of API calls.
	// The valid values are as following:
	//     SECOND
	//     MINUTE
	//     HOUR
	//     DAY
	TimeUnit string `json:"time_unit" required:"true"`
	// Maximum number of times the API can be accessed by an app within the same period.
	// The value of this parameter must be less than that of user_call_limits.
	// This value must be a positive integer and cannot exceed 2,147,483,647.
	AppCallLimits int `json:"app_call_limits,omitempty"`
	// Description of the request throttling policy, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description string `json:"remark,omitempty"`
	// Type of the request throttling policy.
	// 1: exclusive, limiting the maximum number of times a single API bound to the policy can be called within
	// the specified period.
	// 2: shared, limiting the maximum number of times all APIs bound to the policy can be called within the
	// specified period.
	Type int `json:"type,omitempty"`
	// Maximum number of times the API can be accessed by a user within the same period.
	// The value of this parameter must be less than that of api_call_limits.
	// This value must be a positive integer and cannot exceed 2,147,483,647.
	UserCallLimits int `json:"user_call_limits,omitempty"`
	// Maximum number of times the API can be accessed by an IP address within the same period.
	// The value of this parameter must be less than that of api_call_limits.
	// This value must be a positive integer and cannot exceed 2,147,483,647.
	IpCallLimits int `json:"ip_call_limits,omitempty"`
}

type ThrottlingPolicyOptsBuilder interface {
	ToThrottlingPolicyOptsMap() (map[string]interface{}, error)
}

func (opts ThrottlingPolicyOpts) ToThrottlingPolicyOptsMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method by which to create function that create a new throttling policy.
func Create(client *golangsdk.ServiceClient, instanceId string, opts ThrottlingPolicyOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToThrottlingPolicyOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, instanceId), reqBody, &r.Body, nil)
	return
}

// Update is a method by which to udpate an existing throttle policy.
func Update(client *golangsdk.ServiceClient, instanceId, policyId string,
	opts ThrottlingPolicyOptsBuilder) (r UpdateResult) {
	reqBody, err := opts.ToThrottlingPolicyOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, instanceId, policyId), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get is a method to obtain an existing APIG throttling policy by policy ID.
func Get(client *golangsdk.ServiceClient, instanceId, policyId string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, instanceId, policyId), &r.Body, nil)
	return
}

type ListOpts struct {
	// Request throttling policy ID.
	Id string `q:"id"`
	// Request throttling policy name.
	Name string `q:"name"`
	// Offset from which the query starts.
	// If the offset is less than 0, the value is automatically converted to 0. Default to 0.
	Offset int `q:"offset"`
	// Number of items displayed on each page. The valid values are range form 1 to 500, default to 20.
	Limit int `q:"limit"`
	// Parameter name (name) for exact matching.
	PreciseSearch string `q:"precise_search"`
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

// List is a method to obtain an array of one or more throttling policies according to the query parameters.
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
		return ThorttlePage{pagination.SinglePageBase(r)}
	})
}

// Delete is a method to delete an existing throttling policy.
func Delete(client *golangsdk.ServiceClient, instanceId, policyId string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, instanceId, policyId), nil)
	return
}
