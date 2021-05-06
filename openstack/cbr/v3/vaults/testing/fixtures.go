package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/cbr/v3/vaults"
	"github.com/huaweicloud/golangsdk/openstack/common/tags"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedCreateRequest = `
{
  "vault" : {
    "backup_policy_id" : "6dd81d7d-a4cb-443e-b8ed-1af0bd3a261b",
    "billing" : {
      "cloud_type" : "public",
      "consistent_level" : "crash_consistent",
      "object_type" : "server",
      "protect_type" : "backup",
      "size" : 100,
      "charging_mode" : "post_paid",
      "console_url" : "https://console.demo.com/cbr/?agencyId=97fcd896b7914cb98f553a087232e243&region=testregion/cbr/manager/csbs/vaultList"
    },
    "description" : "vault_description",
    "name" : "vault_name",
    "resources" : [ {
  	  "extra_info" : {
  	    "include_volumes" : [ {
  	      "id" : "73ee8446-bce7-4371-9650-b440b5f4c1d0",
  	      "os_version" : "CentOS 7.6 64bit"
  	    } ]
  	  },
  	  "id" : "23a320a5-3efd-4568-b1aa-8dd9183cc64c",
  	  "type" : "OS::Nova::Server"
    } ],
    "tags" : [ {
  	  "key" : "key01",
  	  "value" : "value01"
    } ],
    "enterprise_project_id" : "0"
  }
}`
	expectedCreateResponse = `
{
  "vault" : {
    "provider_id" : "0daac4c5-6707-4851-97ba-169e36266b66",
    "description" : "vault_description",
    "tags" : [ {
  	  "value" : "value01",
  	  "key" : "key01"
    } ],
    "enterprise_project_id" : "0",
    "auto_bind" : false,
    "id" : "ad7627ae-5b0b-492e-b6bd-cd809b745197",
    "user_id" : "38d65be2ecd840d19046e239e841a734",
    "name" : "vault_name",
    "billing" : {
  	  "status" : "available",
  	  "used" : 0,
  	  "protect_type" : "backup",
  	  "object_type" : "server",
  	  "allocated" : 40,
  	  "spec_code" : "vault.backup.server.normal",
  	  "size" : 100,
  	  "cloud_type" : "public",
  	  "consistent_level" : "crash_consistent",
  	  "charging_mode" : "post_paid"
    },
    "created_at" : "2019-05-23T12:51:10.071232",
    "project_id" : "fc347bc64ccd4589ae52e4f44b7433c7",
    "resources" : [ {
      "name" : "ecs-b977-0002",
      "backup_size" : 0,
      "protect_status" : "available",
      "backup_count" : 0,
      "extra_info" : {
        "include_volumes" : [ {
          "os_version" : "CentOS 7.6 64bit",
          "id" : "73ee8446-bce7-4371-9650-b440b5f4c1d0"
        } ]
      },
      "type" : "OS::Nova::Server",
      "id" : "23a320a5-3efd-4568-b1aa-8dd9183cc64c",
      "size" : 40
    } ]
  }
}`
	expectedUpdateRequest = `
{
  "vault" : {
    "billing" : {
      "size" : 100
    },
    "name" : "vault_name"
    }
  }
}`

	expectedListResponse = `
{
  "vaults" : [ {
    "provider_id" : "0daac4c5-6707-4851-97ba-169e36266b66",
    "description" : "vault_description",
    "tags" : [ {
  	  "value" : "value01",
  	  "key" : "key01"
    } ],
    "enterprise_project_id" : "0",
    "auto_bind" : false,
    "id" : "ad7627ae-5b0b-492e-b6bd-cd809b745197",
    "user_id" : "38d65be2ecd840d19046e239e841a734",
    "name" : "vault_name",
    "billing" : {
  	  "status" : "available",
  	  "used" : 0,
  	  "protect_type" : "backup",
  	  "object_type" : "server",
  	  "allocated" : 40,
  	  "spec_code" : "vault.backup.server.normal",
  	  "size" : 100,
  	  "cloud_type" : "public",
  	  "consistent_level" : "crash_consistent",
  	  "charging_mode" : "post_paid"
    },
    "created_at" : "2019-05-23T12:51:10.071232",
    "project_id" : "fc347bc64ccd4589ae52e4f44b7433c7",
    "resources" : [ {
      "name" : "ecs-b977-0002",
      "backup_size" : 0,
      "protect_status" : "available",
      "backup_count" : 0,
      "extra_info" : {
        "include_volumes" : [ {
          "os_version" : "CentOS 7.6 64bit",
          "id" : "73ee8446-bce7-4371-9650-b440b5f4c1d0"
        } ]
      },
      "type" : "OS::Nova::Server",
      "id" : "23a320a5-3efd-4568-b1aa-8dd9183cc64c",
      "size" : 40
    } ]
  } ],
  "count" : 1
}`

	expectedPolicyBindingResponse = `
{
  "associate_policy" : {
    "vault_id" : "ad7627ae-5b0b-492e-b6bd-cd809b745197",
    "policy_id" : "7075c397-25a0-43e2-a83a-bb16eaca3ee5"
  }
}`

	expectedPolicyUnbindingResponse = `
{
  "dissociate_policy" : {
    "vault_id" : "ad7627ae-5b0b-492e-b6bd-cd809b745197",
    "policy_id" : "7075c397-25a0-43e2-a83a-bb16eaca3ee5"
  }
}`

	expectedAssociateResourcesRequest = `
{
  "resources" : [ {
    "extra_info" : {
      "exclude_volumes" : [ "bdef09bb-293f-446a-88a4-86e9f14408c4" ]
    },
    "id" : "97595625-198e-4e4d-879b-9d53f68ba551",
    "type" : "OS::Nova::Server"
  } ]
}`

	expectedAssociateResourcesResponse = `
{
  "add_resource_ids" : [ "97595625-198e-4e4d-879b-9d53f68ba551" ]
}`

	expectedDissociateResourcesRequest = `
{
	"resource_ids" : [ "97595625-198e-4e4d-879b-9d53f68ba551" ]
}`

	expectedDissociateResourcesResponse = `
{
  "remove_resource_ids" : [ "fe578a6c-d1a8-4790-bd52-5954af4d446c" ]
}`

	expectedMigrateResourcesResponse = `
{
  'resource_ids': [ 'abcdde3f-e0e3-403a-b690-fc259dd70008' ],
  'destination_vault_id': 'fe578a6c-d1a8-4790-bd52-5954af4d446c'
}`
)

