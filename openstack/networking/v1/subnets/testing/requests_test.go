package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/networking/v1/common"
	"github.com/huaweicloud/golangsdk/openstack/networking/v1/subnets"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestListSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "subnets": [
        {
            "id": "0345a6ef-9404-487b-87c8-212557a1160d",
            "name": "openlab-subnet",
            "cidr": "192.168.200.0/24",
            "status": "ACTIVE",
            "vpc_id": "58c24204-170e-4ff0-9b42-c53cdea9239a",
            "gateway_ip": "192.168.200.1",
            "dhcp_enable": true,
			"primary_dns": "114.114.114.114",
		    "secondary_dns": "114.114.115.115",
		    "dnsList": [
			  "114.114.114.114",
			  "114.114.115.115"
		    ],
            "neutron_subnet_id": "3d543273-31c3-41f8-b887-ed8c2c837578"
        },
        {
            "id": "134ca339-24dc-44f5-ae6a-cf0404216ed2",
            "name": "openlab-subnet",
            "cidr": "192.168.200.0/24",
            "status": "ACTIVE",
            "vpc_id": "58c24204-170e-4ff0-9b42-c53cdea9239a",
            "gateway_ip": "192.168.200.1",
            "dhcp_enable": true,
			"primary_dns": "114.114.114.114",
		    "secondary_dns": "114.114.115.115",
		    "dnsList": [
			  "114.114.114.114",
			  "114.114.115.115"
		    ],
            "neutron_subnet_id": "3d543273-31c3-41f8-b887-ed8c2c837578"
        }
    ]
}

		`)
	})

	actual, err := subnets.List(fake.ServiceClient(), subnets.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract subnets: %v", err)
	}

	expected := []subnets.Subnet{
		{
			Status:     "ACTIVE",
			CIDR:       "192.168.200.0/24",
			EnableDHCP: true,
			Name:       "openlab-subnet",
			//DnsList:          []string{},
			ID:            "0345a6ef-9404-487b-87c8-212557a1160d",
			GatewayIP:     "192.168.200.1",
			VPC_ID:        "58c24204-170e-4ff0-9b42-c53cdea9239a",
			PRIMARY_DNS:   "114.114.114.114",
			SECONDARY_DNS: "114.114.115.115",
			DnsList:       []string{"114.114.114.114", "114.114.115.115"},
			SubnetId:      "3d543273-31c3-41f8-b887-ed8c2c837578",
		},
		{
			Status:     "ACTIVE",
			CIDR:       "192.168.200.0/24",
			EnableDHCP: true,
			Name:       "openlab-subnet",
			//DnsList:          []string{},
			ID:            "134ca339-24dc-44f5-ae6a-cf0404216ed2",
			GatewayIP:     "192.168.200.1",
			VPC_ID:        "58c24204-170e-4ff0-9b42-c53cdea9239a",
			PRIMARY_DNS:   "114.114.114.114",
			SECONDARY_DNS: "114.114.115.115",
			DnsList:       []string{"114.114.114.114", "114.114.115.115"},
			SubnetId:      "3d543273-31c3-41f8-b887-ed8c2c837578",
		},
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestGetSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/subnets/aab2f0ef-b08b-4f34-9e1a-9f1d8da1afcb", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "subnet": {
        "id": "aab2f0ef-b08b-4f34-9e1a-9f1d8da1afcb",
        "name": "subnet-mgmt",
        "cidr": "10.0.0.0/24",
        "dnsList": [
            "100.125.4.25",
            "8.8.8.8"
        ],
        "status": "ACTIVE",
        "vpc_id": "d4f2c817-d5df-4a66-994a-6571312b470e",
        "gateway_ip": "10.0.0.1",
        "dhcp_enable": true,
        "primary_dns": "100.125.4.25",
        "secondary_dns": "8.8.8.8",
        "neutron_subnet_id": "3d543273-31c3-41f8-b887-ed8c2c837578"
    }
}
		`)
	})

	n, err := subnets.Get(fake.ServiceClient(), "aab2f0ef-b08b-4f34-9e1a-9f1d8da1afcb").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "aab2f0ef-b08b-4f34-9e1a-9f1d8da1afcb", n.ID)
	th.AssertEquals(t, "subnet-mgmt", n.Name)
	th.AssertEquals(t, "10.0.0.0/24", n.CIDR)
	th.AssertEquals(t, "ACTIVE", n.Status)
	th.AssertEquals(t, "d4f2c817-d5df-4a66-994a-6571312b470e", n.VPC_ID)
	th.AssertEquals(t, "3d543273-31c3-41f8-b887-ed8c2c837578", n.SubnetId)
	th.AssertEquals(t, "10.0.0.1", n.GatewayIP)
	th.AssertEquals(t, "100.125.4.25", n.PRIMARY_DNS)
	th.AssertEquals(t, "8.8.8.8", n.SECONDARY_DNS)
	th.AssertEquals(t, true, n.EnableDHCP)
	th.AssertEquals(t, "100.125.4.25", n.DnsList[0])
	th.AssertEquals(t, "8.8.8.8", n.DnsList[1])

}

func TestCreateSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
  "subnet":
         {
          "name": "test_subnets",
          "cidr": "192.168.0.0/16",
          "gateway_ip": "192.168.0.1",
		  "dhcp_enable": true,
          "primary_dns": "8.8.8.8",
          "secondary_dns": "8.8.4.4",
          "availability_zone":"eu-de-02",
          "vpc_id":"3b9740a0-b44d-48f0-84ee-42eb166e54f7",
		  "dnsList": [
             "8.8.8.8",
            "8.8.4.4"
          ]
          }
}
			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "subnet": {
        "id": "6b0cf733-f496-4159-9df1-d74c3584a9f7",
        "name": "test_subnets",
        "cidr": "192.168.0.0/16",
        "dnsList": [
            "8.8.8.8",
            "8.8.4.4"
        ],
        "status": "UNKNOWN",
        "vpc_id": "3b9740a0-b44d-48f0-84ee-42eb166e54f7",
        "gateway_ip": "192.168.0.1",
        "dhcp_enable": true,
        "primary_dns": "8.8.8.8",
        "secondary_dns": "8.8.4.4",
        "availability_zone": "eu-de-02",
        "neutron_subnet_id": "3d543273-31c3-41f8-b887-ed8c2c837578"
    }
}	`)
	})

	options := subnets.CreateOpts{
		Name:             "test_subnets",
		CIDR:             "192.168.0.0/16",
		GatewayIP:        "192.168.0.1",
		PRIMARY_DNS:      "8.8.8.8",
		SECONDARY_DNS:    "8.8.4.4",
		AvailabilityZone: "eu-de-02",
		VPC_ID:           "3b9740a0-b44d-48f0-84ee-42eb166e54f7",
		DnsList:          []string{"8.8.8.8", "8.8.4.4"},
		EnableDHCP:       true,
	}
	n, err := subnets.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "test_subnets", n.Name)
	th.AssertEquals(t, "192.168.0.1", n.GatewayIP)
	th.AssertEquals(t, "192.168.0.0/16", n.CIDR)
	th.AssertEquals(t, true, n.EnableDHCP)
	th.AssertEquals(t, "8.8.8.8", n.PRIMARY_DNS)
	th.AssertEquals(t, "8.8.4.4", n.SECONDARY_DNS)
	th.AssertEquals(t, "eu-de-02", n.AvailabilityZone)
	th.AssertEquals(t, "6b0cf733-f496-4159-9df1-d74c3584a9f7", n.ID)
	th.AssertEquals(t, "UNKNOWN", n.Status)
	th.AssertEquals(t, "3b9740a0-b44d-48f0-84ee-42eb166e54f7", n.VPC_ID)
	th.AssertEquals(t, "3d543273-31c3-41f8-b887-ed8c2c837578", n.SubnetId)
	th.AssertEquals(t, "8.8.8.8", n.DnsList[0])
	th.AssertEquals(t, "8.8.4.4", n.DnsList[1])

}

func TestUpdateSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/vpcs/8f794f06-2275-4d82-9f5a-6d68fbe21a75/subnets/83e3bddc-b9ed-4614-a0dc-8a997095a86c", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
"subnet":
    {
    	"name": "testsubnet",
		"dhcp_enable": false
    }
}
`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "subnet": {
        "id": "83e3bddc-b9ed-4614-a0dc-8a997095a86c",
		"name": "testsubnet",
        "status": "ACTIVE"
    }
}
		`)
	})

	options := subnets.UpdateOpts{Name: "testsubnet"}

	n, err := subnets.Update(fake.ServiceClient(), "8f794f06-2275-4d82-9f5a-6d68fbe21a75", "83e3bddc-b9ed-4614-a0dc-8a997095a86c", options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "testsubnet", n.Name)
	th.AssertEquals(t, "83e3bddc-b9ed-4614-a0dc-8a997095a86c", n.ID)
	th.AssertEquals(t, "ACTIVE", n.Status)
}

func TestDeleteSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/vpcs/8f794f06-2275-4d82-9f5a-6d68fbe21a75/subnets/83e3bddc-b9ed-4614-a0dc-8a997095a86c", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := subnets.Delete(fake.ServiceClient(), "8f794f06-2275-4d82-9f5a-6d68fbe21a75", "83e3bddc-b9ed-4614-a0dc-8a997095a86c")
	th.AssertNoErr(t, res.Err)
}
