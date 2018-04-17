package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/huaweicloud/golangsdk/openstack/antiddos/v1/antiddos"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const CreateOutput string = `
{
  "error_code": "10000000",
  "error_description": "The task has been received and is being handled",
  "task_id": "82463800-70fe-4cba-9a96-06175e246ab3"
}
`

const CreateRequest string = `
{
  "app_type_id": 1,
  "cleaning_access_pos_id": 3,
  "enable_L7": true,
  "http_request_pos_id": 2,
  "traffic_pos_id": 1
}
`

var CreateResponse = antiddos.CreateResponse{
	ErrorCode:        "10000000",
	ErrorDescription: "The task has been received and is being handled",
	TaskId:           "82463800-70fe-4cba-9a96-06175e246ab3",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos/82abaa86-8518-47db-8d63-ddf152824635", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var DeleteResponse = antiddos.DeleteResponse{
	ErrorCode:        "10000000",
	ErrorDescription: "The task has been received and is being handled",
	TaskId:           "f732e7f1-26b2-40f1-85e9-a8a4d3a43038",
}

const DeleteOutput string = `
{
  "error_code": "10000000",
  "error_description": "The task has been received and is being handled",
  "task_id": "f732e7f1-26b2-40f1-85e9-a8a4d3a43038"
}
`

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos/82abaa86-8518-47db-8d63-ddf152824635", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DeleteOutput)
	})
}

var GetResponse = antiddos.GetResponse{
	EnableL7:            true,
	TrafficPosId:        1,
	HttpRequestPosId:    2,
	CleaningAccessPosId: 3,
	AppTypeId:           1,
}

const GetOutput string = `
{
  "enable_L7": true,
  "traffic_pos_id": 1,
  "http_request_pos_id": 2,
  "cleaning_access_pos_id": 3,
  "app_type_id": 1
}
`

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos/82abaa86-8518-47db-8d63-ddf152824635", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

const UpdateOutput string = `
{
  "error_code": "10000000",
  "error_description": "The task has been received and is being handled",
  "task_id": "82463800-70fe-4cba-9a96-06175e246ab3"
}
`

const UpdateRequest string = `
{
  "app_type_id": 1,
  "cleaning_access_pos_id": 3,
  "enable_L7": true,
  "http_request_pos_id": 2,
  "traffic_pos_id": 1
}
`

var UpdateResponse = antiddos.UpdateResponse{
	ErrorCode:        "10000000",
	ErrorDescription: "The task has been received and is being handled",
	TaskId:           "82463800-70fe-4cba-9a96-06175e246ab3",
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos/82abaa86-8518-47db-8d63-ddf152824635", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var ListStatusResponse = antiddos.ListStatusResponse{
	Total: 2,
}

const ListStatusOutput string = `
{
  "total": 2,
  "ddosStatus": [{
    "floating_ip_id": "4d60bba4-0791-4e82-8262-9bdffaeb1d14",
    "floating_ip_address": "49.4.4.36",
    "network_type": "EIP",
    "status": "notConfig",
    "blackhole_endtime": 0
  }]
}
`

func HandleListStatusSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()
		ip := r.Form.Get("ip")
		limit := r.Form.Get("limit")
		offset := r.Form.Get("offset")
		status := r.Form.Get("status")
		if ip == "49." && limit == "2" && offset == "1" && status == "notConfig" {
			fmt.Fprintf(w, ListStatusOutput)
		}
	})

	ListStatusResponse.DdosStatus = append(ListStatusResponse.DdosStatus, struct {
		FloatingIpAddress string `json:"floating_ip_address,"`
		FloatingIpId      string `json:"floating_ip_id,"`
		NetworkType       string `json:"network_type,"`
		Status            string `json:"status,"`
	}{
		FloatingIpId:      "4d60bba4-0791-4e82-8262-9bdffaeb1d14",
		FloatingIpAddress: "49.4.4.36",
		NetworkType:       "EIP",
		Status:            "notConfig",
	})
}

var ListConfigsResponse = antiddos.ListConfigsResponse{
	TrafficLimitedList: []struct {
		TrafficPosId     int `json:"traffic_pos_id,"`
		TrafficPerSecond int `json:"traffic_per_second,"`
		PacketPerSecond  int `json:"packet_per_second,"`
	}{
		{
			TrafficPosId:     1,
			TrafficPerSecond: 10,
			PacketPerSecond:  2000,
		},
		{
			TrafficPosId:     2,
			TrafficPerSecond: 30,
			PacketPerSecond:  6000,
		},
		{
			TrafficPosId:     3,
			TrafficPerSecond: 50,
			PacketPerSecond:  10000,
		},
		{
			TrafficPosId:     4,
			TrafficPerSecond: 70,
			PacketPerSecond:  15000,
		},
		{
			TrafficPosId:     5,
			TrafficPerSecond: 100,
			PacketPerSecond:  20000,
		},
		{
			TrafficPosId:     6,
			TrafficPerSecond: 150,
			PacketPerSecond:  25000,
		},
		{
			TrafficPosId:     7,
			TrafficPerSecond: 200,
			PacketPerSecond:  35000,
		},
		{
			TrafficPosId:     8,
			TrafficPerSecond: 250,
			PacketPerSecond:  50000,
		},
		{
			TrafficPosId:     9,
			TrafficPerSecond: 300,
			PacketPerSecond:  70000,
		},
	},
	HttpLimitedList: []struct {
		HttpRequestPosId    int `json:"http_request_pos_id,"`
		HttpPacketPerSecond int `json:"http_packet_per_second,"`
	}{
		{
			HttpRequestPosId:    1,
			HttpPacketPerSecond: 100,
		},
		{
			HttpRequestPosId:    2,
			HttpPacketPerSecond: 150,
		},
		{
			HttpRequestPosId:    3,
			HttpPacketPerSecond: 240,
		},
		{
			HttpRequestPosId:    4,
			HttpPacketPerSecond: 350,
		},
		{
			HttpRequestPosId:    5,
			HttpPacketPerSecond: 480,
		},
		{
			HttpRequestPosId:    6,
			HttpPacketPerSecond: 550,
		},
		{
			HttpRequestPosId:    7,
			HttpPacketPerSecond: 700,
		},
		{
			HttpRequestPosId:    8,
			HttpPacketPerSecond: 850,
		},
		{
			HttpRequestPosId:    9,
			HttpPacketPerSecond: 1000,
		},
		{
			HttpRequestPosId:    10,
			HttpPacketPerSecond: 1500,
		},
		{
			HttpRequestPosId:    11,
			HttpPacketPerSecond: 2000,
		},
		{
			HttpRequestPosId:    12,
			HttpPacketPerSecond: 3000,
		},
		{
			HttpRequestPosId:    13,
			HttpPacketPerSecond: 5000,
		},
		{
			HttpRequestPosId:    14,
			HttpPacketPerSecond: 10000,
		},
		{
			HttpRequestPosId:    15,
			HttpPacketPerSecond: 20000,
		},
	},
	ConnectionLimitedList: []struct {
		CleaningAccessPosId    int `json:"cleaning_access_pos_id,"`
		NewConnectionLimited   int `json:"new_connection_limited,"`
		TotalConnectionLimited int `json:"total_connection_limited,"`
	}{
		{
			CleaningAccessPosId:    1,
			NewConnectionLimited:   10,
			TotalConnectionLimited: 30,
		},
		{
			CleaningAccessPosId:    2,
			NewConnectionLimited:   20,
			TotalConnectionLimited: 100,
		},
		{
			CleaningAccessPosId:    3,
			NewConnectionLimited:   30,
			TotalConnectionLimited: 200,
		},
		{
			CleaningAccessPosId:    4,
			NewConnectionLimited:   40,
			TotalConnectionLimited: 250,
		},
		{
			CleaningAccessPosId:    5,
			NewConnectionLimited:   50,
			TotalConnectionLimited: 300,
		},
		{
			CleaningAccessPosId:    6,
			NewConnectionLimited:   60,
			TotalConnectionLimited: 500,
		},
		{
			CleaningAccessPosId:    7,
			NewConnectionLimited:   70,
			TotalConnectionLimited: 600,
		},
		{
			CleaningAccessPosId:    8,
			NewConnectionLimited:   80,
			TotalConnectionLimited: 700,
		},
	},
}

var ListConfigsOutput = `
{
  "traffic_limited_list": [{
    "traffic_pos_id": 1,
    "traffic_per_second": 10,
    "packet_per_second": 2000
  }, {
    "traffic_pos_id": 2,
    "traffic_per_second": 30,
    "packet_per_second": 6000
  }, {
    "traffic_pos_id": 3,
    "traffic_per_second": 50,
    "packet_per_second": 10000
  }, {
    "traffic_pos_id": 4,
    "traffic_per_second": 70,
    "packet_per_second": 15000
  }, {
    "traffic_pos_id": 5,
    "traffic_per_second": 100,
    "packet_per_second": 20000
  }, {
    "traffic_pos_id": 6,
    "traffic_per_second": 150,
    "packet_per_second": 25000
  }, {
    "traffic_pos_id": 7,
    "traffic_per_second": 200,
    "packet_per_second": 35000
  }, {
    "traffic_pos_id": 8,
    "traffic_per_second": 250,
    "packet_per_second": 50000
  }, {
    "traffic_pos_id": 9,
    "traffic_per_second": 300,
    "packet_per_second": 70000
  }],
  "http_limited_list": [{
    "http_request_pos_id": 1,
    "http_packet_per_second": 100
  }, {
    "http_request_pos_id": 2,
    "http_packet_per_second": 150
  }, {
    "http_request_pos_id": 3,
    "http_packet_per_second": 240
  }, {
    "http_request_pos_id": 4,
    "http_packet_per_second": 350
  }, {
    "http_request_pos_id": 5,
    "http_packet_per_second": 480
  }, {
    "http_request_pos_id": 6,
    "http_packet_per_second": 550
  }, {
    "http_request_pos_id": 7,
    "http_packet_per_second": 700
  }, {
    "http_request_pos_id": 8,
    "http_packet_per_second": 850
  }, {
    "http_request_pos_id": 9,
    "http_packet_per_second": 1000
  }, {
    "http_request_pos_id": 10,
    "http_packet_per_second": 1500
  }, {
    "http_request_pos_id": 11,
    "http_packet_per_second": 2000
  }, {
    "http_request_pos_id": 12,
    "http_packet_per_second": 3000
  }, {
    "http_request_pos_id": 13,
    "http_packet_per_second": 5000
  }, {
    "http_request_pos_id": 14,
    "http_packet_per_second": 10000
  }, {
    "http_request_pos_id": 15,
    "http_packet_per_second": 20000
  }],
  "connection_limited_list": [{
    "cleaning_access_pos_id": 1,
    "new_connection_limited": 10,
    "total_connection_limited": 30
  }, {
    "cleaning_access_pos_id": 2,
    "new_connection_limited": 20,
    "total_connection_limited": 100
  }, {
    "cleaning_access_pos_id": 3,
    "new_connection_limited": 30,
    "total_connection_limited": 200
  }, {
    "cleaning_access_pos_id": 4,
    "new_connection_limited": 40,
    "total_connection_limited": 250
  }, {
    "cleaning_access_pos_id": 5,
    "new_connection_limited": 50,
    "total_connection_limited": 300
  }, {
    "cleaning_access_pos_id": 6,
    "new_connection_limited": 60,
    "total_connection_limited": 500
  }, {
    "cleaning_access_pos_id": 7,
    "new_connection_limited": 70,
    "total_connection_limited": 600
  }, {
    "cleaning_access_pos_id": 8,
    "new_connection_limited": 80,
    "total_connection_limited": 700
  }],
  "extend_ddos_config": []
}
`

func HandleListConfigsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos/query_config_list", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListConfigsOutput)
	})
}

var WeeklyReportOutput = `
{
  "ddos_intercept_times": 0,
  "weekdata": [{
    "ddos_intercept_times": 0,
    "ddos_blackhole_times": 0,
    "max_attack_bps": 0,
    "max_attack_conns": 0,
    "period_start_date": 1519862400000
  }, {
    "ddos_intercept_times": 0,
    "ddos_blackhole_times": 0,
    "max_attack_bps": 0,
    "max_attack_conns": 0,
    "period_start_date": 1519862400000
  }, {
    "ddos_intercept_times": 0,
    "ddos_blackhole_times": 0,
    "max_attack_bps": 0,
    "max_attack_conns": 0,
    "period_start_date": 1519862400000
  }, {
    "ddos_intercept_times": 0,
    "ddos_blackhole_times": 0,
    "max_attack_bps": 0,
    "max_attack_conns": 0,
    "period_start_date": 1519862400000
  }, {
    "ddos_intercept_times": 0,
    "ddos_blackhole_times": 0,
    "max_attack_bps": 0,
    "max_attack_conns": 0,
    "period_start_date": 1519862400000
  }, {
    "ddos_intercept_times": 0,
    "ddos_blackhole_times": 0,
    "max_attack_bps": 0,
    "max_attack_conns": 0,
    "period_start_date": 1519862400000
  }, {
    "ddos_intercept_times": 0,
    "ddos_blackhole_times": 0,
    "max_attack_bps": 0,
    "max_attack_conns": 0,
    "period_start_date": 1519862400000
  }],
  "top10": [],
  "weeklyCountTableHeader": ["ddos_intercept_times", "max_attack_bps", "max_attack_conns", "period_start_date"],
  "weeklyCountTableBody": [
    ["0", "0", "0", "2018-03-12 07:00:42"],
    ["0", "0", "0", "2018-03-13 07:00:42"],
    ["0", "0", "0", "2018-03-14 07:00:42"],
    ["0", "0", "0", "2018-03-15 07:00:42"],
    ["0", "0", "0", "2018-03-16 07:00:42"],
    ["0", "0", "0", "2018-03-17 07:00:42"],
    ["0", "0", "0", "2018-03-18 07:00:42"]
  ],
  "weeklyTop10TableHeader": ["floating_ip_address", "times"],
  "weeklyTop10TableBody": []
}
`

//init the loc
var responsePeriodTime = time.Date(2018, 3, 1, 0, 0, 0, 0, time.UTC)

var WeeklyReportResponse = antiddos.WeeklyReportResponse{
	DdosInterceptTimes: 0,
	Weekdata: []antiddos.WeekData{
		{
			DdosInterceptTimes: 0,
			DdosBlackholeTimes: 0,
			MaxAttackBps:       0,
			MaxAttackConns:     0,
			PeriodStartDate:    responsePeriodTime,
		},
		{
			DdosInterceptTimes: 0,
			DdosBlackholeTimes: 0,
			MaxAttackBps:       0,
			MaxAttackConns:     0,
			PeriodStartDate:    responsePeriodTime,
		},
		{
			DdosInterceptTimes: 0,
			DdosBlackholeTimes: 0,
			MaxAttackBps:       0,
			MaxAttackConns:     0,
			PeriodStartDate:    responsePeriodTime,
		},
		{
			DdosInterceptTimes: 0,
			DdosBlackholeTimes: 0,
			MaxAttackBps:       0,
			MaxAttackConns:     0,
			PeriodStartDate:    responsePeriodTime,
		},
		{
			DdosInterceptTimes: 0,
			DdosBlackholeTimes: 0,
			MaxAttackBps:       0,
			MaxAttackConns:     0,
			PeriodStartDate:    responsePeriodTime,
		},
		{
			DdosInterceptTimes: 0,
			DdosBlackholeTimes: 0,
			MaxAttackBps:       0,
			MaxAttackConns:     0,
			PeriodStartDate:    responsePeriodTime,
		},
		{
			DdosInterceptTimes: 0,
			DdosBlackholeTimes: 0,
			MaxAttackBps:       0,
			MaxAttackConns:     0,
			PeriodStartDate:    responsePeriodTime,
		},
	},
	Top10: []struct {
		FloatingIpAddress string `json:"floating_ip_address,"`
		Times             int    `json:"times,"`
	}{},
}

func HandleWeeklyReportSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos/weekly", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, WeeklyReportOutput)
	})
}

var ListLogsOutput = `
{
  "tableHeader": ["start_time", "end_time", "status", "trigger_bps", "trigger_pps", "trigger_http_pps"],
  "tableBody": [],
  "total": 1,
  "logs": [{
    "start_time": 1473217200000,
    "end_time": 1473242400000,
    "status": 1,
    "trigger_bps": 51106,
    "trigger_pps": 2600,
    "trigger_http_pps": 3589
  }]
}
`

var ListLogsResponse = antiddos.ListLogsResponse{
	Total: 1,
	Logs: []struct {
		StartTime      int `json:"start_time,"`
		EndTime        int `json:"end_time,"`
		Status         int `json:"status,"`
		TriggerBps     int `json:"trigger_bps,"`
		TriggerPps     int `json:"trigger_pps,"`
		TriggerHttpPps int `json:"trigger_http_pps,"`
	}{
		{
			StartTime:      1473217200000,
			EndTime:        1473242400000,
			Status:         1,
			TriggerBps:     51106,
			TriggerPps:     2600,
			TriggerHttpPps: 3589,
		},
	},
}

func HandleListLogsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos/82abaa86-8518-47db-8d63-ddf152824635/logs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListLogsOutput)
	})
}

var GetStatusOutput = `{"status":"normal"}`

var GetStatusResponse = antiddos.GetStatusResponse{
	Status: "normal",
}

func HandleGetStatusSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos/82abaa86-8518-47db-8d63-ddf152824635/status", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetStatusOutput)
	})
}

var DailyReportOutput = `
{
  "data": [{
    "period_start": 1521650068727,
    "bps_in": 0,
    "bps_attack": 0,
    "total_bps": 0,
    "pps_in": 0,
    "pps_attack": 0,
    "total_pps": 0
  }, {
    "period_start": 1521650368727,
    "bps_in": 0,
    "bps_attack": 0,
    "total_bps": 0,
    "pps_in": 0,
    "pps_attack": 0,
    "total_pps": 0
  }, {
    "period_start": 1521650668727,
    "bps_in": 1,
    "bps_attack": 0,
    "total_bps": 1,
    "pps_in": 2,
    "pps_attack": 0,
    "total_pps": 2
  }],
  "tableHeader": ["period_start", "bps_in", "bps_attack", "total_bps", "pps_in", "pps_attack", "total_pps"],
  "tableBody": [
    ["2018-03-22 00:34:28", "0", "0", "0", "0", "0", "0"],
    ["2018-03-22 00:39:28", "0", "0", "0", "0", "0", "0"],
    ["2018-03-22 00:44:28", "1", "0", "1", "2", "0", "2"]
  ]
}
`

var DailyReportResponse = antiddos.DailyReportResponse{
	Data: []struct {
		PeriodStart int `json:"period_start,"`
		BpsIn       int `json:"bps_in,"`
		BpsAttack   int `json:"bps_attack,"`
		TotalBps    int `json:"total_bps,"`
		PpsIn       int `json:"pps_in,"`
		PpsAttack   int `json:"pps_attack,"`
		TotalPps    int `json:"total_pps,"`
	}{
		{
			PeriodStart: 1521650068727,
			BpsIn:       0,
			BpsAttack:   0,
			TotalBps:    0,
			PpsIn:       0,
			PpsAttack:   0,
			TotalPps:    0,
		},
		{
			PeriodStart: 1521650368727,
			BpsIn:       0,
			BpsAttack:   0,
			TotalBps:    0,
			PpsIn:       0,
			PpsAttack:   0,
			TotalPps:    0,
		},
		{
			PeriodStart: 1521650668727,
			BpsIn:       1,
			BpsAttack:   0,
			TotalBps:    1,
			PpsIn:       2,
			PpsAttack:   0,
			TotalPps:    2,
		},
	},
}

func HandleDailyReportSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/antiddos/82abaa86-8518-47db-8d63-ddf152824635/daily", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DailyReportOutput)
	})
}

var GetTaskOutput = `
{
   "task_status": "running",
   "task_msg": "ABC"
}
`

var GetTaskResponse = antiddos.GetTaskResponse{
	TaskStatus: "running",
	TaskMsg:    "ABC",
}

func HandleGetTaskSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/query_task_status", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetTaskOutput)
	})
}
