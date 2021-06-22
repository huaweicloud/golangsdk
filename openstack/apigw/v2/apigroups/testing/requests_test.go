package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/apigroups"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2Group(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2GroupCreate(t)

	actual, err := apigroups.Create(client.ServiceClient(), "9750f26518a54da8bea1a7c41790c26d", createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2Group(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2GroupGet(t)

	actual, err := apigroups.Get(client.ServiceClient(), "9750f26518a54da8bea1a7c41790c26d", "1c1acdd2f4d14eb886ecd2370cdb9c1a").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestListV2Group(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2GroupList(t)

	pages, err := apigroups.List(client.ServiceClient(), "9750f26518a54da8bea1a7c41790c26d", apigroups.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := apigroups.ExtractGroups(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2Group(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2GroupUpdate(t)

	actual, err := apigroups.Update(client.ServiceClient(), "9750f26518a54da8bea1a7c41790c26d", "1c1acdd2f4d14eb886ecd2370cdb9c1a", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedUpdateResponseData, actual)
}

func TestDeleteV2Group(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2GroupDelete(t)

	err := apigroups.Delete(client.ServiceClient(), "9750f26518a54da8bea1a7c41790c26d", "1c1acdd2f4d14eb886ecd2370cdb9c1a").ExtractErr()
	th.AssertNoErr(t, err)
}
