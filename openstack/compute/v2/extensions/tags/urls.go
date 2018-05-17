package tags

import "github.com/huaweicloud/golangsdk"

func operateTagList(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "tags")
}

var (
	listTags    = operateTagList
	replaceTags = operateTagList
	deleteTags  = operateTagList
)

func operateTag(client *golangsdk.ServiceClient, id, tag string) string {
	return client.ServiceURL("servers", id, "tags", tag)
}

var (
	checkTag  = operateTag
	addTag    = operateTag
	deleteTag = operateTag
)
