package cloudvolumes

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/iec/v1/common"
)

type commonResult struct {
	golangsdk.Result
}

type GetResult struct {
	commonResult
}

// EdgeVolume 云磁盘
type EdgeVolume struct {
	Volume *common.Volume `json:"volume"`
}

func (r GetResult) Extract() (*EdgeVolume, error) {
	var entity EdgeVolume
	err := r.ExtractIntoStructPtr(&entity, "")
	return &entity, err
}

type VolumeType struct {
	VolumeTypes []common.VolumeType `json:"volume_types"`
}

func (r GetResult) ExtractVolumeType() (*VolumeType, error) {
	var entity VolumeType
	err := r.ExtractIntoStructPtr(&entity, "")
	return &entity, err
}
