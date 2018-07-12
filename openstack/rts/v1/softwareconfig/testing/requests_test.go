package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rts/v1/softwareconfig"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateSoftwareConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t, CreateOutput)
	createOpts := softwareconfig.CreateOpts{Name: "test-cong"}
	actual, err := softwareconfig.Create(fake.ServiceClient(), createOpts).Extract()

	th.AssertNoErr(t, err)

	expected := CreateExpected
	th.AssertDeepEquals(t, expected, actual)
}

func TestListSoftwareConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, FullListOutput)

	actual, err := softwareconfig.List(fake.ServiceClient(), softwareconfig.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract vpcs: %v", err)
	}
	th.AssertDeepEquals(t, ListExpected, actual)
	th.AssertNoErr(t, err)
}

func TestGetSoftwareConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t, GetOutput)

	actual, err := softwareconfig.Get(fake.ServiceClient(), "e0be7e37-a581-4b24-bfb1-df4f3048c090").Extract()
	th.AssertNoErr(t, err)

	expected := GetExpected
	th.AssertDeepEquals(t, expected, actual)
}

func TestDeleteSoftwareConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := softwareconfig.Delete(fake.ServiceClient(), "e2fe5553-a481-4549-9d0f-e208de3d98d1").ExtractErr()
	th.AssertNoErr(t, err)
}
