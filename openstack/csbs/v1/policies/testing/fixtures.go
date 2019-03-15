package testing

const (
	policiesEndpoint = "/policies"
	policies_id      = "5af626d2-19b9-4dc4-8e95-ddba008318b3"
)

var getResponse = `
	{
    "policy": {
        "status": "suspended",
        "provider_id": "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
        "description": "My plan",
        "tags": [],
        "scheduled_operations": [
            {
                "description": "My backup policy",
                "enabled": true,
                "trigger_id": "30411091-f206-48e9-8ef9-62be070ea217",
                "trigger": {
                    "properties": {
                        "pattern": "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
                        "start_time": "2018-08-20 07:31:32"
                    },
                    "type": "time",
                    "id": "30411091-f206-48e9-8ef9-62be070ea217",
                    "name": "default"
                },
                "operation_definition": {
                    "max_backups": 20,
                    "provider_id": "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
                    "plan_id": "5af626d2-19b9-4dc4-8e95-ddba008318b3"
                },
                "operation_type": "backup",
                "id": "b70c712d-f48b-43f7-9a0f-3bab86d59149",
                "name": "my-backup-policy"
            }
        ],
        "id": "5af626d2-19b9-4dc4-8e95-ddba008318b3",
        "name": "c2c-policy",
        "parameters": {
            "common": {}
        },
        "created_at": "2018-08-20T07:31:32.718435",
        "project_id": "91d687759aed45d28b5f6084bc2fa8ad",
        "resources": [
            {
                "type": "OS::Nova::Server",
                "id": "cd5955b4-44c0-4f0a-ac57-2401b89cb347",
                "name": "resource1"
            }
        ]
    }
}
`
var createRequest = `
{
  "policy" : {
    "name" : "c2c-policy",
    "description" : "My plan",
    "provider_id" : "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
   "parameters": {
            "common": {}
        },
    "scheduled_operations" : [ {
      "name" : "my-backup-policy",
      "description" : "My backup policy",
      "enabled" : true,
      "operation_definition" : {
        "max_backups" : 20,
		"permanent" : false
      },
      "trigger" : {
        "properties" : {
          "pattern" : "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n"
        }
      },
      "operation_type" : "backup"
    }],
    "resources" : [ {
      "id" : "cd5955b4-44c0-4f0a-ac57-2401b89cb347",			
      "type" : "OS::Nova::Server",
      "name" : "resource1"
     
    }]
  }
}

`

var createResponse = `
{
    "policy": {
        "status": "suspended",
        "provider_id": "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
        "description": "My plan",
        "tags": [],
        "scheduled_operations": [
            {
                "description": "My backup policy",
                "enabled": true,
                "trigger_id": "30411091-f206-48e9-8ef9-62be070ea217",
                "trigger": {
                    "properties": {
                        "pattern": "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
                        "start_time": "2018-08-20 07:31:32"
                    },
                    "type": "time",
                    "id": "30411091-f206-48e9-8ef9-62be070ea217",
                    "name": "default"
                },
                "operation_definition": {
                    "max_backups": "20",
                    "provider_id": "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
                    "plan_id": "5af626d2-19b9-4dc4-8e95-ddba008318b3"
                },
                "operation_type": "backup",
                "id": "b70c712d-f48b-43f7-9a0f-3bab86d59149",
                "name": "my-backup-policy"
            }
        ],
        "id": "5af626d2-19b9-4dc4-8e95-ddba008318b3",
        "name": "c2c-policy",
        "parameters": {
            "common": {}
        },
        "created_at": "2018-08-20T07:31:32.718435",
        "project_id": "91d687759aed45d28b5f6084bc2fa8ad",
        "resources": [
            {
                "type": "OS::Nova::Server",
                "id": "cd5955b4-44c0-4f0a-ac57-2401b89cb347",
                "name": "resource1"
            }
        ]
    }
}
`

var updateRequest = `
{
  "policy" : {
    "name" : "c2c-policy-update",
    "parameters" : {
      "common" : {
      }
    },
    "scheduled_operations" : [ {
      "id" : "b70c712d-f48b-43f7-9a0f-3bab86d59149",
      "name" : "my-backup-policy",
      "description" : "My backup policy",
      "enabled" : true,
      "operation_definition" : {
        "retention_duration_days" : -1,
        "max_backups" : 20,
		"permanent" : false
      },
      "trigger" : {
        "properties" : {
          "pattern" : "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n"
        }
      }
    }
  ]
}
}
`

var updateResponse = `
{
    "policy": {
        "status": "suspended",
        "provider_id": "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
        "description": "My plan",
        "tags": [],
        "scheduled_operations": [
            {
                "description": "My backup policy",
                "enabled": true,
                "trigger": {
                    "properties": {
                        "pattern": "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
                        "start_time": "2018-08-20 07:31:32"
                    }
                },
                "operation_definition": {
                    "max_backups": "20",
                    "provider_id": "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
                    "plan_id": "5af626d2-19b9-4dc4-8e95-ddba008318b3",
                    "retention_duration_days": "-1"
                },
                "operation_type": "backup",
                "id": "b70c712d-f48b-43f7-9a0f-3bab86d59149",
                "name": "my-backup-policy"
            }
        ],
        "id": "5af626d2-19b9-4dc4-8e95-ddba008318b3",
        "user_id": null,
        "name": "c2c-policy-update",
        "parameters": {
            "common": {}
        },
        "created_at": "2018-08-20T07:31:32.718435",
        "project_id": "91d687759aed45d28b5f6084bc2fa8ad",
        "resources": [
            {
                "type": "OS::Nova::Server",
                "id": "cd5955b4-44c0-4f0a-ac57-2401b89cb347",
                "name": "resource1"
            }
        ]
    }
}
`
var listResponse = `
{
    "policies": [
        {
            "status": "suspended",
            "provider_id": "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
            "description": "My plann",
            "scheduled_operations": [
                {
                    "description": "My backup policy",
                    "enabled": true,
                    "trigger_id": "831b5e69-0b75-420c-918e-9cbcb32d97f1",
                    "trigger": {
                        "properties": {
                            "pattern": "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n"
                        },
                        "type": "time",
                        "id": "831b5e69-0b75-420c-918e-9cbcb32d97f1",
                        "name": "default"
                    },
                    "operation_definition": {
                        "max_backups": 5,
                        "provider_id": "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
                        "plan_id": "4d1ce19b-d681-4e44-a87e-c44eb9bfc4c7"
                    },
                    "operation_type": "backup",
                    "id": "e7d50d4c-2f38-40a4-9f9b-c9c355a52417",
                    "name": "my-backupp"
                }
            ],
            "id": "4d1ce19b-d681-4e44-a87e-c44eb9bfc4c7",
            "name": "my-plan-test1",
            "parameters": {
            "common": {}
        },
            "created_at": "2018-08-20T10:43:56.246383",
            "project_id": "91d687759aed45d28b5f6084bc2fa8ad",
            "resources": [
                {
                    "type": "OS::Nova::Server",
                    "id": "9422f270-6fcf-4ba2-9319-a007f2f63a8e",
                    "name": "resource4"
                }
            ]
        }
 	]
}
`
