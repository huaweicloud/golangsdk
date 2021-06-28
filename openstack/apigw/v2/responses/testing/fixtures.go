package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/responses"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"create_time": "2021-06-24T09:33:10.562277766+08:00",
	"default": false,
	"id": "baabc69fdb8f4c458637666c0441e9a4",
	"name": "terraform-test",
	"responses": {
		"ACCESS_DENIED": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": false,
			"status": 402
		},
		"AUTHORIZER_CONF_FAILURE": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 500
		},
		"AUTHORIZER_FAILURE": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 500
		},
		"AUTHORIZER_IDENTITIES_FAILURE": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 401
		},
		"AUTH_FAILURE": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 401
		},
		"AUTH_HEADER_MISSING": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 401
		},
		"BACKEND_TIMEOUT": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 504
		},
		"BACKEND_UNAVAILABLE": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 502
		},
		"DEFAULT_4XX": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true
		},
		"DEFAULT_5XX": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true
		},
		"NOT_FOUND": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 404
		},
		"REQUEST_PARAMETERS_FAILURE": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 400
		},
		"THROTTLED": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 429
		},
		"UNAUTHORIZED": {
			"body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
			"default": true,
			"status": 401
		}
	},
	"update_time": "2021-06-24T09:33:10.562277766+08:00"
}`
	expectedListResponse = `
{
	"responses": [
        {
            "id": "baabc69fdb8f4c458637666c0441e9a4",
            "name": "terraform-test",
            "default": false,
            "create_time": "2021-06-24T01:33:10Z",
            "update_time": "2021-06-24T01:33:10Z"
        },
        {
            "id": "5623b9b3c2154f6ab1a7c0cf5c7c6278",
            "name": "default",
            "default": true,
            "create_time": "2021-06-23T08:29:23Z",
            "update_time": "2021-06-23T08:29:23Z"
        }
    ]
}`
	expectedGetSpecResponse = `
{
	"ACCESS_DENIED": {
        "body": "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
        "default": false,
        "status": 405
    }
}`
)

var (
	createOpts = &responses.ResponseOpts{
		Name: "terraform-test",
		Responses: map[string]responses.ResponseInfo{
			"ACCESS_DENIED": {
				Body:   "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				Status: 402,
			},
		},
		InstanceId: "9b76174b785342078e557f23c01d5e41",
		GroupId:    "d060ade0560a4c01b89bf954ad2e9d6e",
	}

	expectedCreateResponseData = &responses.Response{
		CreateTime: "2021-06-24T09:33:10.562277766+08:00",
		UpdateTime: "2021-06-24T09:33:10.562277766+08:00",
		IsDefault:  false,
		Id:         "baabc69fdb8f4c458637666c0441e9a4",
		Name:       "terraform-test",
		Responses: map[string]responses.ResponseInfo{
			"ACCESS_DENIED": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: false,
				Status:    402,
			},
			"AUTHORIZER_CONF_FAILURE": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    500,
			},
			"AUTHORIZER_FAILURE": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    500,
			},
			"AUTHORIZER_IDENTITIES_FAILURE": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    401,
			},
			"AUTH_FAILURE": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    401,
			},
			"AUTH_HEADER_MISSING": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    401,
			},
			"BACKEND_TIMEOUT": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    504,
			},
			"BACKEND_UNAVAILABLE": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    502,
			},
			"DEFAULT_4XX": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
			},
			"DEFAULT_5XX": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
			},
			"NOT_FOUND": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    404,
			},
			"REQUEST_PARAMETERS_FAILURE": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    400,
			},
			"THROTTLED": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    429,
			},
			"UNAUTHORIZED": {
				Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
				IsDefault: true,
				Status:    401,
			},
		},
	}

	listOpts = &responses.ListOpts{
		InstanceId: "9b76174b785342078e557f23c01d5e41",
		GroupId:    "d060ade0560a4c01b89bf954ad2e9d6e",
	}

	expectedListResponseData = []responses.Response{
		{
			Id:         "baabc69fdb8f4c458637666c0441e9a4",
			Name:       "terraform-test",
			IsDefault:  false,
			CreateTime: "2021-06-24T01:33:10Z",
			UpdateTime: "2021-06-24T01:33:10Z",
		},
		{
			Id:         "5623b9b3c2154f6ab1a7c0cf5c7c6278",
			Name:       "default",
			IsDefault:  true,
			CreateTime: "2021-06-23T08:29:23Z",
			UpdateTime: "2021-06-23T08:29:23Z",
		},
	}

	updateOpts = &responses.ResponseOpts{
		Name:       "terraform-test",
		InstanceId: "9b76174b785342078e557f23c01d5e41",
		GroupId:    "d060ade0560a4c01b89bf954ad2e9d6e",
	}
	specRespOpts = &responses.SpecRespOpts{
		InstanceId: "9b76174b785342078e557f23c01d5e41",
		GroupId:    "d060ade0560a4c01b89bf954ad2e9d6e",
		RespId:     "baabc69fdb8f4c458637666c0441e9a4",
	}

	responseInfoOpts = &responses.ResponseInfo{
		Body:   "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
		Status: 405,
	}

	expectedGetSpecResponseData = &responses.ResponseInfo{
		Body:      "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
		Status:    405,
		IsDefault: false,
	}

	updateSpecRespOpts = responses.ResponseInfo{
		Body:   "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\",\"request_id\":\"$context.requestId\"}",
		Status: 405,
	}
)

func handleV2ResponsesCreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/9b76174b785342078e557f23c01d5e41/api-groups/d060ade0560a4c01b89bf954ad2e9d6e"+
		"/gateway-responses", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV2ResponsesGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/9b76174b785342078e557f23c01d5e41/api-groups/d060ade0560a4c01b89bf954ad2e9d6e"+
		"/gateway-responses/baabc69fdb8f4c458637666c0441e9a4", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV2ResponsesList(t *testing.T) {
	th.Mux.HandleFunc("/instances/9b76174b785342078e557f23c01d5e41/api-groups/d060ade0560a4c01b89bf954ad2e9d6e"+
		"/gateway-responses", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}

func handleV2ResponsesUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/9b76174b785342078e557f23c01d5e41/api-groups/d060ade0560a4c01b89bf954ad2e9d6e"+
		"/gateway-responses/baabc69fdb8f4c458637666c0441e9a4", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV2ResponsesDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/9b76174b785342078e557f23c01d5e41/api-groups/d060ade0560a4c01b89bf954ad2e9d6e"+
		"/gateway-responses/baabc69fdb8f4c458637666c0441e9a4", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}

func handleV2SpecResponseGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/9b76174b785342078e557f23c01d5e41/api-groups/d060ade0560a4c01b89bf954ad2e9d6e"+
		"/gateway-responses/baabc69fdb8f4c458637666c0441e9a4/ACCESS_DENIED", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedGetSpecResponse)
	})
}

func handleV2SpecResponseUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/9b76174b785342078e557f23c01d5e41/api-groups/d060ade0560a4c01b89bf954ad2e9d6e"+
		"/gateway-responses/baabc69fdb8f4c458637666c0441e9a4/ACCESS_DENIED", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedGetSpecResponse)
	})
}

func handleV2SpecResponseDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/9b76174b785342078e557f23c01d5e41/api-groups/d060ade0560a4c01b89bf954ad2e9d6e"+
		"/gateway-responses/baabc69fdb8f4c458637666c0441e9a4/ACCESS_DENIED", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}
