package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/v3/securities"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestUpdateV2DatabaseSSL(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2DatabaseSSLUpdate(t)

	err := securities.UpdateSSL(client.ServiceClient(), "fda30974248d449e9dbdce8ae65d5ba0in01", sslOpts).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUpdateV2DatabasePort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2DatabasePortUpdate(t)

	actual, err := securities.UpdatePort(client.ServiceClient(), "fda30974248d449e9dbdce8ae65d5ba0in01",
		portOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}

func TestUpdateV2SecurityGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2SecurityGroupUpdate(t)

	actual, err := securities.UpdateSecGroup(client.ServiceClient(), "fda30974248d449e9dbdce8ae65d5ba0in01",
		secGroupOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}
