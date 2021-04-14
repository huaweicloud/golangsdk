package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/cbr/v3/policies"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedRequest = `
{
  "policy" : {
    "name" : "policy001",
    "operation_definition" : {
      "retention_duration_days" : 1,
      "timezone" : "UTC+08:00"
    },
    "operation_type" : "backup",
    "trigger" : {
      "properties" : {
        "pattern" : [
			"FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00"
		]
      }
    }
  }
}`
	expectedCreateResponse = `
{
  "policy" : {
    "name" : "policy001",
    "enabled" : true,
    "trigger" : {
      "properties" : {
        "pattern" : [ "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00" ],
        "start_time" : "2019-05-08 06:57:05"
      },
      "type" : "time",
      "id" : "d67269a6-5369-42d7-8150-5254bd446328",
      "name" : "default"
    },
    "operation_definition" : {
  	  "retention_duration_days" : 1,
  	  "year_backups" : 0,
  	  "day_backups" : 0,
  	  "month_backups" : 0,
  	  "week_backups" : 0,
  	  "timezone" : "UTC+08:00"
    },
    "operation_type" : "backup",
    "id" : "cbb3ce6f-3332-4e7c-b98e-77290d8471ff"
  }
}`
	expectedUpdateResponse = `
{
  "policy" : {
    "name" : "policy001",
    "enabled" : true,
    "trigger" : {
      "properties" : {
        "pattern" : [ "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00" ],
        "start_time" : "2019-05-08 06:57:05"
      },
      "type" : "time",
      "id" : "d67269a6-5369-42d7-8150-5254bd446328",
      "name" : "default"
    },
    "operation_definition" : {
      "max_backups" : 0,
      "year_backups" : 0,
      "day_backups" : 0,
      "month_backups" : 0,
      "week_backups" : 0,
      "timezone" : "UTC+08:00",
      "retention_duration_days": 1
    },
    "operation_type" : "backup",
    "id" : "cbb3ce6f-3332-4e7c-b98e-77290d8471ff"
  }
}`
	expectedGetResponse = `
{
  "policy" : {
    "name" : "policy001",
    "associated_vaults" : [ ],
    "enabled" : true,
    "trigger" : {
      "properties" : {
        "pattern" : [ "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00" ],
        "start_time" : "2019-05-08 06:57:05"
      },
      "type" : "time",
      "id" : "d67269a6-5369-42d7-8150-5254bd446328",
      "name" : "default"
    },
    "operation_definition" : {
      "retention_duration_days" : 1,
      "year_backups" : 0,
      "day_backups" : 0,
      "month_backups" : 0,
      "week_backups" : 0,
      "timezone" : "UTC+08:00"
    },
    "operation_type" : "backup",
    "id" : "cbb3ce6f-3332-4e7c-b98e-77290d8471ff"
  }
}`

	expectedListResponse = `
{
  "policies" : [ {
    "name" : "policy001",
    "associated_vaults" : [ ],
    "enabled" : true,
    "trigger" : {
  	"properties" : {
  	  "pattern" : [ "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00" ],
  	  "start_time" : "2019-05-08 06:57:05"
  	},
  	"type" : "time",
  	"id" : "d67269a6-5369-42d7-8150-5254bd446328",
  	"name" : "default"
    },
    "operation_definition" : {
  	  "retention_duration_days" : 1,
  	  "year_backups" : 0,
  	  "day_backups" : 0,
  	  "month_backups" : 0,
  	  "week_backups" : 0,
  	  "timezone" : "UTC+08:00"
    },
    "operation_type" : "backup",
    "id" : "cbb3ce6f-3332-4e7c-b98e-77290d8471ff"
  } ],
  "count" : 10
}`
)

var (
	createOpts = &policies.CreateOpts{
		Name: "policy001",
		OperationDefinition: &policies.PolicyODCreate{
			DailyBackups:          0,
			WeekBackups:           0,
			YearBackups:           0,
			MonthBackups:          0,
			MaxBackups:            0,
			RetentionDurationDays: 1,
			Timezone:              "UTC+08:00",
		},
		OperationType: "backup",
		Trigger: &policies.Trigger{
			Properties: policies.TriggerProperties{
				Pattern: []string{
					"FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00",
				},
			},
		},
	}

	expectedCreateResponseData = &policies.Policy{
		ID:      "cbb3ce6f-3332-4e7c-b98e-77290d8471ff",
		Name:    "policy001",
		Enabled: true,
		OperationDefinition: &policies.PolicyODCreate{
			DailyBackups:          0,
			WeekBackups:           0,
			YearBackups:           0,
			MonthBackups:          0,
			MaxBackups:            0,
			RetentionDurationDays: 1,
			Timezone:              "UTC+08:00",
		},
		OperationType: "backup",
		Trigger: &policies.PolicyTriggerResp{
			TriggerID: "d67269a6-5369-42d7-8150-5254bd446328",
			Name:      "default",
			Type:      "time",
			Properties: policies.PolicyTriggerPropertiesResp{
				Pattern:   []string{"FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00"},
				StartTime: "2019-05-08 06:57:05",
			},
		},
	}
	updateEnabled = true
	updateOpts    = policies.UpdateOpts{
		Enabled: &updateEnabled,
		Name:    "policy001",
		OperationDefinition: &policies.PolicyODCreate{
			DailyBackups:          0,
			WeekBackups:           0,
			YearBackups:           0,
			MonthBackups:          0,
			MaxBackups:            1,
			RetentionDurationDays: 1,
			Timezone:              "UTC+08:00",
		},
		Trigger: &policies.Trigger{
			Properties: policies.TriggerProperties{
				Pattern: []string{
					"FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00",
				},
			},
		},
	}

	expectedGetResponseData = &policies.Policy{
		ID:               "cbb3ce6f-3332-4e7c-b98e-77290d8471ff",
		Name:             "policy001",
		AssociatedVaults: []policies.PolicyAssociateVault{},
		Enabled:          true,
		OperationDefinition: &policies.PolicyODCreate{
			DailyBackups:          0,
			WeekBackups:           0,
			YearBackups:           0,
			MonthBackups:          0,
			MaxBackups:            0,
			RetentionDurationDays: 1,
			Timezone:              "UTC+08:00",
		},
		OperationType: "backup",
		Trigger: &policies.PolicyTriggerResp{
			TriggerID: "d67269a6-5369-42d7-8150-5254bd446328",
			Name:      "default",
			Type:      "time",
			Properties: policies.PolicyTriggerPropertiesResp{
				Pattern:   []string{"FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00"},
				StartTime: "2019-05-08 06:57:05",
			},
		},
	}

	expectedListResponseData = []policies.Policy{
		{
			ID:               "cbb3ce6f-3332-4e7c-b98e-77290d8471ff",
			Name:             "policy001",
			AssociatedVaults: []policies.PolicyAssociateVault{},
			Enabled:          true,
			OperationDefinition: &policies.PolicyODCreate{
				DailyBackups:          0,
				WeekBackups:           0,
				YearBackups:           0,
				MonthBackups:          0,
				MaxBackups:            0,
				RetentionDurationDays: 1,
				Timezone:              "UTC+08:00",
			},
			OperationType: "backup",
			Trigger: &policies.PolicyTriggerResp{
				TriggerID: "d67269a6-5369-42d7-8150-5254bd446328",
				Name:      "default",
				Type:      "time",
				Properties: policies.PolicyTriggerPropertiesResp{
					Pattern:   []string{"FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00"},
					StartTime: "2019-05-08 06:57:05",
				},
			},
		},
	}
)

func handlePolicyCreation(t *testing.T) {
	th.Mux.HandleFunc("/policies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handlePolicyDeletion(t *testing.T) {
	th.Mux.HandleFunc("/policies/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}

func handlePolicyUpdate(t *testing.T) {
	th.Mux.HandleFunc("/policies/cbb3ce6f-3332-4e7c-b98e-77290d8471ff", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedUpdateResponse)
	})
}

func handlePolicyGet(t *testing.T) {
	th.Mux.HandleFunc("/policies/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedGetResponse)
	})
}

func handlePolicyList(t *testing.T) {
	th.Mux.HandleFunc("/policies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}
