package pencake

import "pencake/utils"

func Sleep() {
	switch utils.SystemType() {
	case "darwin":
		utils.RunCommand("pmset displaysleepnow")
		return
	case "windows":
		utils.RunCommand(`%windir%\System32\rundll32.exe powrprof.dll,SetSuspendState 0,1,0`)
		return
	case "linux":
		utils.RunCommand("xset dpms force off")
		return
	default:
		utils.NotSupported("")
		return
	}
}
