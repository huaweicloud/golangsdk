package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/bms/v2/flavors"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestListFlavor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/flavors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "flavors": [
        {
            "name": "physical.h2.large",
            "links": [
                {
                    "href": "https://compute.region.eu-de.otc-tsi.de/v2/91d687759aed45d28b5f6084bc2fa8ad/flavors/physical.h2.large",
                    "rel": "self"
                },
                {
                    "href": "https://compute.region.eu-de.otc-tsi.de/91d687759aed45d28b5f6084bc2fa8ad/flavors/physical.h2.large",
                    "rel": "bookmark"
                }
            ],
            "ram": 196608,
            "OS-FLV-DISABLED:disabled": false,
            "vcpus": 36,
            "os-flavor-access:is_public": true,
            "rxtx_factor": 1,
            "OS-FLV-EXT-DATA:ephemeral": 0,
            "disk": 0,
            "id": "physical.h2.large"
        },
        {
            "name": "physical.m2.medium",
            "links": [
                {
                    "href": "https://compute.region.eu-de.otc-tsi.de/v2/91d687759aed45d28b5f6084bc2fa8ad/flavors/physical.m2.medium",
                    "rel": "self"
                },
                {
                    "href": "https://compute.region.eu-de.otc-tsi.de/91d687759aed45d28b5f6084bc2fa8ad/flavors/physical.m2.medium",
                    "rel": "bookmark"
                }
            ],
            "ram": 2097152,
            "OS-FLV-DISABLED:disabled": false,
            "vcpus": 192,
            "os-flavor-access:is_public": true,
            "rxtx_factor": 1,
            "OS-FLV-EXT-DATA:ephemeral": 0,
            "disk": 17000,
            "id": "physical.m2.medium"
        }
    ]
}
			`)
	})

	//count := 0

	actual, err := flavors.List(fake.ServiceClient(), flavors.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract flavors: %v", err)
	}

	expected := []flavors.Flavor{
		{
			ID:         "physical.h2.large",
			RAM:        196608,
			Disabled:   false,
			Name:       "physical.h2.large",
			VCPUs:      36,
			IsPublic:   true,
			RxTxFactor: 1,
			Ephemeral:  0,
			Disk:       0,
			Links: []golangsdk.Link{
				{
					Href: "https://compute.region.eu-de.otc-tsi.de/v2/91d687759aed45d28b5f6084bc2fa8ad/flavors/physical.h2.large",
					Rel:  "self",
				},
				{
					Href: "https://compute.region.eu-de.otc-tsi.de/91d687759aed45d28b5f6084bc2fa8ad/flavors/physical.h2.large",
					Rel:  "bookmark",
				},
			},
		},
		{
			ID:         "physical.m2.medium",
			RAM:        2097152,
			Disabled:   false,
			Name:       "physical.m2.medium",
			VCPUs:      192,
			IsPublic:   true,
			RxTxFactor: 1,
			Ephemeral:  0,
			Disk:       17000,
			Links: []golangsdk.Link{
				{
					Href: "https://compute.region.eu-de.otc-tsi.de/v2/91d687759aed45d28b5f6084bc2fa8ad/flavors/physical.m2.medium",
					Rel:  "self",
				},
				{
					Href: "https://compute.region.eu-de.otc-tsi.de/91d687759aed45d28b5f6084bc2fa8ad/flavors/physical.m2.medium",
					Rel:  "bookmark",
				},
			},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}
