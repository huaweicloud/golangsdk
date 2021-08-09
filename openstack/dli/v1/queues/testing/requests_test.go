package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/dli/v1/queues"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleList(t)

	listResult := queues.List(client.ServiceClient(), queues.ListOpts{})
	th.AssertNoErr(t, listResult.Err)
	rt := listResult.Body.(*queues.ListResult)
	th.AssertDeepEquals(t, expectedListResponseData, rt.Queues[0])
}

func TestScale(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleScale(t)

	result := queues.ScaleOrRestart(client.ServiceClient(), queues.ActionOpts{
		Action:    "scale_out",
		CuCount:   16,
		QueueName: queueName1,
	})

	th.AssertNoErr(t, result.Err)
}
