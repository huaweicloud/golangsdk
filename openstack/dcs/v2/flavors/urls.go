package flavors

import (
	"github.com/huaweicloud/golangsdk"
)

// listURL will build the get url of List function
func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("flavors")
}
