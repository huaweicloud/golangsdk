package testing

import (
	"fmt"
	"net/http"
	"testing"

	"time"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/csbs/v1/backup"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(backupEndpoint+"/"+checkpoint_item_id, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	s, err := backup.Get(fake.ServiceClient(), checkpoint_item_id).ExtractBackup()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "7b99acfd-18c3-4f26-9d39-b4ebd2ea3e12", s.Id)
	th.AssertEquals(t, "backup-c2c", s.Name)
	th.AssertEquals(t, "available", s.Status)
	th.AssertEquals(t, "2eefe592-8424-4778-8d0d-962c8a5dd6a4", s.CheckpointId)
	th.AssertEquals(t, "backup des", s.Description)
	th.AssertEquals(t, "OS::Nova::Server", s.ResourceType)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/providers/fc4d5750-22e7-4798-8a46-f48f62c4c1da/resources/f8ddc472-cf00-4384-851e-5f2a68c33762/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, createResponse)
		})

	options := &backup.CreateOpts{
		BackupName:  "c2c-backup",
		Description: "mybackup"}
	n, err := backup.Create(fake.ServiceClient(), "f8ddc472-cf00-4384-851e-5f2a68c33762", options).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, n.Id, "92dba83d-cc6f-4883-a20d-de6934510b7e")
	th.AssertEquals(t, n.Status, "protecting")
	th.AssertEquals(t, n.ProtectionPlan.Id, "fake_b94f8b46-b0a1-485a-ad5b-9f8876b85495")
	th.AssertEquals(t, n.ProtectionPlan.Name, "server protect plan for f8ddc472-cf00-4384-851e-5f2a68c33762")
}

func TestQueryResourceCapability(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/providers/fc4d5750-22e7-4798-8a46-f48f62c4c1da/resources/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, queryRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, queryResponse)
		})

	options := &backup.ResourceBackupCapOpts{CheckProtectable: []backup.ResourceCapQueryParams{
		{ResourceId: "069e678a-f1d1-4a38-880b-459bde82fcc6",
			ResourceType: "OS::Nova::Server"}}}
	n, err := backup.QueryResourceBackupCapability(fake.ServiceClient(), options).ExtractQueryResponse()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, n[0].ResourceType, "OS::Nova::Server")
	th.AssertEquals(t, n[0].ResourceId, "069e678a-f1d1-4a38-880b-459bde82fcc6")
	th.AssertEquals(t, n[0].Result, true)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/providers/fc4d5750-22e7-4798-8a46-f48f62c4c1da/checkpoints/fc4d5750-22e7-4798-8a46-f48f62c4c1da",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			w.WriteHeader(http.StatusOK)
		})

	result := backup.Delete(fake.ServiceClient(), "fc4d5750-22e7-4798-8a46-f48f62c4c1da")
	th.AssertNoErr(t, result.Err)
}

func TestList(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/checkpoint_items", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, listResponse)
	})

	actual, err := backup.List(fake.ServiceClient(), backup.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract backups: %v", err)
	}

	var FinishedAt, _ = time.Parse(golangsdk.RFC3339MilliNoZ, "2018-08-14T08:31:08.720800")
	expected := []backup.Backup{
		{
			Status: "available",
			VMMetadata: backup.VMMetadata{
				Eip:              "80.158.17.102",
				CloudServiceType: "QEMU",
				Ram:              8192,
				Vcpus:            4,
				RegionName:       "",
				PrivateIp:        "192.168.0.209",
				Disk:             0,
				ImageType:        "gold",
			},
			Name:         "backup-c2c",
			ResourceId:   "f8ddc472-cf00-4384-851e-5f2a68c33762",
			CheckpointId: "2eefe592-8424-4778-8d0d-962c8a5dd6a4",
			ExtendInfo: backup.ExtendInfo{
				AutoTrigger:          false,
				SpaceSavingRatio:     2,
				ResourceName:         "ecs-ggao",
				FailReason:           "",
				ResourceAz:           "eu-de-02",
				ImageType:            "backup",
				FinishedAt:           FinishedAt,
				AverageSpeed:         19,
				CopyStatus:           "na",
				Incremental:          false,
				TaskId:               "1afcab08-9f97-11e8-9526-286ed488ca8c",
				HypervisorType:       "QEMU",
				SupportedRestoreMode: "backup",
				Progress:             100,
				Supportlld:           true,
				FailOp:               "",
				ResourceType:         "OS::Nova::Server",
				Size:                 146184,
			},
			Id:           "7b99acfd-18c3-4f26-9d39-b4ebd2ea3e12",
			ResourceType: "OS::Nova::Server",
			Description:  "backup des",
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}
