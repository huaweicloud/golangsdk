// +build acceptance compute bootfromvolume

package v2

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	blockstorage "github.com/huaweicloud/golangsdk/acceptance/openstack/blockstorage/v2"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/huaweicloud/golangsdk/openstack/compute/v2/extensions/volumeattach"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestBootFromImage(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	choices, err := clients.AcceptanceTestChoicesFromEnv()
	th.AssertNoErr(t, err)

	blockDevices := []bootfromvolume.BlockDevice{
		{
			BootIndex:           0,
			DeleteOnTermination: true,
			DestinationType:     bootfromvolume.DestinationLocal,
			SourceType:          bootfromvolume.SourceImage,
			UUID:                choices.ImageID,
		},
	}

	server, err := CreateBootableVolumeServer(t, client, blockDevices)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	tools.PrintResource(t, server)

	th.AssertEquals(t, server.Image["id"], choices.ImageID)
}

func TestBootFromNewVolume(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	choices, err := clients.AcceptanceTestChoicesFromEnv()
	th.AssertNoErr(t, err)

	blockDevices := []bootfromvolume.BlockDevice{
		{
			DeleteOnTermination: true,
			DestinationType:     bootfromvolume.DestinationVolume,
			SourceType:          bootfromvolume.SourceImage,
			UUID:                choices.ImageID,
			VolumeSize:          2,
		},
	}

	server, err := CreateBootableVolumeServer(t, client, blockDevices)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	attachPages, err := volumeattach.List(client, server.ID).AllPages()
	th.AssertNoErr(t, err)

	attachments, err := volumeattach.ExtractVolumeAttachments(attachPages)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, server)
	tools.PrintResource(t, attachments)

	if server.Image != nil {
		t.Fatalf("server image should be nil")
	}

	th.AssertEquals(t, len(attachments), 1)

	// TODO: volumes_attached extension
}

func TestBootFromExistingVolume(t *testing.T) {
	clients.RequireLong(t)

	computeClient, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	blockStorageClient, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	volume, err := blockstorage.CreateVolumeFromImage(t, blockStorageClient)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, volume)

	blockDevices := []bootfromvolume.BlockDevice{
		{
			DeleteOnTermination: true,
			DestinationType:     bootfromvolume.DestinationVolume,
			SourceType:          bootfromvolume.SourceVolume,
			UUID:                volume.ID,
		},
	}

	server, err := CreateBootableVolumeServer(t, computeClient, blockDevices)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, computeClient, server)

	attachPages, err := volumeattach.List(computeClient, server.ID).AllPages()
	th.AssertNoErr(t, err)

	attachments, err := volumeattach.ExtractVolumeAttachments(attachPages)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, server)
	tools.PrintResource(t, attachments)

	if server.Image != nil {
		t.Fatalf("server image should be nil")
	}

	th.AssertEquals(t, len(attachments), 1)
	th.AssertEquals(t, attachments[0].VolumeID, volume.ID)
	// TODO: volumes_attached extension
}

func TestBootFromMultiEphemeralServer(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	choices, err := clients.AcceptanceTestChoicesFromEnv()
	th.AssertNoErr(t, err)

	blockDevices := []bootfromvolume.BlockDevice{
		{
			BootIndex:           0,
			DestinationType:     bootfromvolume.DestinationLocal,
			DeleteOnTermination: true,
			SourceType:          bootfromvolume.SourceImage,
			UUID:                choices.ImageID,
			VolumeSize:          5,
		},
		{
			BootIndex:           -1,
			DestinationType:     bootfromvolume.DestinationLocal,
			DeleteOnTermination: true,
			GuestFormat:         "ext4",
			SourceType:          bootfromvolume.SourceBlank,
			VolumeSize:          1,
		},
		{
			BootIndex:           -1,
			DestinationType:     bootfromvolume.DestinationLocal,
			DeleteOnTermination: true,
			GuestFormat:         "ext4",
			SourceType:          bootfromvolume.SourceBlank,
			VolumeSize:          1,
		},
	}

	server, err := CreateMultiEphemeralServer(t, client, blockDevices)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	tools.PrintResource(t, server)
}

func TestAttachNewVolume(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	choices, err := clients.AcceptanceTestChoicesFromEnv()
	th.AssertNoErr(t, err)

	blockDevices := []bootfromvolume.BlockDevice{
		{
			BootIndex:           0,
			DeleteOnTermination: true,
			DestinationType:     bootfromvolume.DestinationLocal,
			SourceType:          bootfromvolume.SourceImage,
			UUID:                choices.ImageID,
		},
		{
			BootIndex:           1,
			DeleteOnTermination: true,
			DestinationType:     bootfromvolume.DestinationVolume,
			SourceType:          bootfromvolume.SourceBlank,
			VolumeSize:          2,
		},
	}

	server, err := CreateBootableVolumeServer(t, client, blockDevices)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	attachPages, err := volumeattach.List(client, server.ID).AllPages()
	th.AssertNoErr(t, err)

	attachments, err := volumeattach.ExtractVolumeAttachments(attachPages)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, server)
	tools.PrintResource(t, attachments)

	th.AssertEquals(t, server.Image["id"], choices.ImageID)
	th.AssertEquals(t, len(attachments), 1)

	// TODO: volumes_attached extension
}

func TestAttachExistingVolume(t *testing.T) {
	clients.RequireLong(t)

	computeClient, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	blockStorageClient, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	choices, err := clients.AcceptanceTestChoicesFromEnv()
	th.AssertNoErr(t, err)

	volume, err := blockstorage.CreateVolume(t, blockStorageClient)
	th.AssertNoErr(t, err)

	blockDevices := []bootfromvolume.BlockDevice{
		{
			BootIndex:           0,
			DeleteOnTermination: true,
			DestinationType:     bootfromvolume.DestinationLocal,
			SourceType:          bootfromvolume.SourceImage,
			UUID:                choices.ImageID,
		},
		{
			BootIndex:           1,
			DeleteOnTermination: true,
			DestinationType:     bootfromvolume.DestinationVolume,
			SourceType:          bootfromvolume.SourceVolume,
			UUID:                volume.ID,
		},
	}

	server, err := CreateBootableVolumeServer(t, computeClient, blockDevices)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, computeClient, server)

	attachPages, err := volumeattach.List(computeClient, server.ID).AllPages()
	th.AssertNoErr(t, err)

	attachments, err := volumeattach.ExtractVolumeAttachments(attachPages)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, server)
	tools.PrintResource(t, attachments)

	th.AssertEquals(t, server.Image["id"], choices.ImageID)
	th.AssertEquals(t, len(attachments), 1)
	th.AssertEquals(t, attachments[0].VolumeID, volume.ID)

	// TODO: volumes_attached extension
}
