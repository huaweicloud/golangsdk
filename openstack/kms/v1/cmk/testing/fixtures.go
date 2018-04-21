package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/kms/v1/cmk"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	"encoding/json"
	"io"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "key_info": {
    "key_id": "30361023-62e0-4609-a5fc-6ff8eb63c186",
    "domain_id": "0984aafba48049a6b9457b89968eeabf"
  }
}
`

var CreateResponse = cmk.CreateResponse{
	KeyInfo: struct {
		// ID of a CMK
		KeyId string `json:"key_id,"`
		// ID of a user domain
		DomainId string `json:"domain_id,"`
	}{
		KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
		DomainId: "0984aafba48049a6b9457b89968eeabf",
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/create-key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var EnableOutput = `
{
  "key_info": {
    "key_id": "30361023-62e0-4609-a5fc-6ff8eb63c186",
    "key_state": "2"
  }
}
`

var EnableResponse = cmk.EnableResponse{
	KeyInfo: struct {
		// ID of a CMK
		KeyId string `json:"key_id,"`

		// Status of a CMK
		KeyState string `json:"key_state,"`
	}{
		KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
		KeyState: "2",
	},
}

func HandleEnableSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/enable-key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, EnableOutput)
	})
}

var DisableOutput = `
{
  "key_info": {
    "key_id": "30361023-62e0-4609-a5fc-6ff8eb63c186",
    "key_state": "3"
  }
}
`

var DisableResponse = cmk.DisableResponse{
	KeyInfo: struct {
		// ID of a CMK
		KeyId string `json:"key_id,"`

		// Status of a CMK
		KeyState string `json:"key_state,"`
	}{
		KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
		KeyState: "3",
	},
}

func HandleDisableSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/disable-key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DisableOutput)
	})
}

var GetOutput = `
{
  "key_info": {
    "key_id": "30361023-62e0-4609-a5fc-6ff8eb63c186",
    "domain_id": "0984aafba48049a6b9457b89968eeabf",
    "key_alias": "TestCMK2",
    "realm": "cn-north-1",
    "key_description": "It's test CMK2.",
    "creation_date": "1522573239000",
    "scheduled_deletion_date": "",
    "expiration_time": "",
    "origin": "kms",
    "key_state": "3",
    "default_key_flag": "0",
    "key_type": "1",
    "key_rotation_enabled": "false"
  }
}
`

var GetResponse = cmk.GetResponse{
	KeyInfo: struct {
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

		//
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
	}{
		//30361023-62e0-4609-a5fc-6ff8eb63c186 0984aafba48049a6b9457b89968eeabf TestCMK2 cn-north-1 It's test CMK2. 1522573239000  3 0 1  kms false
		KeyId:                 "30361023-62e0-4609-a5fc-6ff8eb63c186",
		DomainId:              "0984aafba48049a6b9457b89968eeabf",
		KeyAlias:              "TestCMK2",
		Realm:                 "cn-north-1",
		KeyDescription:        "It's test CMK2.",
		CreationDate:          "1522573239000",
		ScheduledDeletionDate: "",
		KeyState:              "3",
		DefaultKeyFlag:        "0",
		KeyType:               "1",
		ExpirationTime:        "",
		Origin:                "kms",
		KeyRotationEnabled:    "false",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/describe-key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var InstancesOutput = `{"instance_num":2}`

var InstancesResponse = cmk.InstancesResponse{
	InstanceNum: 2,
}

func HandleInstancesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/user-instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, InstancesOutput)
	})
}

var ListFirstResponse = cmk.ListResponse{
	Keys: []string{"30361023-62e0-4609-a5fc-6ff8eb63c186"},
	KeyDetails: []struct {
		KeyId                 string `json:"key_id,"`
		DomainId              string `json:"domain_id,"`
		KeyAlias              string `json:"key_alias,"`
		Realm                 string `json:"realm,"`
		KeyDescription        string `json:"key_description,"`
		CreationDate          string `json:"creation_date,"`
		ScheduledDeletionDate string `json:"scheduled_deletion_date,"`
		KeyState              string `json:"key_state,"`
		DefaultKeyFlag        string `json:"default_key_flag,"`
		KeyType               string `json:"key_type,"`
		ExpirationTime        string `json:"expiration_time,"`
		Origin                string `json:"origin,"`
		KeyRotationEnabled    string `json:"key_rotation_enabled,"`
	}{
		{
			KeyId:                 "30361023-62e0-4609-a5fc-6ff8eb63c186",
			DomainId:              "0984aafba48049a6b9457b89968eeabf",
			KeyAlias:              "TestCMK2",
			Realm:                 "cn-north-1",
			KeyDescription:        "It's test CMK2.",
			CreationDate:          "1522573239000",
			ScheduledDeletionDate: "",
			KeyState:              "3",
			DefaultKeyFlag:        "0",
			KeyType:               "1",
			ExpirationTime:        "",
			Origin:                "kms",
			KeyRotationEnabled:    "false",
		},
	},
	NextMarker: "1",
	Truncated:  "true",
}

var ListSecondResponse = cmk.ListResponse{
	Keys: []string{"e966a300-0c34-4a31-86e1-e67d13e6426a"},
	KeyDetails: []struct {
		KeyId                 string `json:"key_id,"`
		DomainId              string `json:"domain_id,"`
		KeyAlias              string `json:"key_alias,"`
		Realm                 string `json:"realm,"`
		KeyDescription        string `json:"key_description,"`
		CreationDate          string `json:"creation_date,"`
		ScheduledDeletionDate string `json:"scheduled_deletion_date,"`
		KeyState              string `json:"key_state,"`
		DefaultKeyFlag        string `json:"default_key_flag,"`
		KeyType               string `json:"key_type,"`
		ExpirationTime        string `json:"expiration_time,"`
		Origin                string `json:"origin,"`
		KeyRotationEnabled    string `json:"key_rotation_enabled,"`
	}{
		{
			KeyId:                 "e966a300-0c34-4a31-86e1-e67d13e6426a",
			DomainId:              "0984aafba48049a6b9457b89968eeabf",
			KeyAlias:              "TestCMK",
			Realm:                 "cn-north-1",
			KeyDescription:        "It's test CMK.",
			CreationDate:          "1522571717000",
			ScheduledDeletionDate: "",
			KeyState:              "2",
			DefaultKeyFlag:        "0",
			KeyType:               "1",
			ExpirationTime:        "",
			Origin:                "kms",
			KeyRotationEnabled:    "false",
		},
	},
	NextMarker: "1",
	Truncated:  "true",
}

var ListAllPagesResponse = cmk.ListResponse{
	Keys: []string{"30361023-62e0-4609-a5fc-6ff8eb63c186", "e966a300-0c34-4a31-86e1-e67d13e6426a"},
	KeyDetails: []struct {
		KeyId                 string `json:"key_id,"`
		DomainId              string `json:"domain_id,"`
		KeyAlias              string `json:"key_alias,"`
		Realm                 string `json:"realm,"`
		KeyDescription        string `json:"key_description,"`
		CreationDate          string `json:"creation_date,"`
		ScheduledDeletionDate string `json:"scheduled_deletion_date,"`
		KeyState              string `json:"key_state,"`
		DefaultKeyFlag        string `json:"default_key_flag,"`
		KeyType               string `json:"key_type,"`
		ExpirationTime        string `json:"expiration_time,"`
		Origin                string `json:"origin,"`
		KeyRotationEnabled    string `json:"key_rotation_enabled,"`
	}{
		{
			KeyId:                 "30361023-62e0-4609-a5fc-6ff8eb63c186",
			DomainId:              "0984aafba48049a6b9457b89968eeabf",
			KeyAlias:              "TestCMK2",
			Realm:                 "cn-north-1",
			KeyDescription:        "It's test CMK2.",
			CreationDate:          "1522573239000",
			ScheduledDeletionDate: "",
			KeyState:              "3",
			DefaultKeyFlag:        "0",
			KeyType:               "1",
			ExpirationTime:        "",
			Origin:                "kms",
			KeyRotationEnabled:    "false",
		},
		{
			KeyId:                 "e966a300-0c34-4a31-86e1-e67d13e6426a",
			DomainId:              "0984aafba48049a6b9457b89968eeabf",
			KeyAlias:              "TestCMK",
			Realm:                 "cn-north-1",
			KeyDescription:        "It's test CMK.",
			CreationDate:          "1522571717000",
			ScheduledDeletionDate: "",
			KeyState:              "2",
			DefaultKeyFlag:        "0",
			KeyType:               "1",
			ExpirationTime:        "",
			Origin:                "kms",
			KeyRotationEnabled:    "false",
		},
	},
	NextMarker: "",
	Truncated:  "",
}

var ListFirstOutput = `
{
  "keys": ["30361023-62e0-4609-a5fc-6ff8eb63c186"],
  "key_details": [{
    "key_id": "30361023-62e0-4609-a5fc-6ff8eb63c186",
    "domain_id": "0984aafba48049a6b9457b89968eeabf",
    "key_alias": "TestCMK2",
    "realm": "cn-north-1",
    "key_description": "It's test CMK2.",
    "creation_date": "1522573239000",
    "scheduled_deletion_date": "",
    "expiration_time": "",
    "origin": "kms",
    "key_state": "3",
    "default_key_flag": "0",
    "key_type": "1",
    "key_rotation_enabled": "false"
  }],
  "next_marker": "1",
  "truncated": "true"
}
`

var ListSecondOutput = `
{
  "keys": ["e966a300-0c34-4a31-86e1-e67d13e6426a"],
  "key_details": [{
    "key_id": "e966a300-0c34-4a31-86e1-e67d13e6426a",
    "domain_id": "0984aafba48049a6b9457b89968eeabf",
    "key_alias": "TestCMK",
    "realm": "cn-north-1",
    "key_description": "It's test CMK.",
    "creation_date": "1522571717000",
    "scheduled_deletion_date": "",
    "expiration_time": "",
    "origin": "kms",
    "key_state": "2",
    "default_key_flag": "0",
    "key_type": "1",
    "key_rotation_enabled": "false"
  }],
  "next_marker": "",
  "truncated": "false"
}
`

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/list-keys",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			var listOpts cmk.ListOpts
			if reader, ok := r.Body.(io.Reader); ok {
				if readCloser, ok := reader.(io.Closer); ok {
					defer readCloser.Close()
				}
				json.NewDecoder(reader).Decode(&listOpts)
			}
			marker := listOpts.Marker
			switch marker {
			case "1":
				fmt.Fprintf(w, ListSecondOutput)
			case "":
				fmt.Fprintf(w, ListFirstOutput)
			}
		})
}

var ScheduleDeletionResponse = cmk.ScheduleDeletionResponse{
	KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
	KeyState: "4",
}

var ScheduleDeletionOutput = `
{"key_id":"30361023-62e0-4609-a5fc-6ff8eb63c186","key_state":"4"}
`

func HandleScheduleDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/schedule-key-deletion", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ScheduleDeletionOutput)
	})
}

var CancelDeletionResponse = cmk.CancelDeletionResponse{
	KeyId:    "30361023-62e0-4609-a5fc-6ff8eb63c186",
	KeyState: "3",
}

var CancelDeletionOutput = `
{"key_id":"30361023-62e0-4609-a5fc-6ff8eb63c186","key_state":"3"}
`

func HandleCancelDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/cancel-key-deletion", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CancelDeletionOutput)
	})
}

var UpdateAliasResponse = cmk.UpdateAliasResponse{
	KeyInfo: struct {
		KeyId    string `json:"key_id,"`
		KeyAlias string `json:"key_alias,"`
	}{
		KeyId:    "e966a300-0c34-4a31-86e1-e67d13e6426a",
		KeyAlias: "TestABC",
	},
}

var UpdateAliasOutput = `
{"key_info":{"key_id":"e966a300-0c34-4a31-86e1-e67d13e6426a","key_alias":"TestABC"}}
`

func HandleUpdateAliasSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/update-key-alias", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateAliasOutput)
	})
}

var UpdateDescriptionResponse = cmk.UpdateDescriptionResponse{
	KeyInfo: struct {
		KeyId          string `json:"key_id,"`
		KeyDescription string `json:"key_description,"`
	}{
		KeyId:          "e966a300-0c34-4a31-86e1-e67d13e6426a",
		KeyDescription: "TestABC",
	},
}

var UpdateDescriptionOutput = `
{"key_info":{"key_id":"e966a300-0c34-4a31-86e1-e67d13e6426a","key_description":"TestABC"}}
`

func HandleUpdateDescriptionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/update-key-description", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateDescriptionOutput)
	})
}

var QuotasResponse = cmk.QuotasResponse{
	// Quota list
	Quotas: struct {
		// Resource quota list
		Resources []struct {
			// Quota type
			Type string `json:"type,"`

			// Used quotas
			Used int `json:"used,"`

			// Total quotas
			Quota int `json:"quota,"`
		} `json:"resources,"`
	}{
		Resources: []struct {
			// Quota type
			Type string `json:"type,"`

			// Used quotas
			Used int `json:"used,"`

			// Total quotas
			Quota int `json:"quota,"`
		}{
			{
				Type:  "CMK",
				Used:  2,
				Quota: 5,
			},
			{
				Type:  "grant_per_CMK",
				Used:  0,
				Quota: 100,
			},
		},
	},
}

var QuotasOutput = `
{"quotas":{"resources":[{"type":"CMK","used":2,"quota":5},{"type":"grant_per_CMK","used":0,"quota":100}]}}
`

func HandleQuotasSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/kms/user-quotas", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, QuotasOutput)
	})
}
