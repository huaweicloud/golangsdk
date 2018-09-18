package testing

import (
	"github.com/huaweicloud/golangsdk/openstack/vbs/v2/tags"
)

const TagList = `
{
  "total_count":1,
  "resources":[
    {
      "resource_name": "name",
      "resource_id": "0781095c-b8ab-4ce5-99f3-4c5f6ff75319",
      "resource_detail": null,
      "tags": [{
          "key":"key",
          "value":"value"
       }]
    }
  ]
}`

var AddTag = `
{
    "tag":{
        "key":"0f187b65-8d0e-4fc0-9096-3b55d330531e",
        "value":"volume"
    }
}`

var getTags = `{
  "tags": [
    {
      "key": "RUNNING",
      "value":"0781095c-b8ab-4ce5-99f3-4c5f6ff75319"
    },
    {
      "key": "WAITING",
      "values":"2016-12-03T06:24:34.467"
    }
  ]
}`

var batchAction = `{
    "action":"update",
    "tags":[{
        "key":"0f187b65-8d0e-4fc0-9096-3b55d330531e",
        "value":"volume"
        },{
        "key":"0f187b65-8d0e-4fc0-9096-3b55d330531d",
        "value":"volume"
    }]
}`

var ExpectedTags = &tags.Resources{
	TotalCount: 1,
	Resource:   []tags.Resource{{ResourceName: "name", ResourceID: "0781095c-b8ab-4ce5-99f3-4c5f6ff75319", Tags: []tags.Tag{{Key: "key", Value: "value"}}}},
}