var (
	createOpts = &vaults.CreateOpts{
		BackupPolicyID: "6dd81d7d-a4cb-443e-b8ed-1af0bd3a261b",
		Billing: &vaults.BillingCreate{
			ConsistentLevel: "crash_consistent",
			CloudType:       "public",
			ObjectType:      "server",
			ProtectType:     "backup",
			Size:            100,
			ChargingMode:    "post_paid",
			ConsoleURL:      "https://console.demo.com/cbr/?agencyId=97fcd896b7914cb98f553a087232e243&region=testregion/cbr/manager/csbs/vaultList",
		},
		Description:         "vault_description",
		EnterpriseProjectID: "0",
		Name:                "vault_name",
		Resources: []vaults.ResourceCreate{
			{
				ID:   "23a320a5-3efd-4568-b1aa-8dd9183cc64c",
				Type: "OS::Nova::Server",
				ExtraInfo: &vaults.ResourceExtraInfo{
					IncludeVolumes: []vaults.ResourceExtraInfoIncludeVolumes{
						{
							ID:        "73ee8446-bce7-4371-9650-b440b5f4c1d0",
							OSVersion: "CentOS 7.6 64bit",
						},
					},
				},
			},
		},
		Tags: []tags.ResourceTag{
			{
				Key:   "key01",
				Value: "value01",
			},
		},
	}

	expectedCreateResponseData = &vaults.Vault{
		ID:                  "ad7627ae-5b0b-492e-b6bd-cd809b745197",
		Name:                "vault_name",
		UserID:              "38d65be2ecd840d19046e239e841a734",
		AutoBind:            false,
		ProviderID:          "0daac4c5-6707-4851-97ba-169e36266b66",
		ProjectID:           "fc347bc64ccd4589ae52e4f44b7433c7",
		Description:         "vault_description",
		EnterpriseProjectID: "0",
		CreatedAt:           "2019-05-23T12:51:10.071232",
		Billing: vaults.Billing{
			Status:          "available",
			ProtectType:     "backup",
			ObjectType:      "server",
			SpecCode:        "vault.backup.server.normal",
			CloudType:       "public",
			ConsistentLevel: "crash_consistent",
			ChargingMode:    "post_paid",
			Used:            0,
			Allocated:       40,
			Size:            100,
		},
		Tags: []tags.ResourceTag{
			{
				Key:   "key01",
				Value: "value01",
			},
		},
		Resources: []vaults.ResourceResp{
			{
				ID:            "23a320a5-3efd-4568-b1aa-8dd9183cc64c",
				Type:          "OS::Nova::Server",
				Name:          "ecs-b977-0002",
				ProtectStatus: "available",
				BackupSize:    0,
				BackupCount:   0,
				Size:          40,
				ExtraInfo: vaults.ResourceExtraInfo{
					IncludeVolumes: []vaults.ResourceExtraInfoIncludeVolumes{
						{
							ID:        "73ee8446-bce7-4371-9650-b440b5f4c1d0",
							OSVersion: "CentOS 7.6 64bit",
						},
					},
				},
			},
		},
	}

	updateOpts = &vaults.UpdateOpts{
		Billing: &vaults.BillingUpdate{
			Size: 100,
		},
		Name: "vault_name",
	}

	expectedListResponseData = &[]vaults.Vault{
		{
			ID:                  "ad7627ae-5b0b-492e-b6bd-cd809b745197",
			Name:                "vault_name",
			UserID:              "38d65be2ecd840d19046e239e841a734",
			AutoBind:            false,
			ProviderID:          "0daac4c5-6707-4851-97ba-169e36266b66",
			ProjectID:           "fc347bc64ccd4589ae52e4f44b7433c7",
			Description:         "vault_description",
			EnterpriseProjectID: "0",
			CreatedAt:           "2019-05-23T12:51:10.071232",
			Billing: vaults.Billing{
				Status:          "available",
				ProtectType:     "backup",
				ObjectType:      "server",
				SpecCode:        "vault.backup.server.normal",
				CloudType:       "public",
				ConsistentLevel: "crash_consistent",
				ChargingMode:    "post_paid",
				Used:            0,
				Allocated:       40,
				Size:            100,
			},
			Tags: []tags.ResourceTag{
				{
					Key:   "key01",
					Value: "value01",
				},
			},
			Resources: []vaults.ResourceResp{
				{
					ID:            "23a320a5-3efd-4568-b1aa-8dd9183cc64c",
					Type:          "OS::Nova::Server",
					Name:          "ecs-b977-0002",
					ProtectStatus: "available",
					BackupSize:    0,
					BackupCount:   0,
					Size:          40,
					ExtraInfo: vaults.ResourceExtraInfo{
						IncludeVolumes: []vaults.ResourceExtraInfoIncludeVolumes{
							{
								ID:        "73ee8446-bce7-4371-9650-b440b5f4c1d0",
								OSVersion: "CentOS 7.6 64bit",
							},
						},
					},
				},
			},
		},
	}

	bindPolicyOpts = &vaults.BindPolicyOpts{
		PolicyID: "7075c397-25a0-43e2-a83a-bb16eaca3ee5",
	}

	expectedPolicyBindingResponseData = &vaults.PolicyBinding{
		VaultID:  "ad7627ae-5b0b-492e-b6bd-cd809b745197",
		PolicyID: "7075c397-25a0-43e2-a83a-bb16eaca3ee5",
	}

	associateResourcesOpts = &vaults.AssociateResourcesOpts{
		Resources: []vaults.ResourceCreate{
			{
				ExtraInfo: &vaults.ResourceExtraInfo{
					ExcludeVolumes: []string{
						"bdef09bb-293f-446a-88a4-86e9f14408c4",
					},
				},
				ID:   "97595625-198e-4e4d-879b-9d53f68ba551",
				Type: "OS::Nova::Server",
			},
		},
	}

	expectedAssociateResourcesResponseData = []string{
		"97595625-198e-4e4d-879b-9d53f68ba551",
	}

	dissociateResourcesOpts = &vaults.DissociateResourcesOpts{
		ResourceIDs: []string{
			"97595625-198e-4e4d-879b-9d53f68ba551",
		},
	}
)

