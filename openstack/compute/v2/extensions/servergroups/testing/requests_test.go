package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/compute/v2/extensions/servergroups"
	"github.com/huaweicloud/golangsdk/pagination"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	count := 0
	err := servergroups.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := servergroups.ExtractServerGroups(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedServerGroupSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := servergroups.Create(client.ServiceClient(), servergroups.CreateOpts{
		Name:     "test",
		Policies: []string{"anti-affinity"},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreatedServerGroup, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := servergroups.Get(client.ServiceClient(), "4d8c3732-a248-40ed-bebc-539a6ffd25c0").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &FirstServerGroup, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := servergroups.Delete(client.ServiceClient(), "616fb98f-46ca-475e-917e-2563e5a8cd19").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestAddMember(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddMemberSuccessfully(t)

	err := servergroups.UpdateMember(client.ServiceClient(), servergroups.MemberOpts{
		InstanceUUid: "d194d539-07b0-446e-b52c-e639e618e49d",
	}, "add_member", "616fb98f-46ca-475e-917e-2563e5a8cd19").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestRemoveMember(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRemoveMemberSuccessfully(t)

	err := servergroups.UpdateMember(client.ServiceClient(), servergroups.MemberOpts{
		InstanceUUid: "d194d539-07b0-446e-b52c-e639e618e49d",
	}, "remove_member", "616fb98f-46ca-475e-917e-2563e5a8cd19").ExtractErr()
	th.AssertNoErr(t, err)
}
