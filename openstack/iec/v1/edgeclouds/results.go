package edgeclouds

import (
	"time"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/iec/v1/common"
	"github.com/huaweicloud/golangsdk/openstack/iec/v1/servers"
)

// EdgeCloud 边缘业务详情
type EdgeCloud struct {
	ID             string          `json:"id,omitempty"`
	Name           string          `json:"name,omitempty"`
	Description    string          `json:"description,omitempty"`
	StackOptsSlice []StackDetail   `json:"stacks"`
	Coverage       common.Coverage `json:"coverage"`
	DeploymentIDs  []string        `json:"-"`
	ServerCount    int             `json:"-"`
	SuccessNum     int             `json:"success_num"`
	FailedNum      int             `json:"failed_num"`
	Status         string          `json:"status"`
	FailMessage    FailReason      `json:"fail_reason,omitempty"`
	// 边缘业务支持的边缘区域数目，等同于边缘业务下所有实例的区域数目总和
	EdgeRegions int `json:"edge_regions,omitempty"`
	TimeModel
}

// TimeModel record the time of creating, updating, deleting
type TimeModel struct {
	CreateAt time.Time `orm:"column(CREATEAT);type(datetime);auto_now_add" json:"create_at,omitempty"`
	UpdateAt time.Time `orm:"column(UPDATEAT);type(datetime);auto_now" json:"update_at,omitempty"`
	DeleteAt time.Time `orm:"column(DELETEAT);type(datetime)" json:"-"`
	Deleted  string    `orm:"column(DELETED)" json:"-"`
}

// FailReason 失败缘由，只有失败的边缘业务中会显示
type FailReason struct {
	ErrorCode    string `json:"fail_code,omitempty"`
	ErrorMessage string `json:"fail_message,omitempty"`
}

// StackDetail Stack详情
type StackDetail struct {
	//ID
	ID string `json:"id"`

	//NAME
	Name string `json:"name"`

	Resources SliceResourceOptsField `json:"resources"`
}

// SliceResourceOptsField A slice string field.
type SliceResourceOptsField []servers.CreateOpts

type commonResult struct {
	golangsdk.Result
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*EdgeCloud, error) {
	var entity EdgeCloud
	err := r.ExtractIntoStructPtr(&entity, "")
	return &entity, err
}

type DeleteResult struct {
	golangsdk.ErrResult
}
