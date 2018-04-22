package openstack

import (
	"errors"
	"net/url"
	"strings"

	"github.com/huaweicloud/golangsdk"
	tokens2 "github.com/huaweicloud/golangsdk/openstack/identity/v2/tokens"
	tokens3 "github.com/huaweicloud/golangsdk/openstack/identity/v3/tokens"
)

// service have same endpoint address in different location, refer to https://developer.huaweicloud.com/endpoint
var allRegionInOneEndpoint = map[string]struct{}{
	"cdn": struct{}{},
	"dns": struct{}{},
}

/*
V2EndpointURL discovers the endpoint URL for a specific service from a
ServiceCatalog acquired during the v2 identity service.

The specified EndpointOpts are used to identify a unique, unambiguous endpoint
to return. It's an error both when multiple endpoints match the provided
criteria and when none do. The minimum that can be specified is a Type, but you
will also often need to specify a Name and/or a Region depending on what's
available on your OpenStack deployment.
*/
func V2EndpointURL(catalog *tokens2.ServiceCatalog, opts golangsdk.EndpointOpts) (string, error) {
	// Extract Endpoints from the catalog entries that match the requested Type, Name if provided, and Region if provided.
	var endpoints = make([]tokens2.Endpoint, 0, 1)
	for _, entry := range catalog.Entries {
		if (opts.Type == "" || entry.Type == opts.Type) && (opts.Name == "" || entry.Name == opts.Name) {
			for _, endpoint := range entry.Endpoints {
				if opts.Region == "" || endpoint.Region == opts.Region {
					endpoints = append(endpoints, endpoint)
				}
			}
		}
	}

	// Report an error if the options were ambiguous.
	if opts.Type != "" {
		if len(endpoints) > 1 {
			err := &ErrMultipleMatchingEndpointsV2{}
			err.Endpoints = endpoints
			return "", err
		} else if len(endpoints) < 1 {
			return buildUrlIfNotFoundV2(catalog, opts)
		}
	} else {
		if len(endpoints) < 1 {
			return "", &golangsdk.ErrEndpointNotFound{}
		}
	}

	// Extract the appropriate URL from the matching Endpoint.
	for _, endpoint := range endpoints {
		switch opts.Availability {
		case golangsdk.AvailabilityPublic:
			return golangsdk.NormalizeURL(endpoint.PublicURL), nil
		case golangsdk.AvailabilityInternal:
			return golangsdk.NormalizeURL(endpoint.InternalURL), nil
		case golangsdk.AvailabilityAdmin:
			return golangsdk.NormalizeURL(endpoint.AdminURL), nil
		default:
			err := &ErrInvalidAvailabilityProvided{}
			err.Argument = "Availability"
			err.Value = opts.Availability
			return "", err
		}
	}

	// Report an error if there were no matching endpoints.
	err := &golangsdk.ErrEndpointNotFound{}
	return "", err
}

/*
V3EndpointURL discovers the endpoint URL for a specific service from a Catalog
acquired during the v3 identity service.

The specified EndpointOpts are used to identify a unique, unambiguous endpoint
to return. It's an error both when multiple endpoints match the provided
criteria and when none do. The minimum that can be specified is a Type, but you
will also often need to specify a Name and/or a Region depending on what's
available on your OpenStack deployment.
*/
func V3EndpointURL(catalog *tokens3.ServiceCatalog, opts golangsdk.EndpointOpts) (string, error) {
	// Extract Endpoints from the catalog entries that match the requested Type, Interface,
	// Name if provided, and Region if provided.
	var endpoints = make([]tokens3.Endpoint, 0, 1)
	for _, entry := range catalog.Entries {
		if (opts.Type == "" || entry.Type == opts.Type) && (opts.Name == "" || entry.Name == opts.Name) {
			for _, endpoint := range entry.Endpoints {
				if opts.Availability != golangsdk.AvailabilityAdmin &&
					opts.Availability != golangsdk.AvailabilityPublic &&
					opts.Availability != golangsdk.AvailabilityInternal {
					err := &ErrInvalidAvailabilityProvided{}
					err.Argument = "Availability"
					err.Value = opts.Availability
					return "", err
				}
				if (opts.Availability == golangsdk.Availability(endpoint.Interface)) &&
					(opts.Region == "" || endpoint.Region == opts.Region) {
					endpoints = append(endpoints, endpoint)
				}
			}
		}
	}

	// Report an error if the options were ambiguous.
	if opts.Type != "" {
		if len(endpoints) > 1 {
			return "", ErrMultipleMatchingEndpointsV3{Endpoints: endpoints}
		} else if len(endpoints) < 1 {
			return buildUrlIfNotFoundV3(catalog, opts)
		}
	} else {
		if len(endpoints) > 1 {
			return endpoints[0].URL, nil
		} else if len(endpoints) < 1 {
			return "", &golangsdk.ErrEndpointNotFound{}
		}
	}

	// Extract the URL from the matching Endpoint.
	for _, endpoint := range endpoints {
		return golangsdk.NormalizeURL(endpoint.URL), nil
	}

	// Report an error if there were no matching endpoints.
	err := &golangsdk.ErrEndpointNotFound{}
	return "", err
}

/*
 buildUrlIfNotFound builds an endpoint if it is not found in identity service response
*/
func buildUrlIfNotFoundV2(catalog *tokens2.ServiceCatalog, opts golangsdk.EndpointOpts) (string, error) {

	tmpOpts := opts
	tmpOpts.Type = ""

	existingUrl, err := V2EndpointURL(catalog, tmpOpts)
	return generateEndpointUrlWithExisting(existingUrl, opts, err)
}

/*
 buildUrlIfNotFound builds an endpoint if it is not found in identity service response
*/
func buildUrlIfNotFoundV3(catalog *tokens3.ServiceCatalog, opts golangsdk.EndpointOpts) (string, error) {

	tmpOpts := opts
	tmpOpts.Type = ""

	existingUrl, err := V3EndpointURL(catalog, tmpOpts)
	return generateEndpointUrlWithExisting(existingUrl, opts, err)
}

// internal method for extract a valid endpoint address
func generateEndpointUrlWithExisting(existingUrl string, opts golangsdk.EndpointOpts, err error) (string, error) {
	if err != nil || existingUrl == "" {
		return "", errors.New("No suitable endpoint could be found in the service catalog.")
	}

	existingUrl = golangsdk.NormalizeURL(existingUrl)
	u, _ := url.Parse(existingUrl)
	urlDomainParts := strings.Split(u.Host, ".")

	if len(urlDomainParts) < 2 {
		return "", errors.New("No suitable endpoint could be found in the service catalog.")
	}

	var urlParts []string
	urlParts = append(urlParts, opts.Type)

	if _, ok := allRegionInOneEndpoint[opts.Type]; ok {
		urlParts = append(urlParts, urlDomainParts[(len(urlDomainParts)-2):]...)
	} else {
		if len(urlDomainParts) > 3 {
			// such as https://kms.cn-north-1.myhwclouds.com
			urlParts = append(urlParts, urlDomainParts[(len(urlDomainParts)-3):]...)
		} else {
			// such as https://kms.myhwclouds.com
			urlParts = append(urlParts, opts.Region)
			urlParts = append(urlParts, urlDomainParts[len(urlDomainParts)-2:]...)
		}
	}
	return u.Scheme + "://" + strings.Join(urlParts, ".") + "/", nil
}
