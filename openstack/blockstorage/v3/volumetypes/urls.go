package volumetypes

import "github.com/huaweicloud/golangsdk"

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func updateURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}
