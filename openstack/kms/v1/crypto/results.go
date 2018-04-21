package crypto

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type CreateDEKResult struct {
	commonResult
}

func (r CreateDEKResult) Extract() (*CreateDEKResponse, error) {
	var response CreateDEKResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateDEKResponse struct {

	// ID of a CMK
	KeyId string `json:"key_id,"`

	// The plaintext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
	PlainText string `json:"plain_text,"`

	// The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
	CipherText string `json:"cipher_text,"`
}

type CreateDEKWithoutPlainTextResult struct {
	commonResult
}

func (r CreateDEKWithoutPlainTextResult) Extract() (*CreateDEKWithoutPlainTextResponse, error) {
	var response CreateDEKWithoutPlainTextResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateDEKWithoutPlainTextResponse struct {

	// ID of a CMK
	KeyId string `json:"key_id,"`

	// The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
	CipherText string `json:"cipher_text,"`
}

type DecryptDEKResult struct {
	commonResult
}

func (r DecryptDEKResult) Extract() (*DecryptDEKResponse, error) {
	var response DecryptDEKResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DecryptDEKResponse struct {

	// Hexadecimal character string of the plaintext of a DEK
	DataKey string `json:"data_key,"`

	// Number of bytes in the length of the plaintext of a DEK
	DatakeyLength string `json:"datakey_length,"`

	// Hexadecimal character string corresponding to the SHA-256 hash value of the plaintext of a DEK
	DatakeyDgst string `json:"datakey_dgst,"`
}

type DecryptDataResult struct {
	commonResult
}

func (r DecryptDataResult) Extract() (*DecryptDataResponse, error) {
	var response DecryptDataResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DecryptDataResponse struct {

	// ID of a CMK
	KeyId string `json:"key_id,"`

	// Plaintext of data
	PlainText string `json:"plain_text,"`
}

type EncryptDEKResult struct {
	commonResult
}

func (r EncryptDEKResult) Extract() (*EncryptDEKResponse, error) {
	var response EncryptDEKResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type EncryptDEKResponse struct {

	// ID of a CMK
	KeyId string `json:"key_id,"`

	// The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
	CipherText string `json:"cipher_text,"`

	// Number of bytes in the length of a DEK
	DatakeyLength string `json:"datakey_length,"`
}

type EncryptDataResult struct {
	commonResult
}

func (r EncryptDataResult) Extract() (*EncryptDataResponse, error) {
	var response EncryptDataResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type EncryptDataResponse struct {

	// ID of a CMK
	KeyId string `json:"key_id,"`

	// Ciphertext data in Base64 format
	CipherText string `json:"cipher_text,"`
}

type GenerateRandomStringResult struct {
	commonResult
}

func (r GenerateRandomStringResult) Extract() (*GenerateRandomStringResponse, error) {
	var response GenerateRandomStringResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GenerateRandomStringResponse struct {

	// Random numbers are expressed in hexadecimal format. Two characters indicate one byte. Length of a random number must be consistent with the  value entered by a user.
	RandomData string `json:"random_data,"`
}
