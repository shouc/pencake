package pencake

import "pencake/utils"

func Restart() {
	switch utils.SystemType() {
	case "windows":
		utils.RunCommand(`%windir%\System32\shutdown.exe -r`)
		return
	default:
		utils.RunCommand("reboot")
		return
	}
}
