package testing

import "github.com/huaweicloud/golangsdk/openstack/vbs/v2/policies"

const Output = `
{
		  "backup_policy_name": "Test_Policy",
		  "scheduled_policy": {
		    "frequency": 1,
		    "remain_first_backup_of_curMonth": "Y",
		    "rentention_num": 10,
		    "start_time": "12:00",
		    "status": "ON"
		  },
		  "tags": [
		    {
		      "key": "key",
		      "value": "value"
		    }
		  ]
		}`

const UpdateOutput = `
{
    "backup_policy_name": "Test_02",
    "scheduled_policy" : {
        "remain_first_backup_of_curMonth" : "Y",
        "rentention_num" : 10,
        "frequency" : 1,
        "start_time" : "10:00",
        "status" : "ON"
    }
}`

const ListOutput = `{
    "backup_policies" : [
    {
        "backup_policy_id" : "ed8b9f73-4415-494d-a54e-5f3373bc353d",
        "backup_policy_name": "plan01",
        "scheduled_policy" : {
            "remain_first_backup_of_curMonth" : "Y",
            "rentention_num" : 10,
            "frequency" : 1,
            "start_time" : "12:00",
            "status" : "ON"
        },
        "policy_resource_count": 0
    },
    {
        "backup_policy_id" : "8dd473c9-5a80-4ad5-862e-492c9af2b6bd",
        "backup_policy_name": "plan02",
        "scheduled_policy" : {
            "remain_first_backup_of_curMonth" : "Y",
            "rentention_num" : 10,
            "frequency" : 1,
            "start_time" : "14:00",
            "status" : "ON"
        },
        "policy_resource_count": 10
    }]
}`

const AssociateOutput = `
{
    "success_resources": [
        {
            "resource_id": "0f187b65-8d0e-4fc0-9096-3b55d330531e",
            "os_vol_host_attr": "pod01.eu-de-01",
            "availability_zone": "eu-de-01",
            "resource_type": "volume"
        }
    ], 
    "fail_resources": [ ]
}`

const DisssociateOutput = `
{
    "success_resources": [
        {
            "resource_id": "0f187b65-8d0e-4fc0-9096-3b55d330531e"            
        }
    ], 
    "fail_resources": [ ]
}`

var ListPolicies = []policies.Policy{
	{ID: "ed8b9f73-4415-494d-a54e-5f3373bc353d",
		Name: "plan01",
		ScheduledPolicy: policies.ScheduledPolicy{
			Frequency:         1,
			RemainFirstBackup: "Y",
			RententionNum:     10,
			StartTime:         "12:00",
			Status:            "ON",
		},
		ResourceCount: 0,
	},
	{ID: "8dd473c9-5a80-4ad5-862e-492c9af2b6bd",
		Name: "plan02",
		ScheduledPolicy: policies.ScheduledPolicy{
			Frequency:         1,
			RemainFirstBackup: "Y",
			RententionNum:     10,
			StartTime:         "14:00",
			Status:            "ON",
		},
		ResourceCount: 10,
	},
}

var Update = &policies.Policy{
	Name: "Test_02",
	ScheduledPolicy: policies.ScheduledPolicy{
		Frequency:         1,
		RemainFirstBackup: "Y",
		RententionNum:     10,
		StartTime:         "10:00",
		Status:            "ON",
	},
}

var Expected = &policies.Policy{
	Name: "Test_Policy",
	ScheduledPolicy: policies.ScheduledPolicy{
		Frequency:         1,
		RemainFirstBackup: "Y",
		RententionNum:     10,
		StartTime:         "12:00",
		Status:            "ON",
	},
}

var Associate = &policies.ResultResources{
	SuccessResources: []policies.Resource{
		{ResourceID: "0f187b65-8d0e-4fc0-9096-3b55d330531e",
			ResourceType:     "volume",
			Pod:              "pod01.eu-de-01",
			AvailabilityZone: "eu-de-01",
		},
	},
	FailResources: []policies.Resource{},
}

var Disassociate = &policies.ResultResources{
	SuccessResources: []policies.Resource{
		{ResourceID: "0f187b65-8d0e-4fc0-9096-3b55d330531e"},
	},
	FailResources: []policies.Resource{},
}
