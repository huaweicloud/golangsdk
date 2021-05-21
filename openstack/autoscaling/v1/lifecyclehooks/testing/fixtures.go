package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/autoscaling/v1/lifecyclehooks"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedGetResponse = `
{
	"lifecycle_hook_name": "test-hook",
	"default_result": "ABANDON",
	"default_timeout": 3600,
	"notification_topic_urn": "urn:smn:regionId:b53e5554fad0494d96206fb84296510b:gsh",
	"notification_topic_name": "gsh",
	"lifecycle_hook_type": "INSTANCE_LAUNCHING",
	"notification_metadata": null,
	"create_time": "2019-03-18T16:00:11Z"
}`

	expectedUpdateResponse = `
{
	"lifecycle_hook_name": "test-hook",
	"default_result": "CONTINUE",
	"default_timeout": 3600,
	"notification_topic_urn": "urn:smn:regionId:b53e5554fad0494d96206fb84296510b:gsh",
	"notification_topic_name": "gsh",
	"lifecycle_hook_type": "INSTANCE_LAUNCHING",
	"notification_metadata": null,
	"create_time": "2019-03-18T16:00:11Z"
}`

	expectedListResponse = `
{
	"lifecycle_hooks": [
		{
			"lifecycle_hook_name": "test-hook",
			"default_result": "ABANDON",
			"default_timeout": 3600,
			"notification_topic_urn": "urn:smn:regionId:b53e5554fad0494d96206fb84296510b:gsh",
			"notification_topic_name": "gsh",
			"lifecycle_hook_type": "INSTANCE_LAUNCHING",
			"notification_metadata": null,
			"create_time": "2019-03-18T16:00:11Z"
		}
	]
}`
)

var (
	createOpts = lifecyclehooks.CreateOpts{
		Name:                 "test-hook",
		Type:                 "INSTANCE_LAUNCHING",
		DefaultResult:        "ABANDON",
		Timeout:              3600,
		NotificationTopicURN: "urn:smn:regionId:b53e5554fad0494d96206fb84296510b:gsh",
	}

	expectedGetResponseData = &lifecyclehooks.Hook{
		Name:                  "test-hook",
		DefaultResult:         "ABANDON",
		Type:                  "INSTANCE_LAUNCHING",
		Timeout:               3600,
		NotificationTopicURN:  "urn:smn:regionId:b53e5554fad0494d96206fb84296510b:gsh",
		NotificationTopicName: "gsh",
		NotificationMetadata:  "",
		CreateTime:            "2019-03-18T16:00:11Z",
	}

	expectedListResponseData = &[]lifecyclehooks.Hook{
		{
			Name:                  "test-hook",
			DefaultResult:         "ABANDON",
			Type:                  "INSTANCE_LAUNCHING",
			Timeout:               3600,
			NotificationTopicURN:  "urn:smn:regionId:b53e5554fad0494d96206fb84296510b:gsh",
			NotificationTopicName: "gsh",
			NotificationMetadata:  "",
			CreateTime:            "2019-03-18T16:00:11Z",
		},
	}

	updateOpts = lifecyclehooks.UpdateOpts{
		DefaultResult: "CONTINUE",
	}

	expectedUpdateResponseData = &lifecyclehooks.Hook{
		Name:                  "test-hook",
		DefaultResult:         "CONTINUE",
		Type:                  "INSTANCE_LAUNCHING",
		Timeout:               3600,
		NotificationTopicURN:  "urn:smn:regionId:b53e5554fad0494d96206fb84296510b:gsh",
		NotificationTopicName: "gsh",
		NotificationMetadata:  "",
		CreateTime:            "2019-03-18T16:00:11Z",
	}
)

func handleLifeCycleHookCreate(t *testing.T) {
	th.Mux.HandleFunc("/scaling_lifecycle_hook/50ed20b8-9853-4668-a71c-c8c15b5cb85f",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleLifeCycleHookGet(t *testing.T) {
	th.Mux.HandleFunc("/scaling_lifecycle_hook/50ed20b8-9853-4668-a71c-c8c15b5cb85f/test-hook",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleLifeCycleHookList(t *testing.T) {
	th.Mux.HandleFunc("/scaling_lifecycle_hook/50ed20b8-9853-4668-a71c-c8c15b5cb85f/list",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListResponse)
		})
}

func handleLifeCycleHookUpdate(t *testing.T) {
	th.Mux.HandleFunc("/scaling_lifecycle_hook/50ed20b8-9853-4668-a71c-c8c15b5cb85f/test-hook",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedUpdateResponse)
		})
}

func handleLifeCycleHookDelete(t *testing.T) {
	th.Mux.HandleFunc("/scaling_lifecycle_hook/50ed20b8-9853-4668-a71c-c8c15b5cb85f/test-hook",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}
