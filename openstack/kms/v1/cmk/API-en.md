# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/kms/v1.0/cmk"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
Key Management Service (KMS) is a secure, reliable, and easy-to-use service that helps users centrally manage and safeguard their Customer Master Keys (CMKs) and key pairs.

Sample Code, This API enables you to cancel the scheduled deletion of a CMK.

    
    result := cmk.CancelDeletion(client, cmk.CancelDeletionOpts{
        KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    })
    
Sample Code, This API is used to create customer master keys (CMKs) used to encrypt data encryption keys (DEKs).

    
    result, err := cmk.Create(client, cmk.CreateOpts{
        KeyAlias:       "TestCMK2",
        KeyDescription: "It's test CMK2.",
        Origin:         "kms",
        Realm:          "",
        KeyPolicy:      "",
        KeyUsage:       "",
        KeyType:        "",
        Sequence:       "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This API allows you to disable a CMK. A disabled CMK cannot be used.

    
    result, err := cmk.Disable(client, cmk.DisableOpts{
        KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    
Sample Code, This API allows you to enable a CMK. Only an enabled CMK can be used.

    
    result, err := cmk.Enable(client, cmk.EnableOpts{
        KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    
Sample Code, This API allows you to query the details about a CMK.

    
    result, err := cmk.Get(client, cmk.GetOpts{
        KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This API is used to query the number of instances, that is, the number of CMKs created.

    
    result, err := cmk.Instances(client).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This API allows you to query the list of all CMKs.

    
    result, err := cmk.List(client, cmk.ListOpts{
        Limit:  "1",
        Marker: "",
    }).AllPages()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This API enables you to schedule the deletion of a CMK. A CMK can be scheduled to be deleted after 7 to 1096 days.

    
    result, err := cmk.ScheduleDeletion(client, cmk.ScheduleDeletionOpts{
        KeyId:       "30361023-62e0-4609-a5fc-6ff8eb63c186",
        PendingDays: "20",
        Sequence:    "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This API enables you to change the alias of a CMK.

    
    result, err := cmk.UpdateAlias(client, cmk.UpdateAliasOpts{
        KeyId:    "e966a300-0c34-4a31-86e1-e67d13e6426a",
        KeyAlias: "TestABC",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This API enables you to change the description of a CMK.

    
    result, err := cmk.UpdateDescription(client, cmk.UpdateDescriptionOpts{
        KeyId:          "e966a300-0c34-4a31-86e1-e67d13e6426a",
        KeyDescription: "TestABC",
        Sequence:       "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
## Index
**[func CancelDeletion(*golangsdk.ServiceClient, CancelDeletionOptsBuilder) (CancelDeletionResult)](#func-canceldeletion)**  
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Disable(*golangsdk.ServiceClient, DisableOptsBuilder) (DisableResult)](#func-disable)**  
**[func Enable(*golangsdk.ServiceClient, EnableOptsBuilder) (EnableResult)](#func-enable)**  
**[func Get(*golangsdk.ServiceClient, GetOptsBuilder) (GetResult)](#func-get)**  
**[func Grant(*golangsdk.ServiceClient, GrantOptsBuilder) (GrantResult)](#func-grant)**  
**[func Instances(*golangsdk.ServiceClient) (InstancesResult)](#func-instances)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
**[func ListGrants(*golangsdk.ServiceClient, ListGrantsOptsBuilder) (ListGrantsResult)](#func-listgrants)**  
**[func ListRetirable(*golangsdk.ServiceClient, ListRetirableOptsBuilder) (ListRetirableResult)](#func-listretirable)**  
**[func Quotas(*golangsdk.ServiceClient) (QuotasResult)](#func-quotas)**  
**[func Retire(*golangsdk.ServiceClient, RetireOptsBuilder) (RetireResult)](#func-retire)**  
**[func Revoke(*golangsdk.ServiceClient, RevokeOptsBuilder) (RevokeResult)](#func-revoke)**  
**[func ScheduleDeletion(*golangsdk.ServiceClient, ScheduleDeletionOptsBuilder) (ScheduleDeletionResult)](#func-scheduledeletion)**  
**[func UpdateAlias(*golangsdk.ServiceClient, UpdateAliasOptsBuilder) (UpdateAliasResult)](#func-updatealias)**  
**[func UpdateDescription(*golangsdk.ServiceClient, UpdateDescriptionOptsBuilder) (UpdateDescriptionResult)](#func-updatedescription)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|kms|func CancelDeletion(*golangsdk.ServiceClient, CancelDeletionOptsBuilder) (CancelDeletionResult)|POST /v1.0/{project_id}/kms/cancel-key-deletion|
|kms|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v1.0/{project_id}/kms/create-key|
|kms|func Disable(*golangsdk.ServiceClient, DisableOptsBuilder) (DisableResult)|POST /v1.0/{project_id}/kms/disable-key|
|kms|func Enable(*golangsdk.ServiceClient, EnableOptsBuilder) (EnableResult)|POST /v1.0/{project_id}/kms/enable-key|
|kms|func Get(*golangsdk.ServiceClient, GetOptsBuilder) (GetResult)|POST /v1.0/{project_id}/kms/describe-key|
|kms|func Grant(*golangsdk.ServiceClient, GrantOptsBuilder) (GrantResult)|POST /v1.0/{project_id}/kms/create-grant|
|kms|func Instances(*golangsdk.ServiceClient) (InstancesResult)|GET /v1.0/{project_id}/kms/user-instances|
|kms|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|POST /v1.0/{project_id}/kms/list-keys|
|kms|func ListGrants(*golangsdk.ServiceClient, ListGrantsOptsBuilder) (ListGrantsResult)|POST /v1.0/{project_id}/kms/list-grants|
|kms|func ListRetirable(*golangsdk.ServiceClient, ListRetirableOptsBuilder) (ListRetirableResult)|POST /v1.0/{project_id}/kms/list-retirable-grants|
|kms|func Quotas(*golangsdk.ServiceClient) (QuotasResult)|GET /v1.0/{project_id}/kms/user-quotas|
|kms|func Retire(*golangsdk.ServiceClient, RetireOptsBuilder) (RetireResult)|POST /v1.0/{project_id}/kms/retire-grant|
|kms|func Revoke(*golangsdk.ServiceClient, RevokeOptsBuilder) (RevokeResult)|POST /v1.0/{project_id}/kms/revoke-grant|
|kms|func ScheduleDeletion(*golangsdk.ServiceClient, ScheduleDeletionOptsBuilder) (ScheduleDeletionResult)|POST /v1.0/{project_id}/kms/schedule-key-deletion|
|kms|func UpdateAlias(*golangsdk.ServiceClient, UpdateAliasOptsBuilder) (UpdateAliasResult)|POST /v1.0/{project_id}/kms/update-key-alias|
|kms|func UpdateDescription(*golangsdk.ServiceClient, UpdateDescriptionOptsBuilder) (UpdateDescriptionResult)|POST /v1.0/{project_id}/kms/update-key-description|
## Content
## func CancelDeletion
    func CancelDeletion(*golangsdk.ServiceClient, CancelDeletionOptsBuilder) (CancelDeletionResult)  
This API enables you to cancel the scheduled deletion of a CMK.
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
This API is used to create customer master keys (CMKs) used to encrypt data encryption keys (DEKs).
## func Disable
    func Disable(*golangsdk.ServiceClient, DisableOptsBuilder) (DisableResult)  
This API allows you to disable a CMK. A disabled CMK cannot be used.
## func Enable
    func Enable(*golangsdk.ServiceClient, EnableOptsBuilder) (EnableResult)  
This API allows you to enable a CMK. Only an enabled CMK can be used.
## func Get
    func Get(*golangsdk.ServiceClient, GetOptsBuilder) (GetResult)  
This API allows you to query the details about a CMK.
## func Grant
    func Grant(*golangsdk.ServiceClient, GrantOptsBuilder) (GrantResult)  
This API enables you to create a grant to grant permissions on a CMK to a user so that the user can perform operations on the CMK.
## func Instances
    func Instances(*golangsdk.ServiceClient) (InstancesResult)  
This API is used to query the number of instances, that is, the number of CMKs created.
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
This API allows you to query the list of all CMKs.
## func ListGrants
    func ListGrants(*golangsdk.ServiceClient, ListGrantsOptsBuilder) (ListGrantsResult)  
This API enables you to query grants on a CMK.
## func ListRetirable
    func ListRetirable(*golangsdk.ServiceClient, ListRetirableOptsBuilder) (ListRetirableResult)  
This API enables you to query grants that can be retired.
## func Quotas
    func Quotas(*golangsdk.ServiceClient) (QuotasResult)  
This API is used to query the quota of a user, that is, the allocated total number of CMKs that can be created by a user and the number of CMKs that has been created by the user.
## func Retire
    func Retire(*golangsdk.ServiceClient, RetireOptsBuilder) (RetireResult)  
This API enables users to retire a grant.
## func Revoke
    func Revoke(*golangsdk.ServiceClient, RevokeOptsBuilder) (RevokeResult)  
This API allows you to revoke a grant.
## func ScheduleDeletion
    func ScheduleDeletion(*golangsdk.ServiceClient, ScheduleDeletionOptsBuilder) (ScheduleDeletionResult)  
This API enables you to schedule the deletion of a CMK. A CMK can be scheduled to be deleted after 7 to 1096 days.
## func UpdateAlias
    func UpdateAlias(*golangsdk.ServiceClient, UpdateAliasOptsBuilder) (UpdateAliasResult)  
This API enables you to change the alias of a CMK.
## func UpdateDescription
    func UpdateDescription(*golangsdk.ServiceClient, UpdateDescriptionOptsBuilder) (UpdateDescriptionResult)  
This API enables you to change the description of a CMK.
