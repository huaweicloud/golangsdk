package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/identity/v3/services"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

// ListOutput provides a single page of Service results.
const ListOutput = `
{ 
    "services": [ 
        { 
            "name": "service-one", 
            "links": { 
                "self": "https://iamcore_links.com/v3/services/053d21d488d1463c818132d9d08fb617" 
            }, 
            "enabled": true, 
            "type": "compute", 
            "id": "053d21d488d1463c818132d9d08fb617", 
            "description": "Service One",
            "extra": {
                "name": "service-one",
                "description": "Service One"
            }
        }, 
        { 
            "name": "service-two", 
            "links": { 
                "self": "https://iamcore_links.com/v3/services/c2474183dca7453bbd73123a0b78feae" 
            }, 
            "enabled": true, 
            "type": "compute", 
            "id": "c2474183dca7453bbd73123a0b78feae", 
            "description": "Service Two",
            "extra": {
                "name": "service-two",
                "description": "Service Two"
            }
        }
    ], 
    "links": { 
        "self": "https://iamcore_links.com/v3/services?type=compute", 
        "previous": null, 
        "next": null 
    } 
}

`

// GetOutput provides a Get result.
const GetOutput = `
{ 
    "service": { 
		"name": "service-two", 
		"links": { 
			"self": "https://iamcore_links.com/v3/services/c2474183dca7453bbd73123a0b78feae" 
		}, 
		"enabled": true, 
		"type": "compute", 
		"id": "c2474183dca7453bbd73123a0b78feae", 
		"description": "Service Two",
		"extra": {
			"name": "service-two",
			"description": "Service Two"
		}
    } 
}

`

// CreateRequest provides the input to a Create request.
const CreateRequest = `
{
    "service": {
        "description": "Service Two",
        "name": "service-two",
        "type": "compute",
        "email": "service@example.com"
    }
}
`

// UpdateRequest provides the input to as Update request.
const UpdateRequest = `
{
    "service": {
        "type": "compute2",
        "description": "Service Two Updated"
    }
}
`

// UpdateOutput provides an update result.
const UpdateOutput = `
{
    "service": {
        "name": "service-two", 
         "links": { 
             "self": "https://iamcore_links.com/v3/services/c2474183dca7453bbd73123a0b78feae" 
         }, 
         "enabled": true, 
         "type": "compute2", 
         "id": "c2474183dca7453bbd73123a0b78feae", 
         "description": "Service Two Updated",
		"extra": {
			"name": "service-two",
			"description": "Service Two Updated"
		}
    }
}
`

// FirstService is the first service in the List request.
var FirstService = services.Service{
	ID: "053d21d488d1463c818132d9d08fb617",
	Links: map[string]interface{}{
		"self": "https://iamcore_links.com/v3/services/053d21d488d1463c818132d9d08fb617",
	},
	Type:        "compute",
	Enabled:     true,
	Name:        "service-one",
	Description: "Service One",
	Extra: map[string]interface{}{
		"name":        "service-one",
		"description": "Service One",
	},
}

// SecondService is the second service in the List request.
var SecondService = services.Service{
	ID: "c2474183dca7453bbd73123a0b78feae",
	Links: map[string]interface{}{
		"self": "https://iamcore_links.com/v3/services/c2474183dca7453bbd73123a0b78feae",
	},
	Type:        "compute",
	Enabled:     true,
	Name:        "service-two",
	Description: "Service Two",
	Extra: map[string]interface{}{
		"name":        "service-two",
		"description": "Service Two",
	},
}

// SecondServiceUpdated is the SecondService should look after an Update.
var SecondServiceUpdated = services.Service{
	ID: "c2474183dca7453bbd73123a0b78feae",
	Links: map[string]interface{}{
		"self": "https://iamcore_links.com/v3/services/c2474183dca7453bbd73123a0b78feae",
	},
	Type:        "compute2",
	Enabled:     true,
	Name:        "service-two",
	Description: "Service Two Updated",
	Extra: map[string]interface{}{
		"name":        "service-two",
		"description": "Service Two Updated",
	},
}

// ExpectedServicesSlice is the slice of services to be returned from ListOutput.
var ExpectedServicesSlice = []services.Service{FirstService, SecondService}

// HandleListServicesSuccessfully creates an HTTP handler at `/services` on the
// test handler mux that responds with a list of two services.
func HandleListServicesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListOutput)
	})
}

// HandleGetServiceSuccessfully creates an HTTP handler at `/services` on the
// test handler mux that responds with a single service.
func HandleGetServiceSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/services/c2474183dca7453bbd73123a0b78feae", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetOutput)
	})
}

// HandleCreateServiceSuccessfully creates an HTTP handler at `/services` on the
// test handler mux that tests service creation.
func HandleCreateServiceSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, GetOutput)
	})
}

// HandleUpdateServiceSuccessfully creates an HTTP handler at `/services` on the
// test handler mux that tests service update.
func HandleUpdateServiceSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/services/c2474183dca7453bbd73123a0b78feae", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, UpdateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateOutput)
	})
}
