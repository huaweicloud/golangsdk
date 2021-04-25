package tasks

import "github.com/huaweicloud/golangsdk"

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("operation-logs")
}

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("operation-logs", id)
}
