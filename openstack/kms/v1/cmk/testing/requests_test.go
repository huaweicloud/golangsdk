package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/kms/v1/cmk"
	"github.com/huaweicloud/golangsdk/postpagination"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	var createOps = cmk.CreateOpts{
		KeyAlias:       "TestCMK2",
		KeyDescription: "It's test CMK2.",
		Origin:         "kms",
		Realm:          "",
		KeyPolicy:      "",
		KeyUsage:       "",
		KeyType:        "",
		Sequence:       "919c82d4-8046-4722-9094-35c3c6524cff",
	}

	actual, err := cmk.Create(client.ServiceClient(), createOps).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestEnable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleEnableSuccessfully(t)

	var enableOps = cmk.EnableOpts{
		KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
		Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
	}

	actual, err := cmk.Enable(client.ServiceClient(), enableOps).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &EnableResponse, actual)
}

func TestDisable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDisableSuccessfully(t)

	var disableOps = cmk.DisableOpts{
		KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
		Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
	}

	actual, err := cmk.Disable(client.ServiceClient(), disableOps).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &DisableResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	var getOps = cmk.GetOpts{
		KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
		Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
	}

	actual, err := cmk.Get(client.ServiceClient(), getOps).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestInstances(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleInstancesSuccessfully(t)

	actual, err := cmk.Instances(client.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &InstancesResponse, actual)
}

func TestListAllPage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := cmk.List(client.ServiceClient(), cmk.ListOpts{
		Limit: "1",
	}).AllPages()
	th.AssertNoErr(t, err)
	listResponse, err := cmk.ExtractList(actual.(cmk.ListPage))
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListAllPagesResponse, listResponse)
}

func TestListEachPage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	err := cmk.List(client.ServiceClient(), cmk.ListOpts{
		Limit: "1",
	}).EachPage(func(page postpagination.Page) (bool, error) {
		listResponse, err := cmk.ExtractList(page.(cmk.ListPage))
		th.AssertNoErr(t, err)
		isEmpty, err := page.(cmk.ListPage).IsEmpty()
		th.AssertNoErr(t, err)
		th.AssertEquals(t, false, isEmpty)
		if listResponse.NextMarker == "1" {
			th.CheckDeepEquals(t, &ListFirstResponse, listResponse)
			return false, nil
		} else {
			th.CheckDeepEquals(t, &ListSecondResponse, listResponse)
			return true, nil
		}
	})
	th.AssertNoErr(t, err)
}

func TestScheduleDeletion(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleScheduleDeletionSuccessfully(t)

	var scheduleDeletiontOps = cmk.ScheduleDeletionOpts{
		KeyId:       "30361023-62e0-4609-a5fc-6ff8eb63c186",
		PendingDays: "20",
		Sequence:    "919c82d4-8046-4722-9094-35c3c6524cff",
	}

	actual, err := cmk.ScheduleDeletion(client.ServiceClient(), scheduleDeletiontOps).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ScheduleDeletionResponse, actual)
}

func TestCancelDeletion(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCancelDeletionSuccessfully(t)

	var cancelDeletionOps = cmk.CancelDeletionOpts{
		KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
		Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
	}

	actual, err := cmk.CancelDeletion(client.ServiceClient(), cancelDeletionOps).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CancelDeletionResponse, actual)
}

func TestUpdateAlias(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateAliasSuccessfully(t)

	actual, err := cmk.UpdateAlias(client.ServiceClient(), cmk.UpdateAliasOpts{
		KeyId:    "e966a300-0c34-4a31-86e1-e67d13e6426a",
		KeyAlias: "TestABC",
		Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateAliasResponse, actual)
}

func TestUpdateDescription(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateDescriptionSuccessfully(t)

	actual, err := cmk.UpdateDescription(client.ServiceClient(), cmk.UpdateDescriptionOpts{
		KeyId:          "e966a300-0c34-4a31-86e1-e67d13e6426a",
		KeyDescription: "TestABC",
		Sequence:       "919c82d4-8046-4722-9094-35c3c6524cff",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateDescriptionResponse, actual)
}

func TestQuotas(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleQuotasSuccessfully(t)

	actual, err := cmk.Quotas(client.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &QuotasResponse, actual)
}
