package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/fgs/v2/trigger"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2Trigger(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2TriggerCreate(t)

	actual, err := trigger.Create(client.ServiceClient(), createOpts,
		"urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestGetV2Trigger(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2TriggerGet(t)

	actual, err := trigger.Get(client.ServiceClient(),
		"urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test",
		"TIMER", "971f9cff-5d29-42da-ba36-e4e2b9d26664").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestListV2Trigger(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2TriggerList(t)

	pages, err := trigger.List(client.ServiceClient(),
		"urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test").AllPages()
	th.AssertNoErr(t, err)
	actual, err := trigger.ExtractList(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2Trigger(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2TriggerUpdate(t)

	err := trigger.Update(client.ServiceClient(), updateOpts,
		"urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test",
		"TIMER", "971f9cff-5d29-42da-ba36-e4e2b9d26664").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestDeleteV2Trigger(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2TriggerDelete(t)

	err := trigger.Delete(client.ServiceClient(),
		"urn:fss:cn-north-4:0721565481a5d123f6e6c66a2115215a:function:default:test",
		"TIMER", "971f9cff-5d29-42da-ba36-e4e2b9d26664").ExtractErr()
	th.AssertNoErr(t, err)
}
