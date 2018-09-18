package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/vbs/v2/common"
	"github.com/huaweicloud/golangsdk/openstack/vbs/v2/policies"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestCreateV2Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/backuppolicy", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Set("Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `
{
    "backup_policy_name": "Test_Policy",
    "scheduled_policy" : {
        "remain_first_backup_of_curMonth" : "Y",
        "rentention_num" : 10,
        "frequency" : 1,
        "start_time" : "12:00",
        "status" : "ON"
    },
    "tags":[{
      "key":"key",
      "value":"value"
    }]
}
`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, Output)
	})
	createOptions := policies.CreateOpts{Name: "Test_Policy", ScheduledPolicy: policies.ScheduledPolicy{StartTime: "12:00", Status: "ON", Frequency: 1, RententionNum: 10, RemainFirstBackup: "Y"}, Tags: []policies.Tag{{Key: "key", Value: "value"}}}
	actual, err := policies.Create(fake.ServiceClient(), createOptions).Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)
}

func TestUpdateV2Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/backuppolicy/af8a20b0-117d-4fc3-ae53-aa3968a4f870", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `
{
    "backup_policy_name": "Test_02",
    "scheduled_policy" : {
        "remain_first_backup_of_curMonth" : "Y",
        "rentention_num" : 10,
        "frequency" : 1,
        "start_time" : "10:00",
        "status" : "ON"
    }
}
`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateOutput)
	})
	updateOptions := policies.UpdateOpts{Name: "Test_02", ScheduledPolicy: policies.UpdateSchedule{StartTime: "10:00", Status: "ON", Frequency: 1, RententionNum: 10, RemainFirstBackup: "Y"}}
	update, err := policies.Update(fake.ServiceClient(), "af8a20b0-117d-4fc3-ae53-aa3968a4f870", updateOptions).Extract()
	th.AssertNoErr(t, err)
	expected := Update
	th.AssertDeepEquals(t, expected, update)
}

func TestDeleteV2Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/backuppolicy/af8a20b0-117d-4fc3-ae53-aa3968a4f870", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	delete := policies.Delete(fake.ServiceClient(), "af8a20b0-117d-4fc3-ae53-aa3968a4f870")
	th.AssertNoErr(t, delete.Err)
}

func TestListV2Policy(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/backuppolicy", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListOutput)
	})

	actual, err := policies.List(fake.ServiceClient(), policies.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract clusters: %v", err)
	}

	expected := ListPolicies

	th.AssertDeepEquals(t, expected, actual)
}

func TestAssociateV2Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/backuppolicyresources", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `
{
    "backup_policy_id":"915d1fd8-63cb-4054-a2b0-2778210e3a75",
    "resources":[{
        "resource_id":"0f187b65-8d0e-4fc0-9096-3b55d330531e",
        "resource_type":"volume"
        }]
}
`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, AssociateOutput)
	})
	AssoOpts := policies.AssociateOpts{PolicyID: "915d1fd8-63cb-4054-a2b0-2778210e3a75", Resources: []policies.AssociateResource{{ResourceID: "0f187b65-8d0e-4fc0-9096-3b55d330531e", ResourceType: "volume"}}}
	associate, err := policies.Associate(fake.ServiceClient(), AssoOpts).ExtractResource()
	th.AssertNoErr(t, err)
	expected := Associate
	th.AssertDeepEquals(t, expected, associate)
}

func TestDisassociateV2Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/backuppolicyresources/915d1fd8-63cb-4054-a2b0-2778210e3a75/deleted_resources", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `
{
    "resources": [
        {
            "resource_id": "0f187b65-8d0e-4fc0-9096-3b55d330531e"
        }
    ]
}
`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, DisssociateOutput)
	})
	options := policies.DisassociateOpts{Resources: []policies.DisassociateResource{{ResourceID: "0f187b65-8d0e-4fc0-9096-3b55d330531e"}}}
	associate, err := policies.Disassociate(fake.ServiceClient(), "915d1fd8-63cb-4054-a2b0-2778210e3a75", options).ExtractResource()
	th.AssertNoErr(t, err)
	expected := Disassociate
	th.AssertDeepEquals(t, expected, associate)
}
