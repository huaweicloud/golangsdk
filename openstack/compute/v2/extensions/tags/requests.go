package tags

import (
	"github.com/huaweicloud/golangsdk"
)

// ListTags list all tags of server specified by id
func ListTags(client *golangsdk.ServiceClient, id string) (r ListResult) {
	_, r.Err = client.Get(listTags(client, id), &r.Body, nil)
	return
}

// PutTags put tags to server specified by id
func PutTags(client *golangsdk.ServiceClient, id string, tags ...string) error {
	_, err := client.Put(replaceTags(client, id), map[string][]string{"tags": tags}, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}

// ClientTags remove all tags of a server
func CleanTags(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(deleteTags(client, id), nil)
	return err
}

// CheckTag check if the tag exists on the serve
func CheckTag(client *golangsdk.ServiceClient, id, tag string) error {
	_, err := client.Get(checkTag(client, id, tag), nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}

// AddTag add a tag to a server
func AddTag(client *golangsdk.ServiceClient, id, tag string) error {
	_, err := client.Put(addTag(client, id, tag), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}

// DeleteTag delete a tag from a server
func DeleteTag(client *golangsdk.ServiceClient, id, tag string) error {
	_, err := client.Delete(deleteTag(client, id, tag), nil)
	return err
}
