package pencake

import (
	"pencake/utils"
)

func Unsleep() {
	switch utils.SystemType() {
	case "darwin":
		utils.RunCommand("caffeinate -u -t 2")
		return
	case "linux":
		utils.RunCommand("xset dpms force on")
	default:
		utils.NotSupported("")
		return
	}
}
