package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/networking/v1/common"
	"github.com/huaweicloud/golangsdk/openstack/networking/v1/vpcs"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestListVpc(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "vpcs": [
        {
            "id": "14ece7d0-a8d4-4317-982a-041e4f10f442",
            "name": "vpc-elb-l00379969",
            "cidr": "192.168.0.0/16",
            "status": "OK",
            "routes": [],
            "enable_shared_snat": false
        },
        {
            "id": "1e5618c3-89f0-4f58-a14e-33536074ec88",
            "name": "vpc-ops",
            "cidr": "192.168.0.0/16",
            "status": "OK",
            "routes": [],
            "enable_shared_snat": false
        },
        {
            "id": "2140264c-d313-4363-9874-9a5e18aeb516",
            "name": "test",
            "cidr": "192.168.0.0/16",
            "status": "OK",
            "routes": [],
            "enable_shared_snat": false
        }
    ]
}
			`)
	})

	//count := 0

	actual, err := vpcs.List(fake.ServiceClient(), vpcs.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract vpcs: %v", err)
	}

	expected := []vpcs.Vpc{
		{
			Status:           "OK",
			CIDR:             "192.168.0.0/16",
			EnableSharedSnat: false,
			Name:             "vpc-elb-l00379969",
			ID:               "14ece7d0-a8d4-4317-982a-041e4f10f442",
			Routes:           []vpcs.Route{},
		},
		{
			Status:           "OK",
			CIDR:             "192.168.0.0/16",
			EnableSharedSnat: false,
			Name:             "vpc-ops",
			ID:               "1e5618c3-89f0-4f58-a14e-33536074ec88",
			Routes:           []vpcs.Route{},
		},
		{
			Status:           "OK",
			CIDR:             "192.168.0.0/16",
			EnableSharedSnat: false,
			Name:             "test",
			ID:               "2140264c-d313-4363-9874-9a5e18aeb516",
			Routes:           []vpcs.Route{},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestGetVpc(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/vpcs/abda1f6e-ae7c-4ff5-8d06-53425dc11f34", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "vpc": {
        "id": "abda1f6e-ae7c-4ff5-8d06-53425dc11f34",
        "name": "terraform-provider-test-l90006937",
        "cidr": "192.168.0.0/16",
        "status": "OK",
        "routes": [
            {
                "destination": "0.0.0.0/0",
                "nexthop": "192.168.0.5"
            }
        ],
        "enable_shared_snat": false
    }
}
		`)
	})

	n, err := vpcs.Get(fake.ServiceClient(), "abda1f6e-ae7c-4ff5-8d06-53425dc11f34").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "abda1f6e-ae7c-4ff5-8d06-53425dc11f34", n.ID)
	th.AssertEquals(t, "terraform-provider-test-l90006937", n.Name)
	th.AssertEquals(t, "192.168.0.0/16", n.CIDR)
	th.AssertEquals(t, "OK", n.Status)
	th.AssertDeepEquals(t, []vpcs.Route{{DestinationCIDR: "0.0.0.0/0", NextHop: "192.168.0.5"}}, n.Routes)
	th.AssertEquals(t, false, n.EnableSharedSnat)

}

func TestCreateVpc(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
 "vpc":
     {
     "name": "terraform-provider-vpctestcreate",
     "cidr": "192.168.0.0/16"
     }
}
			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "vpc": {
        "id": "97e01fc2-e39e-4cfc-abf6-1d0886d120af",
        "name": "terraform-provider-vpctestcreate",
        "cidr": "192.168.0.0/16",
        "status": "CREATING"
    }
}		`)
	})

	options := vpcs.CreateOpts{
		Name: "terraform-provider-vpctestcreate",
		CIDR: "192.168.0.0/16",
	}
	n, err := vpcs.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "terraform-provider-vpctestcreate", n.Name)
	th.AssertEquals(t, "97e01fc2-e39e-4cfc-abf6-1d0886d120af", n.ID)
	th.AssertEquals(t, "192.168.0.0/16", n.CIDR)
	th.AssertEquals(t, "CREATING", n.Status)
	th.AssertEquals(t, false, n.EnableSharedSnat)
}

func TestUpdateVpc(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/vpcs/97e01fc2-e39e-4cfc-abf6-1d0886d120af", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
"vpc":
    {
    "name": "terraform-provider-new-name",
    "cidr": "192.168.0.0/16"

    }
}
			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "vpc": {
        "id": "97e01fc2-e39e-4cfc-abf6-1d0886d120af",
        "name": "terraform-provider-new-name",
        "cidr": "192.168.0.0/16",
        "status": "OK",
        "routes": [
            {
                "destination": "0.0.0.0/4",
                "nexthop": "192.168.0.4"
            }
        ],
        "enable_shared_snat": false
    }
}
		`)
	})

	options := vpcs.UpdateOpts{Name: "terraform-provider-new-name", CIDR: "192.168.0.0/16"}

	n, err := vpcs.Update(fake.ServiceClient(), "97e01fc2-e39e-4cfc-abf6-1d0886d120af", options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "terraform-provider-new-name", n.Name)
	th.AssertEquals(t, "97e01fc2-e39e-4cfc-abf6-1d0886d120af", n.ID)
	th.AssertEquals(t, "192.168.0.0/16", n.CIDR)
	th.AssertEquals(t, "OK", n.Status)
	th.AssertDeepEquals(t, []vpcs.Route{{DestinationCIDR: "0.0.0.0/4", NextHop: "192.168.0.4"}}, n.Routes)
	th.AssertEquals(t, false, n.EnableSharedSnat)
}

func TestDeleteVpc(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/vpcs/abda1f6e-ae7c-4ff5-8d06-53425dc11f34", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := vpcs.Delete(fake.ServiceClient(), "abda1f6e-ae7c-4ff5-8d06-53425dc11f34")
	th.AssertNoErr(t, res.Err)
}
