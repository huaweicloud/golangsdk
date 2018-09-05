package datastores

import "github.com/huaweicloud/golangsdk"

func baseURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("datastores")
}

func resourceURL(c *golangsdk.ServiceClient, dsID string) string {
	return c.ServiceURL("datastores", dsID)
}

func versionsURL(c *golangsdk.ServiceClient, dsID string) string {
	return c.ServiceURL("datastores", dsID, "versions")
}

func versionURL(c *golangsdk.ServiceClient, dsID, versionID string) string {
	return c.ServiceURL("datastores", dsID, "versions", versionID)
}
