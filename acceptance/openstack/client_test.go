// +build acceptance

package openstack

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestAuthenticatedClient(t *testing.T) {
	// Obtain credentials from the environment.
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		t.Fatalf("Unable to acquire credentials: %v", err)
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		t.Fatalf("Unable to authenticate: %v", err)
	}

	if client.TokenID == "" {
		t.Errorf("No token ID assigned to the client")
	}

	t.Logf("Client successfully acquired a token: %v", client.TokenID)

	// Find the storage service in the service catalog.
	storage, err := openstack.NewObjectStorageV1(client, golangsdk.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err != nil {
		t.Errorf("Unable to locate a storage service: %v", err)
	} else {
		t.Logf("Located a storage service at endpoint: [%s]", storage.Endpoint)
	}
}

func TestReauth(t *testing.T) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		t.Fatalf("Unable to obtain environment auth options: %v", err)
	}

	// Allow reauth
	ao.AllowReauth = true

	provider, err := openstack.NewClient(ao.IdentityEndpoint)
	if err != nil {
		t.Fatalf("Unable to create provider: %v", err)
	}

	err = openstack.Authenticate(provider, ao)
	if err != nil {
		t.Fatalf("Unable to authenticate: %v", err)
	}

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(provider, golangsdk.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err != nil {
		t.Fatalf("Unable to create compute client: %v", err)
	}

	t.Logf("Sleeping for 1 second")
	time.Sleep(1 * time.Second)
	t.Logf("Attempting to reauthenticate")

	err = provider.ReauthFunc()
	if err != nil {
		t.Fatalf("Unable to reauthenticate: %v", err)
	}

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(provider, golangsdk.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err != nil {
		t.Fatalf("Unable to create compute client: %v", err)
	}
}

func TestServiceClientEndpoint(t *testing.T) {
	var err error
	var serviceClient *golangsdk.ServiceClient
	var expectedURL, actualURL string

	apiDomain := "myhuaweicloud.com"
	region := os.Getenv("OS_REGION_NAME")
	projectID := os.Getenv("OS_PROJECT_ID")

	if projectID == "" {
		t.Fatalf("OS_PROJECT_ID must be set for service endpoint acceptance test")
	}

	// Obtain credentials from the environment.
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		t.Fatalf("Unable to acquire credentials: %v", err)
	}

	providerClient, err := openstack.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	if providerClient.TokenID == "" {
		t.Errorf("No token ID assigned to the client")
	}

	// Find the cce service in the service catalog.
	serviceClient, err = openstack.NewCCEV3(providerClient, golangsdk.EndpointOpts{
		Region: region,
	})
	th.AssertNoErr(t, err)
	expectedURL = fmt.Sprintf("https://cce.%s.%s/api/v3/projects/%s/", region, apiDomain, projectID)
	actualURL = serviceClient.ResourceBaseURL()
	th.AssertEquals(t, expectedURL, actualURL)
	t.Logf("cce endpoint:\t %s", actualURL)

	// Find the ces service in the service catalog.
	serviceClient, err = openstack.NewCESClient(providerClient, golangsdk.EndpointOpts{
		Region: region,
	})
	th.AssertNoErr(t, err)
	expectedURL = fmt.Sprintf("https://ces.%s.%s/V1.0/%s/", region, apiDomain, projectID)
	actualURL = serviceClient.ResourceBaseURL()
	th.AssertEquals(t, expectedURL, actualURL)
	t.Logf("ces endpoint:\t %s", actualURL)

	// Find the anti-ddos service in the service catalog.
	serviceClient, err = openstack.NewHwAntiDDoSV1(providerClient, golangsdk.EndpointOpts{
		Region: region,
	})
	th.AssertNoErr(t, err)
	expectedURL = fmt.Sprintf("https://antiddos.%s.%s/v1/%s/", region, apiDomain, projectID)
	actualURL = serviceClient.ResourceBaseURL()
	th.AssertEquals(t, expectedURL, actualURL)
	t.Logf("anti-ddos endpoint:\t %s", actualURL)

	// Find the auto-scaling service in the service catalog.
	serviceClient, err = openstack.NewAutoScalingService(providerClient, golangsdk.EndpointOpts{
		Region: region,
	})
	th.AssertNoErr(t, err)
	expectedURL = fmt.Sprintf("https://as.%s.%s/autoscaling-api/v1/%s/", region, apiDomain, projectID)
	actualURL = serviceClient.ResourceBaseURL()
	th.AssertEquals(t, expectedURL, actualURL)
	t.Logf("as endpoint:\t %s", actualURL)
}
