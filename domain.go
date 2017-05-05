// +build !windows

package serviceresolver

import "errors"

func lookupDomains() ([]string, error) {
	return nil, errors.New("service resolver is only supported on windows machines")
}
