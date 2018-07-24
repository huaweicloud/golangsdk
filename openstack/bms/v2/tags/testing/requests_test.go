package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/bms/v2/tags"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateTag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/2bff7a8a-3934-4f79-b1d6-53dc5540f00e/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "tags": [
        "__type_baremetal"
    ]
}
			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "tags": [
        "__type_baremetal"
    ]
}	`)
	})

	options := tags.CreateOpts{
		Tag: []string{"__type_baremetal"},
	}
	n, err := tags.Create(fake.ServiceClient(), "2bff7a8a-3934-4f79-b1d6-53dc5540f00e", options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "__type_baremetal", n.Tags[0])
}

func TestDeleteTag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/2bff7a8a-3934-4f79-b1d6-53dc5540f00e/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := tags.Delete(fake.ServiceClient(), "2bff7a8a-3934-4f79-b1d6-53dc5540f00e")
	th.AssertNoErr(t, res.Err)
}

func TestGetTags(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/2bff7a8a-3934-4f79-b1d6-53dc5540f00e/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "tags": [
        "__type_baremetal"
    ]
}
		`)
	})

	n, err := tags.Get(fake.ServiceClient(), "2bff7a8a-3934-4f79-b1d6-53dc5540f00e").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "__type_baremetal", n.Tags[0])

}
