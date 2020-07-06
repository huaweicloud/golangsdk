package testing

import (
	"net/http"
	"testing"

	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

// HandleGetAccountSuccessfully creates an HTTP handler at `/` on the test handler mux that
// responds with a `Get` response.
func HandleGetAccountSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "HEAD")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("X-Account-Container-Count", "2")
		w.Header().Set("X-Account-Object-Count", "5")
		w.Header().Set("X-Account-Meta-Quota-Bytes", "42")
		w.Header().Set("X-Account-Bytes-Used", "14")
		w.Header().Set("X-Account-Meta-Subject", "books")
		w.Header().Set("Date", "Fri, 17 Jan 2014 16:09:56 UTC")

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleGetAccountNoQuotaSuccessfully creates an HTTP handler at `/` on the
// test handler mux that responds with a `Get` response.
func HandleGetAccountNoQuotaSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "HEAD")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("X-Account-Container-Count", "2")
		w.Header().Set("X-Account-Object-Count", "5")
		w.Header().Set("X-Account-Bytes-Used", "14")
		w.Header().Set("X-Account-Meta-Subject", "books")
		w.Header().Set("Date", "Fri, 17 Jan 2014 16:09:56 UTC")

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleUpdateAccountSuccessfully creates an HTTP handler at `/` on the test handler mux that
// responds with a `Update` response.
func HandleUpdateAccountSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "X-Account-Meta-Golangsdk-Test", "accounts")

		w.Header().Set("Date", "Fri, 17 Jan 2014 16:09:56 UTC")
		w.WriteHeader(http.StatusNoContent)
	})
}
