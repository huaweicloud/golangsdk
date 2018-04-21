# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/kms/v1.0/crypto"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
Key Management Service (KMS) is a secure, reliable, and easy-to-use service that helps users centrally manage and safeguard their Customer Master Keys (CMKs) and key pairs.

Sample Code, This API allows you to create a DEK. A returned result includes the plaintext and the ciphertext of a DEK.

    
    result, err := crypto.CreateDEK(client, crypto.CreateDEKOpts{
        KeyId:         "e966a300-0c34-4a31-86e1-e67d13e6426a",
        Sequence:      "919c82d4-8046-4722-9094-35c3c6524cff",
        DatakeyLength: "64",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This API allows you to create a plaintext-free DEK, that is, the returned result of this API includes only the plaintext of the DEK.

    
    result, err := crypto.CreateDEKWithoutPlainText(client, crypto.CreateDEKWithoutPlainTextOpts{
        KeyId:         "e966a300-0c34-4a31-86e1-e67d13e6426a",
        Sequence:      "919c82d4-8046-4722-9094-35c3c6524cff",
        DatakeyLength: "512",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This API enables you to decrypt a DEK using a specified CMK.

    
    result, err := crypto.DecryptDEK(client, crypto.DecryptDEKOpts{
        KeyId:               "e966a300-0c34-4a31-86e1-e67d13e6426a",
        CipherText:          "0200980044f1f74e59884b4259ecfdd9149861c93219107895d3aca3afb5ba68991d13679db3736e820d75a17309535b14d6d12796eac84dc4e826ec15ee7db38df0fdb4e97e6c9991f4f043e878387db6d3d48946799f056a8bb9b1952cd73dd1548f2b3939e209df341dd028cb4306925ade0b65393636613330302d306333342d346133312d383665312d65363764313365363432366100000000b90c13a32b15375fbb0f14d6bec4b45d96a328afdb1258747c489e6dbb28a897",
        DatakeyCipherLength: "64",
        Sequence:            "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This API enables you to decrypt data.

    
    result, err := crypto.DecryptData(client, crypto.DecryptDataOpts{
        CipherText: "AgBoAAwOOx+Gi8JbPOvSA0tWgxC11ARcP0ZvqR/izGq+eSMqGlfN8QT3om5xbgoeJ4nfeGK0wcyyvRmpSLvhOyw6J3ZlOTY2YTMwMC0wYzM0LTRhMzEtODZlMS1lNjdkMTNlNjQyNmEAAAAA/XZGoJQFDcRsMwBxoSBuFGb6BwYULbGPN4352ZyZyGw=",
        Sequence:   "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This API enables you to encrypt a DEK using a specified CMK.

    
    result, err := crypto.EncryptDEK(client, crypto.EncryptDEKOpts{
        KeyId:              "e966a300-0c34-4a31-86e1-e67d13e6426a",
        PlainText:          "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000F5A5FD42D16A20302798EF6ED309979B43003D2320D9F0E8EA9831A92759FB4B",
        DatakeyPlainLength: "64",
        Sequence:           "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This API enables you to encrypt data using a specified CMK.

    
    result, err := crypto.EncryptData(client, crypto.EncryptDataOpts{
        KeyId:     "e966a300-0c34-4a31-86e1-e67d13e6426a",
        PlainText: "ABC",
        Sequence:  "919c82d4-8046-4722-9094-35c3c6524cff",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
## Index
**[func CreateDEK(*golangsdk.ServiceClient, CreateDEKOptsBuilder) (CreateDEKResult)](#func-createdek)**  
**[func CreateDEKWithoutPlainText(*golangsdk.ServiceClient, CreateDEKWithoutPlainTextOptsBuilder) (CreateDEKWithoutPlainTextResult)](#func-createdekwithoutplaintext)**  
**[func DecryptDEK(*golangsdk.ServiceClient, DecryptDEKOptsBuilder) (DecryptDEKResult)](#func-decryptdek)**  
**[func DecryptData(*golangsdk.ServiceClient, DecryptDataOptsBuilder) (DecryptDataResult)](#func-decryptdata)**  
**[func EncryptDEK(*golangsdk.ServiceClient, EncryptDEKOptsBuilder) (EncryptDEKResult)](#func-encryptdek)**  
**[func EncryptData(*golangsdk.ServiceClient, EncryptDataOptsBuilder) (EncryptDataResult)](#func-encryptdata)**  
**[func GenerateRandomString(*golangsdk.ServiceClient, GenerateRandomStringOptsBuilder) (GenerateRandomStringResult)](#func-generaterandomstring)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|kms|func CreateDEK(*golangsdk.ServiceClient, CreateDEKOptsBuilder) (CreateDEKResult)|POST /v1.0/{project_id}/kms/create-datakey|
|kms|func CreateDEKWithoutPlainText(*golangsdk.ServiceClient, CreateDEKWithoutPlainTextOptsBuilder) (CreateDEKWithoutPlainTextResult)|POST /v1.0/{project_id}/kms/create-datakey-without-plaintext|
|kms|func DecryptDEK(*golangsdk.ServiceClient, DecryptDEKOptsBuilder) (DecryptDEKResult)|POST /v1.0/{project_id}/kms/decrypt-datakey|
|kms|func DecryptData(*golangsdk.ServiceClient, DecryptDataOptsBuilder) (DecryptDataResult)|POST /v1.0/{project_id}/kms/decrypt-data|
|kms|func EncryptDEK(*golangsdk.ServiceClient, EncryptDEKOptsBuilder) (EncryptDEKResult)|POST /v1.0/{project_id}/kms/encrypt-datakey|
|kms|func EncryptData(*golangsdk.ServiceClient, EncryptDataOptsBuilder) (EncryptDataResult)|POST /v1.0/{project_id}/kms/encrypt-data|
|kms|func GenerateRandomString(*golangsdk.ServiceClient, GenerateRandomStringOptsBuilder) (GenerateRandomStringResult)|POST /v1.0/{project_id}/kms/gen-random|
## Content
## func CreateDEK
    func CreateDEK(*golangsdk.ServiceClient, CreateDEKOptsBuilder) (CreateDEKResult)  
This API allows you to create a DEK. A returned result includes the plaintext and the ciphertext of a DEK.
## func CreateDEKWithoutPlainText
    func CreateDEKWithoutPlainText(*golangsdk.ServiceClient, CreateDEKWithoutPlainTextOptsBuilder) (CreateDEKWithoutPlainTextResult)  
This API allows you to create a plaintext-free DEK, that is, the returned result of this API includes only the plaintext of the DEK.
## func DecryptDEK
    func DecryptDEK(*golangsdk.ServiceClient, DecryptDEKOptsBuilder) (DecryptDEKResult)  
This API enables you to decrypt a DEK using a specified CMK.
## func DecryptData
    func DecryptData(*golangsdk.ServiceClient, DecryptDataOptsBuilder) (DecryptDataResult)  
This API enables you to decrypt data.
## func EncryptDEK
    func EncryptDEK(*golangsdk.ServiceClient, EncryptDEKOptsBuilder) (EncryptDEKResult)  
This API enables you to encrypt a DEK using a specified CMK.
## func EncryptData
    func EncryptData(*golangsdk.ServiceClient, EncryptDataOptsBuilder) (EncryptDataResult)  
This API enables you to encrypt data using a specified CMK.
## func GenerateRandomString
    func GenerateRandomString(*golangsdk.ServiceClient, GenerateRandomStringOptsBuilder) (GenerateRandomStringResult)  
This API allows you to obtain a random number whose length can be up to 512 bits.
