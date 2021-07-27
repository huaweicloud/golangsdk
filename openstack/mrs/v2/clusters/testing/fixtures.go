package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/mrs/v2/clusters"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"cluster_id": "4cb9f483-4464-4b04-b664-b79bf904f441"
}`
)

var (
	createOpts = &clusters.CreateOpts{
		ClusterName:          "terraform_test",
		AvailabilityZone:     "cn-north-4a",
		ClusterType:          "ANALYSIS",
		ClusterVersion:       "MRS 1.9.2",
		Components:           "Tez,Hive,Hadoop",
		EnterpriseProjectId:  "e9ee3f48-f097-406a-aa74-cfece0af3e31",
		LogCollection:        golangsdk.IntToPointer(1),
		LoginMode:            "PASSWORD",
		ManagerAdminPassword: "Terraform!123",
		NodeRootPassword:     "Terraform!123",
		Region:               "cn-north-4",
		SafeMode:             "KERBEROS",
		SubnetId:             "ea6a5175-2ff1-4f72-b5b5-1fe82cea8c19",
		VpcName:              "terraform_test",
		NodeGroups: []clusters.NodeGroupOpts{
			{
				GroupName:       "master_node_default_group",
				NodeSize:        "c6.2xlarge.4.linux.bigdata",
				NodeNum:         2,
				DataVolumeCount: golangsdk.IntToPointer(1),
				RootVolume: &clusters.Volume{
					Size: 600,
					Type: "SAS",
				},
				DataVolume: &clusters.Volume{
					Size: 600,
					Type: "SAS",
				},
			},
			{
				GroupName:       "core_node_analysis_group",
				NodeSize:        "c6.2xlarge.4.linux.bigdata",
				NodeNum:         2,
				DataVolumeCount: golangsdk.IntToPointer(1),
				RootVolume: &clusters.Volume{
					Size: 600,
					Type: "SAS",
				},
				DataVolume: &clusters.Volume{
					Size: 600,
					Type: "SAS",
				},
			},
			{
				GroupName:       "task_node_analysis_group",
				NodeSize:        "c6.2xlarge.4.linux.bigdata",
				NodeNum:         1,
				DataVolumeCount: golangsdk.IntToPointer(1),
				RootVolume: &clusters.Volume{
					Size: 600,
					Type: "SAS",
				},
				DataVolume: &clusters.Volume{
					Size: 600,
					Type: "SAS",
				},
			},
		},
	}

	expectedCreateResponseData = &clusters.Cluster{
		ID: "4cb9f483-4464-4b04-b664-b79bf904f441",
	}
)

func handleV2ClusterCreate(t *testing.T) {
	th.Mux.HandleFunc("/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}
