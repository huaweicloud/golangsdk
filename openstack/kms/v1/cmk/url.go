package cmk

import (
	"github.com/huaweicloud/golangsdk"
)

func CancelDeletionURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "cancel-key-deletion")
}

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "create-key")
}

func DisableURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "disable-key")
}

func EnableURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "enable-key")
}

func GetURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "describe-key")
}

func GrantURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "create-grant")
}

func InstancesURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "user-instances")
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "list-keys")
}

func ListGrantsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "list-grants")
}

func ListRetirableURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "list-retirable-grants")
}

func QuotasURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "user-quotas")
}

func RetireURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "retire-grant")
}

func RevokeURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "revoke-grant")
}

func ScheduleDeletionURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "schedule-key-deletion")
}

func UpdateAliasURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "update-key-alias")
}

func UpdateDescriptionURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "update-key-description")
}
