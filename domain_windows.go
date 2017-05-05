// +build windows

package serviceresolver

import "fmt"

func lookupDomains() (domains []string, err error) {
	info, err := dsRoleGetPrimaryDomainInformation("")
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve primary domain information: %v", err)
	}

	if info.DomainNameDNS != "" {
		domains = append(domains, info.DomainNameDNS)
	}

	if info.DomainForestName != "" && info.DomainForestName != info.DomainNameDNS {
		domains = append(domains, info.DomainForestName)
	}

	return
}
