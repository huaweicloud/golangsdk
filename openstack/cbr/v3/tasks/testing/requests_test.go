package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/cbr/v3/tasks"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestGetV3Task(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleTaskGet(t)

	_, err := tasks.Get(client.ServiceClient(), "4827f2da-b008-4507-ab7d-42d0df5ed912").Extract()
	th.AssertNoErr(t, err)
}

func TestListV3Task(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleTaskList(t)

	pages, err := tasks.List(client.ServiceClient(), tasks.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := tasks.ExtractTasks(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}
