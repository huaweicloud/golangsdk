package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/vbs/v2/backups"
	fake "github.com/huaweicloud/golangsdk/openstack/vbs/v2/common"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

// ListExpected represents the expected object from a List request.
var ListExpected = []backups.Backup{
	{
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
	},
}

// FullListOutput represents the response body from a List request.
const FullListOutput = `
{
    "backups": [
        {
            "status": "available",
            "object_count": 0,
            "os-bak-tenant-attr:tenant_id": "17fbda95add24720a4038ba4b1c705ed",
            "container": "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
            "name": "c2c-test-buckup",
            "links": [
                {
                    "href": "https://vbs.eu-de.otc.t-systems.com/v2/17fbda95add24720a4038ba4b1c705ed/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507",
                    "rel": "self"
                },
                {
                    "href": "https://vbs.eu-de.otc.t-systems.com/17fbda95add24720a4038ba4b1c705ed/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507",
                    "rel": "bookmark"
                }
            ],
            "availability_zone": "eu-de-01",
            "has_dependent_backups": false,
            "snapshot_id": "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
            "volume_id": "5024a06e-6990-4f12-9dcc-8fe26b01a710",
            "is_incremental": false,
            "id": "87566ed6-72cb-4053-aa6e-6f6216b3d507",
            "size": 10
        }
	]
}`

// HandleListSuccessfully creates an HTTP handler at `/backups/detail` on the test handler mux
// that responds with a `List` response.
func HandleListSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/backups/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
	})
}

// getResponse represents the response body from a Get request.
var getResponse = `{
    "backup": {
            "status": "available",
            "object_count": 0,
            "os-bak-tenant-attr:tenant_id": "17fbda95add24720a4038ba4b1c705ed",
            "container": "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
            "name": "c2c-test-buckup",
            "links": [
                {
                    "href": "https://vbs.eu-de.otc.t-systems.com/v2/17fbda95add24720a4038ba4b1c705ed/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507",
                    "rel": "self"
                },
                {
                    "href": "https://vbs.eu-de.otc.t-systems.com/17fbda95add24720a4038ba4b1c705ed/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507",
                    "rel": "bookmark"
                }
            ],
            "availability_zone": "eu-de-01",
            "has_dependent_backups": false,
            "snapshot_id": "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
            "volume_id": "5024a06e-6990-4f12-9dcc-8fe26b01a710",
            "is_incremental": false,
            "id": "87566ed6-72cb-4053-aa6e-6f6216b3d507",
            "size": 10
        }
}`

// HandleGetSuccessfully creates an HTTP handler at `/17fbda95add24720a4038ba4b1c705ed/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507`
// on the test handler mux that responds with a `Get` response.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getResponse)
	})
}

// CreateExpected represents the expected object from a Create request.
var CreateExpected = &backups.JobResponse{
	JobID: "ff8080826576401e01657b99fc444986",
}

// CreateOutput represents the response body from a Create request.
const CreateOutput = `{
    "job_id": "ff8080826576401e01657b99fc444986"
}`

// HandleCreateSuccessfully creates an HTTP handler at `/backups` on the test handler mux
// that responds with a `Create` response.
func HandleCreateSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/cloudbackups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
	})
}

// HandleDeleteSuccessfully creates an HTTP handler at `/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507`
// on the test handler mux that responds with a `Delete` response.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}

// CreateExpected represents the expected object from a Create request.
var RestoreExpected = &backups.BackupRestoreInfo{
	BackupId:   "87566ed6-72cb-4053-aa6e-6f6216b3d507",
	VolumeName: "c2c-test-disk",
	VolumeId:   "5024a06e-6990-4f12-9dcc-8fe26b01a710",
}

// RestoreOutput represents the response body from a Create request.
const RestoreOutput = `
{
    "restore": {
        "backup_id": "87566ed6-72cb-4053-aa6e-6f6216b3d507",
        "volume_name": "c2c-test-disk",
        "volume_id": "5024a06e-6990-4f12-9dcc-8fe26b01a710"
    }
}`

// HandleRestoreSuccessfully creates an HTTP handler at `/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507/restore` on the test handler mux
// that responds with a `Create` response.
func HandleRestoreSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/backups/87566ed6-72cb-4053-aa6e-6f6216b3d507/restore", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, output)
	})
}
