package testing

const (
	backupEndpoint     = "/checkpoint_items"
	checkpoint_item_id = "7b99acfd-18c3-4f26-9d39-b4ebd2ea3e12"
)

var getResponse = `
{
    "checkpoint_item": {
        "status": "available",
        "backup_data": {
            "eip": "80.158.17.102",
            "cloudservicetype": "QEMU",
            "ram": 8192,
            "vcpus": 4,
            "__openstack_region_name": "",
            "private_ip": "192.168.0.209",
            "disk": 0,
            "imagetype": "gold"
        },
        "name": "backup-c2c",
        "resource_id": "f8ddc472-cf00-4384-851e-5f2a68c33762",
        "created_at": "2018-08-14T07:53:15.663766",
        "checkpoint_id": "2eefe592-8424-4778-8d0d-962c8a5dd6a4",
        "updated_at": "2018-08-17T04:33:58.025327",
        "tags": [],
        "extend_info": {
            "auto_trigger": false,
            "space_saving_ratio": 2,
            "resource_name": "ecs-ggao",
            "fail_reason": "",
            "resource_az": "eu-de-02",
            "image_type": "backup",
            "finished_at": "2018-08-14T08:31:08.720800",
            "average_speed": 19,
            "copy_from": null,
            "volume_backups": [
                {
                    "status": "available",
                    "space_saving_ratio": 1,
                    "name": "manualbk_ee6d_ecs-ggao",
                    "bootable": true,
                    "average_speed": 16,
                    "source_volume_size": 24,
                    "source_volume_id": "c14856d0-07e8-453b-a442-444086cbad04",
                    "incremental": false,
                    "snapshot_id": "c18513c3-1ab9-46e3-979b-7cad7c52e516",
                    "source_volume_name": "ecs-ggao",
                    "image_type": "backup",
                    "id": "2422bc5e-4cde-4420-964d-30b7347042a7",
                    "size": 47960
                },
                {
                    "status": "available",
                    "space_saving_ratio": 3,
                    "name": "manualbk_ee6d_ggao-repo-disk",
                    "bootable": false,
                    "average_speed": 22,
                    "source_volume_size": 100,
                    "source_volume_id": "313b9a39-cdc7-4413-8a1b-1888340bdc03",
                    "incremental": false,
                    "snapshot_id": "5c9dd5d5-fc70-42a2-8d58-c9e009ccf418",
                    "source_volume_name": "ggao-repo-disk",
                    "image_type": "backup",
                    "id": "533e5e53-4332-48cd-b920-ae4fd9b3ba94",
                    "size": 98224
                }
            ],
            "fail_code": {},
            "copy_status": "na",
            "incremental": false,
            "taskid": "1afcab08-9f97-11e8-9526-286ed488ca8c",
            "hypervisor_type": "QEMU",
            "supported_restore_mode": "backup",
            "progress": 100,
            "support_lld": true,
            "fail_op": "",
            "resource_type": "OS::Nova::Server",
            "size": 146184
        },
        "id": "7b99acfd-18c3-4f26-9d39-b4ebd2ea3e12",
        "resource_type": "OS::Nova::Server",
        "description": "backup des"
    }
}
		`

var createRequest = `{
    "protect" : {
    "backup_name" : "c2c-backup",
    "description" : "mybackup"
  }
}`

var createResponse = `{
    "checkpoint": {
        "status": "protecting",
        "created_at": "2018-08-17T07:58:56.492307",
        "id": "92dba83d-cc6f-4883-a20d-de6934510b7e",
        "resource_graph": null,
        "project_id": "91d687759aed45d28b5f6084bc2fa8ad",
        "protection_plan": {
            "id": "fake_b94f8b46-b0a1-485a-ad5b-9f8876b85495",
            "resources": [
                {
                    "extra_info": "{}",
                    "type": "OS::Nova::Server",
                    "id": "f8ddc472-cf00-4384-851e-5f2a68c33762",
                    "name": "ecs-ggao"
                }
            ],
            "name": "server protect plan for f8ddc472-cf00-4384-851e-5f2a68c33762"
        }
    }
}`

var queryRequest = `{
  "check_protectable" : [ {
    "resource_id" : "069e678a-f1d1-4a38-880b-459bde82fcc6",
    "resource_type" : "OS::Nova::Server"
  } ]
}`

var queryResponse = `{
    "protectable": [
        {
            "result": true,
            "resource_type": "OS::Nova::Server",
            "resource_id": "069e678a-f1d1-4a38-880b-459bde82fcc6"
        }
    ]
}`

var listResponse = `
{
    "checkpoint_items": [
 {
            "status": "available",
            "backup_data": {
                "eip": "80.158.17.102",
                "cloudservicetype": "QEMU",
                "ram": 8192,
                "vcpus": 4,
                "__openstack_region_name": "",
                "private_ip": "192.168.0.209",
                "disk": 0,
                "imagetype": "gold"
            },
            "name": "backup-c2c",
            "resource_id": "f8ddc472-cf00-4384-851e-5f2a68c33762",
            "checkpoint_id": "2eefe592-8424-4778-8d0d-962c8a5dd6a4",
            "extend_info": {
                "auto_trigger": false,
                "space_saving_ratio": 2,
                "resource_name": "ecs-ggao",
                "fail_reason": "",
                "resource_az": "eu-de-02",
                "image_type": "backup",
                "finished_at": "2018-08-14T08:31:08.720800",
                "average_speed": 19,
                "copy_status": "na",
                "incremental": false,
                "taskid": "1afcab08-9f97-11e8-9526-286ed488ca8c",
                "hypervisor_type": "QEMU",
                "supported_restore_mode": "backup",
                "progress": 100,
                "support_lld": true,
                "fail_op": "",
                "resource_type": "OS::Nova::Server",
                "size": 146184
            },
            "id": "7b99acfd-18c3-4f26-9d39-b4ebd2ea3e12",
            "resource_type": "OS::Nova::Server",
            "description": "backup des"
        }
]
}
`
