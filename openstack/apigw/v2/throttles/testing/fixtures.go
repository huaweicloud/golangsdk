package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/throttles"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"api_call_limits": 50,
	"app_call_limits": 20,
	"bind_num": 0,
	"create_time": "2021-07-19T08:46:39.799181829Z",
	"id": "c481043838f64bcd82c7b0c38907d59d",
	"ip_call_limits": 15,
	"name": "terraform_test",
	"remark": "Created by script",
	"time_interval": 10,
	"time_unit": "MINUTE",
	"type": 1,
	"user_call_limits": 20
}`

	expectedListResponse = `
{
	"throttles": [
		{
			"api_call_limits": 50,
			"app_call_limits": 20,
			"bind_num": 0,
			"create_time": "2021-07-19T08:46:39.799181829Z",
			"id": "c481043838f64bcd82c7b0c38907d59d",
			"ip_call_limits": 15,
			"name": "terraform_test",
			"remark": "Created by script",
			"time_interval": 10,
			"time_unit": "MINUTE",
			"type": 1,
			"user_call_limits": 20
		}
	]
}`
)

var (
	createOpts = &throttles.ThrottlingPolicyOpts{
		Name:           "terraform_test",
		Type:           1,
		TimeInterval:   10,
		TimeUnit:       "MINUTE",
		ApiCallLimits:  50,
		UserCallLimits: 20,
		AppCallLimits:  20,
		IpCallLimits:   15,
		Description:    "Created by script",
	}

	expectedCreateResponseData = &throttles.ThrottlingPolicy{
		Name:           "terraform_test",
		Type:           1,
		TimeInterval:   10,
		TimeUnit:       "MINUTE",
		ApiCallLimits:  50,
		UserCallLimits: 20,
		AppCallLimits:  20,
		IpCallLimits:   15,
		Description:    "Created by script",
		BindNum:        0,
		CreateTime:     "2021-07-19T08:46:39.799181829Z",
		Id:             "c481043838f64bcd82c7b0c38907d59d",
	}

	listOpts = throttles.ListOpts{
		Name: "terraform_test",
	}

	expectedListResponseData = []throttles.ThrottlingPolicy{
		{
			Name:           "terraform_test",
			Type:           1,
			TimeInterval:   10,
			TimeUnit:       "MINUTE",
			ApiCallLimits:  50,
			UserCallLimits: 20,
			AppCallLimits:  20,
			IpCallLimits:   15,
			Description:    "Created by script",
			BindNum:        0,
			CreateTime:     "2021-07-19T08:46:39.799181829Z",
			Id:             "c481043838f64bcd82c7b0c38907d59d",
		},
	}
)

func handleV2ThrottlingPolicyCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/throttles",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleV2ThrottlingPolicyGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/throttles/c481043838f64bcd82c7b0c38907d59d",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleV2ThrottlingPolicyList(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/throttles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}

func handleV2ThrottlingPolicyUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/throttles/c481043838f64bcd82c7b0c38907d59d",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleV2ThrottlingPolicyDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/throttles/c481043838f64bcd82c7b0c38907d59d",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}
