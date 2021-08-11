package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/rds/v3/securities"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedEnableSSLResponse = `{}`

	expectedUpdateResponse = `
{
	"workflowId": "e982d9d7-d96f-4d25-a591-b6be03c93081"
}
`
)

var (
	sslOpts = securities.SSLOpts{
		SSLEnable: golangsdk.Enabled,
	}

	portOpts = securities.PortOpts{
		Port: 3309,
	}

	secGroupOpts = securities.SecGroupOpts{
		SecurityGroupId: "71aa11f4-7d6f-479c-b1b4-123f31412e21",
	}

	expectedGetResponseData = &securities.WorkFlow{
		WorkflowId: "e982d9d7-d96f-4d25-a591-b6be03c93081",
	}
)

func handleV2DatabaseSSLUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/fda30974248d449e9dbdce8ae65d5ba0in01/ssl",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedUpdateResponse)
		})
}

func handleV2DatabasePortUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/fda30974248d449e9dbdce8ae65d5ba0in01/port",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedUpdateResponse)
		})
}

func handleV2SecurityGroupUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/fda30974248d449e9dbdce8ae65d5ba0in01/security-group",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedUpdateResponse)
		})
}
