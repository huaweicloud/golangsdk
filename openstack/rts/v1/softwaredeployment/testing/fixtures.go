package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rts/v1/softwaredeployment"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

// CreateExpected represents the expected object from a Create request.
var CreateExpected = &softwaredeployment.Deployment{
	Id:           "43489279-7b12-4fc5-90ed-320f29e89419",
	ConfigId:     "69070672-d37d-4095-a19c-52ab1fde9a24",
	ServerId:     "a161a111-03a0-4204-b5c9-5df46587df5e",
	Status:       "IN_PROGRESS",
	Action:       "CREATE",
	StatusReason: "Deploy data available",
}

// CreateOutput represents the response body from a Create request.
const CreateOutput = `
{
    "software_deployment": {
        "status": "IN_PROGRESS",
        "server_id": "a161a111-03a0-4204-b5c9-5df46587df5e",
        "config_id": "69070672-d37d-4095-a19c-52ab1fde9a24",
        "action": "CREATE",
        "status_reason": "Deploy data available",
        "id": "43489279-7b12-4fc5-90ed-320f29e89419"
    }
}`

// HandleCreateSuccessfully creates an HTTP handler at `/stacks` on the test handler mux
// that responds with a `Create` response.
func HandleCreateSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/software_deployments", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
	})
}

// CreateExpected represents the expected object from a Create request.
var ListExpected = []softwaredeployment.Deployment{
	{
		Id:           "b357fbe9-4af8-4d00-9499-a13f0d0150bb",
		ConfigId:     "a6ff3598-f2e0-4111-81b0-aa3e1cac2529",
		ServerId:     "e4b191b0-b80b-4782-994c-02abb094480e",
		Status:       "IN_PROGRESS",
		Action:       "CREATE",
		StatusReason: "Deploy data available",
	},
	{
		Id:           "43489279-7b12-4fc5-90ed-320f29e89419",
		ConfigId:     "69070672-d37d-4095-a19c-52ab1fde9a24",
		ServerId:     "a161a111-03a0-4204-b5c9-5df46587df5e",
		Status:       "IN_PROGRESS",
		Action:       "CREATE",
		StatusReason: "Deploy data available",
	},
}

// CreateOutput represents the response body from a Create request.
const ListOutput = `
{
    "software_deployments": [
        {
            "status": "IN_PROGRESS",
            "server_id": "e4b191b0-b80b-4782-994c-02abb094480e",
            "config_id": "a6ff3598-f2e0-4111-81b0-aa3e1cac2529",
            "action": "CREATE",
            "status_reason": "Deploy data available",
            "id": "b357fbe9-4af8-4d00-9499-a13f0d0150bb"
        },
        {
            "status": "IN_PROGRESS",
			"server_id": "a161a111-03a0-4204-b5c9-5df46587df5e",
			"config_id": "69070672-d37d-4095-a19c-52ab1fde9a24",
			"action": "CREATE",
			"status_reason": "Deploy data available",
			"id": "43489279-7b12-4fc5-90ed-320f29e89419"
        }
         ]
}`

func HandleListSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/software_deployments", func(w http.ResponseWriter, r *http.Request) {
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
var GetExpected = &softwaredeployment.Deployment{
	Id:           "43489279-7b12-4fc5-90ed-320f29e89419",
	ConfigId:     "69070672-d37d-4095-a19c-52ab1fde9a24",
	ServerId:     "a161a111-03a0-4204-b5c9-5df46587df5e",
	Status:       "IN_PROGRESS",
	Action:       "CREATE",
	StatusReason: "Deploy data available",
}

// GetOutput represents the response body from a Get request.
const GetOutput = `
{
    "software_deployment": {
        "status": "IN_PROGRESS",
        "server_id": "a161a111-03a0-4204-b5c9-5df46587df5e",
        "config_id": "69070672-d37d-4095-a19c-52ab1fde9a24",
        "action": "CREATE",
        "status_reason": "Deploy data available",
        "id": "43489279-7b12-4fc5-90ed-320f29e89419"
    }
}`

// HandleGetSuccessfully creates an HTTP handler at `/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87`
// on the test handler mux that responds with a `Get` response.
func HandleGetSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/software_deployments/43489279-7b12-4fc5-90ed-320f29e89419", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
	})
}

// GetExpected represents the expected object from a Get request.
var UpdateExpected = &softwaredeployment.Deployment{
	Id:           "43489279-7b12-4fc5-90ed-320f29e89419",
	ConfigId:     "69070672-d37d-4095-a19c-52ab1fde9a24",
	ServerId:     "a161a111-03a0-4204-b5c9-5df46587df5e",
	Status:       "COMPLETE",
	Action:       "CREATE",
	StatusReason: "Outputs received",
	OutputValues: map[string]interface{}{"deploy_stdout": "Writing to /tmp/baaaaa\nWritten to /tmp/baaaaa\n", "deploy_stderr": "+ echo Writing to /tmp/baaaaa\n+ echo fooooo\n+ cat /tmp/baaaaa\n+ echo -n The file /tmp/baaaaa contains fooooo for server ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5 during CREATE\n+ echo Written to /tmp/baaaaa\n+ echo Output to stderr\nOutput to stderr\n",
		"deploy_status_code": "0", "result": "The file /tmp/baaaaa contains fooooo for server ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5 during CREATE"},
}

// GetOutput represents the response body from a Get request.
const UpdateOutput = `
{
    "software_deployment": {
        "status": "COMPLETE",
        "server_id": "a161a111-03a0-4204-b5c9-5df46587df5e",
        "config_id": "69070672-d37d-4095-a19c-52ab1fde9a24",
        "output_values": {
            "deploy_stdout": "Writing to /tmp/baaaaa\nWritten to /tmp/baaaaa\n",
            "deploy_stderr": "+ echo Writing to /tmp/baaaaa\n+ echo fooooo\n+ cat /tmp/baaaaa\n+ echo -n The file /tmp/baaaaa contains fooooo for server ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5 during CREATE\n+ echo Written to /tmp/baaaaa\n+ echo Output to stderr\nOutput to stderr\n",
            "deploy_status_code": "0",
            "result": "The file /tmp/baaaaa contains fooooo for server ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5 during CREATE"
        },
        "action": "CREATE",
        "status_reason": "Outputs received",
        "id": "43489279-7b12-4fc5-90ed-320f29e89419"
    }
}`

// HandleUpdateSuccessfully creates an HTTP handler at `/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87`
// on the test handler mux that responds with a `Get` response.
func HandleUpdateSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/software_deployments/43489279-7b12-4fc5-90ed-320f29e89419", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
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
	th.Mux.HandleFunc("/software_deployments/43489279-7b12-4fc5-90ed-320f29e89419", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}
