package serviceresolver

import (
	"errors"
	"fmt"
	"net"
)

// Service decribes a service address set for a particular domain.
type Service struct {
	Name   string
	Domain string
	CNAME  string // CNAME is the canonical DNS address of the record that was queried for the service set.
	Addrs  []*net.SRV
}

func lookupServices(service string, domains []string) (services []Service, err error) {
	if len(domains) == 0 {
		err = errors.New("no domains detected for service resolution")
		return
	}

	failed := 0
	var lookupErr error

	for _, domain := range domains {
		result := Service{
			Name:   service,
			Domain: domain,
		}

		result.CNAME, result.Addrs, lookupErr = net.LookupSRV(service, "tcp", domain)

		if lookupErr == nil {
			services = append(services, result)
		} else {
			failed++
		}
	}

	if len(services) == 0 && failed > 0 {
		if failed == 1 {
			err = fmt.Errorf("DNS service record lookup failed: %v", lookupErr)
		} else {
			err = fmt.Errorf("DNS service record lookup failed for %d domain(s): %v)", failed, lookupErr)
		}
	}

	return
}
