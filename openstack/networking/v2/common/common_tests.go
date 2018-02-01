package common

import (
	"github.com/huawei-clouds/golangsdk"
	"github.com/huawei-clouds/golangsdk/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
