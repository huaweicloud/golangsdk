package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/cci/v1/persistentvolumeclaims"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreateOptsMarshall(t *testing.T) {
	res, err := createOpts.ToPVCCreateMap()
	th.AssertNoErr(t, err)
	th.AssertJSONEquals(t, expectedRequest, res)
}

func TestCreateV1PersistentVolumeClaim(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePersistentVolumeClaimCreate(t)

	actual, err := persistentvolumeclaims.Create(client.ServiceClient(), createOpts, "terraform-test").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestListV1PersistentVolumeClaim(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePersistentVolumeClaimList(t)

	pages, err := persistentvolumeclaims.List(client.ServiceClient(), listOpts, "terraform-test").AllPages()
	th.AssertNoErr(t, err)
	actual, err := persistentvolumeclaims.ExtractPersistentVolumeClaims(pages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}
