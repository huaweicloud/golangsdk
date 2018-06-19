package stackevents

import "github.com/huaweicloud/golangsdk"

func listURL(c *golangsdk.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "events")
}
