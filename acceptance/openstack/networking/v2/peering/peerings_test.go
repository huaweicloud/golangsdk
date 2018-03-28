package peering

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/peerings"
)

func TestPeeringList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a vpc client: %v", err)
	}

	listOpts := peerings.ListOpts{}
	peering, err := peerings.List(client, listOpts)
	if err != nil {
		t.Fatalf("Unable to list peerings: %v", err)
	}
	for _, peering := range peering {
		tools.PrintResource(t, peering)
	}
}

func TestAcceptPeering(t *testing.T) {

	clientV2, peerClientV2, clientV1, peerClientV1, peeringConn := InitiatePeeringConnCommonTasks(t)

	// Delete a vpc peering connection
	defer DeletePeeringConnNResources(t, clientV2, clientV1, peerClientV1, peeringConn)

	peeringConn1, err := peerings.Accept(peerClientV2, peeringConn.ID).ExtractResult()
	if err != nil {
		t.Fatalf("Unable to accept peering request: %v", err)
	}
	tools.PrintResource(t, peeringConn1)

}

func TestRejectPeering(t *testing.T) {

	clientV2, peerClientV2, clientV1, peerClientV1, peeringConn := InitiatePeeringConnCommonTasks(t)

	// Delete a vpc peering connection
	defer DeletePeeringConnNResources(t, clientV2, clientV1, peerClientV1, peeringConn)

	peerConn1, err := peerings.Reject(peerClientV2, peeringConn.ID).ExtractResult()
	if err != nil {
		t.Fatalf("Unable to Reject peering request: %v", err)
	}
	tools.PrintResource(t, peerConn1)

}

func TestPeeringCRUD(t *testing.T) {

	clientV2, peerClientV2, clientV1, peerClientV1, peeringConn := InitiatePeeringConnCommonTasks(t)

	// Delete a vpc peering connection
	defer DeletePeeringConnNResources(t, clientV2, clientV1, peerClientV1, peeringConn)

	tools.PrintResource(t, peeringConn)
	updateOpts := peerings.UpdateOpts{
		Name: "test2",
	}

	_, err := peerings.Update(clientV2, peeringConn.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update vpc peering connection: %v", err)
	}

	peeringConnGet, err := peerings.Get(peerClientV2, peeringConn.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve vpc peering connection: %v", err)
	}

	tools.PrintResource(t, peeringConnGet)
}
