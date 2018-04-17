package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/antiddos/v1/antiddos"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	createOpt := antiddos.CreateOpts{
		EnableL7:            true,
		TrafficPosId:        1,
		HttpRequestPosId:    2,
		CleaningAccessPosId: 3,
		AppTypeId:           1,
	}

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.Create(client.ServiceClient(), floatingIpId, createOpt).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.Delete(client.ServiceClient(), floatingIpId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &DeleteResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.Get(client.ServiceClient(), floatingIpId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	updateOpt := antiddos.UpdateOpts{
		EnableL7:            true,
		TrafficPosId:        1,
		HttpRequestPosId:    2,
		CleaningAccessPosId: 3,
		AppTypeId:           1,
	}

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.Update(client.ServiceClient(), floatingIpId, updateOpt).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestListStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListStatusSuccessfully(t)

	listOpt := antiddos.ListStatusOpts{
		Limit:  2,
		Offset: 1,
		Status: "notConfig",
		Ip:     "49.",
	}

	actual, err := antiddos.ListStatus(client.ServiceClient(), listOpt).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListStatusResponse, actual)
}

func TestListConfigs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListConfigsSuccessfully(t)

	actual, err := antiddos.ListConfigs(client.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListConfigsResponse, actual)
}

func TestWeeklyReport(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWeeklyReportSuccessfully(t)

	actual, err := antiddos.WeeklyReport(client.ServiceClient(), antiddos.WeeklyReportOpts{PeriodStartDate: responsePeriodTime}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &WeeklyReportResponse, actual)
}

func TestListLogs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListLogsSuccessfully(t)

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.ListLogs(client.ServiceClient(), floatingIpId, antiddos.ListLogsOpts{
		Limit:   2,
		Offset:  1,
		SortDir: "asc",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListLogsResponse, actual)
}

func TestGetStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetStatusSuccessfully(t)

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.GetStatus(client.ServiceClient(), floatingIpId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetStatusResponse, actual)
}

func TestDailyReport(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDailyReportSuccessfully(t)

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.DailyReport(client.ServiceClient(), floatingIpId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &DailyReportResponse, actual)
}

func TestGetTask(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetTaskSuccessfully(t)

	actual, err := antiddos.GetTask(client.ServiceClient(), antiddos.GetTaskOpts{
		TaskId: "4a4fefe7-34a1-40e2-a87c-16932af3ac4a",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetTaskResponse, actual)
}
