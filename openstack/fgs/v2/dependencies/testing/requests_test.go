package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/fgs/v2/dependencies"
	"github.com/huaweicloud/golangsdk/pagination"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestListV2Dependencies(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2DependenciesList(t)

	actual := make([]dependencies.Dependency, 0)
	err := dependencies.List(client.ServiceClient(), listOpts).EachPage(
		func(page pagination.Page) (bool, error) {
			resp, err := dependencies.ExtractDependencies(page)
			th.AssertNoErr(t, err)
			actual = append(actual, resp.Dependencies...)
			return true, nil
		})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, expectedListResponseData, actual)
}
