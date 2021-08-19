package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/fgs/v2/dependencies"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedListPageOneResponse = `{
	"count": 4,
	"dependencies": [
	  {
		"etag": "565a745c27a1f38d5e8af01bdab83578",
		"file_name": "esdk-obs-nodejs-3.x.zip",
		"id": "b1f55dc0-b3aa-4952-b5b9-883d0af27fa5",
		"link": "https://functionstorage-cn-north-4.obs.cn-north-4.myhuaweicloud.com/depends/public/b1f55dc0-b3aa-4952-b5b9-883d0af27fa5.zip",
		"name": "esdk-obs-nodejs-3.x",
		"owner": "public",
		"runtime": "Node.js6.10",
		"size": 368391
	  },
	  {
		"etag": "8d3337ba9f836cbc5014f68e3ee3b901",
		"file_name": "moderation_sdk_nodejs.zip",
		"id": "dabf603a-dcd7-42a8-8014-20e55c92e861",
		"link": "https://functionstorage-cn-north-4.obs.cn-north-4.myhuaweicloud.com/depends/public/dabf603a-dcd7-42a8-8014-20e55c92e861.zip",
		"name": "moderation_sdk_nodejs",
		"owner": "public",
		"runtime": "Node.js6.10",
		"size": 128876
	  }
	],
	"next_marker": 2
}`

	expectedListPageTwoResponse = `{
	"count": 4,
	"dependencies": [
	  {
		"etag": "6fb79329595ef31106e01fab44e45401",
		"file_name": "moment-timezone.zip",
		"id": "b7022a6d-e6cd-4532-8970-313373922591",
		"link": "https://functionstorage-cn-north-4.obs.cn-north-4.myhuaweicloud.com/depends/public/b7022a6d-e6cd-4532-8970-313373922591.zip",
		"name": "moment-timezone",
		"owner": "public",
		"runtime": "Node.js6.10",
		"size": 1359773
	  },
	  {
		"etag": "2635ffb1e8198dd0dcaf6f3390a1c14f-3",
		"id": "3004706f-d26a-4c87-96ae-0f1ed715ca3e",
		"link": "https://testlfk.obs.cn-north-4.myhuaweicloud.com/cloud-ocr-sdk-nodejs-1.x.zip",
		"name": "cloud-ocr-sdk-nodejs-1.x",
		"owner": "public",
		"runtime": "Node.js6.10",
		"size": 21271972
	  }
	],
	"next_marker": 4
}`
)

var (
	listOpts = dependencies.ListOpts{
		Limit: "2",
	}

	expectedListResponseData = []dependencies.Dependency{
		{
			Etag:     "565a745c27a1f38d5e8af01bdab83578",
			FileName: "esdk-obs-nodejs-3.x.zip",
			ID:       "b1f55dc0-b3aa-4952-b5b9-883d0af27fa5",
			Link:     "https://functionstorage-cn-north-4.obs.cn-north-4.myhuaweicloud.com/depends/public/b1f55dc0-b3aa-4952-b5b9-883d0af27fa5.zip",
			Name:     "esdk-obs-nodejs-3.x",
			Owner:    "public",
			Runtime:  "Node.js6.10",
			Size:     368391,
		},
		{
			Etag:     "8d3337ba9f836cbc5014f68e3ee3b901",
			FileName: "moderation_sdk_nodejs.zip",
			ID:       "dabf603a-dcd7-42a8-8014-20e55c92e861",
			Link:     "https://functionstorage-cn-north-4.obs.cn-north-4.myhuaweicloud.com/depends/public/dabf603a-dcd7-42a8-8014-20e55c92e861.zip",
			Name:     "moderation_sdk_nodejs",
			Owner:    "public",
			Runtime:  "Node.js6.10",
			Size:     128876,
		},
		{
			Etag:     "6fb79329595ef31106e01fab44e45401",
			FileName: "moment-timezone.zip",
			ID:       "b7022a6d-e6cd-4532-8970-313373922591",
			Link:     "https://functionstorage-cn-north-4.obs.cn-north-4.myhuaweicloud.com/depends/public/b7022a6d-e6cd-4532-8970-313373922591.zip",
			Name:     "moment-timezone",
			Owner:    "public",
			Runtime:  "Node.js6.10",
			Size:     1359773,
		},
		{
			Etag:    "2635ffb1e8198dd0dcaf6f3390a1c14f-3",
			ID:      "3004706f-d26a-4c87-96ae-0f1ed715ca3e",
			Link:    "https://testlfk.obs.cn-north-4.myhuaweicloud.com/cloud-ocr-sdk-nodejs-1.x.zip",
			Name:    "cloud-ocr-sdk-nodejs-1.x",
			Owner:   "public",
			Runtime: "Node.js6.10",
			Size:    21271972,
		},
	}
)

func handleV2DependenciesList(t *testing.T) {
	th.Mux.HandleFunc("/fgs/dependencies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, expectedListPageOneResponse)
		case "2":
			fmt.Fprintf(w, expectedListPageTwoResponse)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}
