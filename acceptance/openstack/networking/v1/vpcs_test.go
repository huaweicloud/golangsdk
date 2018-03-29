package v1

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/networking/v1/vpcs"
)

func TestVpcList(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	if err != nil {
		t.Fatalf("Unable to create a vpc client: %v", err)
	}

	listOpts := vpcs.ListOpts{}
	allVpcs, err := vpcs.List(client, listOpts)
	if err != nil {
		t.Fatalf("Unable to list vpcs: %v", err)
	}
	for _, vpc := range allVpcs {
		tools.PrintResource(t, vpc)
	}
}

func TestVpcsCRUD(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	if err != nil {
		t.Fatalf("Unable to create a vpc client: %v", err)
	}

	// Create a vpc
	vpc, err := CreateVpc(t, client)
	if err != nil {
		t.Fatalf("Unable to create create: %v", err)
	}
	defer DeleteVpc(t, client, vpc.ID)

	tools.PrintResource(t, vpc)

	newName := tools.RandomString("TESTACC-", 8)
	updateOpts := &vpcs.UpdateOpts{
		Name: newName,
	}

	_, err = vpcs.Update(client, vpc.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update vpc: %v", err)
	}

	newVpc, err := vpcs.Get(client, vpc.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve vpc: %v", err)
	}

	tools.PrintResource(t, newVpc)
}
