package environments

import "github.com/huaweicloud/golangsdk"

const rootPath = "instances"

func rootURL(c *golangsdk.ServiceClient, instanceId string) string {
	return c.ServiceURL(rootPath, instanceId, "envs")
}

func resourceURL(c *golangsdk.ServiceClient, instanceId, envId string) string {
	return c.ServiceURL(rootPath, instanceId, "envs", envId)
}
