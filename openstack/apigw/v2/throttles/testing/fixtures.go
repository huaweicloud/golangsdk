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

	expectedSpecCreateResponse = `
{
	"app_id": "ba985f3d4fd347c3aee686fd749659fc",
	"app_name": "terraform_app",
	"apply_time": "2021-07-21T06:48:03.618383329Z",
	"call_limits": 30,
	"id": "aec8d27e3e034f4293bc766942ed60fd",
	"object_id": "ba985f3d4fd347c3aee686fd749659fc",
	"object_name": "terraform_app",
	"object_type": "APP",
	"throttle_id": "c481043838f64bcd82c7b0c38907d59d"
}`

	expectedSpecListResponse = `
{
	"throttle_specials": [
		{
			"app_id": "ba985f3d4fd347c3aee686fd749659fc",
			"app_name": "terraform_app",
			"apply_time": "2021-07-21T06:48:03.618383329Z",
			"call_limits": 30,
			"id": "aec8d27e3e034f4293bc766942ed60fd",
			"object_id": "ba985f3d4fd347c3aee686fd749659fc",
			"object_name": "terraform_app",
			"object_type": "APP",
			"throttle_id": "c481043838f64bcd82c7b0c38907d59d"
		}
	],
	"size": 1,
	"total": 1
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

	specThrottleCreateOpts = &throttles.SpecThrottleCreateOpts{
		CallLimits: 30,
		ObjectId:   "ba985f3d4fd347c3aee686fd749659fc",
		ObjectType: "APP",
	}

	expectedSpecCreateResponseData = &throttles.SpecThrottle{
		AppId:      "ba985f3d4fd347c3aee686fd749659fc",
		AppName:    "terraform_app",
		ApplyTime:  "2021-07-21T06:48:03.618383329Z",
		CallLimits: 30,
		ID:         "aec8d27e3e034f4293bc766942ed60fd",
		ObjectId:   "ba985f3d4fd347c3aee686fd749659fc",
		ObjectName: "terraform_app",
		ObjectType: "APP",
		ThrottleId: "c481043838f64bcd82c7b0c38907d59d",
	}

	specThrottleListOpts = &throttles.SpecThrottlesListOpts{
		AppName: "terraform_app",
	}

	expectedSpecListResponseData = []throttles.SpecThrottle{
		{
			AppId:      "ba985f3d4fd347c3aee686fd749659fc",
			AppName:    "terraform_app",
			ApplyTime:  "2021-07-21T06:48:03.618383329Z",
			CallLimits: 30,
			ID:         "aec8d27e3e034f4293bc766942ed60fd",
			ObjectId:   "ba985f3d4fd347c3aee686fd749659fc",
			ObjectName: "terraform_app",
			ObjectType: "APP",
			ThrottleId: "c481043838f64bcd82c7b0c38907d59d",
		},
	}

	specThrottleUpdateOpts = &throttles.SpecThrottleUpdateOpts{
		CallLimits: 30,
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

func handleV2SpecThrottlingPolicyCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/throttles/c481043838f64bcd82c7b0c38907d59d/"+
		"throttle-specials",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedSpecCreateResponse)
		})
}

func handleV2SpecThrottlingPolicyList(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/throttles/c481043838f64bcd82c7b0c38907d59d/"+
		"throttle-specials",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedSpecListResponse)
		})
}

func handleV2SpecThrottlingPolicyUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/throttles/c481043838f64bcd82c7b0c38907d59d/"+
		"throttle-specials/aec8d27e3e034f4293bc766942ed60fd",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedSpecCreateResponse)
		})
}

func handleV2SpecThrottlingPolicyDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/throttles/c481043838f64bcd82c7b0c38907d59d/"+
		"throttle-specials/aec8d27e3e034f4293bc766942ed60fd",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}
