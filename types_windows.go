// +build windows

package serviceresolver

// GUID represents a globally unique identifier as used in the Windows APIs.
type GUID [16]byte

// PrimaryDomainInfo contains information about the primary domain of a
// Windows computer.
type PrimaryDomainInfo struct {
	MachineRole      int32
	Flags            uint32
	DomainNameFlat   string
	DomainNameDNS    string
	DomainForestName string
	DomainGUID       GUID
}

type _PrimaryDomainInfo struct {
	MachineRole      int32
	Flags            uint32
	DomainNameFlat   *uint16
	DomainNameDNS    *uint16
	DomainForestName *uint16
	DomainGUID       GUID
}
