package testing

import (
	"fmt"
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/rts/v1/stackresources"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
	"net/http"
	"testing"
	"time"
)

// ListExpected represents the expected object from a List request.
var ListExpected = []stackresources.Resource{
	{
		Name: "hello_world",
		Links: []golangsdk.Link{
			{
				Href: "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world",
				Rel:  "self",
			},
			{
				Href: "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b",
				Rel:  "stack",
			},
		},
		LogicalID:    "hello_world",
		StatusReason: "state changed",
		UpdatedTime:  time.Date(2015, 2, 5, 21, 33, 11, 0, time.UTC),
		CreationTime: time.Date(2015, 2, 5, 21, 33, 10, 0, time.UTC),
		RequiredBy:   []string{},
		Status:       "CREATE_IN_PROGRESS",
		PhysicalID:   "49181cd6-169a-4130-9455-31185bbfc5bf",
		Type:         "OS::Nova::Server",
	},
}

// ListOutput represents the response body from a List request.
const ListOutput = `{
  "resources": [
  {
    "resource_name": "hello_world",
    "links": [
    {
      "href": "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world",
      "rel": "self"
    },
    {
      "href": "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b",
      "rel": "stack"
    }
    ],
    "logical_resource_id": "hello_world",
    "resource_status_reason": "state changed",
    "updated_time": "2015-02-05T21:33:11",
    "required_by": [],
    "resource_status": "CREATE_IN_PROGRESS",
    "physical_resource_id": "49181cd6-169a-4130-9455-31185bbfc5bf",
	"creation_time": "2015-02-05T21:33:10",
    "resource_type": "OS::Nova::Server",
	"attributes": {"SXSW": "atx"},
	"description": "Some resource"
  }
]
}`

// HandleListSuccessfully creates an HTTP handler at `/stacks/hello_world/49181cd6-169a-4130-9455-31185bbfc5bf/resources`
// on the test handler mux that responds with a `List` response.
func HandleListSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/stacks/hello_world/resources", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, output)
		case "49181cd6-169a-4130-9455-31185bbfc5bf":
			fmt.Fprintf(w, `{"resources":[]}`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}
