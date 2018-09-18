package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/vbs/v2/common"
	"github.com/huaweicloud/golangsdk/openstack/vbs/v2/shares"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

// listResponse represents the response body from a List request.
var listResponse = `{
    "shared": [
        {
        "backup_id": "87566ed6-72cb-4053-aa6e-6f6216b3d507",
        "to_project_id": "91d687759aed45d28b5f6084bc2fa8ad",
        "from_project_id": "17fbda95add24720a4038ba4b1c705ed",
        "backup": {
            "status": "available",
            "object_count": 0,
            "container": "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
            "name": "c2c-test-buckup",
            "availability_zone": "eu-de-01",
            "snapshot_id": "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
            "volume_id": "5024a06e-6990-4f12-9dcc-8fe26b01a710",
            "id": "87566ed6-72cb-4053-aa6e-6f6216b3d507",
            "size": 10
        },
        "id": "ac0fb374-a288-4399-ac63-cc080a13a2ee"
    }
  ]
}`

// getResponse represents the response body from a Get request.
var getResponse = `{
    "shared": {
        "backup_id": "87566ed6-72cb-4053-aa6e-6f6216b3d507",
        "to_project_id": "91d687759aed45d28b5f6084bc2fa8ad",
        "from_project_id": "17fbda95add24720a4038ba4b1c705ed",
        "backup": {
            "status": "available",
            "object_count": 0,
            "container": "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
            "name": "c2c-test-buckup",
            "availability_zone": "eu-de-01",
            "snapshot_id": "a704c75f-f0d1-4efa-9fd6-7557fe1ee8d3",
            "volume_id": "5024a06e-6990-4f12-9dcc-8fe26b01a710",
            "id": "87566ed6-72cb-4053-aa6e-6f6216b3d507",
            "size": 10
        },
        "id": "ac0fb374-a288-4399-ac63-cc080a13a2ee"
    }
}`

// HandleGetSuccessfully creates an HTTP handler at `/os-vendor-backup-sharing/ac0fb374-a288-4399-ac63-cc080a13a2ee`
// on the test handler mux that responds with a `Get` response.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-vendor-backup-sharing/ac0fb374-a288-4399-ac63-cc080a13a2ee", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getResponse)
	})
}

// CreateExpected represents the expected object from a Create request.
var CreateExpected = []shares.Share{{
	BackupID:      "87566ed6-72cb-4053-aa6e-6f6216b3d507",
	ToProjectID:   "91d687759aed45d28b5f6084bc2fa8ad",
	FromProjectID: "17fbda95add24720a4038ba4b1c705ed",
	ID:            "34c38ce7-f35c-44f2-a8d8-8d4ebab0cfbb",
},
}

// CreateOutput represents the response body from a Create request.
const CreateOutput = `
{
    "shared": [
        {
            "backup_id": "87566ed6-72cb-4053-aa6e-6f6216b3d507",
            "to_project_id": "91d687759aed45d28b5f6084bc2fa8ad",
            "from_project_id": "17fbda95add24720a4038ba4b1c705ed",
            "id": "34c38ce7-f35c-44f2-a8d8-8d4ebab0cfbb"
        }
    ]
}`

// HandleCreateSuccessfully creates an HTTP handler at `/os-vendor-backup-sharing` on the test handler mux
// that responds with a `Create` response.
func HandleCreateSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/os-vendor-backup-sharing", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
	})
}

// HandleDeleteSuccessfully creates an HTTP handler at `/os-vendor-backup-sharing/87566ed6-72cb-4053-aa6e-6f6216b3d507`
// on the test handler mux that responds with a `Delete` response.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-vendor-backup-sharing/87566ed6-72cb-4053-aa6e-6f6216b3d507", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
	})
}
