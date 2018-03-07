package eips

import "github.com/gophercloud/gophercloud"

const resourcePath = "publicips"
const apiVersion = "v1"

func rootURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath)
}

func resourceURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath, id)
}
