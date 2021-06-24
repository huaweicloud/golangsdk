package responses

import "github.com/huaweicloud/golangsdk"

const rootPath = "instances"

func rootURL(c *golangsdk.ServiceClient, instanceId, groupId string) string {
	return c.ServiceURL(rootPath, instanceId, "api-groups", groupId, "gateway-responses")
}

func resourceURL(c *golangsdk.ServiceClient, instanceId, groupId, respId string) string {
	return c.ServiceURL(rootPath, instanceId, "api-groups", groupId, "gateway-responses", respId)
}

func responseURL(c *golangsdk.ServiceClient, instanceId, groupId, respId, respType string) string {
	return c.ServiceURL(rootPath, instanceId, "api-groups", groupId, "gateway-responses", respId, respType)
}
