package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/apigw/v2/responses"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateV2Responses(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ResponsesCreate(t)

	actual, err := responses.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2Responses(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ResponsesGet(t)

	actual, err := responses.Get(client.ServiceClient(), "9b76174b785342078e557f23c01d5e41",
		"d060ade0560a4c01b89bf954ad2e9d6e", "baabc69fdb8f4c458637666c0441e9a4").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestListV2Responses(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ResponsesList(t)

	pages, err := responses.List(client.ServiceClient(), listOpts).AllPages()
	th.AssertNoErr(t, err)
	actual, err := responses.ExtractResponses(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2Responses(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ResponsesUpdate(t)

	actual, err := responses.Update(client.ServiceClient(), "baabc69fdb8f4c458637666c0441e9a4", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestDeleteV2Responses(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ResponsesDelete(t)

	err := responses.Delete(client.ServiceClient(), "9b76174b785342078e557f23c01d5e41",
		"d060ade0560a4c01b89bf954ad2e9d6e", "baabc69fdb8f4c458637666c0441e9a4").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestGetV2SpecResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2SpecResponseGet(t)

	errorType := "ACCESS_DENIED"
	actual, err := responses.GetSpecResp(client.ServiceClient(), errorType,
		specRespOpts).ExtractSpecResp("ACCESS_DENIED")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetSpecResponseData, actual)
}

func TestUpdateV2SpecResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2SpecResponseUpdate(t)

	errorType := "ACCESS_DENIED"
	actual, err := responses.UpdateSpecResp(client.ServiceClient(), errorType, specRespOpts,
		responseInfoOpts).ExtractSpecResp("ACCESS_DENIED")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetSpecResponseData, actual)
}

func TestDeleteV2SpecResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2SpecResponseDelete(t)

	errorType := "ACCESS_DENIED"
	err := responses.DeleteSpecResp(client.ServiceClient(), errorType, specRespOpts).ExtractErr()
	th.AssertNoErr(t, err)
}
