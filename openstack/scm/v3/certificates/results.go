package certificates

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// This struct defines the authentication information of the domain.
// API that will be used to Obtain Certificate Information
type Authentification struct {
	RecordName  string `json:"record_name"`
	RecordType  string `json:"record_type"`
	RecordValue string `json:"record_value"`
	Domain      string `json:"domain"`
}

// The struct defines the information about the imported certificate.
// This struct is used by the API for obtain certificate information and importing.
type Certificate struct {
	// the importing API only uses certificate_id.
	CertificateId     string             `json:"certificate_id"`
	Id                string             `json:"id"`
	Name              string             `json:"name,omitempty" required:"true"`
	Certificate       string             `json:"certificate" required:"true"`
	CertificateChain  string             `json:"certificate_chain" required:"true"`
	PrivateKey        string             `json:"private_key" required:"true"`
	Status            string             `json:"status"`
	OrderId           string             `json:"order_id"`
	CertificateType   string             `json:"type"`
	Brand             string             `json:"brand"`
	PushSupport       string             `json:"push_support"`
	RevokeReason      string             `json:"revoke_reason"`
	SignatureAlgrithm string             `json:"signature_algrithm"`
	IssueTime         string             `json:"issue_time"`
	NotBefore         string             `json:"not_before"`
	NotAfter          string             `json:"not_after"`
	ValidityPeriod    int8               `json:"validity_period,omitempty"`
	ValidationMethod  string             `json:"validation_method"`
	DomainType        string             `json:"domain_type"`
	Domain            string             `json:"domain"`
	Sans              string             `json:"sans"`
	DomainCount       int8               `json:"domain_count,omitempty"`
	WildcardCount     int8               `json:"wildcard_count,omitempty"`
	Authentifications []Authentification `json:"authentification,omitempty"`
}

// CertificatePage is the page returned by a pager when traversing over a
// collection of certificates.
type CertificatePage struct {
	pagination.SinglePageBase
}

type commonResult struct {
	golangsdk.Result
}

type ImportResult struct {
	commonResult
}

type PushResult struct {
	golangsdk.ErrResult
}

func (r commonResult) Extract() (*Certificate, error) {
	var s Certificate
	err := r.ExtractInto(&s)
	return &s, err
}

type GetResult struct {
	commonResult
}

type ExportResult struct {
	commonResult
}

type DeleteResult struct {
	golangsdk.ErrResult
}
