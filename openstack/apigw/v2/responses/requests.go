package responses

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// CreateOpts allows to create a new custom response or update the existing custom response using given parameters.
type ResponseOpts struct {
	// APIG group name, which can contain 1 to 64 characters, only letters, digits, hyphens (-) and
	// underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Response type definition, which includes a key and value. Options of the key:
	//     AUTH_FAILURE: Authentication failed.
	//     AUTH_HEADER_MISSING: The identity source is missing.
	//     AUTHORIZER_FAILURE: Custom authentication failed.
	//     AUTHORIZER_CONF_FAILURE: There has been a custom authorizer error.
	//     AUTHORIZER_IDENTITIES_FAILURE: The identity source of the custom authorizer is invalid.
	//     BACKEND_UNAVAILABLE: The backend service is unavailable.
	//     BACKEND_TIMEOUT: Communication with the backend service timed out.
	//     THROTTLED: The request was rejected due to request throttling.
	//     UNAUTHORIZED: The app you are using has not been authorized to call the API.
	//     ACCESS_DENIED: Access denied.
	//     NOT_FOUND: No API is found.
	//     REQUEST_PARAMETERS_FAILURE: The request parameters are incorrect.
	//     DEFAULT_4XX: Another 4XX error occurred.
	//     DEFAULT_5XX: Another 5XX error occurred.
	// Each error type is in JSON format.
	Responses map[string]ResponseInfo `json:"responses,omitempty"`
}

type ResponseInfo struct {
	// Response body template.
	Body string `json:"body" required:"true"`
	// HTTP status code of the response. If omitted, the status will be cancelled.
	Status int `json:"status,omitempty"`
	// Indicates whether the response is the default response.
	// Only the response of the API call is supported.
	IsDefault bool `json:"default,omitempty"`
}

type ResponseOptsBuilder interface {
	ToResponseOptsMap() (map[string]interface{}, error)
}

func (opts ResponseOpts) ToResponseOptsMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method by which to create function that create a new custom response.
func Create(client *golangsdk.ServiceClient, instanceId, groupId string, opts ResponseOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToResponseOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, instanceId, groupId), reqBody, &r.Body, nil)
	return
}

// Update is a method by which to create function that udpate the existing custom response.
func Update(client *golangsdk.ServiceClient, instanceId, groupId, respId string,
	opts ResponseOptsBuilder) (r UpdateResult) {
	reqBody, err := opts.ToResponseOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, instanceId, groupId, respId), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get is a method to obtain the specified custom response according to the instanceId, appId and respId.
func Get(client *golangsdk.ServiceClient, instanceId, groupId, respId string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, instanceId, groupId, respId), &r.Body, nil)
	return
}

// ListOpts allows to filter list data using given parameters.
type ListOpts struct {
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

// List is a method to obtain an array of one or more custom reponses according to the query parameters.
func List(client *golangsdk.ServiceClient, instanceId, groupId string, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client, instanceId, groupId)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ResponsePage{pagination.SinglePageBase(r)}
	})
}

// Delete is a method to delete the existing custom response.
func Delete(client *golangsdk.ServiceClient, instanceId, groupId, respId string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, instanceId, groupId, respId), nil)
	return
}

// SpecResp is used to build the APIG response url. All parameters are required.
type SpecResp struct {
	InstanceId string
	GroupId    string
	RespId     string
	RespType   string
}

// GetSpecResp is a method to get the specifies custom response configuration from an group.
func GetSpecResp(client *golangsdk.ServiceClient, spec SpecResp) (r GetSpecRespResult) {
	_, r.Err = client.Get(responseURL(client, spec.InstanceId, spec.GroupId, spec.RespId, spec.RespType), &r.Body, nil)
	return
}

type UpdateSpecRespBuilder interface {
	ToSpecRespUpdateMap() (map[string]interface{}, error)
}

func (opts ResponseInfo) ToSpecRespUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// UpdateSpecResp is a method to update an existing custom response configuration from an group.
func UpdateSpecResp(client *golangsdk.ServiceClient, spec SpecResp, opts UpdateSpecRespBuilder) (r UpdateSpecRespResult) {
	reqBody, err := opts.ToSpecRespUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(responseURL(client, spec.InstanceId, spec.GroupId, spec.RespId, spec.RespType),
		reqBody, &r.Body, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	return
}

// DeleteSpecResp is a method to delete an existing custom response configuration from an group.
func DeleteSpecResp(client *golangsdk.ServiceClient, instanceId, groupId, respId, respType string) (r DeleteResult) {
	_, r.Err = client.Delete(responseURL(client, instanceId, groupId, respId, respType), nil)
	return
}
