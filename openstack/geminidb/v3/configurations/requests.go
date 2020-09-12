package configurations

import (
	"github.com/huaweicloud/golangsdk"
)

type ApplyOpts struct {
	InstanceIds []string `json:"instance_ids" required:"true"`
}

type ApplyBuilder interface {
	ToConfigApplyMap() (map[string]interface{}, error)
}

func (opts ApplyOpts) ToConfigApplyMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Apply(client *golangsdk.ServiceClient, id string, opts ApplyBuilder) (r ApplyResult) {
	b, err := opts.ToConfigApplyMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(applyURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

func GetInstanceConfig(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(instanceConfigURL(client, id), &r.Body, nil)
	return
}
