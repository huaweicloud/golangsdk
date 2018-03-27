package testing

import (
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/identity/v3/services"
	"github.com/huaweicloud/golangsdk/pagination"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateSuccessful(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateServiceSuccessfully(t)

	createOpts := services.CreateOpts{
		Type: "compute",
		Extra: map[string]interface{}{
			"name":        "service-two",
			"description": "Service Two",
			"email":       "service@example.com",
		},
	}

	actual, err := services.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondService, *actual)
}

func TestListServices(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListServicesSuccessfully(t)

	count := 0
	err := services.List(client.ServiceClient(), services.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := services.ExtractServices(page)
		th.AssertNoErr(t, err)

		th.CheckDeepEquals(t, ExpectedServicesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, count, 1)
}

func TestListServicesAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListServicesSuccessfully(t)

	allPages, err := services.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := services.ExtractServices(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedServicesSlice, actual)
}

func TestGetSuccessful(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetServiceSuccessfully(t)

	actual, err := services.Get(client.ServiceClient(), "c2474183dca7453bbd73123a0b78feae").Extract()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondService, *actual)
}

func TestUpdateSuccessful(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateServiceSuccessfully(t)

	updateOpts := services.UpdateOpts{
		Type: "compute2",
		Extra: map[string]interface{}{
			"description": "Service Two Updated",
		},
	}
	actual, err := services.Update(client.ServiceClient(), "c2474183dca7453bbd73123a0b78feae", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondServiceUpdated, *actual)
}

func TestDeleteSuccessful(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/services/053d21d488d1463c818132d9d08fb617", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := services.Delete(client.ServiceClient(), "053d21d488d1463c818132d9d08fb617")
	th.AssertNoErr(t, res.Err)
}
