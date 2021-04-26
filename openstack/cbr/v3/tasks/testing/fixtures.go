package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/cbr/v3/tasks"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedGetResponse = `
{
  "operation_log" : {
    "status" : "success",
    "provider_id" : "0daac4c5-6707-4851-97ba-169e36266b66",
    "checkpoint_id" : "b432511f-d889-428f-8b0e-5f47c524c6b6",
    "updated_at" : "2019-05-23T14:35:23.584418",
    "error_info" : {
      "message" : "",
      "code" : ""
    },
    "started_at" : "2019-05-23T14:31:36.007230",
    "id" : "4827f2da-b008-4507-ab7d-42d0df5ed912",
    "extra_info" : {
      "resource" : {
        "type" : "OS::Nova::Server",
        "id" : "1dab32fa-ebf2-415a-ab0b-eabe6353bc86",
        "name" : "ECS-0001"
      },
      "backup" : {
        "backup_name" : "manualbk_1234",
        "backup_id" : "0e5d0ef6-7f0a-4890-b98c-cb12490e31c1"
      },
      "common" : {
        "progress" : 100,
        "request_id" : "req-cdb98cc4-e87b-4f40-9b4a-57ec036620bc"
      }
    },
    "ended_at" : "2019-05-23T14:35:23.511155",
    "created_at" : "2019-05-23T14:31:36.039365",
    "operation_type" : "backup",
    "project_id" : "04f1829c788037ac2fb8c01eb2b04b95"
  }
}`

	expectedListResponse = `
{
  "operation_logs" : [ {
    "status" : "success",
    "provider_id" : "0daac4c5-6707-4851-97ba-169e36266b66",
    "checkpoint_id" : "b432511f-d889-428f-8b0e-5f47c524c6b6",
    "updated_at" : "2019-05-23T14:35:23.584418",
    "error_info" : {
      "message" : "",
      "code" : ""
    },
    "started_at" : "2019-05-23T14:31:36.007230",
    "id" : "4827f2da-b008-4507-ab7d-42d0df5ed912",
    "extra_info" : {
      "resource" : {
        "type" : "OS::Nova::Server",
        "id" : "1dab32fa-ebf2-415a-ab0b-eabe6353bc86",
        "name" : "ECS-0001"
      },
      "backup" : {
        "backup_name" : "manualbk_1234",
        "backup_id" : "0e5d0ef6-7f0a-4890-b98c-cb12490e31c1"
      },
      "common" : {
        "progress" : 100,
        "request_id" : "req-cdb98cc4-e87b-4f40-9b4a-57ec036620bc"
      }
    },
    "ended_at" : "2019-05-23T14:35:23.511155",
    "created_at" : "2019-05-23T14:31:36.039365",
    "operation_type" : "backup",
    "project_id" : "04f1829c788037ac2fb8c01eb2b04b95"
  } ],
  "count" : 1
}`
)

var (
	expectedGetResponseData = &tasks.OperationLog{
		Status:       "success",
		ProviderID:   "0daac4c5-6707-4851-97ba-169e36266b66",
		CheckpointID: "b432511f-d889-428f-8b0e-5f47c524c6b6",
		UpdatedAt:    "2019-05-23T14:35:23.584418",
		ErrorInfo: tasks.OpErrorInfo{
			Message: "",
			Code:    "",
		},
		StartedAt: "2019-05-23T14:31:36.007230",
		ID:        "4827f2da-b008-4507-ab7d-42d0df5ed912",
		ExtraInfo: tasks.OpExtraInfo{
			Resource: tasks.Resource{
				Type: "OS::Nova::Server",
				ID:   "1dab32fa-ebf2-415a-ab0b-eabe6353bc86",
				Name: "ECS-0001",
			},
			Backup: tasks.OpExtendInfoBackup{
				BackupName: "manualbk_1234",
				BackupID:   "0e5d0ef6-7f0a-4890-b98c-cb12490e31c1",
			},
			Common: tasks.OpExtendInfoCommon{
				Progress:  100,
				RequestID: "req-cdb98cc4-e87b-4f40-9b4a-57ec036620bc",
			},
		},
		EndedAt:       "2019-05-23T14:35:23.511155",
		CreatedAt:     "2019-05-23T14:31:36.039365",
		OperationType: "backup",
		ProjectID:     "04f1829c788037ac2fb8c01eb2b04b95",
	}

	expectedListResponseData = &[]tasks.OperationLog{
		{
			Status:       "success",
			ProviderID:   "0daac4c5-6707-4851-97ba-169e36266b66",
			CheckpointID: "b432511f-d889-428f-8b0e-5f47c524c6b6",
			UpdatedAt:    "2019-05-23T14:35:23.584418",
			ErrorInfo: tasks.OpErrorInfo{
				Message: "",
				Code:    "",
			},
			StartedAt: "2019-05-23T14:31:36.007230",
			ID:        "4827f2da-b008-4507-ab7d-42d0df5ed912",
			ExtraInfo: tasks.OpExtraInfo{
				Resource: tasks.Resource{
					Type: "OS::Nova::Server",
					ID:   "1dab32fa-ebf2-415a-ab0b-eabe6353bc86",
					Name: "ECS-0001",
				},
				Backup: tasks.OpExtendInfoBackup{
					BackupName: "manualbk_1234",
					BackupID:   "0e5d0ef6-7f0a-4890-b98c-cb12490e31c1",
				},
				Common: tasks.OpExtendInfoCommon{
					Progress:  100,
					RequestID: "req-cdb98cc4-e87b-4f40-9b4a-57ec036620bc",
				},
			},
			EndedAt:       "2019-05-23T14:35:23.511155",
			CreatedAt:     "2019-05-23T14:31:36.039365",
			OperationType: "backup",
			ProjectID:     "04f1829c788037ac2fb8c01eb2b04b95",
		},
	}
)

func handleTaskGet(t *testing.T) {
	th.Mux.HandleFunc("/operation-logs/4827f2da-b008-4507-ab7d-42d0df5ed912",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleTaskList(t *testing.T) {
	th.Mux.HandleFunc("/operation-logs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}
