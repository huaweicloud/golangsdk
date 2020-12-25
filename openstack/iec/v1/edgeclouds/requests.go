package edgeclouds

import (
	"net/http"

	"github.com/huaweicloud/golangsdk"
)

func Get(client *golangsdk.ServiceClient, edgeCloudID string) (r GetResult) {
	url := GetURL(client, edgeCloudID)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusOK},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, edgeCloudID string) (r DeleteResult) {
	_, r.Err = client.Delete(DeleteURL(client, edgeCloudID), &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusNoContent},
	})
	return
}
