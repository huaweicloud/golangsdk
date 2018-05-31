package instances

import (
	"github.com/huaweicloud/golangsdk"
)

// InstanceCreate response
type InstanceCreate struct {
	InstanceID string `json:"instance_id"`
}

// CreateResult is a struct that contains all the return parameters of creation
type CreateResult struct {
	golangsdk.Result
}

// Extract from CreateResult
func (r CreateResult) Extract() (*InstanceCreate, error) {
	var s InstanceCreate
	err := r.Result.ExtractInto(&s)
	return &s, err
}

// DeleteResult is a struct which contains the result of deletion
type DeleteResult struct {
	golangsdk.ErrResult
}

// Instance response
type Instance struct {
	Name                 string               `json:"name"`
	Engine               string               `json:"engine"`
	Capacity             int                  `json:"capacity"`
	IP                   string               `json:"ip"`
	Port                 int                  `json:"port"`
	Status               string               `json:"status"`
	Description          string               `json:"description"`
	InstanceID           string               `json:"instance_id"`
	ResourceSpecCode     string               `json:"resource_spec_code"`
	EngineVersion        string               `json:"engine_version"`
	InternalVersion      string               `json:"internal_version"`
	ChargingMode         int                  `json:"charging_mode"`
	VPCID                string               `json:"vpc_id"`
	VPCName              string               `json:"vpc_name"`
	CreatedAt            string               `json:"created_at"`
	ErrorCode            string               `json:"error_code"`
	ProductID            string               `json:"product_id"`
	SecurityGroupID      string               `json:"security_group_id"`
	SecurityGroupName    string               `json:"security_group_name"`
	SubnetID             string               `json:"subnet_id"`
	SubnetName           string               `json:"subnet_name"`
	SubnetCIDR           string               `json:"subnet_cidr"`
	AvailableZones       []string             `json:"available_zones"`
	MaxMemory            int                  `json:"max_memory"`
	UsedMemory           int                  `json:"used_memory"`
	InstanceBackupPolicy InstanceBackupPolicy `json:"instance_backup_policy"`
	UserID               string               `json:"user_id"`
	UserName             string               `json:"user_name"`
	OrderID              string               `json:"order_id"`
	MaintainBegin        string               `json:"maintain_begin"`
	MaintainEnd          string               `json:"maintain_end"`
}

// UpdateResult is a struct from which can get the result of update method
type UpdateResult struct {
	golangsdk.Result
}

// GetResult contains the body of getting detailed
type GetResult struct {
	golangsdk.Result
}

// Extract from GetResult
func (r GetResult) Extract() (*Instance, error) {
	var s Instance
	err := r.Result.ExtractInto(&s)
	return &s, err
}
