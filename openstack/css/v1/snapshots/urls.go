package snapshots

import "github.com/huaweicloud/golangsdk"

func policyURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshot/policy")
}

func createURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshot")
}

func listURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshots")
}

func disableURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshots")
}

func restoreURL(c *golangsdk.ServiceClient, clusterId, snapId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshot", snapId, "restore")
}

func deleteURL(c *golangsdk.ServiceClient, clusterId, snapId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshot", snapId)
}
