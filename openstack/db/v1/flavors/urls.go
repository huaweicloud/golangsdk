package flavors

import "github.com/huaweicloud/golangsdk"

func getURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("flavors")
}
