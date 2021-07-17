package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/authorizers"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedGetResponse = `
{
	"authorizer_type": "FUNC",
	"authorizer_uri": "urn:fss:ae-ad-1:0c22dd73f6005a032f3ec0061de74dbf:function:default:terraform_test",
	"create_time": "2021-07-17T08:33:35Z",
	"id": "3b9a3dda163f4ebb8fdcbbf786bffa20",
	"name": "terraform_test",
	"need_body": false,
	"ttl": 60,
	"type": "BACKEND",
	"user_data": ""
}
`

	expectedListResponse = `
{
	"authorizer_list": [
		{
	    	"authorizer_type": "FUNC",
	    	"authorizer_uri": "urn:fss:ae-ad-1:0c22dd73f6005a032f3ec0061de74dbf:function:default:terraform_test",
	    	"create_time": "2021-07-17T08:33:35Z",
	    	"id": "3b9a3dda163f4ebb8fdcbbf786bffa20",
	    	"name": "terraform_test",
	    	"need_body": false,
	    	"ttl": 60,
	    	"type": "BACKEND",
	    	"user_data": ""
	  	}
	],
	"size": 1,
	"total": 1
}`
)

var (
	createOpts = authorizers.CustomAuthOpts{
		Name:           "terraform_test",
		Type:           "BACKEND",
		AuthorizerType: "FUNC",
		AuthorizerUri:  "urn:fss:ae-ad-1\\:0c22dd73f6005a032f3ec0061de74dbf:function:default:terraform_test",
		IsBodySend:     golangsdk.Disabled,
		CacheAge:       60,
		UserData:       golangsdk.MaybeString(""),
	}

	expectedGetResponseData = &authorizers.CustomAuthorizer{
		Name:           "terraform_test",
		Type:           "BACKEND",
		AuthorizerType: "FUNC",
		AuthorizerUri:  "urn:fss:ae-ad-1:0c22dd73f6005a032f3ec0061de74dbf:function:default:terraform_test",
		CreateTime:     "2021-07-17T08:33:35Z",
		Id:             "3b9a3dda163f4ebb8fdcbbf786bffa20",
		IsBodySend:     false,
		CacheAge:       60,
		UserData:       "",
	}

	listOpts = &authorizers.ListOpts{
		Name: "terraform_test",
	}

	expectedListResponseData = []authorizers.CustomAuthorizer{
		{
			Name:           "terraform_test",
			Type:           "BACKEND",
			AuthorizerType: "FUNC",
			AuthorizerUri:  "urn:fss:ae-ad-1:0c22dd73f6005a032f3ec0061de74dbf:function:default:terraform_test",
			CreateTime:     "2021-07-17T08:33:35Z",
			Id:             "3b9a3dda163f4ebb8fdcbbf786bffa20",
			IsBodySend:     false,
			CacheAge:       60,
			UserData:       "",
		},
	}
)

func handleV2CustomAuthorizerCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/authorizers",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2CustomAuthorizerGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/authorizers/0d2a523974a14fe1a25c1bc2f61b2d9d",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2CustomAuthorizerList(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/authorizers",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListResponse)
		})
}

func handleV2CustomAuthorizerUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/authorizers/0d2a523974a14fe1a25c1bc2f61b2d9d",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2CustomAuthorizerDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/6da953fe33d44650a067e43a4593368b/authorizers/0d2a523974a14fe1a25c1bc2f61b2d9d",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}
