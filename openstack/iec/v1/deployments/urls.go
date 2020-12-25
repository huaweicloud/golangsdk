package deployments

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("deployments")
}

func DeployURL(c *golangsdk.ServiceClient, deploymentID string) string {
	return c.ServiceURL("deployments", deploymentID, "deploy")
}

func ExpandURL(c *golangsdk.ServiceClient, deploymentID string) string {
	return c.ServiceURL("deployments", deploymentID, "expand")
}

func DeleteURL(c *golangsdk.ServiceClient, deploymentID string) string {
	return c.ServiceURL("deployments", deploymentID)
}
