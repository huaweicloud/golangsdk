package eips

import "github.com/huaweicloud/golangsdk"

const resourcePath = "publicips"
const apiVersion = "v1"

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath)
}

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath, id)
}
