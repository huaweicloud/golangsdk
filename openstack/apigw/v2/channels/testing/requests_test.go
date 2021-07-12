package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/channels"

	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2VpcChannel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2VpcChannelCreate(t)

	actual, err := channels.Create(client.ServiceClient(), "b510b8e8ef1442c0a94cdfc551af0ec3",
		createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2VpcChannel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2VpcChannelGet(t)

	actual, err := channels.Get(client.ServiceClient(), "b510b8e8ef1442c0a94cdfc551af0ec3",
		"328d1d563eba4ff084533188b84b9f8d").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestListV2VpcChannel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2VpcChannelList(t)

	pages, err := channels.List(client.ServiceClient(), "b510b8e8ef1442c0a94cdfc551af0ec3",
		channels.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := channels.ExtractInstances(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2VpcChannel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2VpcChannelUpdate(t)

	actual, err := channels.Update(client.ServiceClient(), "b510b8e8ef1442c0a94cdfc551af0ec3",
		"328d1d563eba4ff084533188b84b9f8d", createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestDeleteV2VpcChannel(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2VpcChannelDelete(t)

	err := channels.Delete(client.ServiceClient(), "b510b8e8ef1442c0a94cdfc551af0ec3",
		"328d1d563eba4ff084533188b84b9f8d").ExtractErr()
	th.AssertNoErr(t, err)
}
