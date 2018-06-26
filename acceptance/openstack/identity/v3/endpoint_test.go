// +build acceptance

package v3

import (
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/identity/v3/endpoints"
	"github.com/huaweicloud/golangsdk/openstack/identity/v3/services"
)

func TestEndpointsList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	allPages, err := endpoints.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list endpoints: %v", err)
	}

	allEndpoints, err := endpoints.ExtractEndpoints(allPages)
	if err != nil {
		t.Fatalf("Unable to extract endpoints: %v", err)
	}

	for _, endpoint := range allEndpoints {
		tools.PrintResource(t, endpoint)
	}
}

func TestEndpointsGet(t *testing.T) {

	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	allPages, err := endpoints.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list endpoints: %v", err)
	}

	allEndpoints, err := endpoints.ExtractEndpoints(allPages)
	if err != nil {
		t.Fatalf("Unable to extract endpoints: %v", err)
	}

	if len(allEndpoints) > 0 {
		endpoint := allEndpoints[0]
		p, err := endpoints.Get(client, endpoint.ID).Extract()
		if err != nil {
			t.Fatalf("Unable to get endpoint: %v", err)
		}

		tools.PrintResource(t, p)
	}
}

func TestEndpointsNavigateCatalog(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	// Discover the service we're interested in.
	serviceListOpts := services.ListOpts{
		ServiceType: "compute",
	}

	allPages, err := services.List(client, serviceListOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to lookup compute service: %v", err)
	}

	allServices, err := services.ExtractServices(allPages)
	if err != nil {
		t.Fatalf("Unable to extract service: %v", err)
	}

	if len(allServices) != 1 {
		t.Fatalf("Expected one service, got %d", len(allServices))
	}

	computeService := allServices[0]
	tools.PrintResource(t, computeService)

	// Enumerate the endpoints available for this service.
	endpointListOpts := endpoints.ListOpts{
		Availability: golangsdk.AvailabilityPublic,
		ServiceID:    computeService.ID,
	}

	allPages, err = endpoints.List(client, endpointListOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to lookup compute endpoint: %v", err)
	}

	allEndpoints, err := endpoints.ExtractEndpoints(allPages)
	if err != nil {
		t.Fatalf("Unable to extract endpoint: %v", err)
	}

	if len(allEndpoints) != 1 {
		t.Fatalf("Expected one endpoint, got %d", len(allEndpoints))
	}

	tools.PrintResource(t, allEndpoints[0])

}
