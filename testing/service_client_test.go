package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestServiceURL(t *testing.T) {
	c := &golangsdk.ServiceClient{Endpoint: "http://123.45.67.8/"}
	expected := "http://123.45.67.8/more/parts/here"
	actual := c.ServiceURL("more", "parts", "here")
	th.CheckEquals(t, expected, actual)
}

func TestMoreHeaders(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	c := new(golangsdk.ServiceClient)
	c.MoreHeaders = map[string]string{
		"custom": "header",
	}
	c.ProviderClient = new(golangsdk.ProviderClient)
	resp, err := c.Get(fmt.Sprintf("%s/route", th.Endpoint()), nil, nil)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, resp.Request.Header.Get("custom"), "header")
}
