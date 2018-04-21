/*
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

*/
package cmk
