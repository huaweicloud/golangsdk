package domains

import "github.com/huaweicloud/golangsdk"

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("auth/domains")
}

func getURL(client *golangsdk.ServiceClient, domainID string) string {
	return client.ServiceURL("auth/domains", domainID)
}

func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("auth/domains")
}

func deleteURL(client *golangsdk.ServiceClient, domainID string) string {
	return client.ServiceURL("auth/domains", domainID)
}

func updateURL(client *golangsdk.ServiceClient, domainID string) string {
	return client.ServiceURL("auth/domains", domainID)
}
