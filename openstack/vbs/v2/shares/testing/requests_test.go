package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/vbs/v2/common"
	"github.com/huaweicloud/golangsdk/openstack/vbs/v2/shares"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestListShared(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/os-vendor-backup-sharing/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, listResponse)
	})

	actual, err := shares.List(fake.ServiceClient(), shares.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract shares: %v", err)
	}

	expected := []shares.Share{
		{
			BackupID:      "87566ed6-72cb-4053-aa6e-6f6216b3d507",
			ToProjectID:   "91d687759aed45d28b5f6084bc2fa8ad",
			FromProjectID: "17fbda95add24720a4038ba4b1c705ed",
			ID:            "ac0fb374-a288-4399-ac63-cc080a13a2ee",
			Backup: shares.Backup{Status: "available",
				ObjectCount:      0,
				Container:        "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
				Name:             "c2c-test-buckup",
				AvailabilityZone: "eu-de-01",
				SnapshotID:       "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
				VolumeID:         "5024a06e-6990-4f12-9dcc-8fe26b01a710",
				ID:               "87566ed6-72cb-4053-aa6e-6f6216b3d507",
				Size:             10},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestGetShared(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleGetSuccessfully(t)

	s, err := shares.Get(fake.ServiceClient(), "ac0fb374-a288-4399-ac63-cc080a13a2ee").ExtractShare()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, &shares.Share{
		BackupID:      "87566ed6-72cb-4053-aa6e-6f6216b3d507",
		ToProjectID:   "91d687759aed45d28b5f6084bc2fa8ad",
		FromProjectID: "17fbda95add24720a4038ba4b1c705ed",
		ID:            "ac0fb374-a288-4399-ac63-cc080a13a2ee",
		Backup: shares.Backup{Status: "available",
			ObjectCount:      0,
			Container:        "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
			Name:             "c2c-test-buckup",
			AvailabilityZone: "eu-de-01",
			SnapshotID:       "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
			VolumeID:         "5024a06e-6990-4f12-9dcc-8fe26b01a710",
			ID:               "87566ed6-72cb-4053-aa6e-6f6216b3d507",
			Size:             10},
	})
}

func TestCreateShared(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t, CreateOutput)
	createOpts := shares.CreateOpts{
		BackupID:     "87566ed6-72cb-4053-aa6e-6f6216b3d507",
		ToProjectIDs: []string{"91d687759aed45d28b5f6084bc2fa8ad"}}
	actual, err := shares.Create(fake.ServiceClient(), createOpts).Extract()

	th.AssertNoErr(t, err)

	expected := CreateExpected
	th.AssertDeepEquals(t, expected, actual)
}

func TestDeleteShared(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleDeleteSuccessfully(t)

	deleteOpts := shares.DeleteOpts{
		IsBackupID: true,
	}
	result := shares.Delete(fake.ServiceClient(), "87566ed6-72cb-4053-aa6e-6f6216b3d507", deleteOpts)
	th.AssertNoErr(t, result.Err)
}
