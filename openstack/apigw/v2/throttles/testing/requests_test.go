package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/throttles"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2ThrottlingPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ThrottlingPolicyCreate(t)

	actual, err := throttles.Create(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b", createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2ThrottlingPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ThrottlingPolicyGet(t)

	actual, err := throttles.Get(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b",
		"c481043838f64bcd82c7b0c38907d59d").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestListV2ThrottlingPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ThrottlingPolicyList(t)

	pages, err := throttles.List(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b",
		listOpts).AllPages()
	th.AssertNoErr(t, err)
	actual, err := throttles.ExtractPolicies(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2ThrottlingPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ThrottlingPolicyUpdate(t)

	actual, err := throttles.Update(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b",
		"c481043838f64bcd82c7b0c38907d59d", createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestDeleteV2ThrottlingPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ThrottlingPolicyDelete(t)

	err := throttles.Delete(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b",
		"c481043838f64bcd82c7b0c38907d59d").ExtractErr()
	th.AssertNoErr(t, err)
}
