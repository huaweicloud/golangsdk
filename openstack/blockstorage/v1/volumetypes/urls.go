package volumetypes

import "github.com/huaweicloud/golangsdk"

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func createURL(c *golangsdk.ServiceClient) string {
	return listURL(c)
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return getURL(c, id)
}
