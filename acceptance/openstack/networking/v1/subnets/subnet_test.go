package subnets

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/networking/v1/subnets"
)

func TestSubnetList(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	if err != nil {
		t.Fatalf("Unable to create a subnet : %v", err)
	}
	allPages, err := subnets.List(client, subnets.ListOpts{})
	tools.PrintResource(t, allPages)

}

func TestSubnetsCRUD(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	if err != nil {
		t.Fatalf("Unable to create a subnet : %v", err)
	}

	// Create a subnet
	subnet, err := CreateSubnetNResources(t, client)
	if err != nil {
		t.Fatalf("Unable to create subnet: %v", err)
	}

	// Delete a subnet
	defer DeleteSubnetNResources(t, client, subnet.VPC_ID, subnet.ID)
	tools.PrintResource(t, subnet)

	// Update a subnet
	newName := tools.RandomString("ACPTTEST-", 8)
	updateOpts := &subnets.UpdateOpts{
		Name: newName,
	}
	_, err = subnets.Update(client, subnet.VPC_ID, subnet.ID, updateOpts).Extract()

	// Query a subnet
	newSubnet, err := subnets.Get(client, subnet.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve subnet: %v", err)
	}

	tools.PrintResource(t, newSubnet)
}
