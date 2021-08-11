package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/mrs/v2/jobs"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"job_submit_result": {
		"job_id": "3c0cf394-5da2-46a7-92df-795d998edea7",
		"state": "COMPLETE"
	}
}`

	expectedGetResponse = `
{
	"job_detail": {
		"app_id": "application_1627718181153_0003",
		"arguments": "[--class, com.huawei.bigdata.spark.examples.DriverBehavior, --master, yarn-cluster, obs://obs-demo-analysis-tf/program/driver_behavior.jar, ACCESS_KEY, SECRET_KEY, 1, obs://obs-demo-analysis-tf/input, obs://obs-demo-analysis-tf/output]",
		"job_id": "3c0cf394-5da2-46a7-92df-795d998edea7",
		"job_name": "terraform_test",
		"job_progress": 100,
		"job_result": "FAILED",
		"job_state": "FINISHED",
		"job_type": "SparkSubmit",
		"launcher_id": "application_1627718181153_0002",
		"started_time": 1627719118083,
		"submitted_time": 1627719066265,
		"user": "terraform"
	}
}`

	expectedListResponse = `
{
	"total_record":1,
	"job_list":[
		{
			"app_id": "application_1627718181153_0003",
			"arguments": "[--class, com.huawei.bigdata.spark.examples.DriverBehavior, --master, yarn-cluster, obs://obs-demo-analysis-tf/program/driver_behavior.jar, ACCESS_KEY, SECRET_KEY, 1, obs://obs-demo-analysis-tf/input, obs://obs-demo-analysis-tf/output]",
			"job_id": "3c0cf394-5da2-46a7-92df-795d998edea7",
			"job_name": "terraform_test",
			"job_progress": 100,
			"job_result": "FAILED",
			"job_state": "FINISHED",
			"job_type": "SparkSubmit",
			"launcher_id": "application_1627718181153_0002",
			"started_time": 1627719118083,
			"submitted_time": 1627719066265,
			"user": "terraform"
		}
	]
}`
)

var (
	createOpts = &jobs.CreateOpts{
		JobName: "terraform_test",
		JobType: "SparkSubmit",
		Arguments: []string{
			"--class",
			"com.huawei.bigdata.spark.examples.DriverBehavior",
			"--master",
			"yarn-cluster",
			"obs://obs-demo-analysis-tf/program/driver_behavior.jar",
			"ACCESS_KEY",
			"SECRET_KEY",
			"1",
			"obs://obs-demo-analysis-tf/input",
			"obs://obs-demo-analysis-tf/output",
		},
	}

	expectedCreateResponseData = &jobs.CreateResp{
		JobSubmitResult: jobs.JobResp{
			JobId: "3c0cf394-5da2-46a7-92df-795d998edea7",
			State: "COMPLETE",
		},
	}

	expectedGetResponseData = &jobs.Job{
		AppId:         "application_1627718181153_0003",
		Arguments:     "[--class, com.huawei.bigdata.spark.examples.DriverBehavior, --master, yarn-cluster, obs://obs-demo-analysis-tf/program/driver_behavior.jar, ACCESS_KEY, SECRET_KEY, 1, obs://obs-demo-analysis-tf/input, obs://obs-demo-analysis-tf/output]",
		JobId:         "3c0cf394-5da2-46a7-92df-795d998edea7",
		JobName:       "terraform_test",
		JobProgress:   100,
		JobResult:     "FAILED",
		JobState:      "FINISHED",
		JobType:       "SparkSubmit",
		LauncherId:    "application_1627718181153_0002",
		StartedTime:   1627719118083,
		SubmittedTime: 1627719066265,
		User:          "terraform",
	}

	expectedListResponseData = []jobs.Job{
		{
			AppId:         "application_1627718181153_0003",
			Arguments:     "[--class, com.huawei.bigdata.spark.examples.DriverBehavior, --master, yarn-cluster, obs://obs-demo-analysis-tf/program/driver_behavior.jar, ACCESS_KEY, SECRET_KEY, 1, obs://obs-demo-analysis-tf/input, obs://obs-demo-analysis-tf/output]",
			JobId:         "3c0cf394-5da2-46a7-92df-795d998edea7",
			JobName:       "terraform_test",
			JobProgress:   100,
			JobResult:     "FAILED",
			JobState:      "FINISHED",
			JobType:       "SparkSubmit",
			LauncherId:    "application_1627718181153_0002",
			StartedTime:   1627719118083,
			SubmittedTime: 1627719066265,
			User:          "terraform",
		},
	}

	deleteOpts = jobs.DeleteOpts{
		JobIds: []string{
			"3c0cf394-5da2-46a7-92df-795d998edea7",
		},
	}
)

func handleV2JobCreate(t *testing.T) {
	th.Mux.HandleFunc("/clusters/986416ea-e26b-40f1-b371-cd7be87376a2/job-executions",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleV2JobGet(t *testing.T) {
	th.Mux.HandleFunc("/clusters/986416ea-e26b-40f1-b371-cd7be87376a2/job-executions"+
		"/3c0cf394-5da2-46a7-92df-795d998edea7", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedGetResponse)
	})
}

func handleV2JobList(t *testing.T) {
	th.Mux.HandleFunc("/clusters/986416ea-e26b-40f1-b371-cd7be87376a2/job-executions",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListResponse)
		})
}

func handleV2JobDelete(t *testing.T) {
	th.Mux.HandleFunc("/clusters/986416ea-e26b-40f1-b371-cd7be87376a2/job-executions/batch-delete",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
		})
}
