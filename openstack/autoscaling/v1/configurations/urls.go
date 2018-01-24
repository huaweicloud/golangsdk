package configurations

import (
	"github.com/gophercloud/gophercloud"
)

const resourcePath = "scaling_configuration"

func createURL(client *gophercloud.ServiceClientExtension) string {
	return client.ServiceURL(client.ProjectID, resourcePath)
}

func getURL(client *gophercloud.ServiceClientExtension, id string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, id)
}

func deleteURL(client *gophercloud.ServiceClientExtension, id string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, id)
}

func listURL(client *gophercloud.ServiceClientExtension) string {
	return client.ServiceURL(client.ProjectID, resourcePath)
}
