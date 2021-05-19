package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/autoscaling/v1/lifecyclehooks"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateLifecycleHook(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleLifeCycleHookCreate(t)

	actual, err := lifecyclehooks.Create(client.ServiceClient(), createOpts, "50ed20b8-9853-4668-a71c-c8c15b5cb85f").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestGetLifecycleHook(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleLifeCycleHookGet(t)

	actual, err := lifecyclehooks.Get(client.ServiceClient(), "50ed20b8-9853-4668-a71c-c8c15b5cb85f", "test-hook").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestListLifecycleHook(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleLifeCycleHookList(t)

	actual, err := lifecyclehooks.List(client.ServiceClient(), "50ed20b8-9853-4668-a71c-c8c15b5cb85f").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateLifecycleHook(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleLifeCycleHookUpdate(t)

	actual, err := lifecyclehooks.Update(client.ServiceClient(), updateOpts, "50ed20b8-9853-4668-a71c-c8c15b5cb85f", "test-hook").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedUpdateResponseData, actual)
}

func TestDeleteLifecycleHook(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleLifeCycleHookDelete(t)

	err := lifecyclehooks.Delete(client.ServiceClient(), "50ed20b8-9853-4668-a71c-c8c15b5cb85f", "test-hook").Extract()
	th.AssertNoErr(t, err)
}
