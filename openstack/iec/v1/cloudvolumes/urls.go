package cloudvolumes

import (
	"github.com/huaweicloud/golangsdk"
)

func GetURL(c *golangsdk.ServiceClient, CloudVolumeID string) string {
	return c.ServiceURL("cloudvolumes", CloudVolumeID)
}

func ListVolumeTypeURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("cloudvolumes", "volume-types")
}
