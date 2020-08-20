package addons

import (
	"github.com/huaweicloud/golangsdk"
)

const (
	rootPath = "addons"
)

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootPath)
}

func resourceURL(c *golangsdk.ServiceClient, id, cluster_id string) string {
	return c.ServiceURL(rootPath, id+"?cluster_id="+cluster_id)
}
