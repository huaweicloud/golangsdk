package testing

import (
	"fmt"
	"github.com/huaweicloud/golangsdk/openstack/sfs/v2/shares"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
	"net/http"
	"testing"
	"time"
)

func TestCreateShare(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	options := &shares.CreateOpts{Size: 1, Name: "my_test_share", ShareProto: "NFS"}
	n, err := shares.Create(client.ServiceClient(), options).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, n.Name, "my_test_share")
	th.AssertEquals(t, n.Size, 1)
	th.AssertEquals(t, n.ShareProto, "NFS")
}

func TestDeleteShare(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteResponse(t)

	result := shares.Delete(client.ServiceClient(), shareID)
	th.AssertNoErr(t, result.Err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateResponse(t)

	options := &shares.UpdateOpts{DisplayName: "my_test_share_sfs", DisplayDescription: "test"}
	n, err := shares.Update(client.ServiceClient(), shareID, options).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, n.Name, "my_test_share_sfs")
	th.AssertEquals(t, n.Description, "test")
}

func TestGetShare(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	s, err := shares.Get(client.ServiceClient(), shareID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, &shares.Share{
		AvailabilityZone: "nova",
		ShareNetworkID:   "713df749-aac0-4a54-af52-10f6c991e80c",
		SnapshotID:       "",
		ID:               shareID,
		Size:             1,
		ShareType:        "25747776-08e5-494f-ab40-a64b9d20d8f7",
		ProjectID:        "16e1ab15c35a457e9c2b2aa189f544e1",
		Metadata: map[string]string{
			"project": "my_app",
			"aim":     "doc",
		},
		Status:      "available",
		Description: "My custom share London",
		Host:        "manila2@generic1#GENERIC1",
		Name:        "my_test_share",
		CreatedAt:   time.Date(2015, time.September, 18, 10, 25, 24, 0, time.UTC),
		ShareProto:  "NFS",
		VolumeType:  "default",
		IsPublic:    true,
		Links: []map[string]string{
			{
				"href": "http://172.18.198.54:8786/v2/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
				"rel":  "self",
			},
			{
				"href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
				"rel":  "bookmark",
			},
		},
	})
}

func TestListAccessRights(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, listAccessRightsRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, listAccessRightsResponse)
	})

	c := client.ServiceClient()
	// Client c must have Microversion set; minimum supported microversion for Grant Access is 2.7
	c.Microversion = "2.7"

	s, err := shares.ListAccessRights(c, shareID).ExtractAccessRights()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, []shares.AccessRight{
		{
			AccessType:  "cert",
			AccessTo:    "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
			AccessLevel: "rw",
			State:       "active",
			ID:          "5158f095-4c43-49c0-b5a7-c458e85ed8c8",
		},
	})
}

func TestGrantAcessRight(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, grantAccessRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, grantAccessResponse)
	})

	c := client.ServiceClient()
	// Client c must have Microversion set; minimum supported microversion for Grant Access is 2.7
	c.Microversion = "2.7"

	grantaccOpts := shares.GrantAccessOpts{AccessTo: "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8", AccessType: "cert", AccessLevel: "rw"}
	s, err := shares.GrantAccess(c, shareID, grantaccOpts).ExtractAccess()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, &shares.AccessRight{
		AccessType:  "cert",
		AccessTo:    "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
		AccessLevel: "rw",
		State:       "new",
		ID:          "fc32500f-fa78-4f06-8caf-06ad7fb9726c",
	})
}

func TestDeleteAcess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, deleteAccessRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

	})
	c := client.ServiceClient()
	c.Microversion = "2.7"

	res := shares.DeleteAccessOpts{AccessID: "ea07152b-d08b-4f6b-8785-ce64dce52679"}
	s := shares.DeleteAccess(c, shareID, res)

	th.AssertNoErr(t, s.Err)
}

func TestGetExportLocationsSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/export_locations", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getExportLocationsResponse)
	})

	c := client.ServiceClient()
	s, err := shares.GetExportLocations(c, shareID).ExtractExportLocations()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, []shares.ExportLocation{
		{
			Path: "sfs-nas1.eu-de.otc.t-systems.com:/share-d41ee18b",
			ID:   "fab962ba-4b9a-475e-a380-8e856ed3f92d",
		},
	})
}

func TestListShare(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(shareEndpoint+"/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, listResponse)
	})

	//count := 0

	actual, err := shares.List(fake.ServiceClient(), shares.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract clusters: %v", err)
	}

	expected := []shares.Share{
		{
			Status:           "available",
			AvailabilityZone: "nova",
			ShareNetworkID:   "713df749-aac0-4a54-af52-10f6c991e80c",
			Name:             "my_test_share",
			ID:               "011d21e2-fbc3-4e4a-9993-9ea223f73264",
			Size:             1,
			ShareType:        "25747776-08e5-494f-ab40-a64b9d20d8f7",
			ProjectID:        "16e1ab15c35a457e9c2b2aa189f544e1",
			Description:      "My custom share London",
			Host:             "manila2@generic1#GENERIC1",
			IsPublic:         true,
			ShareProto:       "NFS",
			VolumeType:       "default",
			CreatedAt:        time.Date(2015, time.September, 18, 10, 25, 24, 0, time.UTC),
			Links: []map[string]string{
				{
					"href": "http://172.18.198.54:8786/v2/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
					"rel":  "self",
				},
				{
					"href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
					"rel":  "bookmark",
				},
			},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestExpand(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/"+"action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	})
	options := shares.ExpandOpts{OSExtend: shares.OSExtendOpts{NewSize: 3}}
	resp := shares.Expand(fake.ServiceClient(), shareID, options)
	th.AssertNoErr(t, resp.Err)

}

func TestShrink(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/"+"action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	})
	options := shares.ShrinkOpts{OSShrink: shares.OSShrinkOpts{NewSize: 2}}
	resp := shares.Shrink(fake.ServiceClient(), shareID, options)
	th.AssertNoErr(t, resp.Err)

}
