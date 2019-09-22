package pencake

import "pencake/utils"

func Shutdown() {
	switch utils.SystemType() {
	case "windows":
		utils.RunCommand(`%windir%\System32\shutdown.exe -s`)
		return
	default:
		utils.RunCommand("halt -p")
		return
	}
}
