package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/ecs/v1/cloudservers"
	"github.com/huaweicloud/golangsdk/openstack/ecs/v1/powers"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedPowerOnRequest = `
{
	"os-start": {
		"servers": [
			{
				"id": "42ecde1c-d97d-40c6-ac84-39ed9c35dc90"
			}
		]
	}
}`

	expectedPowerOffRequest = `
{
	"os-stop": {
		"servers": [
			{
				"id": "42ecde1c-d97d-40c6-ac84-39ed9c35dc90"
			}
		]
	}
}`

	expectedPowerOnResponse = `
{
	"job_id": "ff808081787e9f100179693fd6f92e39"
}`

	expectedPowerOffResponse = `
{
	"job_id": "ff808081787ea4b5017969510b0543e6"
}`
)

var (
	powerOpts = &powers.PowerOpts{
		Servers: []powers.ServerInfo{
			{
				ID: "42ecde1c-d97d-40c6-ac84-39ed9c35dc90",
			},
		},
	}

	expectedPowerOnResponseData = &cloudservers.JobResponse{
		JobID: "ff808081787e9f100179693fd6f92e39",
	}

	expectedPowerOffResponseData = &cloudservers.JobResponse{
		JobID: "ff808081787ea4b5017969510b0543e6",
	}
)

func handlePowerOn(t *testing.T) {
	th.Mux.HandleFunc("/cloudservers/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedPowerOnResponse)
	})
}

func handlePowerOff(t *testing.T) {
	th.Mux.HandleFunc("/cloudservers/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedPowerOffResponse)
	})
}
