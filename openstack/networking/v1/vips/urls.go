package vips

import "github.com/huaweicloud/golangsdk"

// endpoint/{tenant_id}/privateips
const resourcePath = "privateips"

// createURL will build the rest query url of creation
func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(client.ProjectID, resourcePath)
}

// deleteURL will build the url of deletion
func deleteURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, id)
}

// getURL will build the get url of get function
func getURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, id)
}
