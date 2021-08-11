package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/fgs/v2/trigger"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedGetResponse = `
{
	"trigger_id": "971f9cff-5d29-42da-ba36-e4e2b9d26664",
	"trigger_type_code": "TIMER",
	"trigger_status": "ACTIVE",
	"event_data": {
		"name": "test",
		"schedule": "1d",
		"schedule_type": "Rate"
	},
	"last_updated_time": "2021-08-11T17:03:20+08:00",
	"created_time": "2021-08-11T17:03:20+08:00"
}`

	expectedListResponse = `
[
	{
		"trigger_id": "971f9cff-5d29-42da-ba36-e4e2b9d26664",
		"trigger_type_code": "TIMER",
		"trigger_status": "ACTIVE",
		"event_data": {
			"name": "test",
			"schedule": "1d",
			"schedule_type": "Rate"
		},
		"last_updated_time": "2021-08-11T17:03:20+08:00",
		"created_time": "2021-08-11T17:03:20+08:00"
	}
]`
)

var (
	createOpts = &trigger.CreateOpts{
		TriggerTypeCode: "TIMER",
		TriggerStatus:   "ACTIVE",
		EventTypeCode:   "MessageCreated",
		EventData: map[string]interface{}{
			"name":          "test",
			"schedule":      "1d",
			"schedule_type": "Rate",
		},
	}

	expectedGetResponseData = &trigger.Trigger{
		CreatedTime:     "2021-08-11T17:03:20+08:00",
		TriggerId:       "971f9cff-5d29-42da-ba36-e4e2b9d26664",
		TriggerTypeCode: "TIMER",
		EventData: map[string]interface{}{
			"name":          "test",
			"schedule":      "1d",
			"schedule_type": "Rate",
		},
		Status:          "ACTIVE",
		LastUpdatedTime: "2021-08-11T17:03:20+08:00",
	}

	expectedListResponseData = []trigger.Trigger{
		{
			CreatedTime:     "2021-08-11T17:03:20+08:00",
			TriggerId:       "971f9cff-5d29-42da-ba36-e4e2b9d26664",
			TriggerTypeCode: "TIMER",
			EventData: map[string]interface{}{
				"name":          "test",
				"schedule":      "1d",
				"schedule_type": "Rate",
			},
			Status:          "ACTIVE",
			LastUpdatedTime: "2021-08-11T17:03:20+08:00",
		},
	}

	updateOpts = trigger.UpdateOpts{
		TriggerStatus: "ACTIVE",
	}
)

func handleV2TriggerCreate(t *testing.T) {
	th.Mux.HandleFunc("/fgs/triggers/urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2TriggerGet(t *testing.T) {
	th.Mux.HandleFunc("/fgs/triggers/urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test"+
		"/TIMER/971f9cff-5d29-42da-ba36-e4e2b9d26664", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedGetResponse)
	})
}

func handleV2TriggerList(t *testing.T) {
	th.Mux.HandleFunc("/fgs/triggers/urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListResponse)
		})
}

func handleV2TriggerUpdate(t *testing.T) {
	th.Mux.HandleFunc("/fgs/triggers/urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test"+
		"/TIMER/971f9cff-5d29-42da-ba36-e4e2b9d26664", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedGetResponse)
	})
}

func handleV2TriggerDelete(t *testing.T) {
	th.Mux.HandleFunc("/fgs/triggers/urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test"+
		"/TIMER/971f9cff-5d29-42da-ba36-e4e2b9d26664", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}
