package cloudservers

import "github.com/huaweicloud/golangsdk"

func createURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL("cloudservers")
}

func deleteURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL("cloudservers", "delete")
}

func getURL(sc *golangsdk.ServiceClient, serverID string) string {
	return sc.ServiceURL("cloudservers", serverID)
}

func jobURL(sc *golangsdk.ServiceClient, jobId string) string {
	return sc.ServiceURL("jobs", jobId)
}

func orderURL(sc *golangsdk.ServiceClient, orderId string) string {
	return sc.ServiceURL(sc.DomainID, "common/order-mgr/orders-resource", orderId)
}

func deleteOrderURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL(sc.DomainID, "common/order-mgr/resources/delete")
}
