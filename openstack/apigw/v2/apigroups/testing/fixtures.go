package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/apigroups"
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
	"groups": [
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
	createDesc = "Created by script"
	createOpts = &apigroups.GroupOpts{
		Name:        "terraform_test",
		Description: &createDesc,
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

	expectedListResponseData = []apigroups.Group{
		{
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
		},
	}

	updateDesc = "Updated by script"
	updateOpts = apigroups.GroupOpts{
		Name:        "terraform_test_update",
		Description: &updateDesc,
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

func handleV2GroupList(t *testing.T) {
	th.Mux.HandleFunc("/instances/9750f26518a54da8bea1a7c41790c26d/api-groups", func(w http.ResponseWriter, r *http.Request) {
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
