package common

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const TokenID = client.TokenID

// Fake project id to use.
const ProjectID = "17fbda95add24720a4038ba4b1c705ed"

func ServiceClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + ProjectID + "/"
	sc.ProjectID = ProjectID
	return sc
}
