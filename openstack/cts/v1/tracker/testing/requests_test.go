package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/cts/v1/tracker"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/tracker",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, createResponse)
		})

	options := &tracker.CreateOptsWithSMN{
		BucketName:     "obs-e51d",
		FilePrefixName: "yO8Q",
		SimpleMessageNotification: tracker.SimpleMessageNotification{
			IsSupportSMN:          true,
			TopicID:               "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:c2c-topic",
			IsSendAllKeyOperation: false,
			Operations:            []string{"login"},
			NeedNotifyUserList:    []string{"user1", "user2"},
		}}
	n, err := tracker.Create(fake.ServiceClient(), options).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, n.BucketName, "obs-e51d")
	th.AssertEquals(t, n.Status, "enabled")
	th.AssertEquals(t, n.TrackerName, "system")
	th.AssertEquals(t, n.SimpleMessageNotification.TopicID, "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:c2c-topic")
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/tracker"+"/"+"system", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, updateResponse)
	})

	options := &tracker.UpdateOptsWithSMN{
		BucketName:     "cirros-img",
		FilePrefixName: "yO8Q",
		Status:         "disabled",
		SimpleMessageNotification: tracker.SimpleMessageNotification{
			IsSupportSMN:          false,
			TopicID:               "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:c2c-topic",
			IsSendAllKeyOperation: false,
			Operations:            []string{"delete", "create", "login"},
			NeedNotifyUserList:    []string{"user1", "user2"},
		},
	}
	n, err := tracker.Update(fake.ServiceClient(), options).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, n.BucketName, "cirros-img")
	th.AssertEquals(t, n.FilePrefixName, "yO8Q")
	th.AssertEquals(t, n.SimpleMessageNotification.TopicID, "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:c2c-topic")
	th.AssertEquals(t, n.Status, "disabled")
	th.AssertEquals(t, n.TrackerName, "system")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/tracker",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			w.WriteHeader(http.StatusNoContent)
		})

	result := tracker.Delete(fake.ServiceClient())
	th.AssertNoErr(t, result.Err)
}

func TestList(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/tracker", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	actual, err := tracker.List(fake.ServiceClient(), tracker.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract cts tracker: %v", err)
	}

	expected := []tracker.Tracker{
		{
			Status:         "enabled",
			BucketName:     "tf-test-bucket",
			FilePrefixName: "yO8Q",
			TrackerName:    "system",
			SimpleMessageNotification: tracker.SimpleMessageNotification{
				IsSupportSMN:          true,
				TopicID:               "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:tf-test-topic",
				IsSendAllKeyOperation: false,
				Operations:            []string{"login"},
				NeedNotifyUserList:    []string{"user1"},
			},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}
