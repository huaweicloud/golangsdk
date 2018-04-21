package cmk

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/postpagination"
)

type commonResult struct {
	golangsdk.Result
}

type CancelDeletionResult struct {
	commonResult
}

func (r CancelDeletionResult) Extract() (*CancelDeletionResponse, error) {
	var response CancelDeletionResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CancelDeletionResponse struct {

	// ID of a CMK
	KeyId string `json:"key_id,"`

	// CMK status: 2 indicates that the CMK is enabled. 3 indicates that the CMK is disabled. 4 indicates that the CMK is scheduled for deletion.
	KeyState string `json:"key_state,"`
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*CreateResponse, error) {
	var response CreateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateResponse struct {

	// ID of a user domain
	KeyInfo struct {

		// ID of a CMK
		KeyId string `json:"key_id,"`

		// ID of a user domain
		DomainId string `json:"domain_id,"`
	} `json:"key_info,"`
}

type DisableResult struct {
	commonResult
}

func (r DisableResult) Extract() (*DisableResponse, error) {
	var response DisableResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DisableResponse struct {

	// ID of a user domain
	KeyInfo struct {

		// ID of a CMK
		KeyId string `json:"key_id,"`

		// CMK status: 2 indicates that the CMK is enabled. 3 indicates that the CMK is disabled. 4 indicates that the CMK is scheduled for deletion.
		KeyState string `json:"key_state,"`
	} `json:"key_info,"`
}

type EnableResult struct {
	commonResult
}

func (r EnableResult) Extract() (*EnableResponse, error) {
	var response EnableResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type EnableResponse struct {

	// ID of a user domain
	KeyInfo struct {

		// ID of a CMK
		KeyId string `json:"key_id,"`

		// CMK status: 2 indicates that the CMK is enabled. 3 indicates that the CMK is disabled. 4 indicates that the CMK is scheduled for deletion.
		KeyState string `json:"key_state,"`
	} `json:"key_info,"`
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*GetResponse, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetResponse struct {

	// ID of a user domain
	KeyInfo struct {

		// ID of a CMK
		KeyId string `json:"key_id,"`

		// ID of a user domain
		DomainId string `json:"domain_id,"`

		// Alias of a CMK
		KeyAlias string `json:"key_alias,"`

		// Region where a CMK resides
		Realm string `json:"realm,"`

		// Description of a CMK
		KeyDescription string `json:"key_description,"`

		// Creation time (time stamp) of a CMK
		CreationDate string `json:"creation_date,"`

		// Scheduled deletion time (time stamp) of a CMK
		ScheduledDeletionDate string `json:"scheduled_deletion_date,"`

		// State of a CMK: 1 indicates that the CMK is waiting to be activated. 2 indicates that the CMK is enabled. 3 indicates that the CMK is disabled. 4 indicates that the CMK is scheduled for deletion. 5 indicates that the CMK is waiting to be imported.
		KeyState string `json:"key_state,"`

		// Identification of a Master Key. The value  indicates a Default Master Key, and the value  indicates a CMK.
		DefaultKeyFlag string `json:"default_key_flag,"`

		// Type of a CMK
		KeyType string `json:"key_type,"`

		// Expiration time
		ExpirationTime string `json:"expiration_time,"`

		// Origin of a CMK. The default value is . The following values are enumerated:
		Origin string `json:"origin,"`

		// Key rotation status. The default value is false, indicating that key rotation is disabled.
		KeyRotationEnabled string `json:"key_rotation_enabled,"`
	} `json:"key_info,"`
}

type GrantResult struct {
	commonResult
}

func (r GrantResult) Extract() (*GrantResponse, error) {
	var response GrantResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GrantResponse struct {
}

type InstancesResult struct {
	commonResult
}

func (r InstancesResult) Extract() (*InstancesResponse, error) {
	var response InstancesResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type InstancesResponse struct {

	// Number of CMKs
	InstanceNum int `json:"instance_num,"`
}

type ListPage struct {
	postpagination.PostMarkerPageBase
}

func (r ListPage) IsEmpty() (bool, error) {
	response, err := ExtractList(r)
	return len(response.Keys) == 0 && len(response.KeyDetails) == 0, err
}

func ExtractList(r ListPage) (*ListResponse, error) {
	var list ListResponse
	r.ExtractInto(&list)
	return &list, r.Err
}

type ListResponse struct {

	// List of CMK IDs ()
	Keys []string `json:"keys,"`

	// List of CMK details. For format details, see .
	KeyDetails []struct {

		// ID of a CMK
		KeyId string `json:"key_id,"`

		// ID of a user domain
		DomainId string `json:"domain_id,"`

		// Alias of a CMK
		KeyAlias string `json:"key_alias,"`

		// Region where a CMK resides
		Realm string `json:"realm,"`

		// Description of a CMK
		KeyDescription string `json:"key_description,"`

		// Creation time (time stamp) of a CMK
		CreationDate string `json:"creation_date,"`

		// Scheduled deletion time (time stamp) of a CMK
		ScheduledDeletionDate string `json:"scheduled_deletion_date,"`

		// State of a CMK: 1 indicates that the CMK is waiting to be activated. 2 indicates that the CMK is enabled. 3 indicates that the CMK is disabled. 4 indicates that the CMK is scheduled for deletion. 5 indicates that the CMK is waiting to be imported.
		KeyState string `json:"key_state,"`

		// Identification of a Master Key. The value  indicates a Default Master Key, and the value  indicates a CMK.
		DefaultKeyFlag string `json:"default_key_flag,"`

		// Type of a CMK
		KeyType string `json:"key_type,"`

		// Expiration time
		ExpirationTime string `json:"expiration_time,"`

		// Origin of a CMK. The default value is . The following values are enumerated:
		Origin string `json:"origin,"`

		// Key rotation status. The default value is false, indicating that key rotation is disabled.
		KeyRotationEnabled string `json:"key_rotation_enabled,"`
	} `json:"key_details,"`

	// This parameter indicates the  value required for obtaining the next page of query results. If the  value is , the  parameter is left blank.
	NextMarker string `json:"next_marker,"`

	// This parameter indicates whether there are more results displayed in another page. If the value is true, there are more results. If the value is false, the current page is the last page.
	Truncated string `json:"truncated,"`
}

type ListGrantsResult struct {
	commonResult
}

func (r ListGrantsResult) Extract() (*ListGrantsResponse, error) {
	var response ListGrantsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListGrantsResponse struct {

	// List of grants
	Grants []struct {

		// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
		KeyId string `json:"key_id,"`

		// 64-byte ID of a grant that meets the regular expression
		GrantId string `json:"grant_id,"`

		// 32-byte ID of a user to which permissions are granted that matches the regular expression
		GranteePrincipal string `json:"grantee_principal,"`

		// Permissions that can be granted. Values: , , , , , , , ,
		Operations int `json:"operations,"`

		// 32-byte ID of a user who created a grant that matches the regular expression
		IssuingPrincipal string `json:"issuing_principal,"`

		// Creation time (time stamp) of a grant
		CreationDate string `json:"creation_date,"`

		// Name of a grant which can be 1 to 255 characters in length and matches the regular expression
		Name string `json:"name,omitempty"`

		// 32-byte ID of a user who can retire a grant that matches the regular expression
		RetiringPrincipal string `json:"retiring_principal,omitempty"`
	} `json:"grants,"`

	// This parameter indicates the  value required for obtaining the next page of query results.
	NextMarker string `json:"next_marker,"`

	//
	Truncated string `json:"truncated,"`

	// This parameter indicates the total number of grants.
	Total int `json:"total,"`
}

type ListRetirableResult struct {
	commonResult
}

func (r ListRetirableResult) Extract() (*ListRetirableResponse, error) {
	var response ListRetirableResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListRetirableResponse struct {

	// List of grants. For format details, see the response returned from the grant querying interface.
	Grants []struct {
	} `json:"grants,"`

	// This parameter indicates the  value required for obtaining the next page of query results.
	NextMarker string `json:"next_marker,"`

	//
	Truncated string `json:"truncated,"`

	// This parameter indicates the total number of grants.
	Total int `json:"total,"`
}

type QuotasResult struct {
	commonResult
}

func (r QuotasResult) Extract() (*QuotasResponse, error) {
	var response QuotasResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type QuotasResponse struct {

	// Quota list
	Quotas struct {

		// Resource quota list
		Resources []struct {

			// Quota type
			Type string `json:"type,"`

			// Used quotas
			Used int `json:"used,"`

			// Total quotas
			Quota int `json:"quota,"`
		} `json:"resources,"`
	} `json:"quotas,"`
}

type RetireResult struct {
	commonResult
}

func (r RetireResult) Extract() (*RetireResponse, error) {
	var response RetireResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type RetireResponse struct {
}

type RevokeResult struct {
	commonResult
}

func (r RevokeResult) Extract() (*RevokeResponse, error) {
	var response RevokeResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type RevokeResponse struct {
}

type ScheduleDeletionResult struct {
	commonResult
}

func (r ScheduleDeletionResult) Extract() (*ScheduleDeletionResponse, error) {
	var response ScheduleDeletionResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ScheduleDeletionResponse struct {

	// ID of a CMK
	KeyId string `json:"key_id,"`

	// CMK status: 2 indicates that the CMK is enabled. 3 indicates that the CMK is disabled. 4 indicates that the CMK is scheduled for deletion.
	KeyState string `json:"key_state,"`
}

type UpdateAliasResult struct {
	commonResult
}

func (r UpdateAliasResult) Extract() (*UpdateAliasResponse, error) {
	var response UpdateAliasResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type UpdateAliasResponse struct {

	// ID of a user domain
	KeyInfo struct {

		// ID of a CMK
		KeyId string `json:"key_id,"`

		// Alias of a CMK
		KeyAlias string `json:"key_alias,"`
	} `json:"key_info,"`
}

type UpdateDescriptionResult struct {
	commonResult
}

func (r UpdateDescriptionResult) Extract() (*UpdateDescriptionResponse, error) {
	var response UpdateDescriptionResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type UpdateDescriptionResponse struct {

	// ID of a user domain
	KeyInfo struct {

		// ID of a CMK
		KeyId string `json:"key_id,"`

		// Description of a CMK
		KeyDescription string `json:"key_description,"`
	} `json:"key_info,"`
}
