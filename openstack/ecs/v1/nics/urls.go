package nics

import "github.com/huaweicloud/golangsdk"

// endpoint/cloudservers/
const (
	rootPath     = "cloudservers"
	resourcePath = "nics"
)

// associateURL will build associate private ip from nic
func associateURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

// disassociateURL will build disassociate private ip from nic
func disassociateURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
