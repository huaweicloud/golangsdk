package instances

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/common/structs"
	"github.com/huaweicloud/golangsdk/pagination"
)

var requestOpts golangsdk.RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

type ChargeInfoOpt structs.ChargeInfo

type DataStoreOpt struct {
	Type    string `json:"type" required:"true"`
	Version string `json:"version" required:"true"`
}

type BackupStrategyOpt struct {
	StartTime string `json:"start_time" required:"true"`
	KeepDays  string `json:"keep_days" required:"true"`
}

type CreateTaurusDBOpts struct {
	Name                string             `json:"name" required:"true"`
	Region              string             `json:"region" required:"true"`
	Mode                string             `json:"mode" required:"true"`
	Flavor              string             `json:"flavor_ref" required:"true"`
	VpcId               string             `json:"vpc_id" required:"true"`
	SubnetId            string             `json:"subnet_id" required:"true"`
	SecurityGroupId     string             `json:"security_group_id" required:"true"`
	Password            string             `json:"password" required:"true"`
	TimeZone            string             `json:"time_zone" required:"true"`
	AZMode              string             `json:"availability_zone_mode" required:"true"`
	SlaveCount          int                `json:"slave_count" required:"true"`
	MasterAZ            string             `json:"master_availability_zone,omitempty"`
	ConfigurationId     string             `json:"configuration_id,omitempty"`
	EnterpriseProjectId string             `json:"enterprise_project_id,omitempty"`
	DataStore           DataStoreOpt       `json:"datastore" required:"true"`
	BackupStrategy      *BackupStrategyOpt `json:"backup_strategy" required:"true"`
	ChargeInfo          *ChargeInfoOpt     `json:"charge_info,omitempty"`
}

type CreateTaurusDBBuilder interface {
	ToInstancesCreateMap() (map[string]interface{}, error)
}

func (opts CreateTaurusDBOpts) ToInstancesCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, opts CreateTaurusDBBuilder) (r CreateResult) {
	b, err := opts.ToInstancesCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: requestOpts.MoreHeaders,
	})

	return
}

func Delete(client *golangsdk.ServiceClient, instanceId string) (r DeleteResult) {
	url := deleteURL(client, instanceId)

	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 202},
		MoreHeaders: requestOpts.MoreHeaders,
	})

	return
}

func Get(client *golangsdk.ServiceClient, instanceId string) (r GetResult) {
	url := getURL(client, instanceId)

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})

	return
}

type ListTaurusDBInstanceOpts struct {
	Id            string `q:"id"`
	Name          string `q:"name"`
	Type          string `q:"type"`
	DataStoreType string `q:"datastore_type"`
	VpcId         string `q:"vpc_id"`
	SubnetId      string `q:"subnet_id"`
	Offset        int    `q:"offset"`
	Limit         int    `q:"limit"`
}

type ListTaurusDBBuilder interface {
	ToTaurusDBListDetailQuery() (string, error)
}

func (opts ListTaurusDBInstanceOpts) ToTaurusDBListDetailQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListTaurusDBBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToTaurusDBListDetailQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TaurusDBPage{pagination.SinglePageBase(r)}
	})
	// Headers supplies additional HTTP headers to populate on each paged request
	pageList.Headers = requestOpts.MoreHeaders

	return pageList
}
