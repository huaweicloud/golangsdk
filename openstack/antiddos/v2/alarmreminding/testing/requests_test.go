package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/antiddos/v2/alarmreminding"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestQueryTraffic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWarnAlertSuccessfully(t)

	actual, err := alarmreminding.WarnAlert(client.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &WarnAlertResponse, actual)
}
