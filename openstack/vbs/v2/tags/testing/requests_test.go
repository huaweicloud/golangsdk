package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/vbs/v2/common"
	"github.com/huaweicloud/golangsdk/openstack/vbs/v2/tags"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestCreateV2Tag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/backuppolicy/ed8b9f73-4415-494d-a54e-5f3373bc353d/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, AddTag)
		w.WriteHeader(http.StatusNoContent)
	})
	options := tags.CreateOpts{Tag: tags.Tag{Key: "0f187b65-8d0e-4fc0-9096-3b55d330531e", Value: "volume"}}
	s := tags.Create(fake.ServiceClient(), "ed8b9f73-4415-494d-a54e-5f3373bc353d", options)
	th.AssertNoErr(t, s.Err)
}

func TestDeleteV2Tag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/backuppolicy/ed8b9f73-4415-494d-a54e-5f3373bc353d/tags/K1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	delete := tags.Delete(fake.ServiceClient(), "ed8b9f73-4415-494d-a54e-5f3373bc353d", "K1")
	th.AssertNoErr(t, delete.Err)
}

func TestGetV2Tag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/backuppolicy/ed8b9f73-4415-494d-a54e-5f3373bc353d/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getTags)
	})

	s, err := tags.Get(fake.ServiceClient(), "ed8b9f73-4415-494d-a54e-5f3373bc353d").Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, &tags.RespTags{Tags: []tags.Tag{{Key: "RUNNING", Value: "0781095c-b8ab-4ce5-99f3-4c5f6ff75319"}, {Key: "WAITING", Value: ""}}})
}

func TestBatchActionsV2Tag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/backuppolicy/ed8b9f73-4415-494d-a54e-5f3373bc353d/tags/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, batchAction)
		w.WriteHeader(http.StatusNoContent)
	})
	options := tags.BatchOpts{Action: tags.ActionUpdate, Tags: []tags.Tag{{Key: "0f187b65-8d0e-4fc0-9096-3b55d330531e", Value: "volume"}, {Key: "0f187b65-8d0e-4fc0-9096-3b55d330531d", Value: "volume"}}}
	s := tags.BatchAction(fake.ServiceClient(), "ed8b9f73-4415-494d-a54e-5f3373bc353d", options)
	th.AssertNoErr(t, s.Err)
}

func TestQueryV2Tags(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/backuppolicy/resource_instances/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `
{ 
  "tags":
      [
        {
           "key": "Tag001",
           "values":["Value001","Value002"]
         }
       ],
  "action":"filter"
}
`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, TagList)
	})
	queryOpts := tags.ListOpts{Action: "filter", Tags: []tags.Tags{{Key: "Tag001", Values: []string{"Value001", "Value002"}}}}
	actual, err := tags.ListResources(fake.ServiceClient(), queryOpts).ExtractResources()
	th.AssertNoErr(t, err)
	expected := ExpectedTags
	th.AssertDeepEquals(t, expected, actual)
}
