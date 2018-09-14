package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/csbs/v1/policies"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(policiesEndpoint+"/"+policies_id, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	s, err := policies.Get(fake.ServiceClient(), policies_id).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "5af626d2-19b9-4dc4-8e95-ddba008318b3", s.ID)
	th.AssertEquals(t, "c2c-policy", s.Name)
	th.AssertEquals(t, "OS::Nova::Server", s.Resources[0].Type)
	th.AssertEquals(t, "resource1", s.Resources[0].Name)
	th.AssertEquals(t, "cd5955b4-44c0-4f0a-ac57-2401b89cb347", s.Resources[0].Id)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc(policiesEndpoint, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, createRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, createResponse)
	})

	options := &policies.CreateOpts{
		Name:        "c2c-policy",
		Description: "My plan",
		ProviderId:  "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
		Parameters: policies.PolicyParam{
			Common: map[string]interface{}{},
		},
		ScheduledOperations: []policies.ScheduledOperation{{
			Name:        "my-backup-policy",
			Description: "My backup policy",
			Enabled:     true,
			OperationDefinition: policies.OperationDefinition{
				MaxBackups: 20,
			},
			Trigger: policies.Trigger{
				Properties: policies.TriggerProperties{
					Pattern: "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
				},
			},
			OperationType: "backup",
		}},
		Resources: []policies.Resource{{
			Id:   "cd5955b4-44c0-4f0a-ac57-2401b89cb347",
			Type: "OS::Nova::Server",
			Name: "resource1"}},
	}
	n, err := policies.Create(fake.ServiceClient(), options).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, n.ID, "5af626d2-19b9-4dc4-8e95-ddba008318b3")
	th.AssertEquals(t, n.Status, "suspended")
	th.AssertEquals(t, "OS::Nova::Server", n.Resources[0].Type)
	th.AssertEquals(t, "resource1", n.Resources[0].Name)
	th.AssertEquals(t, "cd5955b4-44c0-4f0a-ac57-2401b89cb347", n.Resources[0].Id)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(policiesEndpoint+"/"+policies_id, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
	})

	result := policies.Delete(fake.ServiceClient(), policies_id)

	th.AssertNoErr(t, result.Err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(policiesEndpoint+"/"+policies_id, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, updateResponse)
	})

	options := &policies.UpdateOpts{
		Name: "c2c-policy-update",
		Parameters: policies.PolicyParam{
			Common: map[string]interface{}{},
		},
		ScheduledOperations: []policies.ScheduledOperationToUpdate{{
			Name:        "my-backup-policy",
			Description: "My backup policy",
			Enabled:     true,
			Id:          "b70c712d-f48b-43f7-9a0f-3bab86d59149",
			OperationDefinition: policies.OperationDefinition{
				RetentionDurationDays: -1,
				MaxBackups:            20,
			},
			Trigger: policies.Trigger{
				Properties: policies.TriggerProperties{
					Pattern: "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
				},
			},
		}},
	}
	n, err := policies.Update(fake.ServiceClient(), policies_id, options).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, n.Name, "c2c-policy-update")
	th.AssertEquals(t, n.ID, "5af626d2-19b9-4dc4-8e95-ddba008318b3")
}

func TestList(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(policiesEndpoint, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, listResponse)
	})

	actual, err := policies.List(fake.ServiceClient(), policies.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract backup policies: %v", err)
	}

	var CreatedAt, _ = time.Parse(golangsdk.RFC3339MilliNoZ, "2018-08-20T10:43:56.246383")
	expected := []policies.BackupPolicy{
		{
			Status:      "suspended",
			ProviderId:  "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
			Description: "My plann",
			ScheduledOperations: []policies.ScheduledOperationResp{{

				Description: "My backup policy",
				Enabled:     true,
				TriggerID:   "831b5e69-0b75-420c-918e-9cbcb32d97f1",
				Trigger: policies.TriggerResp{
					Properties: policies.TriggerPropertiesResp{
						Pattern: "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
					},
					Type: "time",
					ID:   "831b5e69-0b75-420c-918e-9cbcb32d97f1",
					Name: "default",
				},
				OperationDefinition: policies.OperationDefinitionResp{
					MaxBackups: 5,
					ProviderId: "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
					PlanId:     "4d1ce19b-d681-4e44-a87e-c44eb9bfc4c7",
				},
				OperationType: "backup",
				ID:            "e7d50d4c-2f38-40a4-9f9b-c9c355a52417",
				Name:          "my-backupp",
			},
			},
			ID:   "4d1ce19b-d681-4e44-a87e-c44eb9bfc4c7",
			Name: "my-plan-test1",
			Parameters: policies.PolicyParam{
				Common: map[string]interface{}{},
			},
			CreatedAt: CreatedAt,
			ProjectId: "91d687759aed45d28b5f6084bc2fa8ad",
			Resources: []policies.Resource{
				{
					Type: "OS::Nova::Server",
					Id:   "9422f270-6fcf-4ba2-9319-a007f2f63a8e",
					Name: "resource4",
				},
			},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}
