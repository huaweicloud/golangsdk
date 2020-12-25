package cloudvolumes

import (
	"net/http"

	"github.com/huaweicloud/golangsdk"
)

func Get(client *golangsdk.ServiceClient, CloudVolumeID string) (r GetResult) {
	url := GetURL(client, CloudVolumeID)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	return
}

func ListVolumeType(client *golangsdk.ServiceClient) (r GetResult) {
	url := ListVolumeTypeURL(client)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	return
}
