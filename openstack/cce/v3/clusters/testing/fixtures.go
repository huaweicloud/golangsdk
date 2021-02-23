package testing

import (
	"github.com/huaweicloud/golangsdk/openstack/cce/v3/clusters"
)

const Output = `
{
    "kind": "Cluster",
    "apiVersion": "v3",
    "metadata": {
        "name": "test-cluster",
        "uid": "daa97872-59d7-11e8-a787-0255ac101f54"
    },
    "spec": {
        "type": "VirtualMachine",
        "flavor": "cce.s1.small",
        "version": "v1.7.3-r10",
        "hostNetwork": {
            "vpc": "3305eb40-2707-4940-921c-9f335f84a2ca",
            "subnet": "00e41db7-e56b-4946-bf91-27bb9effd664"
        },
        "containerNetwork": {
            "mode": "overlay_l2"
        },
        "billingMode": 0
    },
    "status": {
        "phase": "Available",
        "endpoints": [
                    {
                        "url": "https://192.168.0.68:5443",
                        "type": "Internal"
                    }
                ]
    }
}`

const OutputOTC = `
{
    "kind": "Cluster",
    "apiVersion": "v3",
    "metadata": {
        "name": "test-cluster",
        "uid": "daa97872-59d7-11e8-a787-0255ac101f54"
    },
    "spec": {
        "type": "VirtualMachine",
        "flavor": "cce.s1.small",
        "version": "v1.7.3-r10",
        "hostNetwork": {
            "vpc": "3305eb40-2707-4940-921c-9f335f84a2ca",
            "subnet": "00e41db7-e56b-4946-bf91-27bb9effd664"
        },
        "containerNetwork": {
            "mode": "overlay_l2"
        },
        "billingMode": 0
    },
    "status": {
        "phase": "Available",
        "endpoints": {
			"internal": "https://192.168.0.68:5443",
            "external": "https://10.34.56.78:5443",
			"external_otc": "https://4d1ecb2c-229a-11e8-9c75-0255ac100ceb.container.eu-de.otc.t-systems.com"
		}
                
    }
}`

var Expected = &clusters.Clusters{
	Kind:       "Cluster",
	ApiVersion: "v3",
	Metadata: clusters.MetaData{
		Name: "test-cluster",
		Id:   "daa97872-59d7-11e8-a787-0255ac101f54",
	},
	Spec: clusters.Spec{
		Type:    "VirtualMachine",
		Flavor:  "cce.s1.small",
		Version: "v1.7.3-r10",
		HostNetwork: clusters.HostNetworkSpec{
			VpcId:    "3305eb40-2707-4940-921c-9f335f84a2ca",
			SubnetId: "00e41db7-e56b-4946-bf91-27bb9effd664",
		},
		ContainerNetwork: clusters.ContainerNetworkSpec{
			Mode: "overlay_l2",
		},
		BillingMode: 0,
	},
	Status: clusters.Status{
		Phase: "Available",
		Endpoints: []clusters.Endpoints{
			{Url: "https://192.168.0.68:5443", Type: "Internal"},
		},
	},
}

var ExpectedOTC = &clusters.Clusters{
	Kind:       "Cluster",
	ApiVersion: "v3",
	Metadata: clusters.MetaData{
		Name: "test-cluster",
		Id:   "daa97872-59d7-11e8-a787-0255ac101f54",
	},
	Spec: clusters.Spec{
		Type:    "VirtualMachine",
		Flavor:  "cce.s1.small",
		Version: "v1.7.3-r10",
		HostNetwork: clusters.HostNetworkSpec{
			VpcId:    "3305eb40-2707-4940-921c-9f335f84a2ca",
			SubnetId: "00e41db7-e56b-4946-bf91-27bb9effd664",
		},
		ContainerNetwork: clusters.ContainerNetworkSpec{
			Mode: "overlay_l2",
		},
		BillingMode: 0,
	},
	Status: clusters.Status{
		Phase: "Available",
		Endpoints: []clusters.Endpoints{
			{Internal: "https://192.168.0.68:5443", External: "https://10.34.56.78:5443",
				ExternalOTC: "https://4d1ecb2c-229a-11e8-9c75-0255ac100ceb.container.eu-de.otc.t-systems.com"},
		},
	},
}

