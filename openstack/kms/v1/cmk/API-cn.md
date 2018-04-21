# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/kms/v1.0/cmk"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
密钥管理服务（Key Management Service）通过简单、便捷的密钥管理界面，提供易用、高安全的云上加密及密钥管理功能，并为您的其它云服务提供加密特性，让您安心使用云服务，专注业务核心的开发。

示例代码, 取消计划删除密钥。

    
    result := cmk.CancelDeletion(client, cmk.CancelDeletionOpts{
        KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    })
    
示例代码, 创建用户主密钥，可用来加密数据密钥。

    
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
    
示例代码, 禁用密钥，密钥禁用后不可以使用。

    
    result, err := cmk.Disable(client, cmk.DisableOpts{
        KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    
示例代码, 启用密钥，密钥启用后才可以使用。

    
    result, err := cmk.Enable(client, cmk.EnableOpts{
        KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    
示例代码, 查询密钥详细信息。

    
    result, err := cmk.Get(client, cmk.GetOpts{
        KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询实例数，获取用户已经创建的用户主密钥数量。

    
    result, err := cmk.Instances(client).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询用户所有密钥列表。

    
    result, err := cmk.List(client, cmk.ListOpts{
        Limit:  "1",
        Marker: "",
    }).AllPages()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 计划多少天后删除密钥，可设置7天～1096天内删除密钥。

    
    result, err := cmk.ScheduleDeletion(client, cmk.ScheduleDeletionOpts{
        KeyId:       "30361023-62e0-4609-a5fc-6ff8eb63c186",
        PendingDays: "20",
        Sequence:    "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 修改用户主密钥别名。

    
    result, err := cmk.UpdateAlias(client, cmk.UpdateAliasOpts{
        KeyId:    "e966a300-0c34-4a31-86e1-e67d13e6426a",
        KeyAlias: "TestABC",
        Sequence: "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
示例代码, 修改用户主密钥描述信息。

    
    result, err := cmk.UpdateDescription(client, cmk.UpdateDescriptionOpts{
        KeyId:          "e966a300-0c34-4a31-86e1-e67d13e6426a",
        KeyDescription: "TestABC",
        Sequence:       "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
## 目录
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
## API对应表
|类别|API|EndPoint|
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
## 开始
## func CancelDeletion
    func CancelDeletion(*golangsdk.ServiceClient, CancelDeletionOptsBuilder) (CancelDeletionResult)  
取消计划删除密钥。
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
创建用户主密钥，可用来加密数据密钥。
## func Disable
    func Disable(*golangsdk.ServiceClient, DisableOptsBuilder) (DisableResult)  
禁用密钥，密钥禁用后不可以使用。
## func Enable
    func Enable(*golangsdk.ServiceClient, EnableOptsBuilder) (EnableResult)  
启用密钥，密钥启用后才可以使用。
## func Get
    func Get(*golangsdk.ServiceClient, GetOptsBuilder) (GetResult)  
查询密钥详细信息。
## func Grant
    func Grant(*golangsdk.ServiceClient, GrantOptsBuilder) (GrantResult)  
创建授权，被授权用户可以对授权密钥进行操作。
## func Instances
    func Instances(*golangsdk.ServiceClient) (InstancesResult)  
查询实例数，获取用户已经创建的用户主密钥数量。
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
查询用户所有密钥列表。
## func ListGrants
    func ListGrants(*golangsdk.ServiceClient, ListGrantsOptsBuilder) (ListGrantsResult)  
查询密钥的授权列表。
## func ListRetirable
    func ListRetirable(*golangsdk.ServiceClient, ListRetirableOptsBuilder) (ListRetirableResult)  
查询用户可以退役的授权列表。
## func Quotas
    func Quotas(*golangsdk.ServiceClient) (QuotasResult)  
查询配额，查询用户可以创建的用户主密钥配额总数及当前使用量信息。
## func Retire
    func Retire(*golangsdk.ServiceClient, RetireOptsBuilder) (RetireResult)  
退役授权，表示被授权用户不再具有授权密钥的操作权。
## func Revoke
    func Revoke(*golangsdk.ServiceClient, RevokeOptsBuilder) (RevokeResult)  
撤销授权，授权用户撤销被授权用户操作密钥的权限。
## func ScheduleDeletion
    func ScheduleDeletion(*golangsdk.ServiceClient, ScheduleDeletionOptsBuilder) (ScheduleDeletionResult)  
计划多少天后删除密钥，可设置7天～1096天内删除密钥。
## func UpdateAlias
    func UpdateAlias(*golangsdk.ServiceClient, UpdateAliasOptsBuilder) (UpdateAliasResult)  
修改用户主密钥别名。
## func UpdateDescription
    func UpdateDescription(*golangsdk.ServiceClient, UpdateDescriptionOptsBuilder) (UpdateDescriptionResult)  
修改用户主密钥描述信息。
