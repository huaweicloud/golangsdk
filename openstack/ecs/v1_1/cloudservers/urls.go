package cloudservers

import "github.com/huaweicloud/golangsdk"

func createURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL("cloudservers")
}

func jobURL(sc *golangsdk.ServiceClient, jobId string) string {
	return sc.ServiceURL("jobs", jobId)
}
