package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/applications"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2Application(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationCreate(t)

	actual, err := applications.Create(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a", appOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2Application(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationGet(t)

	actual, err := applications.Get(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a",
		"50f768cf1c1f4389965aa58d255b2a65").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestListV2Application(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationList(t)

	pages, err := applications.List(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a",
		applications.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := applications.ExtractApplications(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2Application(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationUpdate(t)

	actual, err := applications.Update(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a",
		"50f768cf1c1f4389965aa58d255b2a65", appOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestDeleteV2Application(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationDelete(t)

	err := applications.Delete(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a",
		"50f768cf1c1f4389965aa58d255b2a65").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestAddV2ApplicationCode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationCodeAdd(t)

	actual, err := applications.CreateAppCode(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a",
		"50f768cf1c1f4389965aa58d255b2a65", appCodeOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedAddCodeResponseData, actual)
}

func TestAutoGenerateV2ApplicationCode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationCodeAutoGenerate(t)

	actual, err := applications.AutoGenerateAppCode(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a",
		"50f768cf1c1f4389965aa58d255b2a65").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedAddCodeResponseData, actual)
}

func TestGetV2ApplicationCode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationCodeGet(t)

	actual, err := applications.GetAppCode(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a",
		"50f768cf1c1f4389965aa58d255b2a65", "f284119e19f34d4caead4dd94114a7f4").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedAddCodeResponseData, actual)
}

func TestListV2ApplicationCode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationCodeList(t)

	pages, err := applications.ListAppCode(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a",
		"50f768cf1c1f4389965aa58d255b2a65", applications.ListCodeOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := applications.ExtractAppCodes(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListCodeResponseData, actual)
}

func TestRemove2ApplicationCode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ApplicationCodeRemove(t)

	err := applications.RemoveAppCode(client.ServiceClient(), "c5faacb524d148b59ddd448dd02d016a",
		"50f768cf1c1f4389965aa58d255b2a65", "f284119e19f34d4caead4dd94114a7f4").ExtractErr()
	th.AssertNoErr(t, err)
}
