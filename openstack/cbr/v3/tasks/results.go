package tasks

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type commonResult struct {
	golangsdk.Result
}

type GetResult struct {
	commonResult
}

type TaskPage struct {
	pagination.SinglePageBase
}

type OperationLog struct {
	CheckpointID  string      `json:"checkpoint_id"`
	CreatedAt     string      `json:"created_at"`
	EndedAt       string      `json:"ended_at"`
	ErrorInfo     OpErrorInfo `json:"error_info"`
	ExtraInfo     OpExtraInfo `json:"extra_info"`
	ID            string      `json:"id"`
	OperationType string      `json:"operation_type"`
	PolicyID      string      `json:"policy_id"`
	ProjectID     string      `json:"project_id"`
	ProviderID    string      `json:"provider_id"`
	StartedAt     string      `json:"started_at"`
	Status        string      `json:"status"`
	UpdatedAt     string      `json:"updated_at"`
	VaultID       string      `json:"vault_id"`
	VaultName     string      `json:"vault_name"`
}

type OpErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type OpExtraInfo struct {
	Backup          OpExtendInfoBackup          `json:"backup"`
	Common          OpExtendInfoCommon          `json:"common"`
	Delete          OpExtendInfoDelete          `json:"delete"`
	Sync            OpExtendInfoSync            `json:"sync"`
	RemoveResources OpExtendInfoRemoveResources `json:"remove_resources"`
	Replication     OpExtendInfoReplication     `json:"replication"`
	Resource        Resource                    `json:"resource"`
	Restore         OpExtendInfoRestore         `json:"restore"`
	VaultDelete     OpExtendInfoVaultDelete     `json:"vault_delete"`
}

type OpExtendInfoBackup struct {
	AppConsistencyErrorCode    string `json:"app_consistency_error_code"`
	AppConsistencyErrorMessage string `json:"app_consistency_error_message"`
	AppConsistencyStatus       string `json:"app_consistency_status"`
	BackupID                   string `json:"backup_id"`
	BackupName                 string `json:"backup_name"`
	Incremental                string `json:"incremental"`
}

type OpExtendInfoCommon struct {
	Progress  int    `json:"progress"`
	RequestID string `json:"request_id"`
	TaskID    string `json:"task_id"`
}

type OpExtendInfoDelete struct {
	BackupID   string `json:"backup_id"`
	BackupName string `json:"backup_name"`
}

type OpExtendInfoSync struct {
	SyncBackupNum    int `json:"sync_backup_num"`
	DeleteBackupNum  int `json:"delete_backup_num"`
	ErrSyncBackupNum int `json:"err_sync_backup_num"`
}

type OpExtendInfoRemoveResources struct {
	FailCount  int        `json:"fail_count"`
	TotalCount int        `json:"total_count"`
	Resources  []Resource `json:"resources"`
}

type Resource struct {
	ExtraInfo ResourceExtraInfo `json:"extra_info"`
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Type      string            `json:"type"`
}

type OpExtendInfoReplication struct {
	DestinationBackupID     string `json:"destination_backup_id"`
	DestinationCheckpointID string `json:"destination_checkpoint_id"`
	DestinationProjectID    string `json:"destination_project_id"`
	DestinationRegion       string `json:"destination_region"`
	SourceBackupID          string `json:"source_backup_id"`
	SourceCheckpointID      string `json:"source_checkpoint_id"`
	SourceProjectID         string `json:"source_project_id"`
	SourceRegion            string `json:"source_region"`
	SourceBackupName        string `json:"source_backup_name"`
	DestinationBackupName   string `json:"destination_backup_name"`
}

type ResourceExtraInfo struct {
	ExcludeVolumes []string                          `json:"exclude_volumes"`
	IncludeVolumes []ResourceExtraInfoIncludeVolumes `json:"include_volumes"`
}

type ResourceExtraInfoIncludeVolumes struct {
	ID        string `json:"id"`
	OsVersion string `json:"os_version"`
}

type OpExtendInfoRestore struct {
	BackupID           string `json:"backup_id"`
	BackupName         string `json:"backup_name"`
	TargetResourceId   string `json:"target_resource_id"`
	TargetResourceName string `json:"target_resource_name"`
}

type OpExtendInfoVaultDelete struct {
	FailCount  int `json:"fail_count"`
	TotalCount int `json:"total_count"`
}

func (r commonResult) Extract() (*OperationLog, error) {
	var s struct {
		Operation *OperationLog `json:"operation_log"`
	}
	err := r.ExtractInto(&s)
	return s.Operation, err
}

func ExtractTasks(r pagination.Page) (*[]OperationLog, error) {
	var s struct {
		OperationLog []OperationLog `json:"operation_logs"`
	}
	err := r.(TaskPage).Result.ExtractInto(&s)
	return &s.OperationLog, err
}
