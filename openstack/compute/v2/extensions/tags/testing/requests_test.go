package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/compute/v2/extensions/tags"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const serverID = "{serverId}"

func TestListTags(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	mockListTagsResponse(t, serverID)

	_, err := tags.ListTags(client.ServiceClient(), serverID).Extract()
	th.AssertNoErr(t, err)
}

func TestPutTags(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	mockPutTagsResponse(t, serverID)

	err := tags.PutTags(client.ServiceClient(), serverID, "tag1", "tag2")
	th.AssertNoErr(t, err)
}

func TestCleanTags(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	mockCleanTagsResponse(t, serverID)

	err := tags.CleanTags(client.ServiceClient(), serverID)
	th.AssertNoErr(t, err)
}

func TestCheckTag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	mockCheckTagResponse(t, serverID, "tagN")

	err := tags.CheckTag(client.ServiceClient(), serverID, "tagN")
	th.AssertNoErr(t, err)
}

func TestAddTag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	mockAddTagResponse(t, serverID, "tagN")

	err := tags.AddTag(client.ServiceClient(), serverID, "tagN")
	th.AssertNoErr(t, err)
}

func TestDeleteTag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	mockDeleteTagResponse(t, serverID, "tagN")

	err := tags.DeleteTag(client.ServiceClient(), serverID, "tagN")
	th.AssertNoErr(t, err)
}
