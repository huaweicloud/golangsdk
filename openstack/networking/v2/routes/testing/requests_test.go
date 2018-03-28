package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/networking/v2/common"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/routes"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestListRoutes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/routes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "routes": [
        {
            "destination": "172.31.8.192/26",
            "tenant_id": "87a56a48977e42068f70ad3280c50f0e",
            "nexthop": "283aabd7-dab4-409d-96ff-6c878b9a0219",
            "vpc_id": "93e94d8e-31a6-4c22-bdf7-8b23c7b67329",
            "type": "peering",
            "id": "4ae95cd4-292d-4a27-b3de-1be835eb32e1"
        },
        {
            "destination": "172.31.8.128/26",
            "tenant_id": "87a56a48977e42068f70ad3280c50f0e",
            "nexthop": "283aabd7-dab4-409d-96ff-6c878b9a0219",
            "vpc_id": "93e94d8e-31a6-4c22-bdf7-8b23c7b67329",
            "type": "peering",
            "id": "804d5d09-cee2-418d-8d1b-be29b8e8e9e8"
        },
        {
            "destination": "172.31.8.112/28",
            "tenant_id": "87a56a48977e42068f70ad3280c50f0e",
            "nexthop": "283aabd7-dab4-409d-96ff-6c878b9a0219",
            "vpc_id": "93e94d8e-31a6-4c22-bdf7-8b23c7b67329",
            "type": "peering",
            "id": "9f54e4ac-e052-4198-bb73-51b22ad41035"
        }
    ]
}
			`)
	})

	pages, err := routes.List(fake.ServiceClient(), routes.ListOpts{}).AllPages()
	if err != nil {
		t.Errorf("Failed to get routes: %v", err)
	}

	actual, err := routes.ExtractRoutes(pages)
	if err != nil {
		t.Errorf("Failed to extract routes: %v", err)
	}

	expected := []routes.Route{
		{
			Type:        "peering",
			NextHop:     "283aabd7-dab4-409d-96ff-6c878b9a0219",
			Destination: "172.31.8.192/26",
			VPC_ID:      "93e94d8e-31a6-4c22-bdf7-8b23c7b67329",
			Tenant_Id:   "87a56a48977e42068f70ad3280c50f0e",
			RouteID:     "4ae95cd4-292d-4a27-b3de-1be835eb32e1",
		},
		{
			Type:        "peering",
			NextHop:     "283aabd7-dab4-409d-96ff-6c878b9a0219",
			Destination: "172.31.8.128/26",
			VPC_ID:      "93e94d8e-31a6-4c22-bdf7-8b23c7b67329",
			Tenant_Id:   "87a56a48977e42068f70ad3280c50f0e",
			RouteID:     "804d5d09-cee2-418d-8d1b-be29b8e8e9e8",
		},
		{
			Type:        "peering",
			NextHop:     "283aabd7-dab4-409d-96ff-6c878b9a0219",
			Destination: "172.31.8.112/28",
			VPC_ID:      "93e94d8e-31a6-4c22-bdf7-8b23c7b67329",
			Tenant_Id:   "87a56a48977e42068f70ad3280c50f0e",
			RouteID:     "9f54e4ac-e052-4198-bb73-51b22ad41035",
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestGetRoutes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/routes/39a07dcb-f30e-41c1-97ac-182c8f0d43c1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "route": {
        "destination": "192.168.0.0/16",
        "tenant_id": "87a56a48977e42068f70ad3280c50f0e",
        "nexthop": "d2dea4ba-e988-4e9c-8162-652e74b2560c",
        "vpc_id": "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf",
        "type": "peering",
        "id": "39a07dcb-f30e-41c1-97ac-182c8f0d43c1"
    }
}
		`)
	})

	n, err := routes.Get(fake.ServiceClient(), "39a07dcb-f30e-41c1-97ac-182c8f0d43c1").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "39a07dcb-f30e-41c1-97ac-182c8f0d43c1", n.RouteID)
	th.AssertEquals(t, "192.168.0.0/16", n.Destination)
	th.AssertEquals(t, "87a56a48977e42068f70ad3280c50f0e", n.Tenant_Id)
	th.AssertEquals(t, "d2dea4ba-e988-4e9c-8162-652e74b2560c", n.NextHop)
	th.AssertEquals(t, "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf", n.VPC_ID)
	th.AssertEquals(t, "peering", n.Type)
}

func TestCreateRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/routes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{ 
    "route": { 
        "type": "peering",  
        "nexthop": "d2dea4ba-e988-4e9c-8162-652e74b2560c",  
        "destination": "192.168.0.0/16",  
        "vpc_id": "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf"
    }
}
			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, `
{
    "route": {
        "destination": "192.168.0.0/16",
        "tenant_id": "87a56a48977e42068f70ad3280c50f0e",
        "nexthop": "d2dea4ba-e988-4e9c-8162-652e74b2560c",
        "vpc_id": "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf",
        "type": "peering",
        "id": "39a07dcb-f30e-41c1-97ac-182c8f0d43c1"
    }
}		`)
	})

	options := routes.CreateOpts{
		Type:        "peering",
		NextHop:     "d2dea4ba-e988-4e9c-8162-652e74b2560c",
		Destination: "192.168.0.0/16",
		VPC_ID:      "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf",
	}
	n, err := routes.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "peering", n.Type)
	th.AssertEquals(t, "d2dea4ba-e988-4e9c-8162-652e74b2560c", n.NextHop)
	th.AssertEquals(t, "192.168.0.0/16", n.Destination)
	th.AssertEquals(t, "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf", n.VPC_ID)
	th.AssertEquals(t, "39a07dcb-f30e-41c1-97ac-182c8f0d43c1", n.RouteID)
}

func TestDeleteRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/routes/39a07dcb-f30e-41c1-97ac-182c8f0d43c1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := routes.Delete(fake.ServiceClient(), "39a07dcb-f30e-41c1-97ac-182c8f0d43c1")
	th.AssertNoErr(t, res.Err)
}
