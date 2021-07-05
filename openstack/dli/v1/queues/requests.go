package queues

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/common/tags"
)

// CreateOpts contains the options for create a service. This object is passed to Create().
type CreateOpts struct {
	// Name of a newly created resource queue. The name can contain only digits, letters, and underscores (_), but cannot contain only digits or start with an underscore (_). Length range: 1 to 128 characters.
	QueueName string `json:"queue_name" required:"true"`

	// Indicates the queue type. The options are as follows:
	// sql
	// general
	// all
	// NOTE:
	// If the type is not specified, the default value sql is used.
	QueueType string `json:"queue_type"`

	// Description of a queue.
	Description string `json:"description"`

	// Minimum number of CUs that are bound to a queue. Currently, the value can only be 16, 64, or 256.
	CuCount int `json:"cu_count" required:"true"`

	// Billing mode of a queue. This value can only be set to 1, indicating that the billing is based on the CUH used.
	ChargingMode int `json:"charging_mode"`

	// Enterprise project ID. The value 0 indicates the default enterprise project.
	// NOTE:
	// Users who have enabled Enterprise Management can set this parameter to bind a specified project.
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// CPU architecture of queue computing resources.
	// x86_64 (default)
	// aarch64
	Platform string `json:"platform"`

	// Queue resource mode. The options are as follows:
	// 0: indicates the shared resource mode.
	// 1: indicates the exclusive resource mode.
	ResourceMode int `json:"resource_mode"`

	// Specifies the tag information of the queue to be created, including the JSON character string indicating whether the queue is Dual-AZ. Currently, only the value 2 is supported, indicating that two queues are created.
	Labels map[string]string `json:"labels,omitempty"`

	// Indicates the queue feature. The options are as follows:
	// basic: basic type
	// ai: AI-enhanced (Only the SQL x86_64 dedicated queue supports this option.)
	// The default value is basic.
	// NOTE:
	// For an enhanced AI queue, an AI image is loaded in the background. The image integrates AI algorithm packages based on the basic image. For details, see the Data Lake Insight User Guide.
	Feature string `json:"feature"`

	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToDomainCreateMap() (map[string]interface{}, error)
}

// ToDomainCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToDomainCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create will create a new Domain based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	requstbody, err := opts.ToDomainCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(createURL(c), requstbody, &r.Body, reqOpt)
	return
}

func Delete(c *golangsdk.ServiceClient, queueName string) (r DeleteResult) {
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Delete(resourceURL(c, queueName), reqOpt)
	return
}

func Get(c *golangsdk.ServiceClient, queueName string) (r GetResult) {
	result := new(Queue)
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Get(resourceURL(c, queueName), &result, reqOpt)
	r.Body = result
	return r
}