const ListOutput = `
{
 "items": [
        {
            "kind": "Cluster",
            "apiVersion": "v3",
            "metadata": {
                "name": "test123",
                "uid": "daa97872-59d7-11e8-a787-0255ac101f54"
            },
            "spec": {
                "type": "VirtualMachine",
                "flavor": "cce.s1.small",
                "version": "v1.7.3-r10",
                "hostNetwork": {
                    "vpc": "3305eb40-2707-4940-921c-9f335f84a2ca",
                    "subnet": "00e41db7-e56b-4946-bf91-27bb9effd664"
                },
                "containerNetwork": {
                    "mode": "overlay_l2"
                },
                "billingMode": 0
            },
            "status": {
                "phase": "Available",
                "endpoints": [
                    {
                        "url": "https://192.168.0.68:5443",
                        "type": "Internal"
                    }
                ]
            }
        }
    ]
}
`

const ListOutputOTC = `
{
 "items": [
        {
            "kind": "Cluster",
            "apiVersion": "v3",
            "metadata": {
                "name": "test123",
                "uid": "daa97872-59d7-11e8-a787-0255ac101f54"
            },
            "spec": {
                "type": "VirtualMachine",
                "flavor": "cce.s1.small",
                "version": "v1.7.3-r10",
                "hostNetwork": {
                    "vpc": "3305eb40-2707-4940-921c-9f335f84a2ca",
                    "subnet": "00e41db7-e56b-4946-bf91-27bb9effd664"
                },
                "containerNetwork": {
                    "mode": "overlay_l2"
                },
                "billingMode": 0
            },
            "status": {
                "phase": "Available",
                "endpoints": {
					"internal": "https://192.168.0.68:5443",
            		"external": "https://10.34.56.78:5443",
					"external_otc": "https://4d1ecb2c-229a-11e8-9c75-0255ac100ceb.container.eu-de.otc.t-systems.com"
				}
            }
        }
    ]
}
`

var ListExpected = []clusters.Clusters{
	{
		Kind:       "Cluster",
		ApiVersion: "v3",
		Metadata:   clusters.MetaData{Name: "test123", Id: "daa97872-59d7-11e8-a787-0255ac101f54"},
		Spec: clusters.Spec{Type: "VirtualMachine",
			Flavor:           "cce.s1.small",
			HostNetwork:      clusters.HostNetworkSpec{VpcId: "3305eb40-2707-4940-921c-9f335f84a2ca", SubnetId: "00e41db7-e56b-4946-bf91-27bb9effd664"},
			ContainerNetwork: clusters.ContainerNetworkSpec{Mode: "overlay_l2"},
			BillingMode:      0,
			Version:          "v1.7.3-r10",
		},
		Status: clusters.Status{Phase: "Available", Endpoints: []clusters.Endpoints{{Url: "https://192.168.0.68:5443", Type: "Internal"}}},
	},
}

var ListExpectedOTC = []clusters.Clusters{
	{
		Kind:       "Cluster",
		ApiVersion: "v3",
		Metadata:   clusters.MetaData{Name: "test123", Id: "daa97872-59d7-11e8-a787-0255ac101f54"},
		Spec: clusters.Spec{Type: "VirtualMachine",
			Flavor:           "cce.s1.small",
			HostNetwork:      clusters.HostNetworkSpec{VpcId: "3305eb40-2707-4940-921c-9f335f84a2ca", SubnetId: "00e41db7-e56b-4946-bf91-27bb9effd664"},
			ContainerNetwork: clusters.ContainerNetworkSpec{Mode: "overlay_l2"},
			BillingMode:      0,
			Version:          "v1.7.3-r10",
		},
		Status: clusters.Status{
			Phase: "Available",
			Endpoints: []clusters.Endpoints{
				{Internal: "https://192.168.0.68:5443", External: "https://10.34.56.78:5443",
					ExternalOTC: "https://4d1ecb2c-229a-11e8-9c75-0255ac100ceb.container.eu-de.otc.t-systems.com"},
			},
		},
	},
}
