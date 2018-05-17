package testing

import (
	"net/http"
	"testing"

	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func mockListTagsResponse(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"tags":["tag1","tag2"]}`))
	})
}

func mockPutTagsResponse(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestBody(t, r, `{"tags":["tag1","tag2"]}`)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"tags":["tag1","tag2"]}`))
	})
}

func mockCleanTagsResponse(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}

func mockCheckTagResponse(t *testing.T, id, tag string) {
	th.Mux.HandleFunc("/servers/"+id+"/tags/"+tag, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}

func mockAddTagResponse(t *testing.T, id, tag string) {
	th.Mux.HandleFunc("/servers/"+id+"/tags/"+tag, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}

func mockDeleteTagResponse(t *testing.T, id, tag string) {
	th.Mux.HandleFunc("/servers/"+id+"/tags/"+tag, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}
