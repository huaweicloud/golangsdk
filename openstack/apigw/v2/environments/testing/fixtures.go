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
}`

	expectedCreateVariableResponse = `
{
	"env_id": "3585fce96a5d44f8b445121b9440274a",
	"group_id": "bd7c9608c05e4e93a6b44e47f19b6bed",
	"id": "2dc48632332f4157804175175e71e3e8",
	"variable_name": "Path",
	"variable_value": "/stage/test"
}`

	expectedListVariablesResponse = `
{
	"variables": [
		{
			"env_id": "3585fce96a5d44f8b445121b9440274a",
			"group_id": "bd7c9608c05e4e93a6b44e47f19b6bed",
			"id": "2dc48632332f4157804175175e71e3e8",
			"variable_name": "Path",
			"variable_value": "/stage/test"
		}
	]
}`
)

var (
	createDesc = "Updated by script"
	createOpts = &environments.EnvironmentOpts{
		Name:        "terraform_test",
		Description: &createDesc,
	}

	expectedCreateResponseData = &environments.Environment{
		Id:          "3585fce96a5d44f8b445121b9440274a",
		Name:        "terraform_test",
		Description: "Created by script",
		CreateTime:  "2021-06-22T12:11:28.18948645Z",
	}

	updateDesc = "Updated by script"
	updateOpts = &environments.EnvironmentOpts{
		Name:        "terraform_test",
		Description: &updateDesc,
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

	variableCreateOpts = &environments.CreateVariableOpts{
		EnvId:   "3585fce96a5d44f8b445121b9440274a",
		GroupId: "bd7c9608c05e4e93a6b44e47f19b6bed",
		Name:    "Path",
		Value:   "/stage/test",
	}

	expectedCreateVariableResponseData = &environments.Variable{
		EnvId:   "3585fce96a5d44f8b445121b9440274a",
		GroupId: "bd7c9608c05e4e93a6b44e47f19b6bed",
		Name:    "Path",
		Value:   "/stage/test",
		Id:      "2dc48632332f4157804175175e71e3e8",
	}

	expectedListVariableResponseData = []environments.Variable{
		{
			EnvId:   "3585fce96a5d44f8b445121b9440274a",
			GroupId: "bd7c9608c05e4e93a6b44e47f19b6bed",
			Name:    "Path",
			Value:   "/stage/test",
			Id:      "2dc48632332f4157804175175e71e3e8",
		},
	}
)

func handleV2EnvironmentCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/envs",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleV2EnvironmentList(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/envs",
		func(w http.ResponseWriter, r *http.Request) {
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

func handleV2EnvironmentVariableCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/env-variables",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedCreateVariableResponse)
		})
}

func handleV2EnvironmentVariableGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/env-variables"+
		"/2dc48632332f4157804175175e71e3e8", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateVariableResponse)
	})
}

func handleV2EnvironmentVariableList(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/env-variables",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListVariablesResponse)
		})
}

func handleV2EnvironmentVariableDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/cc4ea721cc6747f7969e06bd21121c52/env-variables"+
		"/2dc48632332f4157804175175e71e3e8", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}
