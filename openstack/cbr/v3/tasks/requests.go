package tasks

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type ListOpts struct {
	EndTime             string `q:"end_time"`
	EnterpriseProjectId string `q:"enterprise_project_id"`
	Limit               int    `q:"limit"`
	Offset              int    `q:"offset"`
	OperationType       string `q:"operation_type"`
	ProviderId          string `q:"provider_id"`
	ResourceId          string `q:"resource_id"`
	ResourceName        string `q:"resource_name"`
	StartTime           string `q:"start_time"`
	Status              string `q:"status"`
	VaultId             string `q:"vault_id"`
	VaultName           string `q:"vault_name"`
}

type ListOptsBuilder interface {
	ToTaskListQuery() (string, error)
}

func (opts ListOpts) ToTaskListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToTaskListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TaskPage{pagination.SinglePageBase(r)}
	})
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(singleURL(client, id), &r.Body, nil)
	return
}
