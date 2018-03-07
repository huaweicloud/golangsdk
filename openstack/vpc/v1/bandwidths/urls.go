package bandwidths

import "github.com/gophercloud/gophercloud"

const resourcePath = "bandwidths"
const apiVersion = "v1"

func resourceURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath, id)
}