func handleVaultCreate(t *testing.T) {
	th.Mux.HandleFunc("/vaults", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleVaultGet(t *testing.T) {
	th.Mux.HandleFunc("/vaults/ad7627ae-5b0b-492e-b6bd-cd809b745197", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleVaultUpdate(t *testing.T) {
	th.Mux.HandleFunc("/vaults/ad7627ae-5b0b-492e-b6bd-cd809b745197", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleVaultDelete(t *testing.T) {
	th.Mux.HandleFunc("/vaults/ad7627ae-5b0b-492e-b6bd-cd809b745197", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}

func handleVaultList(t *testing.T) {
	th.Mux.HandleFunc("/vaults", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}

func handleVaultBindPolicy(t *testing.T) {
	th.Mux.HandleFunc("/vaults/ad7627ae-5b0b-492e-b6bd-cd809b745197/associatepolicy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedPolicyBindingResponse)
		})
}

func handleVaultUnbindPolicy(t *testing.T) {
	th.Mux.HandleFunc("/vaults/ad7627ae-5b0b-492e-b6bd-cd809b745197/dissociatepolicy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedPolicyUnbindingResponse)
		})
}

func handleVaultAssociateResources(t *testing.T) {
	th.Mux.HandleFunc("/vaults/ad7627ae-5b0b-492e-b6bd-cd809b745197/addresources",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedAssociateResourcesResponse)
		})
}

func handleVaultDissociateResources(t *testing.T) {
	th.Mux.HandleFunc("/vaults/ad7627ae-5b0b-492e-b6bd-cd809b745197/removeresources",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedAssociateResourcesResponse)
		})
}
