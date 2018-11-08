package testing

var createRequest = `
{
 "bucket_name": "obs-e51d",
 "file_prefix_name": "yO8Q",
 "smn": {
  "is_support_smn": true,
  "topic_id": "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:c2c-topic",
  "is_send_all_key_operation":false,
  "operations": ["login"],
  "need_notify_user_list": ["user1","user2"]

 }
}
`

var createResponse = `
{
    "bucket_name": "obs-e51d",
    "file_prefix_name": "yO8Q",
    "smn": {
        "is_support_smn": true,
        "topic_id": "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:c2c-topic",
        "is_send_all_key_operation": false,
        "operations": [
            "login"
        ],
        "need_notify_user_list": [
            "user1",
            "user2"
        ]
    },
    "status": "enabled",
    "tracker_name": "system"
}
`

var updateRequest = `
{ 
 "bucket_name" : "cirros-img", 
 "file_prefix_name" : "yO8Q", 
 "smn": {
  "is_support_smn": false,
  "topic_id": "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:c2c-topic",
  "is_send_all_key_operation":false,
  "operations": ["delete","create","login"],
  "need_notify_user_list": ["user1","user2"]
 },
  "status" : "disabled"   
}
`

var updateResponse = `
{
    "bucket_name": "cirros-img",
    "file_prefix_name": "yO8Q",
    "smn": {
        "is_support_smn": false,
        "topic_id": "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:c2c-topic",
        "is_send_all_key_operation": false,
        "operations": [
            "delete",
            "create",
            "login"
        ],
        "need_notify_user_list": [
            "user1",
            "user2"
        ]
    },
    "status": "disabled",
    "tracker_name": "system"
}
`

var getResponse = `
[
    {
        "bucket_name": "tf-test-bucket",
        "file_prefix_name": "yO8Q",
        "smn": {
            "is_send_all_key_operation": false,
            "is_support_smn": true,
            "need_notify_user_list": [
                "user1"
            ],
            "operations": [
                "login"
            ],
            "topic_id": "urn:smn:eu-de:626ce20e52a346c090b09cffc3e038e5:tf-test-topic"
        },
        "status": "enabled",
        "tracker_name": "system"
    }
]
`
