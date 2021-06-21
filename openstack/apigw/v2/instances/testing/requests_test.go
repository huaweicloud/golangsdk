package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/instances"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2Instance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2InstanceCreate(t)

	actual, err := instances.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2Instance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2InstanceGet(t)

	actual, err := instances.Get(client.ServiceClient(), "e6a5871bfb5b47d19c5874790f639ef8").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestListV2Instance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2InstanceList(t)

	pages, err := instances.List(client.ServiceClient(), instances.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := instances.ExtractInstances(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2Instance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2InstanceUpdate(t)

	actual, err := instances.Update(client.ServiceClient(), "e6a5871bfb5b47d19c5874790f639ef8", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestDeleteV2Instance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2InstanceDelete(t)

	err := instances.Delete(client.ServiceClient(), "e6a5871bfb5b47d19c5874790f639ef8").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestDisableEgressV2Instance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2InstanceEgressDisable(t)

	err := instances.DisableEgressAccess(client.ServiceClient(), "e6a5871bfb5b47d19c5874790f639ef8").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUpdateIngressV2Instance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2InstanceIngressUpdate(t)

	actual, err := instances.UpdateIngressAccess(client.ServiceClient(), "e6a5871bfb5b47d19c5874790f639ef8",
		updateIngressOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedUpdateIngressResponseData, actual)
}

func TestDisableIngressV2Instance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2InstanceIngressDisable(t)

	err := instances.DisableIngressAccess(client.ServiceClient(), "e6a5871bfb5b47d19c5874790f639ef8").ExtractErr()
	th.AssertNoErr(t, err)
}
