package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/networking/v2/common"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/peerings"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestListVpcPeerings(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/peerings", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "peerings": [
        {
            "status": "PENDING_ACCEPTANCE",
            "accept_vpc_info": {
              	"vpc_id": "c6efbdb7-dca4-4178-b3ec-692f125c1e25",
               	"tenant_id": "17fbda95add24720a4038ba4b1c705ed"
            	},
            "request_vpc_info": {
                "vpc_id": "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf",
                "tenant_id": "87a56a48977e42068f70ad3280c50f0e"
           		 },
            "name": "test_peering",
            "id": "22a3e5b1-1150-408e-99f7-5e25a391cead"
        },
        {
            "status": "ACTIVE",
            "accept_vpc_info": {
                "vpc_id": "93e94d8e-31a6-4c22-bdf7-8b23c7b67329",
                "tenant_id": "87a56a48977e42068f70ad3280c50f0e"
            	},
            "request_vpc_info": {
                "vpc_id": "b0d686e5-312c-4279-b69c-eedbc779ae69",
                "tenant_id": "bf74229f30c0421fae270386a43315ee"
            	},
            "name": "peering-7750-sunway",
            "id": "283aabd7-dab4-409d-96ff-6c878b9a0219"
        },
        {
            "status": "ACTIVE",
            "accept_vpc_info": {
                "vpc_id": "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf",
                "tenant_id": "87a56a48977e42068f70ad3280c50f0e"
            	},
            "request_vpc_info": {
                "vpc_id": "4117d38e-4c8f-4624-a505-bd96b97d024c",
                "tenant_id": "87a56a48977e42068f70ad3280c50f0e"
            	},
            "name": "test",
            "id": "71d64714-bd4e-44c4-917a-d8d1239e5292"
        }
        ]
 }
			`)
	})

	//count := 0
	actual, err := peerings.List(fake.ServiceClient(), peerings.ListOpts{})

	if err != nil {
		t.Errorf("Failed to extract vpc_peering_connections: %v", err)
	}

	expected := []peerings.Peering{
		{
			ID:             "22a3e5b1-1150-408e-99f7-5e25a391cead",
			Name:           "test_peering",
			Status:         "PENDING_ACCEPTANCE",
			RequestVpcInfo: peerings.VpcInfo{VpcId: "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf", TenantId: "87a56a48977e42068f70ad3280c50f0e"},
			AcceptVpcInfo:  peerings.VpcInfo{VpcId: "c6efbdb7-dca4-4178-b3ec-692f125c1e25", TenantId: "17fbda95add24720a4038ba4b1c705ed"},
		},
		{
			ID:             "283aabd7-dab4-409d-96ff-6c878b9a0219",
			Name:           "peering-7750-sunway",
			Status:         "ACTIVE",
			RequestVpcInfo: peerings.VpcInfo{VpcId: "b0d686e5-312c-4279-b69c-eedbc779ae69", TenantId: "bf74229f30c0421fae270386a43315ee"},
			AcceptVpcInfo:  peerings.VpcInfo{VpcId: "93e94d8e-31a6-4c22-bdf7-8b23c7b67329", TenantId: "87a56a48977e42068f70ad3280c50f0e"},
		},
		{
			ID:             "71d64714-bd4e-44c4-917a-d8d1239e5292",
			Name:           "test",
			Status:         "ACTIVE",
			RequestVpcInfo: peerings.VpcInfo{VpcId: "4117d38e-4c8f-4624-a505-bd96b97d024c", TenantId: "87a56a48977e42068f70ad3280c50f0e"},
			AcceptVpcInfo:  peerings.VpcInfo{VpcId: "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf", TenantId: "87a56a48977e42068f70ad3280c50f0e"},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestCreateVpcPeeringConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/peerings", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "peering": {
        "name": "test",

        "request_vpc_info": {
           "vpc_id": "4117d38e-4c8f-4624-a505-bd96b97d024c"
        },
        "accept_vpc_info": {
            "vpc_id": "c6efbdb7-dca4-4178-b3ec-692f125c1e25",
			"tenant_id": "17fbda95add24720a4038ba4b1c705ed"
        }
    }
}		`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, `
{
    "peering": {
        "status": "PENDING_ACCEPTANCE",
        "accept_vpc_info": {
             "vpc_id": "c6efbdb7-dca4-4178-b3ec-692f125c1e25",
             "tenant_id": "17fbda95add24720a4038ba4b1c705ed"
        },
        "request_vpc_info": {
            "vpc_id": "4117d38e-4c8f-4624-a505-bd96b97d024c",
            "tenant_id": "87a56a48977e42068f70ad3280c50f0e"
        },
        "name": "test",
        "id": "4e6ca99d-8344-4eb2-b2c9-b77368db3704"
    }
}	`)
	})

	options := peerings.CreateOpts{
		Name:           "test",
		RequestVpcInfo: peerings.VpcInfo{VpcId: "4117d38e-4c8f-4624-a505-bd96b97d024c"},
		AcceptVpcInfo:  peerings.VpcInfo{VpcId: "c6efbdb7-dca4-4178-b3ec-692f125c1e25", TenantId: "17fbda95add24720a4038ba4b1c705ed"},
	}
	n, err := peerings.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "test", n.Name)
	th.AssertEquals(t, "4e6ca99d-8344-4eb2-b2c9-b77368db3704", n.ID)

	th.AssertEquals(t, peerings.VpcInfo{VpcId: "4117d38e-4c8f-4624-a505-bd96b97d024c", TenantId: "87a56a48977e42068f70ad3280c50f0e"}, n.RequestVpcInfo)
	th.AssertEquals(t, peerings.VpcInfo{VpcId: "c6efbdb7-dca4-4178-b3ec-692f125c1e25", TenantId: "17fbda95add24720a4038ba4b1c705ed"}, n.AcceptVpcInfo)
	th.AssertEquals(t, "PENDING_ACCEPTANCE", n.Status)
}

func TestUpdateVpcPeeringConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/peerings/4e6ca99d-8344-4eb2-b2c9-b77368db3704", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "peering": {
        "name": "test2"
    }
}`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "peering": {
        "status": "PENDING_ACCEPTANCE",
        "accept_vpc_info": {
            "vpc_id": "c6efbdb7-dca4-4178-b3ec-692f125c1e25",
             "tenant_id": "17fbda95add24720a4038ba4b1c705ed"
        },
        "request_vpc_info": {
            "vpc_id": "4117d38e-4c8f-4624-a505-bd96b97d024c",
            "tenant_id": "87a56a48977e42068f70ad3280c50f0e"
        },
        "name": "test2",
        "id": "4e6ca99d-8344-4eb2-b2c9-b77368db3704"
    }
}
		`)
	})

	options := peerings.UpdateOpts{Name: "test2"}

	n, err := peerings.Update(fake.ServiceClient(), "4e6ca99d-8344-4eb2-b2c9-b77368db3704", options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "test2", n.Name)
	th.AssertEquals(t, "4e6ca99d-8344-4eb2-b2c9-b77368db3704", n.ID)
	th.AssertEquals(t, peerings.VpcInfo{VpcId: "4117d38e-4c8f-4624-a505-bd96b97d024c", TenantId: "87a56a48977e42068f70ad3280c50f0e"}, n.RequestVpcInfo)
	th.AssertEquals(t, peerings.VpcInfo{VpcId: "c6efbdb7-dca4-4178-b3ec-692f125c1e25", TenantId: "17fbda95add24720a4038ba4b1c705ed"}, n.AcceptVpcInfo)
	th.AssertEquals(t, "PENDING_ACCEPTANCE", n.Status)
}

func TestDeleteVpcPeeringConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/peerings/4e6ca99d-8344-4eb2-b2c9-b77368db3704", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := peerings.Delete(fake.ServiceClient(), "4e6ca99d-8344-4eb2-b2c9-b77368db3704")
	th.AssertNoErr(t, res.Err)
}

func TestAcceptVpcPeering(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/peerings/22a3e5b1-1150-408e-99f7-5e25a391cead/accept", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ` {
    "status": "ACTIVE",
    "name": "test_peering",
    "tenant_id": "17fbda95add24720a4038ba4b1c705ed",
    "request_vpc_info": {
        "vpc_id": "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf",
        "tenant_id": "87a56a48977e42068f70ad3280c50f0e"
    },
    "accept_vpc_info": {
        "vpc_id": "c6efbdb7-dca4-4178-b3ec-692f125c1e25",
        "tenant_id": "17fbda95add24720a4038ba4b1c705ed"
    },
    "id": "22a3e5b1-1150-408e-99f7-5e25a391cead"
     }
		`)
	})

	n, err := peerings.Accept(fake.ServiceClient(), "22a3e5b1-1150-408e-99f7-5e25a391cead").ExtractResult()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "22a3e5b1-1150-408e-99f7-5e25a391cead", n.ID)
	th.AssertEquals(t, "test_peering", n.Name)
	th.AssertEquals(t, "ACTIVE", n.Status)
	th.AssertDeepEquals(t, peerings.VpcInfo{VpcId: "c6efbdb7-dca4-4178-b3ec-692f125c1e25", TenantId: "17fbda95add24720a4038ba4b1c705ed"}, n.AcceptVpcInfo)
	th.AssertDeepEquals(t, peerings.VpcInfo{VpcId: "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf", TenantId: "87a56a48977e42068f70ad3280c50f0e"}, n.RequestVpcInfo)

}

func TestRejectVpcPeering(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpc/peerings/22a3e5b1-1150-408e-99f7-5e25a391cead/reject", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ` {
    "status": "ACTIVE",
    "name": "test_peering",
    "tenant_id": "17fbda95add24720a4038ba4b1c705ed",
    "request_vpc_info": {
        "vpc_id": "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf",
        "tenant_id": "87a56a48977e42068f70ad3280c50f0e"
    },
    "accept_vpc_info": {
        "vpc_id": "c6efbdb7-dca4-4178-b3ec-692f125c1e25",
        "tenant_id": "17fbda95add24720a4038ba4b1c705ed"
    },
    "id": "22a3e5b1-1150-408e-99f7-5e25a391cead"
     }
		`)
	})

	n, err := peerings.Reject(fake.ServiceClient(), "22a3e5b1-1150-408e-99f7-5e25a391cead").ExtractResult()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "22a3e5b1-1150-408e-99f7-5e25a391cead", n.ID)
	th.AssertEquals(t, "test_peering", n.Name)
	th.AssertEquals(t, "ACTIVE", n.Status)
	th.AssertDeepEquals(t, peerings.VpcInfo{VpcId: "c6efbdb7-dca4-4178-b3ec-692f125c1e25", TenantId: "17fbda95add24720a4038ba4b1c705ed"}, n.AcceptVpcInfo)
	th.AssertDeepEquals(t, peerings.VpcInfo{VpcId: "3127e30b-5f8e-42d1-a3cc-fdadf412c5bf", TenantId: "87a56a48977e42068f70ad3280c50f0e"}, n.RequestVpcInfo)

}
