package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/mrs/v2/jobs"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2Job(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2JobCreate(t)

	actual, err := jobs.Create(client.ServiceClient(), "986416ea-e26b-40f1-b371-cd7be87376a2", createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2Job(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2JobGet(t)

	actual, err := jobs.Get(client.ServiceClient(), "986416ea-e26b-40f1-b371-cd7be87376a2",
		"3c0cf394-5da2-46a7-92df-795d998edea7").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestListV2Job(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2JobList(t)

	pages, err := jobs.List(client.ServiceClient(), "986416ea-e26b-40f1-b371-cd7be87376a2", jobs.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := jobs.ExtractJobs(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestDeleteV2Job(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2JobDelete(t)

	err := jobs.Delete(client.ServiceClient(), "986416ea-e26b-40f1-b371-cd7be87376a2", deleteOpts).ExtractErr()
	th.AssertNoErr(t, err)
}
