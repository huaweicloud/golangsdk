package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/kms/v1/crypto"

	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

var CreateDEKOutput = `
{
  "key_id": "e966a300-0c34-4a31-86e1-e67d13e6426a",
  "plain_text": "24570315c5c496f51882abc923430549eb0661fdb387953afc4d3f476d0fab21a90ccbdfabfd6b4f19aaf9c4bec3707cac96f55c4f573a75c8e95b53abf5023c",
  "cipher_text": "02009800c959d21a80123c067ef4bedd126f1959cb86de1aa16ca547c68149e03a631f108e7282219bea130b627c2772726685d1368f4ae3657e4bd5bad8a485ee1f5bac18329b33896c37ea72952e1505765b61c4e4d18f577fb3339b52cfebe19dc387cc38a15fd529b77c19ffed7f02ff583f65393636613330302d306333342d346133312d383665312d65363764313365363432366100000000c622f77e48e630a955feb2686acfd7ea0b9de94dbd537a6a552032bd579e1895"
}
`

var CreateDEKResponse = crypto.CreateDEKResponse{
	KeyId:      "e966a300-0c34-4a31-86e1-e67d13e6426a",
	PlainText:  "24570315c5c496f51882abc923430549eb0661fdb387953afc4d3f476d0fab21a90ccbdfabfd6b4f19aaf9c4bec3707cac96f55c4f573a75c8e95b53abf5023c",
	CipherText: "02009800c959d21a80123c067ef4bedd126f1959cb86de1aa16ca547c68149e03a631f108e7282219bea130b627c2772726685d1368f4ae3657e4bd5bad8a485ee1f5bac18329b33896c37ea72952e1505765b61c4e4d18f577fb3339b52cfebe19dc387cc38a15fd529b77c19ffed7f02ff583f65393636613330302d306333342d346133312d383665312d65363764313365363432366100000000c622f77e48e630a955feb2686acfd7ea0b9de94dbd537a6a552032bd579e1895",
}

func HandleCreateDEKSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/create-datakey", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateDEKOutput)
	})
}

var CreateDEKWithoutPlainTextOutput = `
{
  "key_id": "e966a300-0c34-4a31-86e1-e67d13e6426a",
  "cipher_text": "02009800c959d21a80123c067ef4bedd126f1959cb86de1aa16ca547c68149e03a631f108e7282219bea130b627c2772726685d1368f4ae3657e4bd5bad8a485ee1f5bac18329b33896c37ea72952e1505765b61c4e4d18f577fb3339b52cfebe19dc387cc38a15fd529b77c19ffed7f02ff583f65393636613330302d306333342d346133312d383665312d65363764313365363432366100000000c622f77e48e630a955feb2686acfd7ea0b9de94dbd537a6a552032bd579e1895"
}
`

var CreateDEKWithoutPlainTextResponse = crypto.CreateDEKWithoutPlainTextResponse{
	KeyId:      "e966a300-0c34-4a31-86e1-e67d13e6426a",
	CipherText: "02009800c959d21a80123c067ef4bedd126f1959cb86de1aa16ca547c68149e03a631f108e7282219bea130b627c2772726685d1368f4ae3657e4bd5bad8a485ee1f5bac18329b33896c37ea72952e1505765b61c4e4d18f577fb3339b52cfebe19dc387cc38a15fd529b77c19ffed7f02ff583f65393636613330302d306333342d346133312d383665312d65363764313365363432366100000000c622f77e48e630a955feb2686acfd7ea0b9de94dbd537a6a552032bd579e1895",
}

func HandleCreateWithoutPlainTextSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/create-datakey-without-plaintext", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateDEKWithoutPlainTextOutput)
	})
}

