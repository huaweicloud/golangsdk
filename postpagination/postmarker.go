package postpagination

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/huaweicloud/golangsdk"
)

type PageBuilder interface {
	ToServerPageMap() (map[string]interface{}, error)
}

// MarkerPage is a stricter Page interface that describes additional functionality required for use with NewMarkerPager.
// For convenience, embed the MarkedPageBase struct.
type PostMarkerPage interface {
	Page
}

// MarkerPageBase is a page in a collection that's paginated by "limit" and "marker" query parameters.
type PostMarkerPageBase struct {
	PageResult

	ListFieldName string

	Request map[string]interface{}

	// Owner is a reference to the embedding struct.
	Owner PostMarkerPage
}

// NextPageURL generates the URL for the page of results after this one.
func (current PostMarkerPageBase) NextPageURL() (string, error) {
	return current.URL.String(), nil
}

// NextPageRequest generates the body for the page of results after this one.
func (current PostMarkerPageBase) NextPageRequest() (interface{}, error) {
	var nextMarker string
	var truncated string
	var nextRequest map[string]interface{} = nil

	submap, ok := current.Body.(map[string]interface{})
	if !ok {
		err := golangsdk.ErrUnexpectedType{}
		err.Expected = "map[string]interface{}"
		err.Actual = fmt.Sprintf("%v", reflect.TypeOf(current.Body))
		return nil, err
	}

	truncated, ok = submap["truncated"].(string)
	if !ok || truncated == "false" {
		return nil, nil
	}

	nextMarker, ok = submap["next_marker"].(string)
	if !ok {
		return nil, nil
	}

	rawBody, err := json.Marshal(current.Request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawBody, &nextRequest)
	if err != nil {
		return nil, err
	}

	nextRequest["marker"] = nextMarker
	return nextRequest, nil
}

// IsEmpty satisifies the IsEmpty method of the Page interface
func (current PostMarkerPageBase) IsEmpty() (bool, error) {
	if b, ok := current.Body.([]interface{}); ok {
		return len(b) == 0, nil
	}
	err := golangsdk.ErrUnexpectedType{}
	err.Expected = "[]interface{}"
	err.Actual = fmt.Sprintf("%v", reflect.TypeOf(current.Body))
	return true, err
}

// GetBody returns the linked page's body. This method is needed to satisfy the
// Page interface.
func (current PostMarkerPageBase) GetBody() interface{} {
	return current.Body
}
