package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/bms/v2/nics"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestListNIC(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/2bff7a8a-3934-4f79-b1d6-53dc5540f00e/os-interface", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "interfaceAttachments": [
        {
            "port_state": "ACTIVE",
            "fixed_ips": [
                {
                    "subnet_id": "518e34f2-16d4-4242-9378-b7eea505ab9c",
                    "ip_address": "192.168.0.80"
                }
            ],
            "port_id": "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f",
            "net_id": "31e04eb7-7ccc-4116-9f43-a51eb29fa348",
            "mac_addr": "fa:16:3e:00:1a:9a"
        },
        {
            "port_state": "ACTIVE",
            "fixed_ips": [
                {
                    "subnet_id": "f3ef8cb3-9954-4434-a558-70b623b3c69b",
                    "ip_address": "192.168.1.206"
                }
            ],
            "port_id": "a5fff9e7-65e1-4e46-95da-263d66ff4a7a",
            "net_id": "68cf7e29-b770-4951-be66-9b3f16297732",
            "mac_addr": "fa:16:3e:83:dc:08"
        }	
    ]
}
			`)
	})

	//count := 0

	actual, err := nics.List(fake.ServiceClient(), "2bff7a8a-3934-4f79-b1d6-53dc5540f00e", nics.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract nics: %v", err)
	}

	expected := []nics.Nic{
		{
			ID:         "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f",
			Status:     "ACTIVE",
			NetworkID:  "31e04eb7-7ccc-4116-9f43-a51eb29fa348",
			MACAddress: "fa:16:3e:00:1a:9a",
			FixedIP:    []nics.FixedIP{{SubnetID: "518e34f2-16d4-4242-9378-b7eea505ab9c", IPAddress: "192.168.0.80"}},
		},
		{
			ID:         "a5fff9e7-65e1-4e46-95da-263d66ff4a7a",
			Status:     "ACTIVE",
			NetworkID:  "68cf7e29-b770-4951-be66-9b3f16297732",
			MACAddress: "fa:16:3e:83:dc:08",
			FixedIP:    []nics.FixedIP{{SubnetID: "f3ef8cb3-9954-4434-a558-70b623b3c69b", IPAddress: "192.168.1.206"}},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestGetNIC(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/2bff7a8a-3934-4f79-b1d6-53dc5540f00e/os-interface/1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "interfaceAttachment": {
        "port_state": "ACTIVE",
        "fixed_ips": [
            {
                "subnet_id": "518e34f2-16d4-4242-9378-b7eea505ab9c",
                "ip_address": "192.168.0.80"
            }
        ],
        "port_id": "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f",
        "net_id": "31e04eb7-7ccc-4116-9f43-a51eb29fa348",
        "mac_addr": "fa:16:3e:00:1a:9a"
    }
}
		`)
	})

	n, err := nics.Get(fake.ServiceClient(), "2bff7a8a-3934-4f79-b1d6-53dc5540f00e", "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "1d3bf3ae-bc4a-4890-86f8-8c31a6eb764f", n.ID)
	th.AssertEquals(t, "ACTIVE", n.Status)
	th.AssertEquals(t, "fa:16:3e:00:1a:9a", n.MACAddress)
	th.AssertEquals(t, "31e04eb7-7ccc-4116-9f43-a51eb29fa348", n.NetworkID)
	th.AssertDeepEquals(t, []nics.FixedIP{{SubnetID: "518e34f2-16d4-4242-9378-b7eea505ab9c", IPAddress: "192.168.0.80"}}, n.FixedIP)

}
