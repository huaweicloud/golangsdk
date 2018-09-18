package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/vbs/v2/backups"
	fake "github.com/huaweicloud/golangsdk/openstack/vbs/v2/common"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestListBackup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, FullListOutput)

	actual, err := backups.List(fake.ServiceClient(), backups.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract backups: %v", err)
	}
	th.AssertDeepEquals(t, ListExpected, actual)
	th.AssertNoErr(t, err)
}

func TestGetBackup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleGetSuccessfully(t)

	s, err := backups.Get(fake.ServiceClient(), "87566ed6-72cb-4053-aa6e-6f6216b3d507").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, &backups.Backup{
		Name:             "c2c-test-buckup",
		Id:               "87566ed6-72cb-4053-aa6e-6f6216b3d507",
		Status:           "available",
		ObjectCount:      0,
		TenantId:         "17fbda95add24720a4038ba4b1c705ed",
		Container:        "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
		AvailabilityZone: "eu-de-01",
		DependentBackups: false,
		SnapshotId:       "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
		VolumeId:         "5024a06e-6990-4f12-9dcc-8fe26b01a710",
		Incremental:      false,
		Size:             10,
		Links: []golangsdk.Link{
			{
				Href: "https://vbs.eu-de.otc.t-systems.com/v2/17fbda95add24720a4038ba4b1c705ed/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507",
				Rel:  "self",
			},
			{
				Href: "https://vbs.eu-de.otc.t-systems.com/17fbda95add24720a4038ba4b1c705ed/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507",
				Rel:  "bookmark",
			},
		},
	})
}

func TestCreateBackup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t, CreateOutput)

	createOpts := backups.CreateOpts{
		Name:     "backup_test",
		VolumeId: "5024a06e-6990-4f12-9dcc-8fe26b01a710",
	}
	actual, err := backups.Create(fake.ServiceClient(), createOpts).ExtractJobResponse()
	th.AssertNoErr(t, err)

	expected := CreateExpected
	th.AssertDeepEquals(t, expected, actual)
}

func TestDeleteBackup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	result := backups.Delete(fake.ServiceClient(), "87566ed6-72cb-4053-aa6e-6f6216b3d507")
	th.AssertNoErr(t, result.Err)
}

func TestCreateRestore(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRestoreSuccessfully(t, RestoreOutput)

	restoreOpts := backups.BackupRestoreOpts{
		VolumeId: "5024a06e-6990-4f12-9dcc-8fe26b01a710",
	}
	actual, err := backups.CreateBackupRestore(fake.ServiceClient(), "87566ed6-72cb-4053-aa6e-6f6216b3d507", restoreOpts).ExtractBackupRestore()
	th.AssertNoErr(t, err)

	expected := RestoreExpected
	th.AssertDeepEquals(t, expected, actual)
}
