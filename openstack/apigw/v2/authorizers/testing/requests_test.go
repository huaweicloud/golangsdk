package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/authorizers"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2CustomAuthorizer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2CustomAuthorizerCreate(t)

	actual, err := authorizers.Create(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b",
		createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestGetV2CustomAuthorizer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2CustomAuthorizerGet(t)

	actual, err := authorizers.Get(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b",
		"0d2a523974a14fe1a25c1bc2f61b2d9d").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestListV2CustomAuthorizer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2CustomAuthorizerList(t)

	pages, err := authorizers.List(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b",
		listOpts).AllPages()
	th.AssertNoErr(t, err)
	actual, err := authorizers.ExtractCustomAuthorizers(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2CustomAuthorizer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2CustomAuthorizerUpdate(t)

	actual, err := authorizers.Update(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b",
		"0d2a523974a14fe1a25c1bc2f61b2d9d", createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestDeleteV2CustomAuthorizer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2CustomAuthorizerDelete(t)

	err := authorizers.Delete(client.ServiceClient(), "6da953fe33d44650a067e43a4593368b",
		"0d2a523974a14fe1a25c1bc2f61b2d9d").ExtractErr()
	th.AssertNoErr(t, err)
}
