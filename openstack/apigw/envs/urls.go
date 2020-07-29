package envs

import "github.com/huaweicloud/golangsdk"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("envs")
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("envs", id)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("envs")
}
