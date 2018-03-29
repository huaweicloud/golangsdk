package subnets

import (
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/networking/v1/subnets"
	"github.com/huaweicloud/golangsdk/openstack/networking/v1/vpcs"
)

func CreateSubnetNResources(t *testing.T, client *golangsdk.ServiceClient) (*subnets.Subnet, error) {

	vpcName := tools.RandomString("TESTACC-", 8)

	createOpts := vpcs.CreateOpts{
		Name: vpcName,
		CIDR: "192.168.20.0/24",
	}

	t.Logf("Attempting to create vpc: %s", vpcName)

	vpc, err := vpcs.Create(client, createOpts).Extract()
	if err != nil {
		return nil, err
	}
	t.Logf("Created vpc: %s", vpcName)

	subnetName := tools.RandomString("ACPTTEST-", 8)

	createSubnetOpts := subnets.CreateOpts{
		Name:             subnetName,
		CIDR:             "192.168.20.0/24",
		GatewayIP:        "192.168.20.1",
		EnableDHCP:       true,
		AvailabilityZone: "eu-de-02",
		VPC_ID:           vpc.ID,
	}

	t.Logf("Attempting to create subnet: %s", subnetName)

	subnet, err := subnets.Create(client, createSubnetOpts).Extract()
	if err != nil {
		return subnet, err
	}
	t.Logf("Created subnet: %v", subnet)

	return subnet, nil
}

func DeleteSubnetNResources(t *testing.T, client *golangsdk.ServiceClient, vpcID string, id string) {
	t.Logf("Attempting to delete subnet: %s", id)

	err := subnets.Delete(client, vpcID, id).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting subnet: %v", err)
	}

	t.Logf("Deleted subnet: %s", id)

	if err := WaitForSubnetToDelete(client, id, 60); err != nil {
		t.Fatalf("Error deleting subnet: %v", err)
	}

	t.Logf("Attempting to delete vpc: %s", vpcID)

	err = vpcs.Delete(client, vpcID).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting vpc: %v", err)
	}

	t.Logf("Deleted vpc: %s", vpcID)
}

func WaitForSubnetToDelete(client *golangsdk.ServiceClient, subnetID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := subnets.Get(client, subnetID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
		}

		return false, nil
	})
}
