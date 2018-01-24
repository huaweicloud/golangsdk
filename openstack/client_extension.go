package openstack

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	tokens2 "github.com/gophercloud/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/openstack/utils"
)

func GetProjectId(client *gophercloud.ProviderClient) (string, error) {
	versions := []*utils.Version{
		{ID: v2, Priority: 20, Suffix: "/v2.0/"},
		{ID: v3, Priority: 30, Suffix: "/v3/"},
	}

	chosen, endpoint, err := utils.ChooseVersion(client, versions)
	if err != nil {
		return "", err
	}

	switch chosen.ID {
	case v2:
		return getV2ProjectId(client, endpoint)
	case v3:
		return getV3ProjectId(client, endpoint)
	default:
		return "", fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

func getV2ProjectId(client *gophercloud.ProviderClient, endpoint string) (string, error) {
	v2Client, err := NewIdentityV2(client, gophercloud.EndpointOpts{})
	if err != nil {
		return "", err
	}

	if endpoint != "" {
		v2Client.Endpoint = endpoint
	}

	result := tokens2.Get(v2Client, client.TokenID)
	token, err := result.ExtractToken()
	if err != nil {
		return "", err
	}

	return token.Tenant.ID, nil
}

func getV3ProjectId(client *gophercloud.ProviderClient, endpoint string) (string, error) {
	v3Client, err := NewIdentityV3(client, gophercloud.EndpointOpts{})
	if err != nil {
		return "", err
	}

	if endpoint != "" {
		v3Client.Endpoint = endpoint
	}

	result := tokens3.Get(v3Client, client.TokenID)
	project, err := result.ExtractProject()
	if err != nil {
		return "", err
	}

	return project.ID, nil
}

func initClientOptsExtension(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, clientType string) (*gophercloud.ServiceClientExtension, error) {
	pid, e := GetProjectId(client)
	if e != nil {
		return nil, e
	}

	c, e := initClientOpts(client, eo, clientType)
	if e != nil {
		return nil, e
	}

	sc := new(gophercloud.ServiceClientExtension)
	sc.ServiceClient = c
	sc.ProjectID = pid
	return sc, nil
}

//NewAutoScalingService creates a ServiceClient that may be used to access the
//auto-scaling service of huawei public cloud
func NewAutoScalingService(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClientExtension, error) {
	sc, err := initClientOptsExtension(client, eo, "as")
	return sc, err
}
