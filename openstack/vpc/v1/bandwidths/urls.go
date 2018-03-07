package bandwidths

import "github.com/huaweicloud/golangsdk"

const resourcePath = "bandwidths"
const apiVersion = "v1"

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath, id)
}
