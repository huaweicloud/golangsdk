package testing

import (
	"testing"

	az "github.com/huaweicloud/golangsdk/openstack/compute/v2/extensions/availabilityzones"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

// Verifies that availability zones can be listed correctly
func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleGetSuccessfully(t)

	allPages, err := az.List(client.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	actual, err := az.ExtractAvailabilityZones(allPages)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, AZResult, actual)
}

// Verifies that detailed availability zones can be listed correctly
// func TestListDetail(t *testing.T) {
//	th.SetupHTTP()
//	defer th.TeardownHTTP()

//	HandleGetDetailSuccessfully(t)

//	allPages, err := az.ListDetail(client.ServiceClient()).AllPages()
//	th.AssertNoErr(t, err)

//	actual, err := az.ExtractAvailabilityZones(allPages)
//	th.AssertNoErr(t, err)

//	th.CheckDeepEquals(t, AZDetailResult, actual)
// }
