package postpagination

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/huaweicloud/golangsdk"
)

var (
	// ErrPageNotAvailable is returned from a Pager when a next or previous page is requested, but does not exist.
	ErrPageNotAvailable = errors.New("The requested page does not exist.")
)

// Page must be satisfied by the result type of any resource collection.
// It allows clients to interact with the resource uniformly, regardless of whether or not or how it's paginated.
// Generally, rather than implementing this interface directly, implementors should embed one of the concrete PageBase structs,
// instead.
// Depending on the pagination strategy of a particular resource, there may be an additional subinterface that the result type
// will need to implement.
type Page interface {
	// NextPageURL generates the URL for the page of data that follows this collection.
	// Return "" if no such page exists.
	NextPageURL() (string, error)

	// NextPageRequest generate the body for the page for data that follows this collection.
	// Return nil if no such page exists.
	NextPageRequest() (interface{}, error)

	// IsEmpty returns true if this Page has no items in it.
	IsEmpty() (bool, error)

	// GetBody returns the Page Body. This is used in the `AllPages` method.
	GetBody() interface{}
}

// Pager knows how to advance through a specific resource collection, one page at a time.
type Pager struct {
	client *golangsdk.ServiceClient

	method string

	initialURL string

	initialRequest interface{}

	createPage func(r PageResult) Page

	Err error

	// Headers supplies additional HTTP headers to populate on each paged request.
	Headers map[string]string
}

// NewPager constructs a manually-configured pager.
// Supply the URL for the first page, a function that requests a specific page given a URL, and a function that counts a page.
func NewPager(client *golangsdk.ServiceClient, method string, initialURL string, initialRequest interface{}, createPage func(r PageResult) Page) Pager {
	return Pager{
		client:         client,
		method:         method,
		initialURL:     initialURL,
		initialRequest: initialRequest,
		createPage:     createPage,
	}
}

// WithPageCreator returns a new Pager that substitutes a different page creation function. This is
// useful for overriding List functions in delegation.
func (p Pager) WithPageCreator(createPage func(r PageResult) Page) Pager {
	return Pager{
		client:         p.client,
		method:         p.method,
		initialURL:     p.initialURL,
		initialRequest: p.initialRequest,
		createPage:     createPage,
	}
}

func (p Pager) fetchNextPage(method string, url string, request interface{}) (Page, error) {
	resp, err := Request(p.client, method, request, p.Headers, url)
	if err != nil {
		return nil, err
	}

	remembered, err := PageResultFrom(resp)
	if err != nil {
		return nil, err
	}

	return p.createPage(remembered), nil
}

// EachPage iterates over each page returned by a Pager, yielding one at a time to a handler function.
// Return "false" from the handler to prematurely stop iterating.
func (p Pager) EachPage(handler func(Page) (bool, error)) error {
	if p.Err != nil {
		return p.Err
	}
	currentURL := p.initialURL
	currentRequest := p.initialRequest
	for {
		currentPage, err := p.fetchNextPage(p.method, currentURL, currentRequest)
		if err != nil {
			return err
		}

		empty, err := currentPage.IsEmpty()
		if err != nil {
			return err
		}
		if empty {
			return nil
		}

		ok, err := handler(currentPage)
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}

		currentURL, err = currentPage.NextPageURL()
		if err != nil {
			return err
		}
		if currentURL == "" {
			return nil
		}

		currentRequest, err = currentPage.NextPageRequest()
		if err != nil {
			return err
		}
		if currentRequest == nil {
			return nil
		}
	}
}

