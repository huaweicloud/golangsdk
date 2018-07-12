package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/deh/v1/common"
	"github.com/huaweicloud/golangsdk/openstack/deh/v1/hosts"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(dehEndpoint+"/"+HostID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	s, err := hosts.Get(client.ServiceClient(), HostID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "66156a61-27c2-4169-936b-910dd9c73da3", s.ID)
	th.AssertEquals(t, "test-aj2", s.Name)
	th.AssertEquals(t, "eu-de-02", s.Az)
	th.AssertEquals(t, "available", s.State)
	th.AssertDeepEquals(t, hosts.HostPropertiesOpts{
		HostTypeName: "High performance",
		HostType:     "h1",
		Vcpus:        36,
		Memory:       270336,
		Cores:        12,
		Sockets:      2,
	}, s.HostProperties)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(dehEndpoint, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, listResponse)
	})

	//count := 0

	host, err := hosts.List(client.ServiceClient(), hosts.ListOpts{}).AllPages()

	actual, err := hosts.ExtractHosts(host)

	if err != nil {
		t.Errorf("Failed to extract hosts: %v", err)
	}

	expected := []hosts.Host{
		{
			Az:              "eu-de-01",
			Name:            "c2c-deh-test",
			AvailableMemory: 262144,
			AvailableVcpus:  70,
			ID:              "671611d2-b45c-4648-9e78-06eb24522291",
			State:           "available",
			InstanceTotal:   2,
			AutoPlacement:   "off",
			TenantId:        "17fbda95add24720a4038ba4b1c705ed",
			HostProperties: hosts.HostPropertiesOpts{
				HostType:     "general",
				Vcpus:        72,
				Memory:       270336,
				Cores:        12,
				Sockets:      2,
				HostTypeName: "General computing",
			},
			InstanceUuids: []string{"3de1ce75-2550-4a46-a689-dd33ca2b62d6",
				"885dc71d-905d-48b5-bae7-db66801dc175"},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestListServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(dehEndpoint+"/"+HostID+"/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, listserverResponse)
	})

	actual, err := hosts.ListServer(client.ServiceClient(), HostID, hosts.ListServerOpts{})
	th.AssertNoErr(t, err)

	expected := []hosts.Server{
		{
			Status: "ACTIVE",
			Addresses: map[string]interface{}{
				"0b98c646-617f-4d90-9ca5-385f0cd73ea7": []interface{}{
					map[string]interface{}{
						"version": float64(4),
						"addr":    "192.168.3.133",
					},
				},
			},
			Flavor: map[string]interface{}{
				"id": "normal1",
			},
			ID:       "3de1ce75-2550-4a46-a689-dd33ca2b62d6",
			UserID:   "6d78fa8550ae45d6932a1fadfb1fa552",
			Name:     "c2c-ecs-test-2",
			TenantID: "17fbda95add24720a4038ba4b1c705ed",
			Metadata: map[string]string{
				"metering.image_id":         "c0ea3ff1-432e-4650-8a1b-372a80b2d2be",
				"metering.imagetype":        "gold",
				"metering.resourcespeccode": "deh.linux",
				"metering.cloudServiceType": "sys.service.type.ec2",
				"image_name":                "Standard_CentOS_7_latest",
				"metering.resourcetype":     "1",
				"os_bit":                    "64",
				"vpc_id":                    "0b98c646-617f-4d90-9ca5-385f0cd73ea7",
				"os_type":                   "Linux",
				"charging_mode":             "0",
			},
		},
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestAllocateDeH(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc(dehEndpoint, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, allocateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, allocateResponse)
	})

	c := client.ServiceClient()
	allocateOpts := hosts.AllocateOpts{Name: "Test-1",
		Az:            "eu-de-02",
		HostType:      "h1",
		AutoPlacement: "off",
		Quantity:      2}
	s, err := hosts.Allocate(c, allocateOpts).ExtractHost()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, &hosts.AllocatedHosts{
		AllocatedHostIds: []string{"fb4733fd-70a3-44e1-a1cb-0311f028d7e5",
			"7408f985-047d-4313-b3c8-8e12bef01d12"},
	})
}

func TestUpdateDeH(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc(dehEndpoint+"/"+HostID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, updateRequest)
		w.WriteHeader(http.StatusNoContent)
	})

	c := client.ServiceClient()
	updateOpts := hosts.UpdateOpts{Name: "Test-2",
		AutoPlacement: "off",
	}
	s := hosts.Update(c, HostID, updateOpts)
	th.AssertNoErr(t, s.Err)
}

func TestDeleteDeH(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(dehEndpoint+"/"+HostID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})

	result := hosts.Delete(client.ServiceClient(), HostID)
	th.AssertNoErr(t, result.Err)
}
