// +build windows

package serviceresolver

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modnetapi32 = windows.NewLazySystemDLL("netapi32.dll")

	procDsRoleGetPrimaryDomainInformation = modnetapi32.NewProc("DsRoleGetPrimaryDomainInformation")
	procDsRoleFreeMemory                  = modnetapi32.NewProc("DsRoleFreeMemory")
)

func dsRoleGetPrimaryDomainInformation(server string) (info PrimaryDomainInfo, err error) {
	var serverp *uint16
	if server != "" {
		serverp, err = syscall.UTF16PtrFromString(server)
		if err != nil {
			return
		}
	}

	var pInfo *_PrimaryDomainInfo

	r0, _, _ := syscall.Syscall(
		procDsRoleGetPrimaryDomainInformation.Addr(),
		3,
		uintptr(unsafe.Pointer(serverp)),
		uintptr(_DsRolePrimaryDomainInfoBasic),
		uintptr(unsafe.Pointer(&pInfo)),
	)
	if r0 != 0 {
		err = syscall.Errno(r0)
		return
	}
	defer dsRoleFreeMemory((*byte)(unsafe.Pointer(pInfo)))

	info.MachineRole = pInfo.MachineRole
	info.Flags = pInfo.Flags
	info.DomainNameFlat = syscall.UTF16ToString((*[256]uint16)(unsafe.Pointer(pInfo.DomainNameFlat))[:])
	info.DomainNameDNS = syscall.UTF16ToString((*[256]uint16)(unsafe.Pointer(pInfo.DomainNameDNS))[:])
	info.DomainForestName = syscall.UTF16ToString((*[256]uint16)(unsafe.Pointer(pInfo.DomainForestName))[:])
	info.DomainGUID = pInfo.DomainGUID

	return
}

func dsRoleFreeMemory(ptr *byte) (err error) {
	r0, _, _ := syscall.Syscall(
		procDsRoleFreeMemory.Addr(),
		1,
		uintptr(unsafe.Pointer(ptr)),
		0,
		0,
	)
	if r0 != 0 {
		err = syscall.Errno(r0)
	}
	return
}
