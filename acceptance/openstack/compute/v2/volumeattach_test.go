// +build acceptance compute volumeattach

package v2

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	bs "github.com/huaweicloud/golangsdk/acceptance/openstack/blockstorage/v2"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestVolumeAttachAttachment(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	blockClient, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	server, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	volume, err := bs.CreateVolume(t, blockClient)
	th.AssertNoErr(t, err)
	defer bs.DeleteVolume(t, blockClient, volume)

	volumeAttachment, err := CreateVolumeAttachment(t, client, blockClient, server, volume)
	th.AssertNoErr(t, err)
	defer DeleteVolumeAttachment(t, client, blockClient, server, volumeAttachment)

	tools.PrintResource(t, volumeAttachment)

	th.AssertEquals(t, volumeAttachment.ServerID, server.ID)
}
