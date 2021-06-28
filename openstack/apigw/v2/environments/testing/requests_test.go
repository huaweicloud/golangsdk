package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/environments"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2Environment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2EnvironmentCreate(t)

	actual, err := environments.Create(client.ServiceClient(), "cc4ea721cc6747f7969e06bd21121c52",
		createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestListV2Environment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2EnvironmentList(t)

	pages, err := environments.List(client.ServiceClient(), "cc4ea721cc6747f7969e06bd21121c52",
		environments.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := environments.ExtractEnvironments(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2Environment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2EnvironmentUpdate(t)

	actual, err := environments.Update(client.ServiceClient(), "cc4ea721cc6747f7969e06bd21121c52",
		"3585fce96a5d44f8b445121b9440274a", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedUpdateResponseData, actual)
}

func TestDeleteV2Environment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2EnvironmentDelete(t)

	err := environments.Delete(client.ServiceClient(), "cc4ea721cc6747f7969e06bd21121c52",
		"3585fce96a5d44f8b445121b9440274a").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestCreateV2EnvironmentVariable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2EnvironmentVariableCreate(t)

	actual, err := environments.CreateVariable(client.ServiceClient(), "cc4ea721cc6747f7969e06bd21121c52",
		variableCreateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateVariableResponseData, actual)
}

func TestGetV2EnvironmentVariable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2EnvironmentVariableGet(t)

	actual, err := environments.GetVariable(client.ServiceClient(), "cc4ea721cc6747f7969e06bd21121c52",
		"2dc48632332f4157804175175e71e3e8").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateVariableResponseData, actual)
}

func TestListV2EnvironmentVariable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2EnvironmentVariableList(t)

	pages, err := environments.ListVariables(client.ServiceClient(), "cc4ea721cc6747f7969e06bd21121c52",
		environments.ListVariablesOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := environments.ExtractVariables(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListVariableResponseData, actual)
}

func TestDeleteV2EnvironmentVariable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2EnvironmentVariableDelete(t)

	err := environments.DeleteVariable(client.ServiceClient(), "cc4ea721cc6747f7969e06bd21121c52",
		"2dc48632332f4157804175175e71e3e8").ExtractErr()
	th.AssertNoErr(t, err)
}
