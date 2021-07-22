package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/dli/v1/queues"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedListResponse = `
	{
		"is_success": true,
		"message": "",
		"queues": [
			{
				"charging_mode": 1,
				"cidr_in_mgntsubnet": "172.24.128.0/18",
				"cidr_in_subnet": "172.24.0.0/18",
				"cidr_in_vpc": "172.16.0.0/12",
				"create_time": 1626771444081,
				"cu_count": 16,
				"description": "",
				"feature": "basic",
				"is_restarting": false,
				"labels": "",
				"max_cu_count": 100,
				"owner": "niuzhenguo",
				"platform": "x86_64",
				"queue_name": "tf_acc_test_dli_queue_h8yr3",
				"queue_resource_type": "vm",
				"queue_type": "sql",
				"resource_id": "b483aa39-ab65-442b-8c20-2f578a9d868b",
				"resource_mode": 1
			}
		]
	}`
)

var (
	expectedListResponseData = queues.Queue{
		QueueName:           "tf_acc_test_dli_queue_h8yr3",
		Description:         "",
		Owner:               "niuzhenguo",
		CreateTime:          1626771444081,
		QueueType:           "sql",
		CuCount:             16,
		ChargingMode:        1,
		ResourceId:          "b483aa39-ab65-442b-8c20-2f578a9d868b",
		EnterpriseProjectId: "",
		CidrInVpc:           "172.16.0.0/12",
		CidrInMgntsubnet:    "172.24.128.0/18",
		CidrInSubnet:        "172.24.0.0/18",
		ResourceMode:        1,
		Platform:            "x86_64",
		IsRestarting:        false,
		Feature:             "basic",
		QueueResourceType:   "vm",
	}
)

func handleList(t *testing.T) {
	th.Mux.HandleFunc("/queues", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}
