package policies

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type BackupPolicy struct {
	CreatedAt           string                   `json:"created_at"`
	Description         string                   `json:"description"`
	ID                  string                   `json:"id"`
	Name                string                   `json:"name"`
	Parameters          PolicyParam              `json:"parameters"`
	ProjectId           string                   `json:"project_id"`
	ProviderId          string                   `json:"provider_id"`
	Resources           []Resource               `json:"resources"`
	ScheduledOperations []ScheduledOperationResp `json:"scheduled_operations"`
	Status              string                   `json:"status"`
	Tags                []ResourceTag            `json:"tags"`
}

type ScheduledOperationResp struct {
	Description         string              `json:"description"`
	Enabled             bool                `json:"enabled"`
	Name                string              `json:"name"`
	OperationType       string              `json:"operation_type"`
	OperationDefinition OperationDefinition `json:"operation_definition"`
	Trigger             TriggerResp         `json:"trigger"`
	ID                  string              `json:"id"`
	TriggerID           string              `json:"trigger_id"`
}

type TriggerResp struct {
	Properties TriggerPropertiesResp `json:"properties"`
	Name       string                `json:"name"`
	ID         string                `json:"id"`
	Type       string                `json:"type"`
}

type TriggerPropertiesResp struct {
	Pattern   string `json:"pattern"`
	StartTime string `json:"start_time"`
}

// Extract will get the backup policies object from the commonResult
func (r commonResult) Extract() (*BackupPolicy, error) {
	var s struct {
		BackupPolicy *BackupPolicy `json:"policy"`
	}

	err := r.ExtractInto(&s)
	return s.BackupPolicy, err
}

// BackupPolicyPage is the page returned by a pager when traversing over a
// collection of backup policies.
type BackupPolicyPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of backup policies has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r BackupPolicyPage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"policies_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a BackupPolicyPage struct is empty.
func (r BackupPolicyPage) IsEmpty() (bool, error) {
	is, err := ExtractBackupPolicies(r)
	return len(is) == 0, err
}

// ExtractBackupPolicies accepts a Page struct, specifically a BackupPolicyPage struct,
// and extracts the elements into a slice of Policy structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractBackupPolicies(r pagination.Page) ([]BackupPolicy, error) {
	var s struct {
		BackupPolicies []BackupPolicy `json:"policies"`
	}
	err := (r.(BackupPolicyPage)).ExtractInto(&s)
	return s.BackupPolicies, err
}

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type DeleteResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type ListResult struct {
	commonResult
}
