package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/apigroups"
	"github.com/huaweicloud/golangsdk/openstack/dli/v1/queues"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"id": "1c1acdd2f4d14eb886ecd2370cdb9c1a",
	"is_default": 2,
	"name": "terraform_test",
	"on_sell_status": 2,
	"register_time": "2021-06-22T07:02:20.133688796Z",
	"remark": "Created by script",
	"sl_domain": "1c1acdd2f4d14eb886ecd2370cdb9c1a.apigw-cn-north-4-myhuaweicloud.com",
	"sl_domains": [
		"1c1acdd2f4d14eb886ecd2370cdb9c1a.apigw-cn-north-4-myhuaweicloud.com"
	],
	"status": 1,
	"update_time": "2021-06-22T07:02:20.133688906Z"
}`

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
			},
			{
				"queue_name": "test",
				"owner": "testuser",
				"description": "",
				"create_time": 1562221422671,
				"queue_type": "spark",
				"cu_count": 16,
				"charging_mode": 2,
				"resource_id": "26afb850-d3c9-42c1-81c0-583d1163e80f",
				"enterprise_project_id": "0",
				"cidr_in_vpc": "10.0.0.0/8",
				"cidr_in_subnet": "10.0.0.0/24",
				"cidr_in_mgntsubnet": "10.23.128.0/24",
				"resource_mode": 1,
				"platform": "x86_64",
				"is_restarting": false
			}
		]
	}`

	expectedUpdateResponse = `
{
	"id": "1c1acdd2f4d14eb886ecd2370cdb9c1a",
	"is_default": 2,
	"name": "terraform_test_update",
	"on_sell_status": 2,
	"register_time": "2021-06-22T07:02:20.133688796Z",
	"remark": "Updated by script",
	"sl_domain": "1c1acdd2f4d14eb886ecd2370cdb9c1a.apigw-cn-north-4-myhuaweicloud.com",
	"sl_domains": [
		"1c1acdd2f4d14eb886ecd2370cdb9c1a.apigw-cn-north-4-myhuaweicloud.com"
	],
	"status": 1,
	"update_time": "2021-06-22T07:02:20.133688906Z"
}`
)

var (
	createOpts = &apigroups.GroupOpts{
		Name:        "terraform_test",
		Description: "Created by script",
	}

	expectedCreateResponseData = &apigroups.Group{
		Id:              "1c1acdd2f4d14eb886ecd2370cdb9c1a",
		IsDefault:       2,
		Name:            "terraform_test",
		OnSellStatus:    2,
		RegistraionTime: "2021-06-22T07:02:20.133688796Z",
		Description:     "Created by script",
		Subdomain:       "1c1acdd2f4d14eb886ecd2370cdb9c1a.apigw-cn-north-4-myhuaweicloud.com",
		Subdomains: []string{
			"1c1acdd2f4d14eb886ecd2370cdb9c1a.apigw-cn-north-4-myhuaweicloud.com",
		},
		Status:     1,
		UpdateTime: "2021-06-22T07:02:20.133688906Z",
	}

	expectedUpdateResponseData = &apigroups.Group{
		Id:              "1c1acdd2f4d14eb886ecd2370cdb9c1a",
		IsDefault:       2,
		Name:            "terraform_test_update",
		OnSellStatus:    2,
		RegistraionTime: "2021-06-22T07:02:20.133688796Z",
		Description:     "Updated by script",
		Subdomain:       "1c1acdd2f4d14eb886ecd2370cdb9c1a.apigw-cn-north-4-myhuaweicloud.com",
		Subdomains: []string{
			"1c1acdd2f4d14eb886ecd2370cdb9c1a.apigw-cn-north-4-myhuaweicloud.com",
		},
		Status:     1,
		UpdateTime: "2021-06-22T07:02:20.133688906Z",
	}

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

	updateOpts = apigroups.GroupOpts{
		Name:        "terraform_test_update",
		Description: "Updated by script",
	}
)

func handleV2GroupCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/9750f26518a54da8bea1a7c41790c26d/api-groups",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleV2GroupGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/9750f26518a54da8bea1a7c41790c26d/api-groups/1c1acdd2f4d14eb886ecd2370cdb9c1a",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleList(t *testing.T) {
	th.Mux.HandleFunc("/queues", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}

func handleV2GroupUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/9750f26518a54da8bea1a7c41790c26d/api-groups/1c1acdd2f4d14eb886ecd2370cdb9c1a",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedUpdateResponse)
		})
}

func handleV2GroupDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/9750f26518a54da8bea1a7c41790c26d/api-groups/1c1acdd2f4d14eb886ecd2370cdb9c1a",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}
