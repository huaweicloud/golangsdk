package alarmreminding

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type WarnAlertResult struct {
	commonResult
}

func (r WarnAlertResult) Extract() (*WarnAlertResponse, error) {
	var response WarnAlertResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type WarnAlertResponse struct {
	// Alarm configuration
	WarnConfig struct {
		// DDoS attacks
		AntiDDoS bool `json:"antiDDoS,"`

		// Brute force cracking (system logins, FTP, and DB)
		BruceForce bool `json:"bruce_force,omitempty"`

		// Alarms about remote logins
		RemoteLogin bool `json:"remote_login,omitempty"`

		// Weak passwords (system and database)
		WeakPassword bool `json:"weak_password,omitempty"`

		// Overly high rights of a database process
		HighPrivilege bool `json:"high_privilege,omitempty"`

		// Web page backdoors
		BackDoors bool `json:"back_doors,omitempty"`

		// Reserved
		Waf bool `json:"waf,omitempty"`

		// Possible values: 0: indicates that alarms are sent once a day. 1: indicates that alarms are sent once every half hour. This parameter is mandatory for the Host Intrusion Detection (HID) service.
		SendFrequency int `json:"send_frequency,omitempty"`
	} `json:"warn_config,"`

	// ID of an alarm group
	TopicUrn string `json:"topic_urn,"`

	// Description of an alarm group
	DisplayName string `json:"display_name,"`
}
