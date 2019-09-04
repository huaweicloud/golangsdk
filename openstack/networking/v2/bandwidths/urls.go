package bandwidths

import "github.com/huaweicloud/golangsdk"

func PostURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("bandwidths")
}

func BatchPostURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("batch-bandwidths")
}
func UpdateURL(c *golangsdk.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID)
}

func DeleteURL(c *golangsdk.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID)
}

func InsertURL(c *golangsdk.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID, "insert")
}

func RemoveURL(c *golangsdk.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID, "remove")
}
