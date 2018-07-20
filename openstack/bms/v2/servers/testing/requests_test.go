package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/bms/v2/servers"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestListServers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "servers": [
{
        "tenant_id": "17fbda95add24720a4038ba4b1c705ed",  
        "id": "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f",
        "user_id": "aec83183a5b54cf2bc28812d1dc5509e",
        "name": "BMS-1",  	
"key_name": "KeyPair-click2cloud",
        "flavor": {
            "links": [
                {
                    "rel": "bookmark",
                    "href": "https://ecs.eu-de.otc.t-systems.com:443/17fbda95add24720a4038ba4b1c705ed/flavors/physical.o2.medium"
                }
            ],
            "id": "physical.o2.medium"
        },       
        "status": "ACTIVE"

    },
{
        "tenant_id": "17fbda95add24720a4038ba4b1c705ed",  
        "id": "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764r",
        "user_id": "aec83183a5b54cf2bc28812d1dc5509e",
        "name": "BMS-2",    		
"key_name": "KeyPair-click2cloud",
        "flavor": {
            "links": [
                {
                    "rel": "bookmark",
                    "href": "https://ecs.eu-de.otc.t-systems.com:443/17fbda95add24720a4038ba4b1c705ed/flavors/physical.o2.medium"
                }
            ],
            "id": "physical.o2.medium"
        },       
        "status": "ACTIVE"
    }
]
}
			`)
	})

	serverOpts := servers.ListOpts{}
	actual, err := servers.List(fake.ServiceClient(), serverOpts)
	if err != nil {
		t.Errorf("Failed to extract server list: %v", err)
	}

	expected := []servers.Server{
		{
			ID:       "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f",
			Status:   "ACTIVE",
			UserID:   "aec83183a5b54cf2bc28812d1dc5509e",
			Name:     "BMS-1",
			TenantID: "17fbda95add24720a4038ba4b1c705ed",
			KeyName:  "KeyPair-click2cloud",
			Flavor: servers.Flavor{ID: "physical.o2.medium", Links: []servers.Links{{Rel: "bookmark",
				Href: "https://ecs.eu-de.otc.t-systems.com:443/17fbda95add24720a4038ba4b1c705ed/flavors/physical.o2.medium"}}},
		},
		{
			ID:       "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764r",
			Status:   "ACTIVE",
			UserID:   "aec83183a5b54cf2bc28812d1dc5509e",
			Name:     "BMS-2",
			TenantID: "17fbda95add24720a4038ba4b1c705ed",
			KeyName:  "KeyPair-click2cloud",
			Flavor: servers.Flavor{ID: "physical.o2.medium", Links: []servers.Links{{Rel: "bookmark",
				Href: "https://ecs.eu-de.otc.t-systems.com:443/17fbda95add24720a4038ba4b1c705ed/flavors/physical.o2.medium"}}},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestGetServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/2bff7a8a-3934-4f79-b1d6-53dc5540f00e", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "server": {
        "tenant_id": "17fbda95add24720a4038ba4b1c705ed",  
        "id": "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f",
        "user_id": "aec83183a5b54cf2bc28812d1dc5509e",
        "name": "BMS-1",        
        "status": "ACTIVE"
    }
}
		`)
	})

	n, err := servers.Get(fake.ServiceClient(), "2bff7a8a-3934-4f79-b1d6-53dc5540f00e").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f", n.ID)
	th.AssertEquals(t, "ACTIVE", n.Status)
	th.AssertEquals(t, "BMS-1", n.Name)
	th.AssertEquals(t, "aec83183a5b54cf2bc28812d1dc5509e", n.UserID)
	th.AssertEquals(t, "17fbda95add24720a4038ba4b1c705ed", n.TenantID)

}
