package keys

import "github.com/huaweicloud/golangsdk"

const (
	resourcePath = "kms"
)

func getURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "describe-key")
}

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "create-key")
}

func deleteURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "schedule-key-deletion")
}

func updateAliasURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "update-key-alias")
}

func updateDesURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "update-key-description")
}

func dataEncryptURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "create-datakey")
}

func encryptDEKURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "encrypt-datakey")
}

func enableKeyURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "enable-key")
}

func disableKeyURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "disable-key")
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "list-keys")
}
