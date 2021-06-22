package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/environments"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"create_time": "2021-06-22T12:11:28.18948645Z",
	"id": "3585fce96a5d44f8b445121b9440274a",
	"name": "terraform_test",
	"remark": "Created by script"
}`
	expectedListResponse = `
{
	"envs": [
		{
			"create_time": "2020-10-14T16:15:19.417962Z",
			"id": "DEFAULT_ENVIRONMENT_RELEASE_ID",
			"name": "RELEASE",
			"remark": "生产环境"
		},
		{
			"create_time": "2021-06-22T11:40:18Z",
			"id": "3585fce96a5d44f8b445121b9440274a",
			"name": "terraform_test_update",
			"remark": "Updated by script"
		}
	]
}`
	expectedUpdateResponse = `
{
	"create_time": "2021-06-22T11:40:18Z",
	"id": "3585fce96a5d44f8b445121b9440274a",
	"name": "terraform_test_update",
	"remark": "Updated by script"
}
`
)

var (
	createOpts = &environments.EnvironmentOpts{
		Name:        "terraform_test",
		Description: "Created by script",
	}

	expectedCreateResponseData = &environments.Environment{
		Id:          "3585fce96a5d44f8b445121b9440274a",
		Name:        "terraform_test",
		Description: "Created by script",
		CreateTime:  "2021-06-22T12:11:28.18948645Z",
	}

	updateOpts = &environments.EnvironmentOpts{
		Name:        "terraform_test",
		Description: "Created by script",
	}

	expectedUpdateResponseData = &environments.Environment{
		Id:          "3585fce96a5d44f8b445121b9440274a",
		Name:        "terraform_test_update",
		Description: "Updated by script",
		CreateTime:  "2021-06-22T11:40:18Z",
	}

	expectedListResponseData = []environments.Environment{
		{
			CreateTime:  "2020-10-14T16:15:19.417962Z",
			Id:          "DEFAULT_ENVIRONMENT_RELEASE_ID",
			Name:        "RELEASE",
			Description: "生产环境",
		},
		{
			CreateTime:  "2021-06-22T11:40:18Z",
			Id:          "3585fce96a5d44f8b445121b9440274a",
			Name:        "terraform_test_update",
			Description: "Updated by script",
		},
	}
)

func handleV2EnvironmentCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/envs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV2EnvironmentList(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/envs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}

func handleV2EnvironmentUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/envs/3585fce96a5d44f8b445121b9440274a",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedUpdateResponse)
		})
}

func handleV2EnvironmentDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/envs/3585fce96a5d44f8b445121b9440274a",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}
