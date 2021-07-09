package queues

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/common/tags"
)

type Queue struct {
	IsSuccess bool `json:"is_success"`

	Message string `json:"message"`

	// Name of a newly created resource queue. The name can contain only digits, letters, and underscores (_),
	// but cannot contain only digits or start with an underscore (_). Length range: 1 to 128 characters.
	QueueName string `json:"queueName" required:"true"`

	// Description of a queue.
	Description string `json:"description"`

	Owner string `json:"owner"`

	CreateTime int64 `json:"create_time"`

	// Indicates the queue type. The options are as follows:
	// sql
	// general
	// all
	// NOTE:
	// If the type is not specified, the default value sql is used.
	QueueType string `json:"queueType"`

	// Minimum number of CUs that are bound to a queue. Currently, the value can only be 16, 64, or 256.
	CuCount int `json:"cuCount" required:"true"`

	// Billing mode of a queue. This value can only be set to 1, indicating that the billing is based on the CUH used.
	ChargingMode int `json:"chargingMode"`

	//
	ResourceId string `json:"resource_id"`

	// Queue resource mode. The options are as follows:
	// 0: indicates the shared resource mode.
	// 1: indicates the exclusive resource mode.
	ResourceMode int `json:"resource_mode"`

	// Enterprise project ID. The value 0 indicates the default enterprise project.
	// NOTE:
	// Users who have enabled Enterprise Management can set this parameter to bind a specified project.
	EnterpriseProjectId string `json:"enterprise_project_id"`

	ResourceType string `json:"resource_type"`

	// CPU architecture of queue computing resources.
	// x86_64 (default)
	// aarch64
	Platform string `json:"platform"`

	// Specifies the tag information of the queue to be created, including the JSON character string indicating
	//whether the queue is Dual-AZ. Currently, only the value 2 is supported, indicating that two queues are created.
	Labels map[string]string `json:"labels"`

	// Indicates the queue feature. The options are as follows:
	// basic: basic type
	// ai: AI-enhanced (Only the SQL x86_64 dedicated queue supports this option.)
	// The default value is basic.
	// NOTE:
	// For an enhanced AI queue, an AI image is loaded in the background.
	// The image integrates AI algorithm packages based on the basic image.
	Feature string `json:"feature"`

	Tags []tags.ResourceTag `json:"tags"`
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	golangsdk.Result
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	golangsdk.Result
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	golangsdk.Result
}
