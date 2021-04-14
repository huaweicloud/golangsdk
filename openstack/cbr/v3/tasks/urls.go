package tasks

import "github.com/huaweicloud/golangsdk"

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("operation-logs")
}

func singleURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("operation-logs", id)
}
