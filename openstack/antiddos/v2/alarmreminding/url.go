package alarmreminding

import (
	"github.com/huaweicloud/golangsdk"
)

func WarnAlertURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("warnalert", "alertconfig", "query")
}
