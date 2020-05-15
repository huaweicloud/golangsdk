package tags

import (
	"github.com/huaweicloud/golangsdk"
)

// supported resourceType: "DNS-public_zone", "DNS-private_zone",
// "DNS-public_recordset", "DNS-private_recordset", "DNS-ptr_record"
func actionURL(c *golangsdk.ServiceClient, resourceType, id string) string {
	return c.ServiceURL(c.ProjectID, resourceType, id, "tags/action")
}

func getURL(c *golangsdk.ServiceClient, resourceType, id string) string {
	return c.ServiceURL(c.ProjectID, resourceType, id, "tags")
}

func listURL(c *golangsdk.ServiceClient, resourceType string) string {
	return c.ServiceURL(c.ProjectID, resourceType, "tags")
}
