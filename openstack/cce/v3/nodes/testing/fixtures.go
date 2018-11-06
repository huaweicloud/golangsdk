package testing

import "github.com/huaweicloud/golangsdk/openstack/cce/v3/nodes"

const Output = `{
    "kind": "Host",
    "apiVersion": "v3",
    "metadata": {
        "name": "test-node"
    },
    "spec": {
        "flavor": "s3.large.2",
        "az": "cn-east-2a",
        "login": {
            "sshKey": "test-keypair"
        },
        "rootVolume": {
            "volumetype": "SATA",
            "size": 40
        },
        "publicIP": {
		      "eip": {
		        "bandwidth": {}
		      }
		    },
        "dataVolumes": [
            {
                "volumetype": "SATA",
                "size": 100
            }
        ]
}
}`

var Expected = &nodes.Nodes{
	Kind:       "Host",
	Apiversion: "v3",
	Metadata:   nodes.Metadata{Name: "test-node"},
	Spec: nodes.Spec{
		Flavor: "s3.large.2",
		Az:     "cn-east-2a",
		Login: nodes.LoginSpec{
			SshKey: "test-keypair",
		},
		ExtendParam: nodes.ExtendParam{},
		PublicIP:    nodes.PublicIPSpec{Eip: nodes.EipSpec{Bandwidth: nodes.BandwidthOpts{}, IpType: ""}},
		RootVolume: nodes.VolumeSpec{
			VolumeType: "SATA",
			Size:       40,
		},
		DataVolumes: []nodes.VolumeSpec{
			{
				VolumeType: "SATA",
				Size:       100,
			},
		},
	},
}

const JobOutput = `{
    "kind": "Job",
    "metadata": {
        "uid": "73ce052c-8b1b-11e8-8f9d-0255ac10193f"    
    },
    "spec": {
        "type": "ScaleupCluster",
        "clusterUID": "6951aa4d-88ef-11e8-b196-0255ac101c43",
        "resourceName": "cluster-test",
        "subJobs": [
            {
                "kind": "Job",
                "metadata": {
                    "uid": "73cc28df-8b1b-11e8-8f9d-0255ac10193f"
                },
                "spec": {
                    "type": "CreateNode",
                    "clusterUID": "6951aa4d-88ef-11e8-b196-0255ac101c43",
                    "resourceName": "myhost",
                    "subJobs": [
                        {
                            "kind": "Job",
                            "metadata": {
                                "uid": "73ce03fd-8b1b-11e8-8f9d-0255ac10193f" 
                            },
                            "spec": {
                                "type": "GetPSMCert",
                                "clusterUID": "6951aa4d-88ef-11e8-b196-0255ac101c43"
                            },
                            "status": {
                                "phase": "Success"
                            }
                        },
                        {
                            "kind": "Job",
                            "metadata": {
                                "uid": "73ce0473-8b1b-11e8-8f9d-0255ac10193f"
                            },
                            "spec": {
                                "type": "InstallNode",
                                "clusterUID": "6951aa4d-88ef-11e8-b196-0255ac101c43",
                                "resourceID": "73bd7e31-8b1b-11e8-8f9d-0255ac10193f"
                            },
                            "status": {
                                "phase": "Success"
                            }
                        }
                    ]
                }
            }
        ]
    },
    "status": {
        "phase": "Success"
    }
}`

var ExpectedJob = &nodes.Job{
	Kind:     "Job",
	Status:   nodes.JobStatus{Phase: "Success"},
	Metadata: nodes.JobMetadata{ID: "73ce052c-8b1b-11e8-8f9d-0255ac10193f"},
	Spec: nodes.JobSpec{Type: "ScaleupCluster",
		ClusterID:    "6951aa4d-88ef-11e8-b196-0255ac101c43",
		ResourceName: "cluster-test",
		SubJobs: []nodes.Job{{Kind: "Job",
			Metadata: nodes.JobMetadata{ID: "73cc28df-8b1b-11e8-8f9d-0255ac10193f"},
			Spec: nodes.JobSpec{Type: "CreateNode",
				ClusterID:    "6951aa4d-88ef-11e8-b196-0255ac101c43",
				ResourceName: "myhost",
				SubJobs: []nodes.Job{{Kind: "Job",
					Metadata: nodes.JobMetadata{ID: "73ce03fd-8b1b-11e8-8f9d-0255ac10193f"},
					Spec: nodes.JobSpec{Type: "GetPSMCert",
						ClusterID: "6951aa4d-88ef-11e8-b196-0255ac101c43"},
					Status: nodes.JobStatus{Phase: "Success"}},
					{Kind: "Job",
						Metadata: nodes.JobMetadata{ID: "73ce0473-8b1b-11e8-8f9d-0255ac10193f"},
						Spec: nodes.JobSpec{Type: "InstallNode",
							ClusterID:  "6951aa4d-88ef-11e8-b196-0255ac101c43",
							ResourceID: "73bd7e31-8b1b-11e8-8f9d-0255ac10193f"},
						Status: nodes.JobStatus{Phase: "Success"}},
				},
			},
		}},
	},
}
