package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/antiddos/v2/alarmreminding"

	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

var WarnAlertOutput = `
{
  "warn_config": {
    "antiDDoS": false,
    "bruce_force": true,
    "remote_login": false,
    "weak_password": false,
    "high_privilege": false,
    "back_doors": false,
    "waf": true,
    "send_frequency": 1
  },
  "topic_urn": "ABC",
  "display_name": "123"
}
`

var WarnAlertResponse = alarmreminding.WarnAlertResponse{
	WarnConfig: struct {
		AntiDDoS      bool `json:"antiDDoS,"`
		BruceForce    bool `json:"bruce_force,omitempty"`
		RemoteLogin   bool `json:"remote_login,omitempty"`
		WeakPassword  bool `json:"weak_password,omitempty"`
		HighPrivilege bool `json:"high_privilege,omitempty"`
		BackDoors     bool `json:"back_doors,omitempty"`
		Waf           bool `json:"waf,omitempty"`
		SendFrequency int  `json:"send_frequency,omitempty"`
	}{
		AntiDDoS:      false,
		BruceForce:    true,
		RemoteLogin:   false,
		WeakPassword:  false,
		HighPrivilege: false,
		BackDoors:     false,
		Waf:           true,
		SendFrequency: 1,
	},
	TopicUrn:    "ABC",
	DisplayName: "123",
}

func HandleWarnAlertSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/warnalert/alertconfig/query", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, WarnAlertOutput)
	})
}
