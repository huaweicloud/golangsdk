package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rts/v1/softwaredeployment"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateSoftwareDeployment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t, CreateOutput)

	createOpts := softwaredeployment.CreateOpts{
		ConfigId:     "031e3891-d183-4f8e-a836-589a5dce541c",
		ServerId:     "b7653627-9b2a-4b61-b18f-e20ea88ee924",
		TenantId:     "17fbda95add24720a4038ba4b1c705ed",
		Status:       "IN_PROGRESS",
		Action:       "CREATE",
		StatusReason: "Deploy data available",
	}

	actual, err := softwaredeployment.Create(fake.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	expected := CreateExpected
	th.CheckDeepEquals(t, expected, actual)
}

func TestListSoftwareDeployment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, ListOutput)

	listOpts := softwaredeployment.ListOpts{}
	actual, err := softwaredeployment.List(fake.ServiceClient(), listOpts)

	th.AssertNoErr(t, err)
	expected := ListExpected
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetSoftwareDeployment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t, GetOutput)
	actual, err := softwaredeployment.Get(fake.ServiceClient(), "43489279-7b12-4fc5-90ed-320f29e89419").Extract()

	th.AssertNoErr(t, err)
	expected := GetExpected
	th.CheckDeepEquals(t, expected, actual)
}

func TestUpdateSoftwareDeployment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t, UpdateOutput)
	ouput := map[string]interface{}{
		"deploy_stdout":      "Writing to /tmp/baaaaa\nWritten to /tmp/baaaaa\n",
		"deploy_stderr":      "+ echo Writing to /tmp/baaaaa\n+ echo fooooo\n+ cat /tmp/baaaaa\n+ echo -n The file /tmp/baaaaa contains fooooo for server ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5 during CREATE\n+ echo Written to /tmp/baaaaa\n+ echo Output to stderr\nOutput to stderr\n",
		"deploy_status_code": "0", "result": "The file /tmp/baaaaa contains fooooo for server ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5 during CREATE"}
	updateOpts := softwaredeployment.UpdateOpts{
		Status:       "COMPLETE",
		ConfigId:     "a6ff3598-f2e0-4111-81b0-aa3e1cac2529",
		OutputValues: ouput,
		StatusReason: "Outputs received"}
	actual, err := softwaredeployment.Update(fake.ServiceClient(), "43489279-7b12-4fc5-90ed-320f29e89419", updateOpts).Extract()

	th.AssertNoErr(t, err)
	expected := UpdateExpected
	th.CheckDeepEquals(t, expected, actual)
}

func TestDeleteSoftwareConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := softwaredeployment.Delete(fake.ServiceClient(), "43489279-7b12-4fc5-90ed-320f29e89419").ExtractErr()
	th.AssertNoErr(t, err)
}
