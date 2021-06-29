package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/channels"

	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"balance_strategy": 1,
	"create_time": "2021-07-05T09:29:38Z",
	"id": "328d1d563eba4ff084533188b84b9f8d",
	"member_type": "ecs",
	"name": "terraform_test",
	"port": 8080,
	"status": 1,
	"type": 2
}`
	expectedGetResponse = `
{
	"balance_strategy": 1,
	"create_time": "2021-07-05T09:29:38Z",
	"id": "328d1d563eba4ff084533188b84b9f8d",
	"member_type": "ecs",
	"members": [
		{
			"ecs_id": "dfe1f776-dd07-4ec0-912f-47525679d76a",
			"ecs_name": "terraform_test",
			"host": "192.168.9.195",
			"weight": 1
		}
	],
	"name": "terraform_test",
	"port": 8080,
	"status": 1,
	"type": 2,
	"vpc_health_config": {
		"http_code": "201,202,203",
		"method": "GET",
		"path": "/",
		"port": 8080,
		"protocol": "https",
		"threshold_abnormal": 5,
		"threshold_normal": 2,
		"time_interval": 10,
		"timeout": 5
	}
}`
	expectedListResponse = `
{
	"vpc_channels": [
		{
			"balance_strategy": 1,
			"create_time": "2021-07-05T09:29:38Z",
			"id": "328d1d563eba4ff084533188b84b9f8d",
			"member_type": "ecs",
			"members": [
				{
					"create_time": "2021-07-05T09:29:38Z",
					"ecs_id": "dfe1f776-dd07-4ec0-912f-47525679d76a",
					"ecs_name": "terraform_test",
					"host": "192.168.9.195",
					"id": "a4ffcf94477a4e3697deb5700313d861",
					"status": 1,
					"vpc_channel_id": "328d1d563eba4ff084533188b84b9f8d",
					"weight": 1
				}
			],
			"name": "terraform_test",
			"port": 8080,
			"status": 1,
			"type": 2,
			"vpc_health_config": {
				"create_time": "2021-07-05T09:29:38Z",
				"http_code": "201,202,203",
				"id": "caa3f8122eef4ce98794e84cf9fc5543",
				"method": "GET",
				"path": "/",
				"port": 8080,
				"protocol": "https",
				"threshold_abnormal": 5,
				"threshold_normal": 2,
				"time_interval": 10,
				"timeout": 5,
				"vpc_channel_id": "328d1d563eba4ff084533188b84b9f8d"
			}
		}
	]
}
`
)

var (
	createOpts = &channels.ChannelOpts{
		Name: "terraform_test",
		Type: 2,
		Members: []channels.MemberInfo{
			{
				EcsId:   "dfe1f776-dd07-4ec0-912f-47525679d76a",
				EcsName: "terraform_test",
				Weight:  1,
			},
		},
		VpcHealthConfig: channels.VpcHealthConfig{
			Protocol:          "https",
			Path:              "/",
			Method:            "GET",
			Port:              8080,
			ThresholdAbnormal: 5,
			ThresholdNormal:   2,
			TimeInterval:      10,
			Timeout:           5,
			HttpCodes:         "201,202,203",
			EnableClientSsl:   false,
		},
		Port:            8080,
		BalanceStrategy: 1,
		MemberType:      "ecs",
	}

	expectedCreateResponseData = &channels.VpcChannel{
		Id:              "328d1d563eba4ff084533188b84b9f8d",
		Name:            "terraform_test",
		CreateTime:      "2021-07-05T09:29:38Z",
		MemberType:      "ecs",
		Status:          1,
		Type:            2,
		BalanceStrategy: 1,
		Port:            8080,
	}

	expectedGetResponseData = &channels.VpcChannel{
		Id:              "328d1d563eba4ff084533188b84b9f8d",
		BalanceStrategy: 1,
		CreateTime:      "2021-07-05T09:29:38Z",
		MemberType:      "ecs",
		Members: []channels.MemberInfo{
			{
				EcsId:   "dfe1f776-dd07-4ec0-912f-47525679d76a",
				EcsName: "terraform_test",
				Weight:  1,
				Host:    "192.168.9.195",
			},
		},
		Name:   "terraform_test",
		Port:   8080,
		Status: 1,
		Type:   2,
		VpcHealthConfig: channels.VpcHealthConfig{
			Protocol:          "https",
			Path:              "/",
			Method:            "GET",
			Port:              8080,
			ThresholdNormal:   2,
			ThresholdAbnormal: 5,
			HttpCodes:         "201,202,203",
			TimeInterval:      10,
			Timeout:           5,
		},
	}

	expectedListResponseData = []channels.VpcChannel{
		{
			Id:              "328d1d563eba4ff084533188b84b9f8d",
			BalanceStrategy: 1,
			CreateTime:      "2021-07-05T09:29:38Z",
			MemberType:      "ecs",
			Members: []channels.MemberInfo{
				{
					EcsId:   "dfe1f776-dd07-4ec0-912f-47525679d76a",
					EcsName: "terraform_test",
					Weight:  1,
					Host:    "192.168.9.195",
				},
			},
			Name:   "terraform_test",
			Port:   8080,
			Status: 1,
			Type:   2,
			VpcHealthConfig: channels.VpcHealthConfig{
				Protocol:          "https",
				Path:              "/",
				Method:            "GET",
				Port:              8080,
				ThresholdNormal:   2,
				ThresholdAbnormal: 5,
				HttpCodes:         "201,202,203",
				TimeInterval:      10,
				Timeout:           5,
			},
		},
	}
)

func handleV2VpcChannelCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/b510b8e8ef1442c0a94cdfc551af0ec3/vpc-channels",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleV2VpcChannelGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/b510b8e8ef1442c0a94cdfc551af0ec3/vpc-channels/328d1d563eba4ff084533188b84b9f8d",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2VpcChannelList(t *testing.T) {
	th.Mux.HandleFunc("/instances/b510b8e8ef1442c0a94cdfc551af0ec3/vpc-channels",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListResponse)
		})
}

func handleV2VpcChannelUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/b510b8e8ef1442c0a94cdfc551af0ec3/vpc-channels/328d1d563eba4ff084533188b84b9f8d",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2VpcChannelDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/b510b8e8ef1442c0a94cdfc551af0ec3/vpc-channels/328d1d563eba4ff084533188b84b9f8d",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}