var EncryptDataOutput = `
{
  "key_id": "e966a300-0c34-4a31-86e1-e67d13e6426a",
  "cipher_text": "AgBoADOno72EHR/JjaUrSX7IKbS11ARcP0ZvqR/izGq+eSMqGlfN8QT3om5xbgoeJ4nfeC5n8bxo2XGoAIpcJaiYzj9lOTY2YTMwMC0wYzM0LTRhMzEtODZlMS1lNjdkMTNlNjQyNmEAAAAAoufe9kr8XlgVUSamOMOtsOIpw0Shy1pF1GnL2Wreobo="
}
`

var EncryptDataResponse = crypto.EncryptDataResponse{
	KeyId:      "e966a300-0c34-4a31-86e1-e67d13e6426a",
	CipherText: "AgBoADOno72EHR/JjaUrSX7IKbS11ARcP0ZvqR/izGq+eSMqGlfN8QT3om5xbgoeJ4nfeC5n8bxo2XGoAIpcJaiYzj9lOTY2YTMwMC0wYzM0LTRhMzEtODZlMS1lNjdkMTNlNjQyNmEAAAAAoufe9kr8XlgVUSamOMOtsOIpw0Shy1pF1GnL2Wreobo=",
}

func HandleEncryptDataSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/encrypt-data", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, EncryptDataOutput)
	})
}

var DecryptDataOutput = `
{
    "key_id":"e966a300-0c34-4a31-86e1-e67d13e6426a",
    "plain_text":"ABC"
}
`

var DecryptDataResponse = crypto.DecryptDataResponse{
	KeyId:     "e966a300-0c34-4a31-86e1-e67d13e6426a",
	PlainText: "ABC",
}

func HandleDecryptDataSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/decrypt-data", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DecryptDataOutput)
	})
}

var EncryptDEKOutput = `
{
  "key_id": "e966a300-0c34-4a31-86e1-e67d13e6426a",
  "cipher_text": "AgBoADOno72EHR/JjaUrSX7IKbS11ARcP0ZvqR/izGq+eSMqGlfN8QT3om5xbgoeJ4nfeC5n8bxo2XGoAIpcJaiYzj9lOTY2YTMwMC0wYzM0LTRhMzEtODZlMS1lNjdkMTNlNjQyNmEAAAAAoufe9kr8XlgVUSamOMOtsOIpw0Shy1pF1GnL2Wreobo="
}
`

var EncryptDEKResponse = crypto.EncryptDEKResponse{
	KeyId:      "e966a300-0c34-4a31-86e1-e67d13e6426a",
	CipherText: "AgBoADOno72EHR/JjaUrSX7IKbS11ARcP0ZvqR/izGq+eSMqGlfN8QT3om5xbgoeJ4nfeC5n8bxo2XGoAIpcJaiYzj9lOTY2YTMwMC0wYzM0LTRhMzEtODZlMS1lNjdkMTNlNjQyNmEAAAAAoufe9kr8XlgVUSamOMOtsOIpw0Shy1pF1GnL2Wreobo=",
}

func HandleEncryptDEKSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/encrypt-datakey", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, EncryptDEKOutput)
	})
}

var DecryptDEKOutput = `
{
    "data_key":"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "datakey_length":"64",
    "datakey_dgst":"f5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b"
}
`

var DecryptDEKResponse = crypto.DecryptDEKResponse{
	DataKey:       "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	DatakeyLength: "64",
	DatakeyDgst:   "f5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b",
}

func HandleDecryptDEKSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/decrypt-datakey", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DecryptDEKOutput)
	})
}

var GenerateRandomStringOutput = `
{"random_data":"e9b6bc9e7223390eb16f220e4f4be8741b1e936bd52b4347a52f689e485d939e91b78e31be39c04a57c2ac7d24678c1805619704768af45dcfca028fe83e650c"}
`

var GenerateRandomStringResponse = crypto.GenerateRandomStringResponse{
	RandomData: "e9b6bc9e7223390eb16f220e4f4be8741b1e936bd52b4347a52f689e485d939e91b78e31be39c04a57c2ac7d24678c1805619704768af45dcfca028fe83e650c",
}

func HandleGenerateRandomStringSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/gen-random", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GenerateRandomStringOutput)
	})
}
