package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/instances"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateRequest = `
{
	"available_zone_ids": [
		"ap-southeast-2a",
		"ap-southeast-2b"
	],
	"bandwidth_size": 6,
	"description": "Created by script",
	"eip_id": "9361e906-340d-4603-9e67-dadecc6c95ad",
	"enterprise_project_id": "4a52b236-2644-4c10-b31c-7d53a58f75a3",
	"instance_name": "terraform-test",
	"maintain_begin": "22:00:00",
	"maintain_end": "02:00:00",
	"security_group_id": "0f3fa672-e3f1-4e8d-b4d1-d160d408eca8",
	"spec_id": "BASIC",
	"subnet_id": "8149c60b-939f-4eec-a667-6d40aa1e14df",
	"vpc_id": "faf4d1ae-350a-45b5-9d4b-5754b6720cce"
}`
	expectedCreateResponse = `
{
	"instance_id": "e6a5871bfb5b47d19c5874790f639ef8"
}`
	expectedGetResponse = `
{
	"available_zone_ids": "[cn-north-4a]",
	"bandwidth_size": 5,
	"charging_mode": 0,
	"create_time": 1623742314452,
	"enterprise_project_id": "4a52b236-2644-4c10-b31c-7d53a58f75",
	"id": "3fb0bdaa8b27480c971f78734497cd17",
	"ingress_ip": "192.168.155.102",
	"instance_name": "terraform-test",
	"instance_status": 6,
	"instance_version": "fe97b022e6cf401c9e6679f35cf7b130",
	"maintain_begin": "22:00:00",
	"maintain_end": "02:00:00",
	"nat_eip_address": "94.74.112.175",
	"security_group_id": "4b23c9b9-f941-477b-92f0-3adc6eb76124",
	"spec": "BASIC",
	"status": "Running",
	"subnet_id": "7fd65ff0-3ccb-4756-a98c-e2bfe91d3c69",
	"supported_features": [
		"gateway_responses",
		"ratelimit",
		"request_body_size",
		"backend_timeout",
		"app_token",
		"app_basic",
		"app_secret",
		"multi_auth",
		"route",
		"sign_basic",
		"app_route",
		"backend_client_certificate",
		"ssl_ciphers",
		"cors",
		"app_quota",
		"app_acl",
		"real_ip_from_xff",
		"set_resp_headers"
	],
	"user_id": "05602623488025011f3bc015b70b16c3",
	"vpc_id": "1c105033-5b68-4cfe-b58b-b1000e517908"
}`
	expectedListResponse = `
{
	"instances": [
		{
			"charging_mode": 0,
			"create_time": 1623816488875,
			"eip_address": null,
			"enterprise_project_id": "4a52b236-2644-4c10-b31c-7d53a58f75a3",
			"id": "de379eed30aa4d31a84f426ea3c7ef4e",
			"instance_name": "tf-acc-test-0616",
			"instance_status": 6,
			"project_id": "0581b95a0b8010e32f81c015009f6587",
			"spec": "BASIC",
			"status": "Running",
			"type": "apig"
		}
	]
}`
)

var (
	createOpts = &instances.CreateOpts{
		AvailableZoneIds:    []string{"ap-southeast-2a", "ap-southeast-2b"},
		BandwidthSize:       6,
		Description:         "Created by script",
		EipId:               "9361e906-340d-4603-9e67-dadecc6c95ad",
		EnterpriseProjectId: "4a52b236-2644-4c10-b31c-7d53a58f75a3",
		Name:                "terraform-test",
		MaintainBegin:       "22:00:00",
		MaintainEnd:         "02:00:00",
		SecurityGroupId:     "0f3fa672-e3f1-4e8d-b4d1-d160d408eca8",
		Edition:             "BASIC",
		SubnetId:            "8149c60b-939f-4eec-a667-6d40aa1e14df",
		VpcId:               "faf4d1ae-350a-45b5-9d4b-5754b6720cce",
	}

	expectedCreateResponseData = &instances.CreateResp{
		Id: "e6a5871bfb5b47d19c5874790f639ef8",
	}

	expectedListResponseData = []instances.BaseInstance{
		{
			ChargeMode:          0,
			CreateTimestamp:     1623816488875,
			EnterpriseProjectId: "4a52b236-2644-4c10-b31c-7d53a58f75a3",
			Id:                  "de379eed30aa4d31a84f426ea3c7ef4e",
			Name:                "tf-acc-test-0616",
			StatusId:            6,
			ProjectId:           "0581b95a0b8010e32f81c015009f6587",
			Edition:             "BASIC",
			Status:              "Running",
			Type:                "apig",
		},
	}

	updateOpts = instances.UpdateOpts{
		Description:   "Updated by script",
		Name:          "terraform-update",
		MaintainBegin: "18:00:00",
		MaintainEnd:   "22:00:00",
	}

	updateEgressOpts = instances.EgressAccessOpts{
		BandwidthSize: "10",
	}

	updateIngressOpts = instances.IngressAccessOpts{
		EipId: "706673d2-e36b-4577-87bc-e6d6e71812f7",
	}

	expectedGetResponseData = &instances.Instance{
		AvailableZoneIds:      "[cn-north-4a]",
		BandwidthSize:         5,
		ChargeMode:            0,
		CreateTimestamp:       1623742314452,
		EnterpriseProjectId:   "4a52b236-2644-4c10-b31c-7d53a58f75",
		Id:                    "3fb0bdaa8b27480c971f78734497cd17",
		Ipv4VpcIngressAddress: "192.168.155.102",
		Name:                  "terraform-test",
		StatusId:              6,
		Version:               "fe97b022e6cf401c9e6679f35cf7b130",
		MaintainBegin:         "22:00:00",
		MaintainEnd:           "02:00:00",
		Ipv4EgressAddress:     "94.74.112.175",
		SecurityGroupId:       "4b23c9b9-f941-477b-92f0-3adc6eb76124",
		Edition:               "BASIC",
		Status:                "Running",
		SubnetId:              "7fd65ff0-3ccb-4756-a98c-e2bfe91d3c69",
		SupportedFeatures: []string{
			"gateway_responses",
			"ratelimit",
			"request_body_size",
			"backend_timeout",
			"app_token",
			"app_basic",
			"app_secret",
			"multi_auth",
			"route",
			"sign_basic",
			"app_route",
			"backend_client_certificate",
			"ssl_ciphers",
			"cors",
			"app_quota",
			"app_acl",
			"real_ip_from_xff",
			"set_resp_headers",
		},
		VpcId:  "1c105033-5b68-4cfe-b58b-b1000e517908",
		UserId: "05602623488025011f3bc015b70b16c3",
	}
)

func handleV2InstanceCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV2InstanceGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/e6a5871bfb5b47d19c5874790f639ef8", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedGetResponse)
	})
}

func handleV2InstanceList(t *testing.T) {
	th.Mux.HandleFunc("/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}

func handleV2InstanceUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/e6a5871bfb5b47d19c5874790f639ef8", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedGetResponse)
	})
}

func handleV2InstanceDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/e6a5871bfb5b47d19c5874790f639ef8", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}

func handleV2InstanceEgressDisable(t *testing.T) {
	th.Mux.HandleFunc("/instances/e6a5871bfb5b47d19c5874790f639ef8/nat-eip", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		_, _ = fmt.Fprint(w, expectedGetResponse)
	})
}

func handleV2InstanceIngressDisable(t *testing.T) {
	th.Mux.HandleFunc("/instances/e6a5871bfb5b47d19c5874790f639ef8/eip", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}
