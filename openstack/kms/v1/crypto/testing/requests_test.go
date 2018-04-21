package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/kms/v1/crypto"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateDEK(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateDEKSuccessfully(t)

	actual, err := crypto.CreateDEK(client.ServiceClient(), crypto.CreateDEKOpts{
		KeyId:         "e966a300-0c34-4a31-86e1-e67d13e6426a",
		Sequence:      "919c82d4-8046-4722-9094-35c3c6524cff",
		DatakeyLength: "512",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateDEKResponse, actual)
}

func TestCreateDEKWithoutPlainText(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateWithoutPlainTextSuccessfully(t)

	actual, err := crypto.CreateDEKWithoutPlainText(client.ServiceClient(), crypto.CreateDEKWithoutPlainTextOpts{
		KeyId:         "e966a300-0c34-4a31-86e1-e67d13e6426a",
		Sequence:      "919c82d4-8046-4722-9094-35c3c6524cff",
		DatakeyLength: "512",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateDEKWithoutPlainTextResponse, actual)
}

func TestEncryptData(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleEncryptDataSuccessfully(t)

	actual, err := crypto.EncryptData(client.ServiceClient(), crypto.EncryptDataOpts{
		KeyId:     "e966a300-0c34-4a31-86e1-e67d13e6426a",
		PlainText: "ABC",
		Sequence:  "919c82d4-8046-4722-9094-35c3c6524cff",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &EncryptDataResponse, actual)
}

func TestDecryptData(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDecryptDataSuccessfully(t)

	actual, err := crypto.DecryptData(client.ServiceClient(), crypto.DecryptDataOpts{
		CipherText: "AgBoAAwOOx+Gi8JbPOvSA0tWgxC11ARcP0ZvqR/izGq+eSMqGlfN8QT3om5xbgoeJ4nfeGK0wcyyvRmpSLvhOyw6J3ZlOTY2YTMwMC0wYzM0LTRhMzEtODZlMS1lNjdkMTNlNjQyNmEAAAAA/XZGoJQFDcRsMwBxoSBuFGb6BwYULbGPN4352ZyZyGw=",
		Sequence:   "919c82d4-8046-4722-9094-35c3c6524cff",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &DecryptDataResponse, actual)
}

func TestEncryptDEK(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleEncryptDEKSuccessfully(t)

	actual, err := crypto.EncryptDEK(client.ServiceClient(), crypto.EncryptDEKOpts{
		KeyId:              "e966a300-0c34-4a31-86e1-e67d13e6426a",
		PlainText:          "ABC",
		DatakeyPlainLength: "64",
		Sequence:           "919c82d4-8046-4722-9094-35c3c6524cff",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &EncryptDEKResponse, actual)
}

func TestDecryptDEK(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDecryptDEKSuccessfully(t)

	actual, err := crypto.DecryptDEK(client.ServiceClient(), crypto.DecryptDEKOpts{
		KeyId:               "e966a300-0c34-4a31-86e1-e67d13e6426a",
		CipherText:          "0200980044f1f74e59884b4259ecfdd9149861c93219107895d3aca3afb5ba68991d13679db3736e820d75a17309535b14d6d12796eac84dc4e826ec15ee7db38df0fdb4e97e6c9991f4f043e878387db6d3d48946799f056a8bb9b1952cd73dd1548f2b3939e209df341dd028cb4306925ade0b65393636613330302d306333342d346133312d383665312d65363764313365363432366100000000b90c13a32b15375fbb0f14d6bec4b45d96a328afdb1258747c489e6dbb28a897",
		DatakeyCipherLength: "64",
		Sequence:            "919c82d4-8046-4722-9094-35c3c6524cff",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &DecryptDEKResponse, actual)
}

func TestGenerateRandomString(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGenerateRandomStringSuccessfully(t)

	actual, err := crypto.GenerateRandomString(client.ServiceClient(), crypto.GenerateRandomStringOpts{
		RandomDataLength: "512",
		Sequence:         "919c82d4-8046-4722-9094-35c3c6524cff",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GenerateRandomStringResponse, actual)
}
