package routes

import (
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/peerings"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/routes"
)

func CreateRoute(t *testing.T, clientV2 *golangsdk.ServiceClient, peeringConn *peerings.Peering) (*routes.Route, error) {

	createRouteOpts := routes.CreateOpts{
		NextHop:     peeringConn.ID,
		Destination: "192.168.0.0/16",
		VPC_ID:      peeringConn.RequestVpcInfo.VpcId,
		Type:        "peering",
	}

	t.Logf("Attempting to create route")

	route, err := routes.Create(clientV2, createRouteOpts).Extract()
	if err != nil {
		return route, err
	}
	t.Logf("Created route: %s", route)

	return route, nil
}

func DeleteRoute(t *testing.T, clientV2 *golangsdk.ServiceClient, routeID string) {
	t.Logf("Attempting to delete route: %s", routeID)

	err := routes.Delete(clientV2, routeID).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting route: %v", err)
	}

	t.Logf("Deleted route: %s", routeID)
}
