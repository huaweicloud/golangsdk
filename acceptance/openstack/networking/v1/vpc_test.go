package v1

import (
	"testing"

	"github.com/huaweicloud/golangsdk"
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
	vpc, err := createVpc(t, client)
	if err != nil {
		t.Fatalf("Unable to create create: %v", err)
	}
	defer deleteVpc(t, client, vpc.ID)

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

func createVpc(t *testing.T, client *golangsdk.ServiceClient) (*vpcs.Vpc, error) {

	vpcName := tools.RandomString("TESTACC-", 8)

	createOpts := vpcs.CreateOpts{
		Name: vpcName,
		CIDR: "192.168.20.0/24",
	}

	t.Logf("Attempting to create vpc: %s", vpcName)

	vpc, err := vpcs.Create(client, createOpts).Extract()
	if err != nil {
		return vpc, err
	}
	t.Logf("Created vpc: %s", vpcName)

	return vpc, nil
}

func deleteVpc(t *testing.T, client *golangsdk.ServiceClient, vpcID string) {
	t.Logf("Attempting to delete vpc: %s", vpcID)

	err := vpcs.Delete(client, vpcID).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting vpc: %v", err)
	}

	t.Logf("Deleted vpc: %s", vpcID)
}
