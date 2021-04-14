package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/cbr/v3/policies"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV3PolicyMarshall(t *testing.T) {
	res, err := createOpts.ToPolicyCreateMap()
	th.AssertNoErr(t, err)
	th.AssertJSONEquals(t, expectedRequest, res)
}

func TestCreateV3Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePolicyCreation(t)

	actual, err := policies.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestDeleteV3Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePolicyDeletion(t)

	err := policies.Delete(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUpdateV3Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePolicyUpdate(t)

	updateId := "cbb3ce6f-3332-4e7c-b98e-77290d8471ff"
	actual, err := policies.Update(client.ServiceClient(), updateId, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV3Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePolicyGet(t)

	actual, err := policies.Get(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestListV3Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePolicyList(t)

	pages, err := policies.List(client.ServiceClient(), policies.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := policies.ExtractPolicies(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}
