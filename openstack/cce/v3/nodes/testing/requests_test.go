package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/cce/v3/common"
	"github.com/huaweicloud/golangsdk/openstack/cce/v3/nodes"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestListNode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "kind": "List",
    "apiVersion": "v3",
	"items":
	[ 
       {
            "kind": "Host",
            "apiVersion": "v3",
            "metadata": {
                "name": "test-node-1234",
                "uid": "b99acd73-5d7c-11e8-8e76-0255ac101929"               
            },
            "spec": {
                "flavor": "s1.medium",
                "az": "cn-east-2a",
                "login": {
                    "sshKey": "test-keypair",
                    "userPassword": {}
                },
                "rootVolume": {
                    "volumetype": "SATA",
                    "size": 40
                },
                "dataVolumes": [
                    {
                        "volumetype": "SATA",
                        "size": 100
                    }
                ],
                "publicIP": {
                    "eip": {
                        "bandwidth": {}
                    }
                },
                "billingMode": 0
            },
            "status": {
                "phase": "Active",
                "serverId": "41748e56-33d4-46a1-aa57-2c8c29907995",
                "privateIP": "192.168.0.3"
            }
        }
	]
}
		`)
	})

	listNodes := nodes.ListOpts{Name: "test-node-1234"}
	actual, err := nodes.List(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", listNodes)

	if err != nil {
		t.Errorf("Failed to extract nodes: %v", err)
	}

	expected := []nodes.Nodes{
		{
			Kind:       "Host",
			Apiversion: "v3",
			Metadata: nodes.Metadata{Name: "test-node-1234",
				Id: "b99acd73-5d7c-11e8-8e76-0255ac101929"},
			Spec: nodes.Spec{Az: "cn-east-2a",
				Login:       nodes.LoginSpec{SshKey: "test-keypair"},
				RootVolume:  nodes.VolumeSpec{Size: 40, VolumeType: "SATA"},
				BillingMode: 0,
				DataVolumes: []nodes.VolumeSpec{
					{
						VolumeType: "SATA",
						Size:       100,
					}},
				Flavor: "s1.medium",
			},
			Status: nodes.Status{Phase: "Active", ServerID: "41748e56-33d4-46a1-aa57-2c8c29907995", PrivateIP: "192.168.0.3"},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestGetV3Node(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes/cf4bc001-58f1-11e8-ad73-0255ac101926", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, Output)
	})

	actual, err := nodes.Get(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", "cf4bc001-58f1-11e8-ad73-0255ac101926").Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)

}

func TestCreateV3Node(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `
			{
	  "apiversion": "v3",
	  "kind": "Node",
	  "metadata": {
	    "name": "test-node"
	  },
	  "spec": {
	    "az": "cn-east-2a",
	    "count": 1,
        "extendParam": {
        },
	    "dataVolumes": [
	      {
	        "size": 100,
	        "volumetype": "SATA"
	      }
	    ],
	    "flavor": "s3.large.2",
	    "login": {
	      "sshKey": "test-keypair",
	      "userPassword": {
			"password": "",
            "username": ""
		  }
	    },
		"nodeNicSpec": {
		  "primaryNic": {}
		},
	    "publicIP": {
		      "eip": {
		        "bandwidth": {}
		      }
		 },
	    "rootVolume": {
	      "size": 40,
	      "volumetype": "SATA"
	    }
	  }
	}
`)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, Output)
	})
	options := nodes.CreateOpts{Kind: "Node",
		ApiVersion: "v3",
		Metadata:   nodes.CreateMetaData{Name: "test-node"},
		Spec: nodes.Spec{Flavor: "s3.large.2", Az: "cn-east-2a",
			Login:       nodes.LoginSpec{SshKey: "test-keypair"},
			RootVolume:  nodes.VolumeSpec{Size: 40, VolumeType: "SATA"},
			DataVolumes: []nodes.VolumeSpec{{Size: 100, VolumeType: "SATA"}},
			Count:       1},
	}
	actual, err := nodes.Create(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", options).Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)

}

func TestUpdateV3Node(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes/cf4bc001-58f1-11e8-ad73-0255ac101926", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "metadata": {
        "name": "test-node"
    }
}			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, Output)
	})
	options := nodes.UpdateOpts{Metadata: nodes.UpdateMetadata{Name: "test-node"}}
	actual, err := nodes.Update(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", "cf4bc001-58f1-11e8-ad73-0255ac101926", options).Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)
}

func TestDeleteNode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes/cf4bc001-58f1-11e8-ad73-0255ac101926", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
	})

	err := nodes.Delete(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", "cf4bc001-58f1-11e8-ad73-0255ac101926").ExtractErr()
	th.AssertNoErr(t, err)

}

func TestGetV3Job(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/jobs/73ce03fd-8b1b-11e8-8f9d-0255ac10193f", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, JobOutput)
	})

	actual, err := nodes.GetJobDetails(fake.ServiceClient(), "73ce03fd-8b1b-11e8-8f9d-0255ac10193f").ExtractJob()
	th.AssertNoErr(t, err)
	expected := ExpectedJob
	th.AssertDeepEquals(t, expected, actual)

}
