package routes

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/openstack/networking/v2/peering"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/peerings"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/routes"
)

func TestRouteList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a route client: %v", err)
	}

	listOpts := routes.ListOpts{}
	pages, err := routes.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to list routers: %v", err)
	}

	allRoutes, err := routes.ExtractRoutes(pages)
	if err != nil {
		t.Errorf("Failed to extract routes: %v", err)
	}

	for _, router := range allRoutes {
		tools.PrintResource(t, router)
	}
}

func TestRoutesCRUD(t *testing.T) {

	clientV2, peerClientV2, clientV1, peerClientV1, peeringConn := peering.InitiatePeeringConnCommonTasks(t)

	_, err := peerings.Accept(peerClientV2, peeringConn.ID).ExtractResult()
	if err != nil {
		t.Fatalf("Unable to accept peering request: %v", err)
	}

	// Create a Route
	route, err := CreateRoute(t, clientV2, peeringConn)

	if err != nil {
		t.Fatalf("Unable to create route: %v", err)
	}

	defer peering.DeletePeeringConnNResources(t, clientV2, clientV1, peerClientV1, peeringConn)
	defer DeleteRoute(t, clientV2, route.RouteID)

	tools.PrintResource(t, route)

	newRoute, err := routes.Get(clientV2, route.RouteID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve route: %v", err)
	}

	tools.PrintResource(t, newRoute)
}
