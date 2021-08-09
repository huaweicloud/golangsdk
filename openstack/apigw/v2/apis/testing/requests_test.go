package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/apis"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2API(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2APICreate(t)

	actual, err := apis.Create(client.ServiceClient(), "33fc92ffb7e749df952ecc7729d972bc",
		createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2API(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2APIGet(t)

	actual, err := apis.Get(client.ServiceClient(), "33fc92ffb7e749df952ecc7729d972bc",
		"cded6d80fc9f442c9842eaf854f10525").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestListV2API(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2APIList(t)

	pages, err := apis.List(client.ServiceClient(), "33fc92ffb7e749df952ecc7729d972bc",
		listOpts).AllPages()
	th.AssertNoErr(t, err)
	actual, err := apis.ExtractApis(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2API(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2APIUpdate(t)

	actual, err := apis.Update(client.ServiceClient(), "33fc92ffb7e749df952ecc7729d972bc",
		"cded6d80fc9f442c9842eaf854f10525", createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestDeleteV2API(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2APIDelete(t)

	err := apis.Delete(client.ServiceClient(), "33fc92ffb7e749df952ecc7729d972bc",
		"cded6d80fc9f442c9842eaf854f10525").ExtractErr()
	th.AssertNoErr(t, err)
}
