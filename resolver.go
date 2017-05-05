package serviceresolver

import (
	"context"
	"fmt"
)

// Resolver attempts to locate service addresses in a zero configuration
// environment by querying DNS service records in the host's domain(s).
//
// The zero value of Resolver is safe for use.
type Resolver struct {
}

// DefaultResolver is a service resolver with default configuration.
var DefaultResolver = &Resolver{}

// Resolve returns an ordered slice of services by looking up DNS service
// records for the given service name.
//
// A host may be affiliated with more than one domain; each service in the
// returned service set represents a successful query against a possible domain.
//
// The slice of addresses contained in each service are ordered according to the
// same rules applied by net.LookupSRV.
func (r *Resolver) Resolve(ctx context.Context, service string) (services []Service, err error) {
	domains, err := lookupDomains()
	if err != nil {
		return nil, fmt.Errorf("serviceresolver: %v", err)
	}
	services, err = lookupServices(service, domains)
	if err != nil {
		return nil, fmt.Errorf("serviceresolver: %v", err)
	}
	return
}
