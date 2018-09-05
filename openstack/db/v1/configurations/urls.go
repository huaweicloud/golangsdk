package configurations

import "github.com/huaweicloud/golangsdk"

func baseURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("configurations")
}

func resourceURL(c *golangsdk.ServiceClient, configID string) string {
	return c.ServiceURL("configurations", configID)
}

func instancesURL(c *golangsdk.ServiceClient, configID string) string {
	return c.ServiceURL("configurations", configID, "instances")
}

func listDSParamsURL(c *golangsdk.ServiceClient, datastoreID, versionID string) string {
	return c.ServiceURL("datastores", datastoreID, "versions", versionID, "parameters")
}

func getDSParamURL(c *golangsdk.ServiceClient, datastoreID, versionID, paramID string) string {
	return c.ServiceURL("datastores", datastoreID, "versions", versionID, "parameters", paramID)
}

func listGlobalParamsURL(c *golangsdk.ServiceClient, versionID string) string {
	return c.ServiceURL("datastores", "versions", versionID, "parameters")
}

func getGlobalParamURL(c *golangsdk.ServiceClient, versionID, paramID string) string {
	return c.ServiceURL("datastores", "versions", versionID, "parameters", paramID)
}
