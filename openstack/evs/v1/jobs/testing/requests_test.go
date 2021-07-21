package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/evs/v1/jobs"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const expectedResponse = `
{
    "begin_time": "2021-07-21T06:53:00.431Z",
    "end_time": "2021-07-21T06:53:07.443Z",
    "entities": {
        "name": "volume_1",
        "size": 50,
        "volume_id": "4c3b7a01-989a-4c34-876c-4699f3766457",
        "volume_type": "SSD"
    },
    "error_code": null,
    "fail_reason": null,
    "job_id": "ff8080817aa5637b017ac7d6ca9005d5",
    "job_type": "createVolume",
    "status": "SUCCESS"
}`

var expectedJobResponseData = &jobs.Job{
	Status: "SUCCESS",
	Entities: jobs.JobEntity{
		Name:       "volume_1",
		Size:       50,
		VolumeID:   "4c3b7a01-989a-4c34-876c-4699f3766457",
		VolumeType: "SSD",
	},
	JobID:     "ff8080817aa5637b017ac7d6ca9005d5",
	JobType:   "createVolume",
	BeginTime: "2021-07-21T06:53:00.431Z",
	EndTime:   "2021-07-21T06:53:07.443Z",
}

func handleGetJobDetail(t *testing.T) {
	th.Mux.HandleFunc("/jobs/ff8080817aa5637b017ac7d6ca9005d5",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedResponse)
		})
}

func TestGetJobDetail(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleGetJobDetail(t)

	actual, err := jobs.GetJobDetails(client.ServiceClient(), "ff8080817aa5637b017ac7d6ca9005d5").ExtractJob()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedJobResponseData, actual)
}
