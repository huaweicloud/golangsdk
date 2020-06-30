package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/geminidb/v3/instances"
	"github.com/huaweicloud/golangsdk/pagination"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, CreateResponse)
	})

	options := instances.CreateGeminiDBOpts{
		Name: "test-cassandra-01",
		DataStore: instances.DataStore{
			Type:          "GeminiDB-Cassandra",
			Version:       "3.11",
			StorageEngine: "rocksDB",
		},
		Region:           "aaa",
		AvailabilityZone: "bbb",
		VpcId:            "674e9b42-cd8d-4d25-a2e6-5abcc565b961",
		SubnetId:         "f1df08c5-71d1-406a-aff0-de435a51007b",
		SecurityGroupId:  "7aa51dbf-5b63-40db-9724-dad3c4828b58",
		Password:         "Test@123",
		Mode:             "Cluster",
		Flavor: []instances.FlavorOpt{
			{
				Num:      "3",
				Size:     500,
				Storage:  "ULTRAHIGH",
				SpecCode: "nosql.cassandra.4xlarge.4",
			},
		},
		BackupStrategy: &instances.BackupStrategyOpt{
			StartTime: "08:15-09:15",
			KeepDays:  "8",
		},
		EnterpriseProjectId: "0",
	}

	actual, err := instances.Create(client.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	expected := instances.CreateResponse{
		AvailabilityZone: "bbb",
		Flavor: []instances.Flavor{
			{
				Num:      "3",
				Size:     "500",
				Storage:  "ULTRAHIGH",
				SpecCode: "nosql.cassandra.4xlarge.4",
			},
		},
		JobId: "c010abd0-48cf-4fa8-8cbc-090f093eaa2f",
		BackupStrategy: instances.BackupStrategyOpt{
			StartTime: "08:15-09:15",
			KeepDays:  "8",
		},
	}

	expected.GeminiDBBase = instances.GeminiDBBase{
		Id:   "39b6a1a278844ac48119d86512e0000bin06",
		Name: "test-cassandra-01",
		DataStore: instances.DataStore{
			Type:          "GeminiDB-Cassandra",
			Version:       "3.11",
			StorageEngine: "rocksDB",
		},
		Status:              "creating",
		Region:              "aaa",
		VpcId:               "674e9b42-cd8d-4d25-a2e6-5abcc565b961",
		SubnetId:            "f1df08c5-71d1-406a-aff0-de435a51007b",
		SecurityGroupId:     "7aa51dbf-5b63-40db-9724-dad3c4828b58",
		Mode:                "Cluster",
		EnterpriseProjectId: "0",
	}
	th.AssertDeepEquals(t, expected, *actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/instances/4e8e5957", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
	res := instances.Delete(client.ServiceClient(), "4e8e5957")
	th.AssertNoErr(t, res.Err)
}

func TestGetAllInstances(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, AllInstancesResponse)
	})

	ExpectedAllInstancesResponse := instances.ListGeminiDBResponse{
		TotalCount: 2,
		Instances: []instances.GeminiDBInstance{
			{
				Port:              "8635",
				Engine:            " rocksDB ",
				DbUserName:        "rwuser",
				PayMode:           "0",
				TimeZone:          "",
				MaintenanceWindow: "02:00-06:00",
				Actions:           []string{"CREATE"},
				Groups: []instances.Groups{
					{
						Id:     "0b0ff12541794e1084f6827e424be2d6gr06",
						Status: "creating",
						Volume: instances.Volume{
							Size: "100",
							Used: "0.003",
						},
						Nodes: []instances.Nodes{
							{
								Id:               "233eaac9c6f245c0bb9c2d21eea12d1bno06",
								Name:             "GeminiDB-efa6_priam_node_2",
								Status:           "normal",
								PrivateIp:        "192.168.0.174",
								SpecCode:         "nosql.cassandra.xlarge.4",
								AvailabilityZone: "bbb",
							},
							{
								Id:               "d57d76d6320a4a7b86db82c317550c4ano02",
								Name:             "GeminiDB-efa6_priam_node_1",
								Status:           "normal",
								PrivateIp:        "192.168.0.39",
								SpecCode:         "nosql.cassandra.xlarge.4",
								AvailabilityZone: "bbb",
							},
							{
								Id:               "f46b0a1cf4d9400e9fd7af17f8742d37no02",
								Name:             "GeminiDB-efa6_prima_node_3",
								Status:           "normal",
								PrivateIp:        "192.168.0.176",
								SpecCode:         "nosql.cassandra.xlarge.4",
								AvailabilityZone: "bbb",
							},
						},
					},
				},
				BackupStrategy: instances.BackupStrategy{
					StartTime: "16:00-17:00",
					KeepDays:  7,
				},
			},

			{
				Port:              "8635",
				Engine:            " rocksDB ",
				DbUserName:        "rwuser",
				PayMode:           "0",
				TimeZone:          "",
				MaintenanceWindow: "02:00-06:00",
				Actions:           []string{"CREATE"},
				Groups: []instances.Groups{
					{
						Id:     "0b0ff12541794e1084f6827e424be2d1gr07",
						Status: "creating",
						Volume: instances.Volume{
							Size: "100",
							Used: "0.003",
						},
						Nodes: []instances.Nodes{
							{
								Id:               "233eaac9c6f245c0bb9c2d21eea12d1bno06",
								Name:             "GeminiDB-efa7_priam_node_2",
								Status:           "normal",
								PrivateIp:        "192.168.0.174",
								SpecCode:         "nosql.cassandra.xlarge.4",
								AvailabilityZone: "bbb",
							},
							{
								Id:               "d57d76d6320a4a7b86db82c317550c4ano02",
								Name:             "GeminiDB-efa7_priam_node_1",
								Status:           "normal",
								PrivateIp:        "192.168.0.39",
								SpecCode:         "nosql.cassandra.xlarge.4",
								AvailabilityZone: "bbb",
							},
							{
								Id:               "f46b0a1cf4d9400e9fd7af17f8742d37no02",
								Name:             "GeminiDB-efa7_prima_node_3",
								Status:           "normal",
								PrivateIp:        "192.168.0.176",
								SpecCode:         "nosql.cassandra.xlarge.4",
								AvailabilityZone: "bbb",
							},
						},
					},
				},
				BackupStrategy: instances.BackupStrategy{
					StartTime: "16:00-17:00",
					KeepDays:  7,
				},
			},
		},
	}

	ExpectedAllInstancesResponse.Instances[0].GeminiDBBase = instances.GeminiDBBase{
		Id:              "8436a91546294036b75931e879882200in06",
		Name:            "GeminiDB-efa6",
		Status:          "normal",
		Region:          "aaa",
		Mode:            " Cluster",
		VpcId:           "674e9b42-cd8d-4d25-a2e6-5abcc565b961",
		SubnetId:        "f1df08c5-71d1-406a-aff0-de435a51007b",
		SecurityGroupId: "7aa51dbf-5b63-40db-9724-dad3c4828b58",
		DataStore: instances.DataStore{
			Type:    "GeminiDB-Cassandra",
			Version: "3.11",
		},
		EnterpriseProjectId: "0",
	}

	ExpectedAllInstancesResponse.Instances[1].GeminiDBBase = instances.GeminiDBBase{
		Id:              "1236a91546294036b75931e879882200in06",
		Name:            "GeminiDB-efa7",
		Status:          "normal",
		Region:          "aaa",
		Mode:            " Cluster ",
		VpcId:           "674e9b42-cd8d-4d25-a2e6-5abcc565b961",
		SubnetId:        "f1df08c5-71d1-406a-aff0-de435a51007b",
		SecurityGroupId: "7aa51dbf-5b63-40db-9724-dad3c4828b58",
		DataStore: instances.DataStore{
			Type:    "GeminiDB-Cassandra",
			Version: "3.11",
		},
		EnterpriseProjectId: "0",
	}

	options := instances.ListGeminiDBInstanceOpts{
		Id: "ed7cc6166ec24360a5ed5c5c9c2ed726in06",
	}

	count := 0
	err := instances.List(client.ServiceClient(), options).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := instances.ExtractGeminiDBInstances(page)
		if err != nil {
			t.Errorf("Failed to extract instances: %v", err)
			return false, err
		}
		th.CheckDeepEquals(t, ExpectedAllInstancesResponse, actual)
		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, count)
}
