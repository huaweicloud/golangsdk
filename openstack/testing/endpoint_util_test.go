package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestContainsProjectId(t *testing.T) {
	endpointContains := []string{"https://as.eu-de.otc.t-systems.com/autoscaling-api/v1/f9842224f84e44f99c2878eddc7f9ef5",
		"https://elb.t-systems.com/rds/v1.0/c9842224f84e44f99c2878eddc7f9ef5/",
		"https://elb.eu-de.otc.t-systems.com/v1.1/c9842224f84e44f99c2878eddc7f9ef5",
		"https://elb.eu-de.otc.t-systems.com/v2/c9842224f84e44f99c2878eddc7f9ef5",
		"https://elb.eu-de.otc.t-systems.com/v2.0/c9842224f84e44f99c2878eddc7f9ef5",
		"https://elb.eu-de.otc.t-systems.com/V2.0/c9842224f84e44f99c2878eddc7f9ef5/list",
		"https://as.eu-de.otc.t-systems.com/autoscaling-api/v1/c9842224f84e44f99c2878eddc7f9ef5/abc",
		"https://as.eu-de.otc.t-systems.com/autoscaling-api/V11/c9842224f84e44f99c2878eddc7f9ef5",
		"https://as.eu-de.otc.t-systems.com/autoscaling-api/v2/c9842224f84e44f99c2878eddc7f9ef5",
		"https://as.eu-de.otc.t-systems.com/autoscaling-api/V2/c9842224f84e44f99c2878eddc7f9ef5",
		"http://as.eu-de.otc.t-systems.com/autoscaling-api/V2/c9842224f84e44f99c2878eddc7f9ef5"}

	for _, enpoint := range endpointContains {
		th.AssertEquals(t, true, openstack.ContainsProjectId(enpoint))
	}
}

func TestNotContainsProjectId(t *testing.T) {
	endpointContains := []string{"https://as.eu-de.otc.t-systems.com/autoscaling-api/v1",
		"https://as.eu-de.otc.t-systems.com/autoscaling-api/v1/",
		"https://as.eu-de.otc.t-systems.com/autoscaling-api/v1/abc",
		"https://as.eu-de.otc.t-systems.com/autoscaling-api/V1"}

	for _, enpoint := range endpointContains {
		th.AssertEquals(t, false, openstack.ContainsProjectId(enpoint))
	}
}
