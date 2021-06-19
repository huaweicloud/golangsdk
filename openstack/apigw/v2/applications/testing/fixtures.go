package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/applications"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"app_key": "6a4b1efab9374f599059ff3fe4382089",
	"app_secret": "44e380875cf046feb29ab14694566239",
	"app_type": "apig",
	"creator": "USER",
	"id": "50f768cf1c1f4389965aa58d255b2a65",
	"name": "terraform_test",
	"register_time": "2021-06-19T03:15:23.181079903Z",
	"remark": "Created by script",
	"status": 1,
	"update_time": "2021-06-19T03:15:23.181080043Z"
}`
	expectedGetResponse = `
{
	"app_key": "6a4b1efab9374f599059ff3fe4382089",
	"app_secret": "44e380875cf046feb29ab14694566239",
	"app_type": "apig",
	"creator": "USER",
	"id": "50f768cf1c1f4389965aa58d255b2a65",
	"name": "terraform_test",
	"register_time": "2021-06-19T03:15:23.181079903Z",
	"remark": "Created by script",
	"status": 1,
	"update_time": "2021-06-19T03:15:23.181080043Z"
}`
	expectedListResponse = `
{
	"apps": [
		{
			"app_key": "6a4b1efab9374f599059ff3fe4382089",
			"app_secret": "44e380875cf046feb29ab14694566239",
			"app_type": "apig",
			"bind_num": 0,
			"creator": "USER",
			"id": "50f768cf1c1f4389965aa58d255b2a65",
			"name": "terraform_test",
			"register_time": "2021-06-19T03:15:23.181079903Z",
			"remark": "Created by script",
			"status": 1,
			"update_time": "2021-06-19T03:15:23.181080043Z"
		}
	]
}`
	expectedAddCodeResponse = `
{
	"app_code": "dGVzdEFQSUNvZGVPZlRoZUFQSUdBcHBsaWNhdGlvbjEyMzQ1Njc4OTAhQCMkJSstLz10b0Jhc2U2NA==",
	"app_id": "5b9ac7b0b2434860a264e05d8c31470e",
	"create_time": "2021-06-19T06:21:36.446251055Z",
	"id": "5103639943e648599ef082285001c34e"
}`

	expectedListCodeResponse = `
{
	"app_codes": [
		{
		  "app_code": "dGVzdEFQSUNvZGVPZlRoZUFQSUdBcHBsaWNhdGlvbjEyMzQ1Njc4OTAhQCMkJSstLz10b0Jhc2U2NA==",
		  "app_id": "5b9ac7b0b2434860a264e05d8c31470e",
		  "create_time": "2021-06-19T06:21:36.446251055Z",
		  "id": "5103639943e648599ef082285001c34e"
		}
	]
}`
)

var (
	appOpts = &applications.AppOpts{
		Name:        "tf_acc_test_0618",
		Description: "Created by script",
	}

	expectedCreateResponseData = &applications.Application{
		AppKey:          "6a4b1efab9374f599059ff3fe4382089",
		AppSecret:       "44e380875cf046feb29ab14694566239",
		Type:            "apig",
		Creator:         "USER",
		Id:              "50f768cf1c1f4389965aa58d255b2a65",
		Name:            "terraform_test",
		RegistraionTime: "2021-06-19T03:15:23.181079903Z",
		Description:     "Created by script",
		Status:          1,
		UpdateTime:      "2021-06-19T03:15:23.181080043Z",
	}

	expectedGetResponseData = &applications.Application{
		AppKey:          "6a4b1efab9374f599059ff3fe4382089",
		AppSecret:       "44e380875cf046feb29ab14694566239",
		Type:            "apig",
		Creator:         "USER",
		Id:              "50f768cf1c1f4389965aa58d255b2a65",
		Name:            "terraform_test",
		RegistraionTime: "2021-06-19T03:15:23.181079903Z",
		Description:     "Created by script",
		Status:          1,
		UpdateTime:      "2021-06-19T03:15:23.181080043Z",
	}

	expectedListResponseData = []applications.Application{
		{
			AppKey:          "6a4b1efab9374f599059ff3fe4382089",
			AppSecret:       "44e380875cf046feb29ab14694566239",
			Type:            "apig",
			Creator:         "USER",
			Id:              "50f768cf1c1f4389965aa58d255b2a65",
			Name:            "terraform_test",
			RegistraionTime: "2021-06-19T03:15:23.181079903Z",
			Description:     "Created by script",
			Status:          1,
			UpdateTime:      "2021-06-19T03:15:23.181080043Z",
		},
	}

	appCodeOpts = applications.AppCodeOpts{
		AppCode: "dGVzdEFQSUNvZGVPZlRoZUFQSUdBcHBsaWNhdGlvbjEyMzQ1Njc4OTAhQCMkJSstLz10b0Jhc2U2NA==",
	}

	expectedAddCodeResponseData = &applications.AppCode{
		AppId:      "5b9ac7b0b2434860a264e05d8c31470e",
		Code:       "dGVzdEFQSUNvZGVPZlRoZUFQSUdBcHBsaWNhdGlvbjEyMzQ1Njc4OTAhQCMkJSstLz10b0Jhc2U2NA==",
		CreateTime: "2021-06-19T06:21:36.446251055Z",
		Id:         "5103639943e648599ef082285001c34e",
	}

	expectedListCodeResponseData = []applications.AppCode{
		{
			AppId:      "5b9ac7b0b2434860a264e05d8c31470e",
			Code:       "dGVzdEFQSUNvZGVPZlRoZUFQSUdBcHBsaWNhdGlvbjEyMzQ1Njc4OTAhQCMkJSstLz10b0Jhc2U2NA==",
			CreateTime: "2021-06-19T06:21:36.446251055Z",
			Id:         "5103639943e648599ef082285001c34e",
		},
	}
)

func handleV2ApplicationCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleV2ApplicationGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps/50f768cf1c1f4389965aa58d255b2a65",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2ApplicationList(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}

func handleV2ApplicationUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps/50f768cf1c1f4389965aa58d255b2a65",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2ApplicationDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps/50f768cf1c1f4389965aa58d255b2a65",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}

func handleV2ApplicationCodeAdd(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps/50f768cf1c1f4389965aa58d255b2a65/app-codes",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedAddCodeResponse)
		})
}

func handleV2ApplicationCodeAutoGenerate(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps/50f768cf1c1f4389965aa58d255b2a65/app-codes",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedAddCodeResponse)
		})
}

func handleV2ApplicationCodeGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps/50f768cf1c1f4389965aa58d255b2a65/app-codes/f284119e19f34d4caead4dd94114a7f4", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, expectedAddCodeResponse)
	})
}

func handleV2ApplicationCodeList(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps/50f768cf1c1f4389965aa58d255b2a65/app-codes",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListCodeResponse)
		})
}

func handleV2ApplicationCodeRemove(t *testing.T) {
	th.Mux.HandleFunc("/instances/c5faacb524d148b59ddd448dd02d016a/apps/50f768cf1c1f4389965aa58d255b2a65/app-codes"+
		"/f284119e19f34d4caead4dd94114a7f4", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}
