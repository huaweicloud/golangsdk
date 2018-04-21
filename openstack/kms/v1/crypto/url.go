package crypto

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateDEKURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "create-datakey")
}

func CreateDEKWithoutPlainTextURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "create-datakey-without-plaintext")
}

func DecryptDEKURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "decrypt-datakey")
}

func DecryptDataURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "decrypt-data")
}

func EncryptDEKURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "encrypt-datakey")
}

func EncryptDataURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "encrypt-data")
}

func GenerateRandomStringURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("kms", "gen-random")
}
