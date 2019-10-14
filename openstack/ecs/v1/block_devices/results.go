package block_devices

import "github.com/huaweicloud/golangsdk"

type VolumeAttachment struct {
	PciAddress string `json:"pciAddress"`
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*VolumeAttachment, error) {
	s := &VolumeAttachment{}
	return s, r.ExtractInto(s)
}
