package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/identity/v3/groups"
	"github.com/huaweicloud/golangsdk/openstack/identity/v3/projects"
	"github.com/huaweicloud/golangsdk/openstack/identity/v3/users"
	"github.com/huaweicloud/golangsdk/pagination"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestListUsers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListUsersSuccessfully(t)

	count := 0
	err := users.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := users.ExtractUsers(page)
		th.AssertNoErr(t, err)

		th.CheckDeepEquals(t, ExpectedUsersSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, count, 1)
}

func TestListUsersAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListUsersSuccessfully(t)

	allPages, err := users.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := users.ExtractUsers(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedUsersSlice, actual)
}

func TestGetUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetUserSuccessfully(t)

	actual, err := users.Get(client.ServiceClient(), "9fe1d3").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondUser, *actual)
}

func TestCreateUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateUserSuccessfully(t)

	iTrue := true
	createOpts := users.CreateOpts{
		Name:             "jsmith",
		DomainID:         "1789d1",
		Enabled:          &iTrue,
		Password:         "secretsecret",
		DefaultProjectID: "263fd9",
	}

	actual, err := users.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondUser, *actual)
}

func TestCreateNoOptionsUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateNoOptionsUserSuccessfully(t)

	iTrue := true
	createOpts := users.CreateOpts{
		Name:             "jsmith",
		DomainID:         "1789d1",
		Enabled:          &iTrue,
		Password:         "secretsecret",
		DefaultProjectID: "263fd9",
	}

	actual, err := users.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondUserNoOptions, *actual)
}

func TestUpdateUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateUserSuccessfully(t)

	iFalse := false
	updateOpts := users.UpdateOpts{
		Enabled: &iFalse,
	}

	actual, err := users.Update(client.ServiceClient(), "9fe1d3", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondUserUpdated, *actual)
}

func TestDeleteUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteUserSuccessfully(t)

	res := users.Delete(client.ServiceClient(), "9fe1d3")
	th.AssertNoErr(t, res.Err)
}

func TestListUserGroups(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListUserGroupsSuccessfully(t)
	allPages, err := users.ListGroups(client.ServiceClient(), "9fe1d3").AllPages()
	th.AssertNoErr(t, err)
	actual, err := groups.ExtractGroups(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedGroupsSlice, actual)
}

func TestListUserProjects(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListUserProjectsSuccessfully(t)
	allPages, err := users.ListProjects(client.ServiceClient(), "9fe1d3").AllPages()
	th.AssertNoErr(t, err)
	actual, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedProjectsSlice, actual)
}

func TestListInGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListInGroupSuccessfully(t)

	iTrue := true
	listOpts := users.ListOpts{
		Enabled: &iTrue,
	}

	allPages, err := users.ListInGroup(client.ServiceClient(), "ea167b", listOpts).AllPages()
	th.AssertNoErr(t, err)
	actual, err := users.ExtractUsers(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedUsersSlice, actual)
}
