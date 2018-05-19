package flavors

import (
	"encoding/json"
	"strconv"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type commonResult struct {
	golangsdk.Result
}

// GetResult is the response of a Get operations. Call its Extract method to
// interpret it as a Flavor.
type GetResult struct {
	commonResult
}

// Extract provides access to the individual Flavor returned by the Get and
// Create functions.
func (r commonResult) Extract() (*Flavor, error) {
	var s struct {
		Flavor *Flavor `json:"flavor"`
	}
	err := r.ExtractInto(&s)
	return s.Flavor, err
}

// Flavor represent (virtual) hardware configurations for server resources
// in a region.
type Flavor struct {
	// ID is the flavor's unique ID.
	ID string `json:"id"`

	// Disk is the amount of root disk, measured in GB.
	Disk int `json:"disk"`

	// RAM is the amount of memory, measured in MB.
	RAM int `json:"ram"`

	// Name is the name of the flavor.
	Name string `json:"name"`

	// RxTxFactor describes bandwidth alterations of the flavor.
	RxTxFactor float64 `json:"rxtx_factor"`

	// Swap is the amount of swap space, measured in MB.
	Swap int `json:"swap"`

	// VCPUs indicates how many (virtual) CPUs are available for this flavor.
	VCPUs int `json:"vcpus"`

	// IsPublic indicates whether the flavor is public.
	IsPublic bool `json:"os-flavor-access:is_public"`

	// Ephemeral is the amount of ephemeral disk space, measured in GB.
	Ephemeral int `json:"OS-FLV-EXT-DATA:ephemeral"`

	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (r *Flavor) UnmarshalJSON(b []byte) error {
	type tmp Flavor
	var s struct {
		tmp
		Swap interface{} `json:"swap"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = Flavor(s.tmp)

	switch t := s.Swap.(type) {
	case float64:
		r.Swap = int(t)
	case string:
		switch t {
		case "":
			r.Swap = 0
		default:
			swap, err := strconv.ParseFloat(t, 64)
			if err != nil {
				return err
			}
			r.Swap = int(swap)
		}
	}

	return nil
}

// FlavorPage contains a single page of all flavors from a ListDetails call.
type FlavorPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines if a FlavorPage contains any results.
func (page FlavorPage) IsEmpty() (bool, error) {
	flavors, err := ExtractFlavors(page)
	return len(flavors) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (page FlavorPage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"flavors_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

// ExtractFlavors provides access to the list of flavors in a page acquired
// from the ListDetail operation.
func ExtractFlavors(r pagination.Page) ([]Flavor, error) {
	var s struct {
		Flavors []Flavor `json:"flavors"`
	}
	err := (r.(FlavorPage)).ExtractInto(&s)
	return s.Flavors, err
}

// Extract interprets any extraSpecsResult as ExtraSpecs, if possible.
func (r extraSpecsResult) Extract() (map[string]string, error) {
	var s struct {
		ExtraSpecs map[string]string `json:"extra_specs"`
	}
	err := r.ExtractInto(&s)
	return s.ExtraSpecs, err
}

// extraSpecsResult contains the result of a call for (potentially) multiple
// key-value pairs. Call its Extract method to interpret it as a
// map[string]interface.
type extraSpecsResult struct {
	golangsdk.Result
}

// ListExtraSpecsResult contains the result of a Get operation. Call its Extract
// method to interpret it as a map[string]interface.
type ListExtraSpecsResult struct {
	extraSpecsResult
}
