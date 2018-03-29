package peering

import (
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/networking/v1/vpcs"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/peerings"
)

func CreatePeeringResourcesNConn(t *testing.T, clientV2 *golangsdk.ServiceClient, peerClientV2 *golangsdk.ServiceClient,
	clientV1 *golangsdk.ServiceClient, peerClientV1 *golangsdk.ServiceClient) (*peerings.Peering, error) {

	vpcName := tools.RandomString("TESTACC-vpc", 8)
	peerVpcName := tools.RandomString("TESTACC-peervpc", 8)

	createOpts := vpcs.CreateOpts{
		Name: vpcName,
		CIDR: "192.168.20.0/24",
	}

	t.Logf("Attempting to create vpc: %s and peer vpc: %s", vpcName, peerVpcName)

	vpc, err := vpcs.Create(clientV1, createOpts).Extract()
	if err != nil {
		return nil, err
	}

	peerVpc, err := vpcs.Create(peerClientV1, createOpts).Extract()
	if err != nil {
		return nil, err
	}

	t.Logf("Created vpcs: %s %s", vpcName, peerVpcName)

	peeringConnName := tools.RandomString("TESTACC-", 8)

	peerCreateOpts := peerings.CreateOpts{
		Name:           peeringConnName,
		RequestVpcInfo: peerings.VpcInfo{VpcId: vpc.ID},
		AcceptVpcInfo:  peerings.VpcInfo{VpcId: peerVpc.ID, TenantId: peerClientV2.ProjectID},
	}

	t.Logf("Attempting to create vpc peering connection: %s", peeringConnName)

	peeringConns, err := peerings.Create(clientV2, peerCreateOpts).Extract()
	if err != nil {
		return peeringConns, err
	}

	if err := WaitForPeeringConnToCreate(clientV2, peeringConns.ID, 60); err != nil {
		return peeringConns, err
	}

	t.Logf("Created vpc peering connection: %s", peeringConnName)

	return peeringConns, nil
}

func DeletePeeringConnNResources(t *testing.T, clientV2 *golangsdk.ServiceClient, clientV1 *golangsdk.ServiceClient,
	peerClientV1 *golangsdk.ServiceClient, peeringConn *peerings.Peering) {
	t.Logf("Attempting to delete vpc peering connection: %s", peeringConn.ID)

	err := peerings.Delete(clientV2, peeringConn.ID).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting vpc peering connection: %v", err)
	}

	t.Logf("Deleted vpc peering connection: %s", peeringConn.ID)

	t.Logf("Attempting to delete vpc: %s", peeringConn.RequestVpcInfo.VpcId)

	err = vpcs.Delete(clientV1, peeringConn.RequestVpcInfo.VpcId).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting vpc: %v", err)
	}

	err = vpcs.Delete(peerClientV1, peeringConn.AcceptVpcInfo.VpcId).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting vpc: %v", err)
	}

	t.Logf("Deleted vpcs: %s and %s", peeringConn.RequestVpcInfo.VpcId, peeringConn.AcceptVpcInfo.VpcId)
}

func InitiatePeeringConnCommonTasks(t *testing.T) (*golangsdk.ServiceClient, *golangsdk.ServiceClient,
	*golangsdk.ServiceClient, *golangsdk.ServiceClient, *peerings.Peering) {

	clientV2, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network v2 client: %v", err)
	}

	clientV1, err := clients.NewNetworkV1Client()
	if err != nil {
		t.Fatalf("Unable to create a network v1 client: %v", err)
	}

	peerClientV2, err := clients.NewPeerNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network v2 client for peer: %v", err)
	}

	peerClientV1, err := clients.NewPeerNetworkV1Client()
	if err != nil {
		t.Fatalf("Unable to create a network v1 client for peer: %v", err)
	}

	// Create a vpc peering connection
	peeringConn, err := CreatePeeringResourcesNConn(t, clientV2, peerClientV2, clientV1, peerClientV1)
	if err != nil {
		t.Fatalf("Unable to create vpc peering connection: %v", err)
	}

	return clientV2, peerClientV2, clientV1, peerClientV1, peeringConn
}

func WaitForPeeringConnToCreate(client *golangsdk.ServiceClient, peeringConnID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		conn, err := peerings.Get(client, peeringConnID).Extract()
		if err != nil {
			return false, err
		}

		if conn.Status == "PENDING_ACCEPTANCE" {
			return true, nil
		}

		return false, nil
	})
}
