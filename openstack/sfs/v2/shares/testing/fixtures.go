package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	shareEndpoint = "/shares"
	shareID       = "011d21e2-fbc3-4e4a-9993-9ea223f73264"
)

var createRequest = `{
		"share": {
			"name": "my_test_share",
			"size": 1,
			"share_proto": "NFS"
		}
	}`

var createResponse = `{
		"share": {
			"name": "my_test_share",
			"share_proto": "NFS",
			"size": 1,
			"status": null,
			"share_server_id": null,
			"project_id": "16e1ab15c35a457e9c2b2aa189f544e1",
			"share_type": "25747776-08e5-494f-ab40-a64b9d20d8f7",
			"share_type_name": "default",
			"availability_zone": null,
			"created_at": "2015-09-18T10:25:24.533287",
			"export_location": null,
			"links": [
				{
					"href": "http://172.18.198.54:8786/v1/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
					"rel": "self"
				},
				{
					"href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
					"rel": "bookmark"
				}
			],
			"share_network_id": null,
			"export_locations": [],
			"host": null,
			"access_rules_status": "active",
			"task_state": null,
			"snapshot_support": true,
			"consistency_group_id": "9397c191-8427-4661-a2e8-b23820dc01d4",
			"source_cgsnapshot_member_id": null,
			"volume_type": "default",
			"snapshot_id": null,
			"is_public": true,
			"metadata": {
				"project": "my_app",
				"aim": "doc"
			},
			"id": "011d21e2-fbc3-4e4a-9993-9ea223f73264",
			"description": "My custom share London"
		}
	}`

// MockCreateResponse creates a mock response
func MockCreateResponse(t *testing.T) {
	th.Mux.HandleFunc(shareEndpoint, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, createRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, createResponse)
	})
}

var updateRequest = `{
    "share": {
        "display_name": "my_test_share_sfs",
        "display_description": "test"
    }
}`

var updateResponse = `{
		"share": {
			"name": "my_test_share_sfs",
			"share_proto": "NFS",
			"size": 1,
			"status": null,
			"share_server_id": null,
			"project_id": "16e1ab15c35a457e9c2b2aa189f544e1",
			"share_type": "25747776-08e5-494f-ab40-a64b9d20d8f7",
			"share_type_name": "default",
			"availability_zone": null,
			"created_at": "2015-09-18T10:25:24.533287",
			"export_location": null,
			"links": [
				{
					"href": "http://172.18.198.54:8786/v1/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
					"rel": "self"
				},
				{
					"href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
					"rel": "bookmark"
				}
			],
			"share_network_id": null,
			"export_locations": [],
			"host": null,
			"access_rules_status": "active",
			"task_state": null,
			"snapshot_support": true,
			"consistency_group_id": "9397c191-8427-4661-a2e8-b23820dc01d4",
			"source_cgsnapshot_member_id": null,
			"volume_type": "default",
			"snapshot_id": null,
			"is_public": true,
			"metadata": {
				"project": "my_app",
				"aim": "doc"
			},
			"id": "011d21e2-fbc3-4e4a-9993-9ea223f73264",
			"description": "test"
		}
	}`

// MockCreateResponse creates a mock response
func MockUpdateResponse(t *testing.T) {
	th.Mux.HandleFunc(shareEndpoint+"/"+shareID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, updateResponse)
	})
}

// MockDeleteResponse creates a mock delete response
func MockDeleteResponse(t *testing.T) {
	th.Mux.HandleFunc(shareEndpoint+"/"+shareID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}

var getResponse = `{
    "share": {
        "links": [
            {
                "href": "http://172.18.198.54:8786/v2/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
                "rel": "self"
            },
            {
                "href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
                "rel": "bookmark"
            }
        ],
        "availability_zone": "nova",
        "share_network_id": "713df749-aac0-4a54-af52-10f6c991e80c",
        "snapshot_id": null,
        "id": "011d21e2-fbc3-4e4a-9993-9ea223f73264",
        "size": 1,
        "share_type": "25747776-08e5-494f-ab40-a64b9d20d8f7",
        "project_id": "16e1ab15c35a457e9c2b2aa189f544e1",
        "metadata": {
            "project": "my_app",
            "aim": "doc"
        },
        "status": "available",
        "description": "My custom share London",
        "host": "manila2@generic1#GENERIC1",
        "is_public": true,
        "name": "my_test_share",
        "created_at": "2015-09-18T10:25:24.000000",
        "share_proto": "NFS",
        "volume_type": "default"
    }
}`

// MockGetResponse creates a mock get response
func MockGetResponse(t *testing.T) {
	th.Mux.HandleFunc(shareEndpoint+"/"+shareID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getResponse)
	})
}

var grantAccessRequest = `{
    "os-allow_access": {
        "access_to": "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
        "access_type": "cert",
        "access_level": "rw"
    }
}`

var grantAccessResponse = `{
    "access": {
        "share_id": "1b8facf8-b822-4349-a033-e078b2a84b7f",
        "access_type": "cert",
        "access_to": "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
        "access_level": "rw",
        "state": "new",
        "id": "fc32500f-fa78-4f06-8caf-06ad7fb9726c"
    }
}`

var listAccessRightsRequest = `{
    "os-access_list": null
}`

var listAccessRightsResponse = `{
    "access_list": [
        {
            "access_level": "rw",
            "state": "active",
            "id": "5158f095-4c43-49c0-b5a7-c458e85ed8c8",
            "access_type": "cert",
            "access_to": "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8"
        }
    ]
}`

var deleteAccessRequest = `{
    "os-deny_access": {
        "access_id": "ea07152b-d08b-4f6b-8785-ce64dce52679"
    }
}`

var getExportLocationsResponse = `{
    "export_locations": [
        {
		"path": "sfs-nas1.eu-de.otc.t-systems.com:/share-d41ee18b",
		"id": "fab962ba-4b9a-475e-a380-8e856ed3f92d"		
	}
    ]
}`

var listResponse = `{
    "shares": [
		{
        "links": [
            {
                "href": "http://172.18.198.54:8786/v2/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
                "rel": "self"
            },
            {
                "href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/011d21e2-fbc3-4e4a-9993-9ea223f73264",
                "rel": "bookmark"
            }
        ],
        "availability_zone": "nova",
        "share_network_id": "713df749-aac0-4a54-af52-10f6c991e80c",
        "id": "011d21e2-fbc3-4e4a-9993-9ea223f73264",
        "size": 1,
        "share_type": "25747776-08e5-494f-ab40-a64b9d20d8f7",
        "project_id": "16e1ab15c35a457e9c2b2aa189f544e1",
        "status": "available",
        "description": "My custom share London",
        "host": "manila2@generic1#GENERIC1",
        "is_public": true,
		"created_at": "2015-09-18T10:25:24.000000",
        "name": "my_test_share",
        "share_proto": "NFS",
        "volume_type": "default"
	}
    ]
}`
