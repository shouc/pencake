package pencake

import (
	"pencake/utils"
	"strings"
)

var vmNames = [...]string{
	"VirtualBox",
	"Oracle", "VMWare", "Parallels", "Qemu",
	"Microsoft VirtualPC", "Virtuozzo", "Xen",
}

func VmScan() string {
	switch utils.SystemType() {
	case "darwin":
		return checkMacVM()
	default:
		utils.NotSupported("")
		return ""
	}

}

func checkMacVM() string {
	output := utils.RunCommand("ioreg -l | grep -e Manufacturer -e 'Vendor Name'")
	for _, vmName := range vmNames {
		if strings.Contains(output, vmName) {
			return vmName
		}
	}
	return ""
}
