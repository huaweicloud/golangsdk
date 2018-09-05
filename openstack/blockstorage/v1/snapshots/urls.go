package snapshots

import "github.com/huaweicloud/golangsdk"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("snapshots")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func listURL(c *golangsdk.ServiceClient) string {
	return createURL(c)
}

func metadataURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id, "metadata")
}

func updateMetadataURL(c *golangsdk.ServiceClient, id string) string {
	return metadataURL(c, id)
}
