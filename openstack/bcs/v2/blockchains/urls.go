package blockchains

import "github.com/huaweicloud/golangsdk"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("blockchains")
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("blockchains")
}

func getURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("blockchains", instanceID)
}

func updateURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("blockchains", instanceID)
}

func deleteURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("blockchains", instanceID)
}

func extraURL(c *golangsdk.ServiceClient, instanceID string, extra string) string {
	return c.ServiceURL("blockchains", instanceID, extra)
}
