package crypto

import (
	"github.com/huaweicloud/golangsdk"
)

type CreateDEKOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// The value of this parameter must be a series of "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}
	EncryptionContext string `json:"encryption_context,omitempty"`

	// Number of bits in the length of a DEK (The maximum number is 512.)
	DatakeyLength string `json:"datakey_length,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type CreateDEKOptsBuilder interface {
	ToCreateDEKMap() (map[string]interface{}, error)
}

func (opts CreateDEKOpts) ToCreateDEKMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateDEK(client *golangsdk.ServiceClient, opts CreateDEKOptsBuilder) (r CreateDEKResult) {
	b, err := opts.ToCreateDEKMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateDEKURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type CreateDEKWithoutPlainTextOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// The value of this parameter must be a series of "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}
	EncryptionContext string `json:"encryption_context,omitempty"`

	// Number of bits in the length of a DEK (The maximum number is 512.)
	DatakeyLength string `json:"datakey_length,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type CreateDEKWithoutPlainTextOptsBuilder interface {
	ToCreateDEKWithoutPlainTextMap() (map[string]interface{}, error)
}

func (opts CreateDEKWithoutPlainTextOpts) ToCreateDEKWithoutPlainTextMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateDEKWithoutPlainText(client *golangsdk.ServiceClient, opts CreateDEKWithoutPlainTextOptsBuilder) (r CreateDEKWithoutPlainTextResult) {
	b, err := opts.ToCreateDEKWithoutPlainTextMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateDEKWithoutPlainTextURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type DecryptDEKOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// The value of this parameter must be a series of "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}
	EncryptionContext string `json:"encryption_context,omitempty"`

	// This parameter indicates the hexadecimal character string of the DEK ciphertext and the metadata. The value is the  value in the encryption result of a DEK.
	CipherText string `json:"cipher_text,"`

	// Number of bytes in the length of a DEK (The maximum number is 64.)
	DatakeyCipherLength string `json:"datakey_cipher_length,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type DecryptDEKOptsBuilder interface {
	ToDecryptDEKMap() (map[string]interface{}, error)
}

func (opts DecryptDEKOpts) ToDecryptDEKMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func DecryptDEK(client *golangsdk.ServiceClient, opts DecryptDEKOptsBuilder) (r DecryptDEKResult) {
	b, err := opts.ToDecryptDEKMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(DecryptDEKURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type DecryptDataOpts struct {

	// Ciphertext of encrypted data. The value is the  value in the data encryption result that matches the regular expression .
	CipherText string `json:"cipher_text,"`

	// The value of this parameter must be a series of "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive information and must be within 8192 characters in length.
	EncryptionContext string `json:"encryption_context,omitempty"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type DecryptDataOptsBuilder interface {
	ToDecryptDataMap() (map[string]interface{}, error)
}

func (opts DecryptDataOpts) ToDecryptDataMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func DecryptData(client *golangsdk.ServiceClient, opts DecryptDataOptsBuilder) (r DecryptDataResult) {
	b, err := opts.ToDecryptDataMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(DecryptDataURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type EncryptDEKOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// The value of this parameter must be a series of "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}
	EncryptionContext string `json:"encryption_context,omitempty"`

	// Both the plaintext (64 bytes) of a DEK and the SHA-256 hash value (32 bytes) of the plaintext are expressed as a hexadecimal character string.
	PlainText string `json:"plain_text,"`

	// Number of bytes of the plaintext of a DEK (The maximum number is 64.)
	DatakeyPlainLength string `json:"datakey_plain_length,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type EncryptDEKOptsBuilder interface {
	ToEncryptDEKMap() (map[string]interface{}, error)
}

func (opts EncryptDEKOpts) ToEncryptDEKMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func EncryptDEK(client *golangsdk.ServiceClient, opts EncryptDEKOptsBuilder) (r EncryptDEKResult) {
	b, err := opts.ToEncryptDEKMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(EncryptDEKURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type EncryptDataOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// The value of this parameter must be a series of "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}
	EncryptionContext string `json:"encryption_context,omitempty"`

	// Plaintext data which is 1 to 4096 bytes in length and matches the regular expression . After being converted into a byte array, it is still 1 to 4096 bytes in length.
	PlainText string `json:"plain_text,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type EncryptDataOptsBuilder interface {
	ToEncryptDataMap() (map[string]interface{}, error)
}

func (opts EncryptDataOpts) ToEncryptDataMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func EncryptData(client *golangsdk.ServiceClient, opts EncryptDataOptsBuilder) (r EncryptDataResult) {
	b, err := opts.ToEncryptDataMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(EncryptDataURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type GenerateRandomStringOpts struct {

	// Number of bits of a random number (The maximum length of a random number is 512 bits.)
	RandomDataLength string `json:"random_data_length,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type GenerateRandomStringOptsBuilder interface {
	ToGenerateRandomStringMap() (map[string]interface{}, error)
}

func (opts GenerateRandomStringOpts) ToGenerateRandomStringMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomString(client *golangsdk.ServiceClient, opts GenerateRandomStringOptsBuilder) (r GenerateRandomStringResult) {
	b, err := opts.ToGenerateRandomStringMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(GenerateRandomStringURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
