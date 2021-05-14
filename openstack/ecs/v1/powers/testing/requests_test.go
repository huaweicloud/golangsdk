package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/ecs/v1/powers"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestPowerOnV1(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePowerOn(t)

	actual, err := powers.PowerAction(client.ServiceClient(), powerOpts, "on-start").ExtractJobResponse()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedPowerOnResponseData, actual)
}

func TestPowerOffV1(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePowerOff(t)

	actual, err := powers.PowerAction(client.ServiceClient(), powerOpts, "on-stop").ExtractJobResponse()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedPowerOffResponseData, actual)
}