// AllPages returns all the pages from a `List` operation in a single page,
// allowing the user to retrieve all the pages at once.
func (p Pager) AllPages() (Page, error) {
	// pagesSlice holds all the pages until they get converted into as Page Body.
	var pagesSlice []interface{}
	// body will contain the final concatenated Page body.
	var body reflect.Value

	// Grab a test page to ascertain the page body type.
	testPage, err := p.fetchNextPage(p.method, p.initialURL, p.initialRequest)
	if err != nil {
		return nil, err
	}
	// Store the page type so we can use reflection to create a new mega-page of
	// that type.
	pageType := reflect.TypeOf(testPage)

	// if it's a single page, just return the testPage (first page)
	if _, found := pageType.FieldByName("SinglePageBase"); found {
		return testPage, nil
	}

	// Switch on the page body type. Recognized types are `map[string]interface{}`,
	// `[]byte`, and `[]interface{}`.
	switch pb := testPage.GetBody().(type) {
	case map[string]interface{}:
		// key is the map key for the page body if the body type is `map[string]interface{}`.
		var mapPagesSlice = make(map[string]interface{})
		// Iterate over the pages to concatenate the bodies.
		err = p.EachPage(func(page Page) (bool, error) {
			b := page.GetBody().(map[string]interface{})
			for k, v := range b {
				// If it's a linked page, we don't want the `links`, we want the other one.
				if !strings.HasSuffix(k, "links") && !strings.HasSuffix(k, "next_marker") && !strings.HasSuffix(k, "truncated") {
					// check the field's type. we only want []interface{} (which is really []map[string]interface{})
					switch vt := v.(type) {
					case []interface{}:
						slice := mapPagesSlice[k]
						if slice == nil {
							var temp []interface{}
							slice = temp
						}
						slice = append(slice.([]interface{}), vt...)
						mapPagesSlice[k] = slice
					}
				}
			}
			return true, nil
		})
		if err != nil {
			return nil, err
		}
		var dummyKey = ""
		var dummyValue []interface{}
		body = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(dummyKey), reflect.TypeOf(dummyValue)))
		for mapKey, mapValue := range mapPagesSlice {
			body.SetMapIndex(reflect.ValueOf(mapKey), reflect.ValueOf(mapValue))
		}
	case []byte:
		// Iterate over the pages to concatenate the bodies.
		err = p.EachPage(func(page Page) (bool, error) {
			b := page.GetBody().([]byte)
			pagesSlice = append(pagesSlice, b)
			// seperate pages with a comma
			pagesSlice = append(pagesSlice, []byte{10})
			return true, nil
		})
		if err != nil {
			return nil, err
		}
		if len(pagesSlice) > 0 {
			// Remove the trailing comma.
			pagesSlice = pagesSlice[:len(pagesSlice)-1]
		}
		var b []byte
		// Combine the slice of slices in to a single slice.
		for _, slice := range pagesSlice {
			b = append(b, slice.([]byte)...)
		}
		// Set body to value of type `bytes`.
		body = reflect.New(reflect.TypeOf(b)).Elem()
		body.SetBytes(b)
	case []interface{}:
		// Iterate over the pages to concatenate the bodies.
		err = p.EachPage(func(page Page) (bool, error) {
			b := page.GetBody().([]interface{})
			pagesSlice = append(pagesSlice, b...)
			return true, nil
		})
		if err != nil {
			return nil, err
		}
		// Set body to value of type `[]interface{}`
		body = reflect.MakeSlice(reflect.TypeOf(pagesSlice), len(pagesSlice), len(pagesSlice))
		for i, s := range pagesSlice {
			body.Index(i).Set(reflect.ValueOf(s))
		}
	default:
		err := golangsdk.ErrUnexpectedType{}
		err.Expected = "map[string]interface{}/[]byte/[]interface{}"
		err.Actual = fmt.Sprintf("%T", pb)
		return nil, err
	}

	// Each `Extract*` function is expecting a specific type of page coming back,
	// otherwise the type assertion in those functions will fail. pageType is needed
	// to create a type in this method that has the same type that the `Extract*`
	// function is expecting and set the Body of that object to the concatenated
	// pages.
	page := reflect.New(pageType)
	// Set the page body to be the concatenated pages.
	page.Elem().FieldByName("Body").Set(body)
	// Set any additional headers that were pass along. The `objectstorage` pacakge,
	// for example, passes a Content-Type header.
	h := make(http.Header)
	for k, v := range p.Headers {
		h.Add(k, v)
	}
	page.Elem().FieldByName("Header").Set(reflect.ValueOf(h))
	// Type assert the page to a Page interface so that the type assertion in the
	// `Extract*` methods will work.
	return page.Elem().Interface().(Page), err
}
