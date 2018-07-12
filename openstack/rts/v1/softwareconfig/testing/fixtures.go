package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rts/v1/softwareconfig"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

// CreateExpected represents the expected object from a Create request.
var CreateExpected = &softwareconfig.SoftwareConfig{
	Inputs: []map[string]interface{}{{"type": "String", "name": "foo"},
		{"type": "String", "name": "bar"}},
	Group:   "script",
	Name:    "test-cong",
	Outputs: []map[string]interface{}{{"type": "String", "name": "result", "error_output": "false"}},
	Config:  "#!/bin/sh -x\necho \"Writing to /tmp/$bar\"\necho $foo > /tmp/$bar\necho -n \"The file /tmp/$bar contains cat /tmp/$bar for server $deploy_server_id during $deploy_action\" > $heat_outputs_path.result\necho \"Written to /tmp/$bar\"\necho \"Output to stderr\" 1>&2",
	Id:      "e0be7e37-a581-4b24-bfb1-df4f3048c090",
}

// CreateOutput represents the response body from a Create request.
const CreateOutput = `
{
    "software_config": {
        "inputs": [
            {
                "type": "String",
                "name": "foo"
            },
            {
                "type": "String",
                "name": "bar"
            }
        ],
        "group": "script",
        "name": "test-cong",
        "outputs": [
            {
                "type": "String",
                "name": "result",
                "error_output": "false"
            }
        ],
        "config": "#!/bin/sh -x\necho \"Writing to /tmp/$bar\"\necho $foo > /tmp/$bar\necho -n \"The file /tmp/$bar contains cat /tmp/$bar for server $deploy_server_id during $deploy_action\" > $heat_outputs_path.result\necho \"Written to /tmp/$bar\"\necho \"Output to stderr\" 1>&2",
        "id": "e0be7e37-a581-4b24-bfb1-df4f3048c090"
    }
}`

// HandleCreateSuccessfully creates an HTTP handler at `/stacks` on the test handler mux
// that responds with a `Create` response.
func HandleCreateSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/software_configs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
	})
}

// ListExpected represents the expected object from a List request.
var ListExpected = []softwareconfig.SoftwareConfig{
	{
		Group: "script",
		Name:  "test-cong",
		Id:    "e0be7e37-a581-4b24-bfb1-df4f3048c090",
	},
	{
		Group: "script",
		Name:  "test-cong1",
		Id:    "743a15f6-9a55-49fe-80bb-a9188f39fc07",
	},
	{
		Group: "script",
		Name:  "a-config-we5zpvyu7b5o",
		Id:    "a6ff3598-f2e0-4111-81b0-aa3e1cac2529",
	},
}

// FullListOutput represents the response body from a List request without a marker.
const FullListOutput = `
{
    "software_configs": [
        {
            "group": "script",
            "id": "e0be7e37-a581-4b24-bfb1-df4f3048c090",
            "name": "test-cong"
        },
        {
            "group": "script",
            "id": "743a15f6-9a55-49fe-80bb-a9188f39fc07",
            "name": "test-cong1"
        },
        {
            "group": "script",
            "id": "a6ff3598-f2e0-4111-81b0-aa3e1cac2529",
            "name": "a-config-we5zpvyu7b5o"
        }
    ]
}`

// HandleListSuccessfully creates an HTTP handler at `/stacks` on the test handler mux
// that responds with a `List` response.
func HandleListSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/software_configs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
		//r.ParseForm()
	})
}

// GetExpected represents the expected object from a Get request.
var GetExpected = &softwareconfig.SoftwareConfig{
	Inputs: []map[string]interface{}{{"type": "String", "name": "foo"},
		{"type": "String", "name": "bar"}},
	Group:   "script",
	Name:    "test-cong",
	Outputs: []map[string]interface{}{{"type": "String", "name": "result", "error_output": "false"}},
	Config:  "#!/bin/sh -x\necho \"Writing to /tmp/$bar\"\necho $foo > /tmp/$bar\necho -n \"The file /tmp/$bar contains cat /tmp/$bar for server $deploy_server_id during $deploy_action\" > $heat_outputs_path.result\necho \"Written to /tmp/$bar\"\necho \"Output to stderr\" 1>&2",
	Id:      "e0be7e37-a581-4b24-bfb1-df4f3048c090",
}

// GetOutput represents the response body from a Get request.
const GetOutput = `
{
    "software_config": {
        "inputs": [
            {
                "type": "String",
                "name": "foo"
            },
            {
                "type": "String",
                "name": "bar"
            }
        ],
        "group": "script",
        "name": "test-cong",
        "outputs": [
            {
                "type": "String",
                "name": "result",
                "error_output": "false"
            }
        ],
        "config": "#!/bin/sh -x\necho \"Writing to /tmp/$bar\"\necho $foo > /tmp/$bar\necho -n \"The file /tmp/$bar contains cat /tmp/$bar for server $deploy_server_id during $deploy_action\" > $heat_outputs_path.result\necho \"Written to /tmp/$bar\"\necho \"Output to stderr\" 1>&2",
        "id": "e0be7e37-a581-4b24-bfb1-df4f3048c090"
    }
}`

// HandleGetSuccessfully creates an HTTP handler at `/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87`
// on the test handler mux that responds with a `Get` response.
func HandleGetSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/software_configs/e0be7e37-a581-4b24-bfb1-df4f3048c090", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
	})
}

// HandleDeleteSuccessfully creates an HTTP handler at `/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87`
// on the test handler mux that responds with a `Delete` response.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/software_configs/e2fe5553-a481-4549-9d0f-e208de3d98d1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}
