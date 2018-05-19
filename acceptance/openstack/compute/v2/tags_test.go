// +build acceptance compute tags

package v2

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/compute/v2/extensions/tags"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestServersTags(t *testing.T) {

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	server, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	tools.PrintResource(t, server)

	err = tags.PutTags(client, server.ID, "tag1", "tag2")
	th.AssertNoErr(t, err)
	t.Logf("Update tags succcessfully")

	tagList, err := tags.ListTags(client, server.ID).Extract()
	th.AssertNoErr(t, err)

	expectedTags := []string{"tag1", "tag2"}
	th.AssertDeepEquals(t, expectedTags, tagList.Tags)

	err = tags.DeleteTag(client, server.ID, "tag1")
	th.AssertNoErr(t, err)

	tagList, err = tags.ListTags(client, server.ID).Extract()
	th.AssertNoErr(t, err)

	expectedTags = []string{"tag2"}
	th.AssertDeepEquals(t, expectedTags, tagList.Tags)

	err = tags.AddTag(client, server.ID, "tag1")
	th.AssertNoErr(t, err)
	t.Logf("Add tag successfully")

	tagList, err = tags.ListTags(client, server.ID).Extract()
	th.AssertNoErr(t, err)

	expectedTags = []string{"tag1", "tag2"}

	th.AssertDeepEquals(t, expectedTags, tagList.Tags)

	err = tags.CleanTags(client, server.ID)
	th.AssertNoErr(t, err)
	t.Logf("Clean tags successfully")

	tagList, err = tags.ListTags(client, server.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, map[string]string{}, tagList)
}
