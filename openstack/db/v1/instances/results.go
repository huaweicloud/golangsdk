package instances

import (
	"encoding/json"
	"time"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/db/v1/datastores"
	"github.com/huaweicloud/golangsdk/openstack/db/v1/users"
	"github.com/huaweicloud/golangsdk/pagination"
)

// Volume represents information about an attached volume for a database instance.
type Volume struct {
	// The size in GB of the volume
	Size int

	Used float64
}

// Flavor represents (virtual) hardware configurations for server resources in a region.
type Flavor struct {
	// The flavor's unique identifier.
	ID string
	// Links to access the flavor.
	Links []golangsdk.Link
}

// Fault describes the fault reason in more detail when a database instance has errored
type Fault struct {
	// Indicates the time when the fault occured
	Created time.Time `json:"-"`

	// A message describing the fault reason
	Message string

	// More details about the fault, for example a stack trace. Only filled
	// in for admin users.
	Details string
}

func (r *Fault) UnmarshalJSON(b []byte) error {
	type tmp Fault
	var s struct {
		tmp
		Created golangsdk.JSONRFC3339NoZ `json:"created"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Fault(s.tmp)

	r.Created = time.Time(s.Created)

	return nil
}

// Instance represents a remote MySQL instance.
type Instance struct {
	// Indicates the datetime that the instance was created
	Created time.Time `json:"-"`

	// Indicates the most recent datetime that the instance was updated.
	Updated time.Time `json:"-"`

	// Indicates the hardware flavor the instance uses.
	Flavor Flavor

	// A DNS-resolvable hostname associated with the database instance (rather
	// than an IPv4 address). Since the hostname always resolves to the correct
	// IP address of the database instance, this relieves the user from the task
	// of maintaining the mapping. Note that although the IP address may likely
	// change on resizing, migrating, and so forth, the hostname always resolves
	// to the correct database instance.
	Hostname string

	// The IP addresses associated with the database instance
	// Is empty if the instance has a hostname
	IP []string

	// Indicates the unique identifier for the instance resource.
	ID string

	// Exposes various links that reference the instance resource.
	Links []golangsdk.Link

	// The human-readable name of the instance.
	Name string

	// The build status of the instance.
	Status string

	// Fault information (only available when the instance has errored)
	Fault *Fault

	// Information about the attached volume of the instance.
	Volume Volume

	// Indicates how the instance stores data.
	Datastore datastores.DatastorePartial
}

func (r *Instance) UnmarshalJSON(b []byte) error {
	type tmp Instance
	var s struct {
		tmp
		Created golangsdk.JSONRFC3339NoZ `json:"created"`
		Updated golangsdk.JSONRFC3339NoZ `json:"updated"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Instance(s.tmp)

	r.Created = time.Time(s.Created)
	r.Updated = time.Time(s.Updated)

	return nil
}

type commonResult struct {
	golangsdk.Result
}

// CreateResult represents the result of a Create operation.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a Get operation.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
type DeleteResult struct {
	golangsdk.ErrResult
}

// ConfigurationResult represents the result of a AttachConfigurationGroup/DetachConfigurationGroup operation.
type ConfigurationResult struct {
	golangsdk.ErrResult
}

// Extract will extract an Instance from various result structs.
func (r commonResult) Extract() (*Instance, error) {
	var s struct {
		Instance *Instance `json:"instance"`
	}
	err := r.ExtractInto(&s)
	return s.Instance, err
}

// InstancePage represents a single page of a paginated instance collection.
type InstancePage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks to see whether the collection is empty.
func (page InstancePage) IsEmpty() (bool, error) {
	instances, err := ExtractInstances(page)
	return len(instances) == 0, err
}

// NextPageURL will retrieve the next page URL.
func (page InstancePage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"instances_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

// ExtractInstances will convert a generic pagination struct into a more
// relevant slice of Instance structs.
func ExtractInstances(r pagination.Page) ([]Instance, error) {
	var s struct {
		Instances []Instance `json:"instances"`
	}
	err := (r.(InstancePage)).ExtractInto(&s)
	return s.Instances, err
}

// EnableRootUserResult represents the result of an operation to enable the root user.
type EnableRootUserResult struct {
	golangsdk.Result
}

// Extract will extract root user information from a UserRootResult.
func (r EnableRootUserResult) Extract() (*users.User, error) {
	var s struct {
		User *users.User `json:"user"`
	}
	err := r.ExtractInto(&s)
	return s.User, err
}

// ActionResult represents the result of action requests, such as: restarting
// an instance service, resizing its memory allocation, and resizing its
// attached volume size.
type ActionResult struct {
	golangsdk.ErrResult
}

// IsRootEnabledResult is the result of a call to IsRootEnabled. To see if
// root is enabled, call the type's Extract method.
type IsRootEnabledResult struct {
	golangsdk.Result
}

// Extract is used to extract the data from a IsRootEnabledResult.
func (r IsRootEnabledResult) Extract() (bool, error) {
	return r.Body.(map[string]interface{})["rootEnabled"] == true, r.Err
}
