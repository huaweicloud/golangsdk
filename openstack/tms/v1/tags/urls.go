package tags

import "github.com/huaweicloud/golangsdk"

const (
	rootPath   = "predefine_tags"
	actionPath = "predefine_tags/action"
)

func actionURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(actionPath)
}

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath)
}
